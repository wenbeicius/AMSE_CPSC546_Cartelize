package service

import (
	"context"

	"github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/models"
)

// AccountsService describes the service.
type AccountsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	ListAdmins(ctx context.Context) ([]models.Admin, error)
	GetAdmin(ctx context.Context, id int) (*models.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (*models.Admin, error)
	CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, bool, error)
}

type basicAccountsService struct{}

func (b *basicAccountsService) ListAdmins(ctx context.Context) (m0 []models.Admin, e1 error) {
	// TODO implement the business logic of ListAdmins
	return m0, e1
}
func (b *basicAccountsService) GetAdmin(ctx context.Context, id int) (m0 *models.Admin, e1 error) {
	// TODO implement the business logic of GetAdmin
	return m0, e1
}
func (b *basicAccountsService) GetAdminByEmail(ctx context.Context, email string) (m0 *models.Admin, e1 error) {
	// TODO implement the business logic of GetAdminByEmail
	return m0, e1
}
func (b *basicAccountsService) CreateAdmin(ctx context.Context, admin *models.Admin) (m0 *models.Admin, b1 bool, e2 error) {
	// TODO implement the business logic of CreateAdmin
	return m0, b1, e2
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService() AccountsService {
	return &basicAccountsService{}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountsService {
	var svc AccountsService = NewBasicAccountsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
