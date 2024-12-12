package auth

import (
	"API/constants"
	"API/pkg/helper"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func SigninController(c *fiber.Ctx) error {
	var loginBody LoginDTO

	if err := c.BodyParser(&loginBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Something went wrong",
			Data:       fiber.Map{},
		})
	}

	token, refreshToken, err := Signin(loginBody.Email, loginBody.Password)
	if err != nil {
		e := err.Error()
		switch e {
		case "1":
			return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "username or password is incorrect",
				Data:       fiber.Map{},
			})
		default:
			log.Err(err)
			return c.Status(fiber.StatusInternalServerError).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Something went wrong",
				Data:       fiber.Map{},
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(constants.ApiResponse{
		StatusCode: fiber.StatusOK,
		Message:    "SignIn Successful",
		Data: fiber.Map{
			"AccessToken":  token,
			"RefreshToken": refreshToken,
		},
	})
}

func SignupController(c *fiber.Ctx) error {
	var signupBody SignupDTO

	if err := c.BodyParser(&signupBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Something went wrong",
			Data:       fiber.Map{},
		})
	}

	//check if any null values are provided
	if signupBody.Email == "" || signupBody.Password == "" || signupBody.ConfirmPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "No feild can be blank",
			Data:       fiber.Map{},
		})
	}

	//validate the recieved email address
	_, err := mail.ParseAddress(signupBody.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid Email",
			Data:       fiber.Map{},
		})
	}

	//check if the given password and confirm password are same
	if signupBody.Password != signupBody.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Password does not match with confirm password",
			Data:       fiber.Map{},
		})
	}

	err = Signup(signupBody)
	if err != nil {
		e := err.Error()
		switch e {
		case "1":
			return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "User with this emailid already exists",
				Data:       fiber.Map{},
			})
		default:
			log.Err(err)
			return c.Status(fiber.StatusInternalServerError).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Something went wrong",
				Data:       fiber.Map{},
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(constants.ApiResponse{
		StatusCode: fiber.StatusOK,
		Message:    "User Added Successfully",
		Data:       fiber.Map{},
	})
}

func VerifyController(c *fiber.Ctx) error {
	_, err := helper.Authenticate(c)
	if err != nil {
		log.Err(err)
		return c.Status(fiber.StatusUnauthorized).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
			Data:       fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(constants.ApiResponse{
		StatusCode: fiber.StatusOK,
		Message:    "User Authorized",
		Data:       fiber.Map{},
	})

}

func RefreshController(c *fiber.Ctx) error {
	email, err := helper.AuthenticateRefresh(c)
	if err != nil {
		e := err.Error()
		switch e {
		case "1":
			return c.Status(fiber.StatusUnauthorized).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Use refresh token only",
				Data:       fiber.Map{},
			})
		default:
			return c.Status(fiber.StatusUnauthorized).JSON(constants.ApiResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       fiber.Map{},
			})
		}
	}

	token, err := RefreshToken(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Something went wrong",
			Data:       fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(constants.ApiResponse{
		StatusCode: fiber.StatusOK,
		Message:    "new access token generated",
		Data:       fiber.Map{"AccessToken": token},
	})
}

func RevokeTokenController(c *fiber.Ctx) error {
	_, err := helper.Authenticate(c)
	if err != nil {
		log.Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(constants.ApiResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "token already expired or invalid",
			Data:       fiber.Map{},
		})
	}

	RevokeToken(c.Get("Authorization"))

	return c.Status(fiber.StatusOK).JSON(constants.ApiResponse{
		StatusCode: fiber.StatusOK,
		Message:    "token revoked",
		Data:       fiber.Map{},
	})
}
