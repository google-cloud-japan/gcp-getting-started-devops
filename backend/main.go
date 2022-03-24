/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"golang.org/x/oauth2/google"
)

var projectID = ""
var onGoogleCloud = true

const (
	service        = "devops-demo"
	serviceVersion = "1.0.0"
	metricPrefix   = "devops-"
)

func logRequestWithTrace(r *http.Request, traceID string, spanID string, msg string) {
	log.WithFields(log.Fields{
		"UserAgent":                     r.UserAgent(),
		"RequestURL":                    r.Host,
		"RequestURI":                    r.RequestURI,
		"RequestMethod":                 r.Method,
		"Host":                          r.Host,
		"Proto":                         r.Proto,
		"ProtoMajor":                    r.ProtoMajor,
		"ProtoMinor":                    r.ProtoMinor,
		"RemoteAddr":                    r.RemoteAddr,
		"logging.googleapis.com/trace":  "projects/" + projectID + "/traces/" + traceID,
		"logging.googleapis.com/spanId": spanID,
	}).Info(msg)
}

func logRequest(r *http.Request, msg string) {
	log.WithFields(log.Fields{
		"UserAgent":                     r.UserAgent(),
		"RequestURL":                    r.Host,
		"RequestURI":                    r.RequestURI,
		"RequestMethod":                 r.Method,
		"Host":                          r.Host,
		"Proto":                         r.Proto,
		"ProtoMajor":                    r.ProtoMajor,
		"ProtoMinor":                    r.ProtoMinor,
		"RemoteAddr":                    r.RemoteAddr,
	}).Info(msg)
}

func logMethod(traceID string, spanID string, msg string) {
	log.WithFields(log.Fields{
		"logging.googleapis.com/trace":  "projects/" + projectID + "/traces/" + traceID,
		"logging.googleapis.com/spanId": spanID,
	}).Info(msg)
}

func normalHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if onGoogleCloud {
			_, span := trace.StartSpan(context.Background(), "/normal")
			defer span.End()
			s := span.SpanContext()

			addAttributesToSpan(span, r)
			logRequestWithTrace(r, s.TraceID.String(), s.SpanID.String(), "Access to normal")
		} else {
			logRequest(r, "Access to normal")
		}

		returnElapsedTimeAsJSON(w, start)
	})
}

func benchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if onGoogleCloud {
			ctx, span := trace.StartSpan(context.Background(), "/bench")
			defer span.End()
			s := span.SpanContext()

			addAttributesToSpan(span, r)
			logRequestWithTrace(r, s.TraceID.String(), s.SpanID.String(), "Access to bench")

			// Stress
			fibonacci(ctx, rand.Intn(3000000000))
			fibonacci(ctx, rand.Intn(3000000000))
			fibonacci(ctx, rand.Intn(3000000000))

			log.Printf("context: %v\n", ctx)
		} else {
			fibonacciOnLocal(rand.Intn(3000000000))
			fibonacciOnLocal(rand.Intn(3000000000))
			fibonacciOnLocal(rand.Intn(3000000000))

			logRequest(r, "Access to bench")
		}

		returnElapsedTimeAsJSON(w, start)
	})
}

func returnElapsedTimeAsJSON(w http.ResponseWriter, start time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"elapsed": %d}`, time.Since(start).Milliseconds())))
}

func addAttributesToSpan(span *trace.Span, r *http.Request) {
	span.AddAttributes(
		trace.StringAttribute("Method", r.Method),
		trace.StringAttribute("Host", r.Host),
		trace.StringAttribute("Proto", r.Proto),
		trace.Int64Attribute("ProtoMajor", int64(r.ProtoMajor)),
		trace.Int64Attribute("ProtoMinor", int64(r.ProtoMinor)),
		trace.Int64Attribute("ContentLength", int64(r.ContentLength)),
		trace.StringAttribute("RequestURI", r.RequestURI),
		trace.StringAttribute("RemoteAddr", r.RemoteAddr),
		trace.StringAttribute("UserAgent", r.UserAgent()),
		trace.StringAttribute("Referer", r.Referer()),
	)
}

func fibonacci(ctx context.Context, number int) int {
	_, span := trace.StartSpan(ctx, "fibonacci")
	s := span.SpanContext()
	defer span.End()

	result := fibonacciOnLocal(number)

	span.AddAttributes(trace.Int64Attribute("loopNum", int64(number)))
	logMethod(s.TraceID.String(), s.SpanID.String(), fmt.Sprintf("Fibonacci calculation completed: %v loops", number))

	return result
}

func fibonacciOnLocal(number int) int {
	prev, next := 0, 1
	if number < 2 {
		return number
	}
	for i := 1; i < number; i++ {
		prev, next = next, prev+next
	}
	return next
}

func initLogSettings() {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "time",
			log.FieldKeyLevel: "severity",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	initLogSettings()

	// Initialize the credential and projectID with Application Default Credentials (ADC)
	ctx := context.Background()
	cred, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	projectID = cred.ProjectID
	if len(projectID) == 0 {
		log.Println("Failed to get the credential. Running as local mode...")
		onGoogleCloud = false
	} else {
		log.Infof("Successfully retrieved GCP Project ID: %v", projectID)
	}

	if onGoogleCloud {
		// Initialize Profiler
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: serviceVersion,
			ProjectID:      projectID,
		}); err != nil {
			log.Fatal(err.Error())
		}
		log.Info("Successfully initialized profiler")

		// Initialize exporter
		exporter, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:    projectID,
			MetricPrefix: metricPrefix,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Info("Successfully initialized exporter")
		defer exporter.Flush()
		trace.RegisterExporter(exporter)
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	}

	// Serve Frontend scripts under /static/ directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Register handlers
	http.Handle("/normal", normalHandler())
	http.Handle("/bench", benchHandler())

	log.Info("Start listening : 8080...")
	http.ListenAndServe(":8080", nil)
}
