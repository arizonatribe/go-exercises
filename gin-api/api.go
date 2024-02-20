package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func handleUserName(c *gin.Context) {
	user := c.Params.ByName("name")
	val, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": val})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func handleAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		_, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			db[user] = json.Value
			c.JSON(http.StatusCreated, gin.H{"status": "created"})
		}
	}
}

/* example curl for /admin with basicauth header
  dGltOlBAc3N3MHJk is base64("tim:P@ssw0rd")

curl -X POST \
  http://localhost:8080/admin \
  -H 'authorization: Basic dGltOlBAc3N3MHJk' \
  -H 'content-type: application/json' \
  -d '{"value":"P@ssw0rd"}'
*/

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlePing)
	r.GET("/user/:name", handleUserName)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"tim":  "P@ssw0rd", // dGltOlBAc3N3MHJk
		"jane": "H311o",    // amFuZTpIMzExbw==
	}))

	authorized.POST("/admin", handleAdmin)

	return r
}
