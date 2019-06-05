package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/models"
)

// AccountsService describes the service.
type AccountsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	ListAdmins(ctx context.Context) (*[]models.Admin, error)
	GetAdmin(ctx context.Context, id int) (*models.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (*models.Admin, error)
	CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, bool, error)
}

type basicAccountsService struct {
	db *sqlx.DB
}

var errNotFound = errors.New("Not found")

func (b *basicAccountsService) ListAdmins(ctx context.Context) (*[]models.Admin, error) {
	// TODO implement the business logic of ListAdmins
	admins, err := models.ListAdmins(b.db)
	if err == sql.ErrNoRows {
		return nil, errNotFound
	}
	return admins, err
}
func (b *basicAccountsService) GetAdmin(ctx context.Context, id int) (*models.Admin, error) {
	// TODO implement the business logic of GetAdmin
	admin, err := models.GetAdmin(b.db, id)

	if err == sql.ErrNoRows {
		return nil, errNotFound
	}

	return admin, nil
}
func (b *basicAccountsService) GetAdminByEmail(ctx context.Context, email string) (m0 *models.Admin, e1 error) {
	// TODO implement the business logic of GetAdminByEmail
	admin, err := models.GetAdminByEmail(b.db, email)

	if err == sql.ErrNoRows {
		return nil, errNotFound
	}

	return admin, err
}
func (b *basicAccountsService) CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, bool, error) {
	// TODO implement the business logic of CreateAdmin
	admin, created, err := admin.CreateAdmin(b.db)

	if err == sql.ErrNoRows {
		return nil, false, errNotFound
	}

	return admin, created, err
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService(db *sqlx.DB) AccountsService {
	return &basicAccountsService{
		db: db,
	}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware, db *sqlx.DB) AccountsService {
	svc := NewBasicAccountsService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
