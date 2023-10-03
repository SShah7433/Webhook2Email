package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Email struct {
	From string `json:"from" binding:"required"`

	To  []string `json:"to" binding:"required"`
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`

	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.Default()
	r.POST("/sendemails", func(c *gin.Context) {

		var f []Email
		if err := c.BindJSON(&f); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation of payload failed"})
			return
		}

		user := os.Getenv("W2E_USER")
		password := os.Getenv("W2E_PASS")

		host := os.Getenv("W2E_HOST")
		port, _ := strconv.Atoi(os.Getenv("W2E_PORT"))

		d := gomail.NewDialer(host, port, user, password)
		s, err := d.Dial()
		if err != nil {
			panic(err)
		}

		m := gomail.NewMessage()
		for _, value := range f {
			m.SetHeader("From", value.From)
			m.SetHeader("To", value.To...)
			m.SetHeader("Cc", value.Cc...)
			m.SetHeader("Bcc", value.Bcc...)
			m.SetHeader("Subject", value.Subject)
			m.SetBody("text/plain", value.Message)

			if err := gomail.Send(s, m); err != nil {
				c.Status(http.StatusInternalServerError)
			}
			m.Reset()

		}

		c.Status(http.StatusOK)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
