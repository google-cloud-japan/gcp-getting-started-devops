/*
Copyright 2021 Google LLC

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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Response struct {
	Elapsed *int `json:"elapsed"`
}

func TestMain(m *testing.M) {
	// Run tests on local
	isLocal = true
	status := m.Run()
	os.Exit(status)
}

func TestFibonacciOnLocal(t *testing.T) {
	testcases := []struct {
		name string
		arg  int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"10", 10, 55},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciOnLocal(tt.arg); got != tt.want {
				t.Errorf(`fibonacciOnLocal(%d) = %v, want %v`, tt.arg, got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	testcases := []struct {
		name string
		arg  int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"10", 10, 55},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if got := fibonacci(ctx, tt.arg); got != tt.want {
				t.Errorf(`fibonacci(%d) = %v, want %v`, tt.arg, got, tt.want)
			}
		})
	}
}

func TestNormalHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/normal", nil)
	rr := httptest.NewRecorder()

	normalHandler().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(`Normal handler returned wrong status code: got %v want %v`,
			status, http.StatusOK)
	}
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf(`Normal handler returned wrong Content-Type : got %v want application/json`, contentType)
	}
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf(`Normal handler must return JSON string: got %v`, rr.Body.String())
	}
	if response.Elapsed == nil {
		t.Errorf(`Normal handler returned wrong JSON format: got %v`, rr.Body.String())
	}
}

func TestBenchHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/bench", nil)
	rr := httptest.NewRecorder()

	benchHandler().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(`Bench handler returned wrong status code: got %v want %v`,
			status, http.StatusOK)
	}
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf(`Bench handler returned wrong Content-Type : got %v want "application/json"`, contentType)
	}
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf(`Bench handler must return JSON string: got %v`, rr.Body.String())
	}
	if response.Elapsed == nil {
		t.Errorf(`Bench handler returned wrong JSON format: got %v`, rr.Body.String())
	}
}
