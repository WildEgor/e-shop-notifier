package validators

import (
	"encoding/json"
	core_dtos "github.com/WildEgor/e-shop-gopack/pkg/core/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// TODO: add custom validators

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// FIXME
	//if err != nil {
	//	// Make error message for each invalid field.
	//	for _, err := range err.(validator.ValidationErrors) {
	//		fields[err.Field()] = err.Error()
	//	}
	//}

	return fields
}

// ParseAndValidateHttp parser
func ParseAndValidateHttp(ctx fiber.Ctx, out interface{}) error {
	resp := core_dtos.NewResponse(ctx)

	// Checking received data from JSON body. Return status 400 and error message.
	if err := ctx.Bind().Body(&out); err != nil {
		resp.SetStatus(fiber.StatusBadRequest)
		resp.SetMessage(err.Error())
		return nil
	}

	// Create a new validator for a RegistrationRequestDto.
	validate := NewValidator()

	// Validate fields.
	if err := validate.Struct(&out); err != nil {
		resp.SetStatus(fiber.StatusBadRequest)
		resp.SetMessage(err.Error()) // ValidatorErrors(err)
		return nil
	}

	return nil
}

func ParseAndValidateBytes(in []byte, out any) error {
	if err := json.Unmarshal(in, out); err != nil {
		return err
	}

	validate := NewValidator()

	// Validate fields.
	if err := validate.Struct(out); err != nil {
		return err
	}

	return nil
}
