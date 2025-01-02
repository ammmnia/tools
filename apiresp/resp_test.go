package apiresp

import (
	"github.com/ammmnia/tools/errs"
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

func TestName(t *testing.T) {
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
	GinError(c, errs.ErrTokenNotExist.WrapLocal())

}
