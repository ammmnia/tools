// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errs

import (
	"github.com/ammmnia/tools/i18n"
	"strconv"
	"strings"
)

var DefaultCodeRelation = newCodeRelation()

type CodeError interface {
	Code() int
	Msg() string
	Detail() string
	WithDetail(detail string) CodeError
	Error
}

func NewCodeError(code int, msg string) CodeError {
	return &codeError{
		code: code,
		msg:  msg,
	}
}

type codeError struct {
	code   int
	msg    string
	detail string
}

func (e *codeError) Code() int {
	return e.code
}

func (e *codeError) Msg() string {
	return e.msg
}

func (e *codeError) Detail() string {
	return e.detail
}

func (e *codeError) WithDetail(detail string) CodeError {
	var d string
	if e.detail == "" {
		d = detail
	} else {
		d = e.detail + ", " + detail
	}
	return &codeError{
		code:   e.code,
		msg:    e.msg,
		detail: d,
	}
}

func (e *codeError) Wrap() error {
	return Wrap(e)
}

func (e *codeError) WrapMsg(msg string, kv ...any) error {
	return WrapMsg(e, msg, kv...)
}

func (e *codeError) WrapLocal(kv ...any) error {
	tmp := map[string]interface{}{}
	for i := 0; i < len(kv); i += 2 {
		tmp[kv[i].(string)] = kv[i+1]
	}
	message, err := i18n.Locale.Localize(&i18n.LocalizeConfig{
		MessageID:    e.msg,
		TemplateData: tmp,
	})
	if err != nil {
		return Wrap(e)
	}
	e.msg = message
	return WrapMessage(e, message)
}

func (e *codeError) Is(err error) bool {
	codeErr, ok := Unwrap(err).(CodeError)
	if !ok {
		if err == nil && e == nil {
			return true
		}
		return false
	}
	if e == nil {
		return false
	}
	code := codeErr.Code()
	if e.code == code {
		return true
	}
	return DefaultCodeRelation.Is(e.code, code)
}

const initialCapacity = 3

func (e *codeError) Error() string {
	v := make([]string, 0, initialCapacity)
	v = append(v, strconv.Itoa(e.code), e.msg)

	if e.detail != "" {
		v = append(v, e.detail)
	}

	return strings.Join(v, " ")
}
