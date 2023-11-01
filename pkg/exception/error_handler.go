package exception

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gocron.com/m/pkg/helper"
	"strings"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if ok {
		return ctx.Status(fiber.StatusOK).JSON(helper.ApiResponseFail(ValidationErrors(err), "Validation error"))
	}

	var validationError ValidationError
	ok = errors.As(err, &validationError)
	if ok {
		return ctx.Status(fiber.StatusOK).JSON(helper.ApiResponseFail(err.Error(), "Validation error"))
	}

	var errorInterface error
	ok = errors.As(err, &errorInterface)
	if ok {
		return ctx.Status(fiber.StatusOK).JSON(helper.ApiResponseFail("", err.Error()))
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ApiResponseFail("", "Something wrong, please contact admin"))
}

func PanicIfNeeded(err interface{}) {
	if err != nil {
		helper.SaveLogError(err.(error).Error(), err)
		panic(err)
	}
}

func IsError(err error, function string) {
	//check if error is not from gorm or validator
	if err != nil &&
		!errors.Is(err, validator.ValidationErrors{}) &&
		!errors.Is(err, ValidationError{}) {
		helper.SaveLogError(fmt.Sprintf("Error from %s: %s", function, err.Error()), map[string]interface{}{
			"function": function,
			"error":    err,
		})
	}
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", strings.ToLower(fe.Field()))
	case "email":
		return "Invalid email address"
	case "gte":
		return fmt.Sprintf("Field %s must be greater than %s", strings.ToLower(fe.Field()), fe.Param())
	case "lte":
		return fmt.Sprintf("Field %s must be lower than %s", strings.ToLower(fe.Field()), fe.Param())
	case "eqfield":
		return fmt.Sprintf("Field %s must be equal with %s", strings.ToLower(fe.Field()), fe.Param())
	case "nefield":
		return fmt.Sprintf("Field %s must be not equal with %s", strings.ToLower(fe.Field()), fe.Param())
	case "gtfield":
		return fmt.Sprintf("Field %s must be greater than %s", strings.ToLower(fe.Field()), fe.Param())
	case "gtefield":
		return fmt.Sprintf("Field %s must be greater than or equal with %s", strings.ToLower(fe.Field()), fe.Param())
	case "ltfield":
		return fmt.Sprintf("Field %s must be lower than %s", strings.ToLower(fe.Field()), fe.Param())
	case "ltefield":
		return fmt.Sprintf("Field %s must be lower than or equal with %s", strings.ToLower(fe.Field()), fe.Param())
	case "oneof":
		return fmt.Sprintf("Field %s must be one of %s", strings.ToLower(fe.Field()), fe.Param())
	case "unique":
		return fmt.Sprintf("Field %s must be unique", strings.ToLower(fe.Field()))
	case "numeric":
		return fmt.Sprintf("Field %s must be numeric", strings.ToLower(fe.Field()))
	case "alphanum":
		return fmt.Sprintf("Field %s must be alphanumeric", strings.ToLower(fe.Field()))
	case "alphanumunicode":
		return fmt.Sprintf("Field %s must be alphanumeric unicode", strings.ToLower(fe.Field()))
	case "datetime":
		return fmt.Sprintf("Field %s must be datetime", strings.ToLower(fe.Field()))
	}
	return fe.Error() // default error
}

// ValidationErrors is a function to return validation error
func ValidationErrors(err error) any {

	//validate if error is validator.ValidationErrors
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		return err.Error()
	}
	errMap := make(map[string]string)
	for _, e := range errs {
		errMap[strings.ToLower(e.Field())] = msgForTag(e)
	}
	return errMap
}
