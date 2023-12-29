/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package problem

import (
	"encoding/json"
	"net/http"
)

// Details represents a problem detail as defined in RFC 7807.
// It provides a standardized format for HTTP API responses to provide detailed information about errors.
type Details struct {
	Type     string  `json:"type"`               // A URI that identifies the problem type.
	Title    string  `json:"title"`              // A short, human-readable title that describes the problem type.
	Status   *int    `json:"status,omitempty"`   // The HTTP status code generated by the server.
	Detail   *string `json:"detail,omitempty"`   // A human-readable explanation providing more context about the problem.
	Instance *string `json:"instance,omitempty"` // A URI that identifies the specific instance of the problem.
}

// NewProblemDetails creates a new ProblemDetails instance and sets the required fields.
func NewProblemDetails(t string, title string) *Details {
	return &Details{
		Type:  t,
		Title: title,
	}
}

// SetStatus sets the HTTP status code for the problem details.
func (p *Details) SetStatus(s int) {
	p.Status = &s
}

// SetDetail sets the detailed explanation for the problem details.
func (p *Details) SetDetail(d string) {
	p.Detail = &d
}

// SetInstance sets the instance URI for the problem details.
func (p *Details) SetInstance(i string) {
	p.Instance = &i
}

// Respond writes the ProblemDetails instance to the HTTP response writer as JSON.
func (p *Details) Respond(w http.ResponseWriter, status int) error {
	// Set the HTTP status code and content type
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(status)

	// Convert the ProblemDetails instance to JSON
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		return err
	}

	return nil
}
