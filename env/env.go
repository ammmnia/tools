// Copyright © 2024 OpenIM open source community. All rights reserved.
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

package env

import (
	"os"
	"strconv"

	"github.com/ammmnia/tools/errs"
)

// GetString returns the env variable for the given key
// and falls back to the given defaultValue if not set.
func GetString(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}

// GetInt returns the env variable (parsed as integer) for
// the given key and falls back to the given defaultValue if not set.
func GetInt(key string, defaultValue int) (int, error) {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue, errs.WrapMsg(err, "Atoi failed", "value", v)
		}
		return value, nil
	}
	return defaultValue, nil
}

// GetFloat64 returns the env variable (parsed as float64) for
// the given key and falls back to the given defaultValue if not set.
func GetFloat64(key string, defaultValue float64) (float64, error) {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue, errs.WrapMsg(err, "ParseFloat failed", "value", v)
		}
		return value, nil
	}
	return defaultValue, nil
}

// GetBool returns the env variable (parsed as bool) for
// the given key and falls back to the given defaultValue if not set.
func GetBool(key string, defaultValue bool) (bool, error) {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue, errs.WrapMsg(err, "ParseBool failed", "value", v)
		}
		return value, nil
	}
	return defaultValue, nil
}
