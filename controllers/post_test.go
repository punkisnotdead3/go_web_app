package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	url := "/api/v1/post"
	r := gin.Default()
	r.POST(url, CreatePostHandler)
	w := httptest.NewRecorder()
	body := `{
    "title":"新增呢2",
    "content":"vivo比华1231为好多了",
    "community_id":2
         }`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// 判断响应中 是不是包含了 未登录的 错误信息
	assert.Contains(t, w.Body.String(), "未登录")
	// 也可以继续用下面的方法来判定
	res := new(Response)
	err := json.Unmarshal(w.Body.Bytes(), res)
	if err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}
	assert.Equal(t, res.Code, CodeNoLogin)
}
