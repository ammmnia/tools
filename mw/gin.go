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
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ammmnia/tools/apiresp"
	"github.com/ammmnia/tools/constant"
	"github.com/ammmnia/tools/errs"
	"github.com/ammmnia/tools/log"
	"github.com/ammmnia/tools/tokenverify"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CorsHandler gin cross-domain configuration.
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header(
			"Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
		) // Cross-domain key settings allow browsers to resolve.
		c.Header(
			"Access-Control-Max-Age",
			"172800",
		) // Cache request information in seconds.
		c.Header(
			"Access-Control-Allow-Credentials",
			"false",
		) //  Whether cross-domain requests need to carry cookie information, the default setting is true.
		c.Header(
			"content-type",
			"application/json",
		) // Set the return format to json.
		// Release all option pre-requests
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, "Options Request!")
			c.Abort()
			return
		}
		c.Next()
	}
}

func GinParseOperationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			operationID := c.Request.Header.Get(string(constant.OperationID))
			if operationID == "" {
				err := errs.New("header must have operationID")
				apiresp.GinError(c, errs.ErrArgs.WrapMsg(err.Error()))
				c.Abort()
				return
			}
			c.Set(string(constant.OperationID), operationID)
		}
		c.Next()
	}
}

func GinParseToken(secretKey jwt.Keyfunc, whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			for _, wApi := range whitelist {
				if strings.HasPrefix(c.Request.URL.Path, wApi) {
					c.Next()
					return
				}
			}

			token := c.Request.Header.Get(string(constant.Token))
			if token == "" {
				apiresp.GinError(c, errs.ErrArgs.WrapMsg("header must have token"))
				c.Abort()
				return
			}

			claims, err := tokenverify.GetClaimFromToken(token, secretKey)
			if err != nil {
				log.ZWarn(c, "header get token error", errs.ErrArgs.WrapMsg("header must have token"))
				apiresp.GinError(c, errs.ErrArgs.WrapMsg("header must have token"))
				c.Abort()
				return
			}

			c.Set(string(constant.OpUserPlatform), constant.PlatformIDToName(claims.PlatformID))
			c.Set(string(constant.OpUserID), claims.UserID)
			c.Next()
		}
	}
}

func CreateToken(userID string, accessSecret string, accessExpire int64, platformID constant.PlatformID) (string, error) {
	claims := tokenverify.BuildClaims(userID, platformID, accessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", errs.WrapMsg(err, "token.SignedString")
	}
	return tokenString, nil
}

// Logger instance a Logger middleware with config.
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		//raw := c.Request.URL.RawQuery

		// Process request
		c.Next()
		if c.Writer.Status() == http.StatusOK {
			log.ZInfo(c, path, "ClientIP", c.ClientIP(), "Method", c.Request.Method, "StatusCode", c.Writer.Status(), "Latency", time.Since(start))
		} else {
			log.ZError(c, path, errors.New(c.Errors.ByType(gin.ErrorTypePrivate).String()), "ClientIP", c.ClientIP(), "Method", c.Request.Method, "StatusCode", c.Writer.Status(), "Latency", time.Since(start))
		}

	}
}
