package routes

import (
	"errors"
	"fmt"

	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
	"gorm.io/gorm"
)

/*Handlers for Auth Views*/

func HandleViewHome(c *fiber.Ctx) error {
	if c.Locals("userId") != nil && c.Locals("userId") != "" {
		return c.Redirect("todo/list")
	}
	return c.Render("home", fiber.Map{})
}

func HandleViewLogin(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	if c.Method() == "POST" {
		var (
			user models.User
			err  error
		)
		fm := fiber.Map{
			"type": "alert-error",
		}

		if user, err = models.CheckUserName(c.FormValue("username"), db); err != nil {
			fm["message"] = "There is no user with that username"

			return flash.WithError(c, fm).Redirect("/login")
		}
		session, err := store.Get(c)
		if err != nil {
			fm["message"] = fmt.Sprintf("something went wrong: %s", err)

			return flash.WithError(c, fm).Redirect("/login")
		}

		session.Set(USER_ID, user.ID)

		err = session.Save()
		if err != nil {
			fm["message"] = fmt.Sprintf("something went wrong: %s", err)

			return flash.WithError(c, fm).Redirect("/login")
		}

		c.Locals("userId", user.ID)

		fm = fiber.Map{
			"type":    "alert-succes",
			"message": "You have successfully logged in!!",
			"User":    user,
		}

		return flash.WithSuccess(c, fm).Redirect("/todo/list")
	}

	return c.Render("login", fiber.Map{
		"Page":    "login",
		"Message": flash.Get(c),
	})
}

func HandleViewRegister(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	if c.Method() == "POST" {
		// Verificar si el username ya existe
		if _, err := models.CheckUserName(c.FormValue("username"), db); err == nil {
			// El username ya está en uso
			fm := fiber.Map{
				"type":    "alert-error",
				"message": "The username is already in use",
			}
			return flash.WithError(c, fm).Redirect("/register")
		}

		// Crear nuevo usuario solo con el username
		user := &models.User{
			Username: c.FormValue("username"),
		}

		err := models.CreateUser(user, db)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: users.username" {
				err = errors.New("the username is already in use")
			}
			fm := fiber.Map{
				"type":    "alert-error",
				"message": fmt.Sprintf("something went wrong: %s", err),
			}
			return flash.WithError(c, fm).Redirect("/register")
		}

		fm := fiber.Map{
			"type":    "alert-success",
			"message": "You have successfully registered!!",
		}

		return flash.WithSuccess(c, fm).Redirect("/login")
	}

	return c.Render("register", fiber.Map{
		"Page":    "Register",
		"Message": flash.Get(c),
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	fm := fiber.Map{
		"type": "alert-error",
	}

	session, err := store.Get(c)
	if err != nil {
		fm["message"] = "You are not authorized"
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	userId := session.Get(USER_ID)
	if userId == nil {
		fm["message"] = "You are not authorized"
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user, err := models.GetUserById(userId.(uint64), db)
	if err != nil {
		fm["message"] = "You are not authorized"
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	c.Locals("userId", userId)
	c.Locals("username", user.Username)

	return c.Next()
}

func MaybeAuthMiddleware(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	session, err := store.Get(c)
	if err != nil {
		return c.Next()
	}

	userId := session.Get(USER_ID)
	if userId == nil {
		return c.Next()
	}

	user, err := models.GetUserById(userId.(uint64), db)
	if err != nil {
		return c.Next()
	}

	c.Locals("userId", userId)
	c.Locals("username", user.Username)

	return c.Next()
}

func HandleLogout(c *fiber.Ctx) error {
	fm := fiber.Map{
		"type": "alert-error",
	}

	session, err := store.Get(c)
	if err != nil {
		fm["message"] = "Logged out (no session)"
		return flash.WithError(c, fm).Redirect("/login")
	}

	err = session.Destroy()
	if err != nil {
		fm["message"] = fmt.Sprintf("Something went wrong: %s", err)
		return flash.WithError(c, fm).Redirect("/login")
	}

	fm = fiber.Map{
		"type":    "alert-success",
		"message": "You have successfully logged out!!",
	}

	return flash.WithSuccess(c, fm).Redirect("/login")
}
