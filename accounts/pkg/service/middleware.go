package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	"github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/models"
)

// Middleware describes a service middleware.
type Middleware func(AccountsService) AccountsService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountsService) AccountsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ListAdmins(ctx context.Context) (m0 *[]models.Admin, e1 error) {
	defer func() {
		l.logger.Log("method", "ListAdmins", "m0", m0, "e1", e1)
	}()
	return l.next.ListAdmins(ctx)
}
func (l loggingMiddleware) GetAdmin(ctx context.Context, id int) (m0 *models.Admin, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAdmin", "id", id, "m0", m0, "e1", e1)
	}()
	return l.next.GetAdmin(ctx, id)
}
func (l loggingMiddleware) GetAdminByEmail(ctx context.Context, email string) (m0 *models.Admin, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAdminByEmail", "email", email, "m0", m0, "e1", e1)
	}()
	return l.next.GetAdminByEmail(ctx, email)
}
func (l loggingMiddleware) CreateAdmin(ctx context.Context, admin *models.Admin) (m0 *models.Admin, b1 bool, e2 error) {
	defer func() {
		l.logger.Log("method", "CreateAdmin", "admin", admin, "m0", m0, "b1", b1, "e2", e2)
	}()
	return l.next.CreateAdmin(ctx, admin)
}
