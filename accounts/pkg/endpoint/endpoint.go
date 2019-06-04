package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/service"
)

// ListAdminsRequest collects the request parameters for the ListAdmins method.
type ListAdminsRequest struct{}

// ListAdminsResponse collects the response parameters for the ListAdmins method.
type ListAdminsResponse struct {
	M0 []models.Admin `json:"m0"`
	E1 error          `json:"e1"`
}

// MakeListAdminsEndpoint returns an endpoint that invokes ListAdmins on the service.
func MakeListAdminsEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		m0, e1 := s.ListAdmins(ctx)
		return ListAdminsResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r ListAdminsResponse) Failed() error {
	return r.E1
}

// GetAdminRequest collects the request parameters for the GetAdmin method.
type GetAdminRequest struct {
	Id int `json:"id"`
}

// GetAdminResponse collects the response parameters for the GetAdmin method.
type GetAdminResponse struct {
	M0 *models.Admin `json:"m0"`
	E1 error         `json:"e1"`
}

// MakeGetAdminEndpoint returns an endpoint that invokes GetAdmin on the service.
func MakeGetAdminEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminRequest)
		m0, e1 := s.GetAdmin(ctx, req.Id)
		return GetAdminResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminResponse) Failed() error {
	return r.E1
}

// GetAdminByEmailRequest collects the request parameters for the GetAdminByEmail method.
type GetAdminByEmailRequest struct {
	Email string `json:"email"`
}

// GetAdminByEmailResponse collects the response parameters for the GetAdminByEmail method.
type GetAdminByEmailResponse struct {
	M0 *models.Admin `json:"m0"`
	E1 error         `json:"e1"`
}

// MakeGetAdminByEmailEndpoint returns an endpoint that invokes GetAdminByEmail on the service.
func MakeGetAdminByEmailEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminByEmailRequest)
		m0, e1 := s.GetAdminByEmail(ctx, req.Email)
		return GetAdminByEmailResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminByEmailResponse) Failed() error {
	return r.E1
}

// CreateAdminRequest collects the request parameters for the CreateAdmin method.
type CreateAdminRequest struct {
	Admin *models.Admin `json:"admin"`
}

// CreateAdminResponse collects the response parameters for the CreateAdmin method.
type CreateAdminResponse struct {
	M0 *models.Admin `json:"m0"`
	B1 bool          `json:"b1"`
	E2 error         `json:"e2"`
}

// MakeCreateAdminEndpoint returns an endpoint that invokes CreateAdmin on the service.
func MakeCreateAdminEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAdminRequest)
		m0, b1, e2 := s.CreateAdmin(ctx, req.Admin)
		return CreateAdminResponse{
			B1: b1,
			E2: e2,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAdminResponse) Failed() error {
	return r.E2
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ListAdmins implements Service. Primarily useful in a client.
func (e Endpoints) ListAdmins(ctx context.Context) (m0 []models.Admin, e1 error) {
	request := ListAdminsRequest{}
	response, err := e.ListAdminsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListAdminsResponse).M0, response.(ListAdminsResponse).E1
}

// GetAdmin implements Service. Primarily useful in a client.
func (e Endpoints) GetAdmin(ctx context.Context, id int) (m0 *models.Admin, e1 error) {
	request := GetAdminRequest{Id: id}
	response, err := e.GetAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminResponse).M0, response.(GetAdminResponse).E1
}

// GetAdminByEmail implements Service. Primarily useful in a client.
func (e Endpoints) GetAdminByEmail(ctx context.Context, email string) (m0 *models.Admin, e1 error) {
	request := GetAdminByEmailRequest{Email: email}
	response, err := e.GetAdminByEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminByEmailResponse).M0, response.(GetAdminByEmailResponse).E1
}

// CreateAdmin implements Service. Primarily useful in a client.
func (e Endpoints) CreateAdmin(ctx context.Context, admin *models.Admin) (m0 *models.Admin, b1 bool, e2 error) {
	request := CreateAdminRequest{Admin: admin}
	response, err := e.CreateAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAdminResponse).M0, response.(CreateAdminResponse).B1, response.(CreateAdminResponse).E2
}
