package controllers

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strings"
)

type ApiError struct {
	Param   string
	Message string
}

func ReturnErrorResponse(c echo.Context, code int, status string, err error) error {
	return c.JSON(code, map[string]interface{}{
		"status":  status,
		"message": err.Error(),
		"data":    map[string]interface{}{},
	})
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}

func ReturnFirstErrorValidation(c echo.Context, code int, status string, err error) error {
	//var errs []error
	////validate := validator.New()
	//english := en.New()
	//uni := ut.New(english, english)
	//trans, _ := uni.GetTranslator("en")
	////_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	//validatorErrs := err.(validator.ValidationErrors)
	//for _, e := range validatorErrs {
	//	translatedErr := fmt.Errorf(e.Translate(trans))
	//	errs = append(errs, translatedErr)
	//}

	trimmedValidation := strings.Trim(err.Error(), "Key: ")
	trimmedValidation = strings.Trim(trimmedValidation, "Error:")
	splittedValidations := strings.Split(trimmedValidation, " tag")
	if len(splittedValidations) <= 0 {
		return c.JSON(code, map[string]interface{}{
			"status":  status,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
	}
	firstValidation := splittedValidations[0]
	valueValidations := strings.Split(firstValidation, "Field validation for ")
	//for _, validation := range splittedValidations {
	//}
	errString := strings.Replace(valueValidations[1], "failed on the", "is", 1)
	normalizedErrString := strings.ToLower(strings.Replace(errString, "'", "", 4))
	normalizedErrString = strings.Replace(normalizedErrString, "is required", "cannot be null", 1)
	if strings.Contains(normalizedErrString, "activitygroupid") {
		normalizedErrString = strings.Replace(normalizedErrString, "activitygroupid", "activity_group_id", 1)
	}

	// return first validation error value
	return c.JSON(code, map[string]interface{}{
		"status":  status,
		"message": normalizedErrString,
	})
}
func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
