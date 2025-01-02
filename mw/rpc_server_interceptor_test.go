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

package mw

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 示例处理函数
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func TestHandleError(t *testing.T) {
	// 创建一个新的 Gin 引擎
	router := gin.New()
	router.GET("/hello", HelloHandler)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	// 创建一个新的 gin.Context 实例
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// 调用处理函数
	router.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Hello, World!"}`, w.Body.String())
	//HandleError(c, "dxtest", "req", errs.ErrTokenNotExist.WrapLocal())
}
