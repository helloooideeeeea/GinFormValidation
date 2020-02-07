package GinFormValidation
import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

type ErrorJsonSchema struct {
	Errors []ErrorJsonSchemaChild `json:"errors"`
}

type ErrorJsonSchemaChild struct {
	Column		string		`json:"column"`
	Contents	[]string	`json:"messages"`
}

func ErrorsToJson(form interface{}, err error, messageTranslationCallback func(message string) string) interface{} {

	errors := err.(validator.ValidationErrors)
	f := reflect.TypeOf(form)

	parent := ErrorJsonSchema{}

	for _, e := range errors {
		field, _ := f.FieldByName(e.Field())
		tags := []string{}
		items := strings.Split(field.Tag.Get("bind_error"), ",")
		for _, item := range items {
			keyValue := strings.Split(item,"=")
			if e.Tag() == strings.TrimSpace(keyValue[0]) {
				var translatedValue string
				if messageTranslationCallback != nil {
					translatedValue = messageTranslationCallback(keyValue[1])
				} else {
					translatedValue = keyValue[1]
				}
				tags = append(tags, strings.TrimSpace(translatedValue))
			}
		}

		child := ErrorJsonSchemaChild{}
		child.Column = e.Field()
		child.Contents = tags

		parent.Errors = append(parent.Errors, child)
	}

	return parent
}
