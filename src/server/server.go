package server

import (
	"net/http"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"

	stats "github.com/lyft/gostats"
	"google.golang.org/grpc"
)

type Server interface {
	/**
	 * Starts the HTTP and gRPC servers. This should be done after
	 * all endpoints have been registered through 'AddHttpEndpoint'
	 * and 'GrpcServer'.
	 */
	Start()

	/**
	 * Returns the root of the stats tree for the server
	 */
	Scope() stats.Scope

	/**
	 * Add an HTTP endpoint to the local debug port.
	 */
	AddDebugHttpEndpoint(path string, help string, handler http.HandlerFunc)
	AddJsonHandler(pb.RateLimitServiceServer)

	/**
	 * Returns the embedded gRPC server to be used for registering gRPC endpoints.
	 */
	GrpcServer() *grpc.Server

	// /**
	//  * Returns the runtime configuration for the server.
	//  */
	// Runtime() loader.IFace

	/**
	 *  Stops serving the grpc port (for integration testing).
	 */
	Stop()

	HealthCheckFail()
	HealthCheckOK()
}
