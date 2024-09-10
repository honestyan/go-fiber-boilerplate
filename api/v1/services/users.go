package services

import (
	"context"
	"database/sql"
	"time"

	M "github.com/honestyan/go-fiber-boilerplate/models"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(dbTrx *sql.Tx, ctx context.Context) ([]*M.User, *T.ServiceError) {
	users, err := M.Users().All(ctx, dbTrx)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error retrieving users",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	return users, nil
}

func GetUser(dbTrx *sql.Tx, ctx context.Context, id int) (*M.User, *T.ServiceError) {
	user, err := M.FindUser(ctx, dbTrx, id)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error retrieving user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	return user, nil
}

func CreateUser(dbTrx *sql.Tx, ctx context.Context, body *T.UserBody) (*M.User, *T.ServiceError) {
	existingUserByUsername, err := M.Users(qm.Where("username = ?", body.Username)).One(ctx, dbTrx)
	if err != nil && err != sql.ErrNoRows {
		return nil, &T.ServiceError{
			Message: "Error checking username",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}
	if existingUserByUsername != nil {
		return nil, &T.ServiceError{
			Message: "Username already exists",
			Code: fiber.StatusConflict,
		}
	}

	existingUserByEmail, err := M.Users(qm.Where("email = ?", body.Email)).One(ctx, dbTrx)
	if err != nil && err != sql.ErrNoRows {
		return nil, &T.ServiceError{
			Message: "Error checking email",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}
	if existingUserByEmail != nil {
		return nil, &T.ServiceError{
			Message: "Email already exists",
			Code: fiber.StatusConflict,
		}
	}

	user := &M.User{
		Username: body.Username,
		Email:    body.Email,
		Name:     body.Name,
		Gender:   M.GenderEnum(body.Gender),
		Password: body.Password,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error hashing password",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}
	user.Password = string(hashedPassword)

	err = user.Insert(ctx, dbTrx, boil.Infer())
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error creating user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	return user, nil
}

func UpdateUser(dbTrx *sql.Tx, ctx context.Context, id int, body *T.UserBody) (*M.User, *T.ServiceError) {
	existingUser, err := M.FindUser(ctx, dbTrx, id)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error retrieving user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	existingUser.Email 	= body.Email
	existingUser.Name 	= body.Name
	existingUser.Gender = M.GenderEnum(body.Gender)
	if body.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, &T.ServiceError{
				Message: "Error hashing password",
				Code: fiber.StatusInternalServerError,
				Error: err,
			}
		}
		existingUser.Password = string(hashedPassword)
	}
	existingUser.Modified = time.Now()

	rowsAffected, err := existingUser.Update(ctx, dbTrx, boil.Infer())
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error updating user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	if rowsAffected == 0 {
		return nil, &T.ServiceError{
			Message: "No rows affected",
			Code: fiber.StatusNotFound,
		}
	}

	return existingUser, nil
}

func DeleteUser(dbTrx *sql.Tx, ctx context.Context, id int) *T.ServiceError {
	user, err := M.FindUser(ctx, dbTrx, id)

	if err != nil {
		return &T.ServiceError{
			Message: "Error retrieving user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	_, err = user.Delete(ctx, dbTrx)

	if err != nil {
		return &T.ServiceError{
			Message: "Error deleting user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	return nil
}

func Login(dbTrx *sql.Tx, ctx context.Context, body *T.LoginBody) (*M.User, *T.ServiceError) {
	user, err := M.Users(qm.Where("username = ? OR email = ?", body.Username, body.Username)).One(ctx, dbTrx)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Error retrieving user",
			Code: fiber.StatusInternalServerError,
			Error: err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Invalid username or password",
			Code: fiber.StatusUnauthorized,
		}
	}

	return user, nil
}