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

package problem_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitub.com/sixwaaaay/problem"
)

func TestNewProblemDetails(t *testing.T) {
	http.StatusText(404)
	details := problem.NewProblemDetails("type", "title")

	assert.Equal(t, "type", details.Type)
	assert.Equal(t, "title", details.Title)
}

func TestSetStatus(t *testing.T) {
	details := problem.NewProblemDetails("type", "title")
	status := 404
	details.SetStatus(status)

	assert.Equal(t, status, *details.Status)
}

func TestSetDetail(t *testing.T) {
	details := problem.NewProblemDetails("type", "title")
	detail := "detail"
	details.SetDetail(detail)

	assert.Equal(t, detail, *details.Detail)
}

func TestSetInstance(t *testing.T) {
	details := problem.NewProblemDetails("type", "title")
	instance := "instance"
	details.SetInstance(instance)

	assert.Equal(t, instance, *details.Instance)
}

func TestRespond(t *testing.T) {
	details := problem.NewProblemDetails("type", "title")
	details.SetStatus(404)
	details.SetDetail("detail")
	details.SetInstance("instance")

	recorder := httptest.NewRecorder()
	err := details.Respond(recorder, http.StatusNotFound)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, "application/problem+json; charset=utf-8", recorder.Header().Get("Content-Type"))

	var response problem.Details
	err = json.NewDecoder(recorder.Body).Decode(&response)

	assert.Nil(t, err)
	assert.Equal(t, details, &response)
}
