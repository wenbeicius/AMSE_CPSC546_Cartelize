package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/terasurfer/AMSE_CPSC546_Cartelize/accounts/pkg/endpoint"
	pb "github.com/terasurfer/AMSE_CPSC546_Cartelize/accounts/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeListAdminsHandler creates the handler logic
func makeListAdminsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListAdminsEndpoint, decodeListAdminsRequest, encodeListAdminsResponse, options...)
}

// decodeListAdminsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListAdmins request.
// TODO implement the decoder
func decodeListAdminsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Decoder is not impelemented")
}

// encodeListAdminsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListAdminsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Encoder is not impelemented")
}
func (g *grpcServer) ListAdmins(ctx context1.Context, req *pb.ListAdminsRequest) (*pb.ListAdminsReply, error) {
	_, rep, err := g.listAdmins.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAdminsReply), nil
}

// makeGetAdminHandler creates the handler logic
func makeGetAdminHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetAdminEndpoint, decodeGetAdminRequest, encodeGetAdminResponse, options...)
}

// decodeGetAdminResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetAdmin request.
// TODO implement the decoder
func decodeGetAdminRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Decoder is not impelemented")
}

// encodeGetAdminResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Encoder is not impelemented")
}
func (g *grpcServer) GetAdmin(ctx context1.Context, req *pb.GetAdminRequest) (*pb.GetAdminReply, error) {
	_, rep, err := g.getAdmin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAdminReply), nil
}

// makeGetAdminByEmailHandler creates the handler logic
func makeGetAdminByEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetAdminByEmailEndpoint, decodeGetAdminByEmailRequest, encodeGetAdminByEmailResponse, options...)
}

// decodeGetAdminByEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetAdminByEmail request.
// TODO implement the decoder
func decodeGetAdminByEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Decoder is not impelemented")
}

// encodeGetAdminByEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetAdminByEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Encoder is not impelemented")
}
func (g *grpcServer) GetAdminByEmail(ctx context1.Context, req *pb.GetAdminByEmailRequest) (*pb.GetAdminByEmailReply, error) {
	_, rep, err := g.getAdminByEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAdminByEmailReply), nil
}

// makeCreateAdminHandler creates the handler logic
func makeCreateAdminHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateAdminEndpoint, decodeCreateAdminRequest, encodeCreateAdminResponse, options...)
}

// decodeCreateAdminResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateAdmin request.
// TODO implement the decoder
func decodeCreateAdminRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Decoder is not impelemented")
}

// encodeCreateAdminResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Accounts' Encoder is not impelemented")
}
func (g *grpcServer) CreateAdmin(ctx context1.Context, req *pb.CreateAdminRequest) (*pb.CreateAdminReply, error) {
	_, rep, err := g.createAdmin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateAdminReply), nil
}
