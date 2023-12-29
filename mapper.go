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

import "net/http"

// StatusRFC returns a rfc text for the HTTP status code. It returns the empty
// string if the code is unknown.
// support HTTP status codes registered with IANA.
// https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func StatusRFC(code int) string {
	switch code {
	case http.StatusContinue:
		return "https://tools.ietf.org/html/rfc9110#section-15.2.1"
	case http.StatusSwitchingProtocols:
		return "https://tools.ietf.org/html/rfc9110#section-15.2.2"
	case http.StatusProcessing:
		return "https://tools.ietf.org/html/rfc2518#section-10.1"
	case http.StatusEarlyHints:
		return "https://tools.ietf.org/html/rfc8297#section-2"
	case http.StatusOK:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.1"
	case http.StatusCreated:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.2"
	case http.StatusAccepted:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.3"
	case http.StatusNonAuthoritativeInfo:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.4"
	case http.StatusNoContent:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.5"
	case http.StatusResetContent:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.6"
	case http.StatusPartialContent:
		return "https://tools.ietf.org/html/rfc9110#section-15.3.7"
	case http.StatusMultiStatus:
		return "https://tools.ietf.org/html/rfc4918#section-11.1"
	case http.StatusAlreadyReported:
		return "https://tools.ietf.org/html/rfc5842#section-7.1"
	case http.StatusIMUsed:
		return "https://tools.ietf.org/html/rfc3229#section-10.4.1"
	case http.StatusMultipleChoices:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.1"
	case http.StatusMovedPermanently:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.2"
	case http.StatusFound:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.3"
	case http.StatusSeeOther:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.4"
	case http.StatusNotModified:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.5"
	case http.StatusUseProxy:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.6"
	case http.StatusTemporaryRedirect:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.8"
	case http.StatusPermanentRedirect:
		return "https://tools.ietf.org/html/rfc9110#section-15.4.9"
	case http.StatusBadRequest:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.1"
	case http.StatusUnauthorized:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.2"
	case http.StatusPaymentRequired:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.3"
	case http.StatusForbidden:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.4"
	case http.StatusNotFound:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.5"
	case http.StatusMethodNotAllowed:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.6"
	case http.StatusNotAcceptable:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.7"
	case http.StatusProxyAuthRequired:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.8"
	case http.StatusRequestTimeout:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.9"
	case http.StatusConflict:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.10"
	case http.StatusGone:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.11"
	case http.StatusLengthRequired:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.12"
	case http.StatusPreconditionFailed:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.13"
	case http.StatusRequestEntityTooLarge:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.14"
	case http.StatusRequestURITooLong:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.15"
	case http.StatusUnsupportedMediaType:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.16"
	case http.StatusRequestedRangeNotSatisfiable:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.17"
	case http.StatusExpectationFailed:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.18"
	case http.StatusTeapot:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.19"
	case http.StatusMisdirectedRequest:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.20"
	case http.StatusUnprocessableEntity:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.21"
	case http.StatusLocked:
		return "https://tools.ietf.org/html/rfc4918#section-11.3"
	case http.StatusFailedDependency:
		return "https://tools.ietf.org/html/rfc4918#section-11.4"
	case http.StatusTooEarly:
		return "https://tools.ietf.org/html/rfc8470#section-5.2"
	case http.StatusUpgradeRequired:
		return "https://tools.ietf.org/html/rfc9110#section-15.5.22"
	case http.StatusPreconditionRequired:
		return "https://tools.ietf.org/html/rfc6585#section-3"
	case http.StatusTooManyRequests:
		return "https://tools.ietf.org/html/rfc6585#section-4"
	case http.StatusRequestHeaderFieldsTooLarge:
		return "https://tools.ietf.org/html/rfc6585#section-5"
	case http.StatusUnavailableForLegalReasons:
		return "https://tools.ietf.org/html/rfc7725#section-3"
	case http.StatusInternalServerError:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.1"
	case http.StatusNotImplemented:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.2"
	case http.StatusBadGateway:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.3"
	case http.StatusServiceUnavailable:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.4"
	case http.StatusGatewayTimeout:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.5"
	case http.StatusHTTPVersionNotSupported:
		return "https://tools.ietf.org/html/rfc9110#section-15.6.6"
	case http.StatusVariantAlsoNegotiates:
		return "https://tools.ietf.org/html/rfc2295#section-8.1"
	case http.StatusInsufficientStorage:
		return "https://tools.ietf.org/html/rfc4918#section-11.5"
	case http.StatusLoopDetected:
		return "https://tools.ietf.org/html/rfc5842#section-7.2"
	case http.StatusNotExtended:
		return "https://tools.ietf.org/html/rfc2774#section-7"
	case http.StatusNetworkAuthenticationRequired:
		return "https://tools.ietf.org/html/rfc6585#section-6"

	}
	return ""
}
