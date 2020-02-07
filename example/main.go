package main

import (
	"github.com/helloooideeeeea/GinFormValidation"
	"github.com/gin-gonic/gin"
	"log"
)

type signInForm struct {
	Email		string		`binding:"required,email" bind_error:"required=required_email,email=invalid_email"`
	Password	string		`binding:"required,min=6,max=20" bind_error:"required=required_passwd,max=passwd_max, min=passwd_min"`
}
/*
 this library support validation message library for gin

 this library be able to attach error message-id to binding methods(ex. required)  by struct's meta tag
 `binding: "validate-attribute" bind_error: "validate-attribute=message-id"`
 for example,
 `binding: "required" bind_error: "required=message-id"`

local test..
 $ curl -vvv -X POST -d '{"Emai":"hogehgo", "Password":"ho"}' http://localhost:8080/signIn | jq

{
  "errors": [
    {
      "column": "Email",
      "messages": [
        "required_email"
      ]
    },
    {
      "column": "Password",
      "messages": [
        "passwd_min"
      ]
    }
  ]
}
*/
func main() {
	r := gin.Default()
	r.POST("/signIn", func(c *gin.Context) {

		var form signInForm
		if err := c.BindJSON(&form); err != nil {

			//errJson := GinFormValidation.ErrorsToJson(form, err, func(message_id string) string {
			//	return translateMessageFromMessageId(message_id) <- if you translate Error message-id, you can do filter message. example, i18n
			//})

			errJson := GinFormValidation.ErrorsToJson(form, err, nil)
			c.JSON(400, errJson)
			return
		}

		c.String(200, "OK")
		return
	})
	if err := r.Run(); err != nil {
		log.Fatal("Web Server error", err)
	}
}


