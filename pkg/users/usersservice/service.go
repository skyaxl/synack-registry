package userservice

import (
	"context"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/skyaxl/synack-registry/db/pkg/models"
	"github.com/skyaxl/synack-registry/pkg/apierrors"
	"github.com/skyaxl/synack-registry/pkg/users/userscontracts"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Service struct {
	db boil.ContextExecutor
}

//New Service
func New(db boil.ContextExecutor) *Service {
	return &Service{db}
}

//Save user
func (svc Service) Save(ctx context.Context, user userscontracts.User) (err error) {
	userD := models.User{
		UserName: null.StringFrom(user.Username),
		Password: null.StringFrom(user.Password),
		Name:     null.StringFrom(user.Name),
	}
	golog.Infof("Inserting user: %s", user.Name)
	if err = userD.Insert(ctx, svc.db, boil.Infer()); err != nil {
		golog.Errorf("Error on inserting user: %s, err: %v", user.Name, err)
		return err
	}
	return
}

//Get Get
func (svc Service) Get(ctx context.Context, username string) (u userscontracts.User, err error) {
	var userD *models.User
	if userD, err = models.FindUser(ctx, svc.db, null.StringFrom(username)); err != nil {
		golog.Errorf("Error on finding user: %s, err: %v", username, err)
		err = errors.WithMessagef(apierrors.ErrNotFound, "Error on get %s, err: %v", username, err)
		return
	}

	return userscontracts.User{
		Username: userD.UserName.String,
		Name:     userD.Name.String,
	}, nil
}

//Update Update
func (svc Service) Update(ctx context.Context, user userscontracts.User) (err error) {
	var userD *models.User
	if userD, err = models.FindUser(ctx, svc.db, null.StringFrom(user.Username)); err != nil {
		golog.Errorf("Error on finding user: %s, err: %v", user.Name, err)
		err = errors.WithMessagef(apierrors.ErrNotFound, "Error on get %s, err: %v", user.Name, err)
		return
	}

	userD.Name = null.StringFrom(user.Name)
	userD.Password = null.StringFrom(user.Password)

	if _, err = userD.Update(ctx, svc.db, boil.Infer()); err != nil {
		golog.Errorf("Error on updating user: %s, err: %v", user.Username, err)
	}
	return
}

//Delete user
func (svc Service) Delete(ctx context.Context, username string) (err error) {
	var userD *models.User
	if userD, err = models.FindUser(ctx, svc.db, null.StringFrom(username)); err != nil {
		golog.Errorf("Error on finding user: %s, err: %v", username, err)
		err = errors.WithMessagef(apierrors.ErrNotFound, "Error on get %s, err: %v", username, err)
		return
	}
	if _, err = userD.Delete(ctx, svc.db); err != nil {
		golog.Errorf("Error on deleting user: %s, err: %v", username, err)
	}
	return
}

//Autenticate autenticate
func (svc Service) Autenticate(ctx context.Context, user userscontracts.AuthenticationRequest) (u userscontracts.User, err error) {
	var userD *models.User
	if userD, err = models.FindUser(ctx, svc.db, null.StringFrom(user.Username)); err != nil {
		golog.Errorf("Error on finding user: %s, err: %v", user.Username, err)
		err = errors.WithMessagef(apierrors.ErrNotFound, "Error on get %s, err: %v", user.Username, err)
		return
	}

	if user.Password == userD.Password.String {
		return userscontracts.User{
			Username: userD.UserName.String,
			Name:     userD.Name.String,
		}, nil
	}
	return u, errors.WithMessagef(apierrors.ErrNotAuthorized, "Invalid user %s or password", user.Username)
}
