package error

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

func ParseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return out
}

func parseFieldError(e validator.FieldError) string {
	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]
	log.Print(tag)
	switch tag {
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	default:
		return fmt.Errorf("%v", e).Error()
	}
}

func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
