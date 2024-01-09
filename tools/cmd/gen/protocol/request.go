// Copyright 2024 The langsvr Authors
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package protocol

// Request represents a LSP request
type Request struct {
	// Whether the request is deprecated or not. If deprecated the property contains the deprecation message
	Deprecated string
	// An optional documentation
	Documentation string
	// An optional error data type
	ErrorData Type
	// The direction in which this request is sent in the protocol
	MessageDirection MessageDirection
	// The request's method name
	Method string
	// The parameter type(s) if any
	Params []Type
	// Optional partial result type if the request supports partial result reporting
	PartialResult Type
	// Whether this is a proposed feature. If omitted the feature is final
	Proposed bool
	// Optional a dynamic registration method if it different from the request's method
	RegistrationMethod string
	// Optional registration options if the request supports dynamic registration
	RegistrationOptions Type
	// The result type
	Result Type
	// Since when (release number) this request is available. Is undefined if not known
	Since string

	// C++ class name for the request
	Name string
}
