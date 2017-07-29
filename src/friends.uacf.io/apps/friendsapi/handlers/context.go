package handlers

import (
	"strconv"
	"time"

	"friends.uacf.io/services/auth"

	"go.uacf.io/apierrors/grpcmapper"
	log "go.uacf.io/logging"
	"go.uacf.io/metrics"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// This is an example of extracting metadata from context and annotating your own context.
// This is definitely not a secure auth mechanism for an -external- GRPC service but might
// be enough for internal services to provide contextual authn
func AuthContext(ctx context.Context) context.Context {
	// metadata.FromContext gets the grpc metadata (a dictionary) sent from the client.
	if md, ok := metadata.FromContext(ctx); ok && len(md["User-Id"]) > 0 {
		// Would normally come as a header like Grpc-Metadata-User-Id
		if uid, err := strconv.ParseInt(md["User-Id"][0], 10, 64); err == nil {
			ctx = auth.NewContext(ctx, uid)
		}
	}
	return ctx
}

// a "middleware" to do some common tasks to most GRPC methods
//  - Extract auth from context
//  - timing
//  - request count metrics
//  - debug logging
//  - error code mapping back to GRPC errors
func Middleware(ctx context.Context, name string, handler func(ctx context.Context) error) error {
	// extract the auth from context
	ctx = AuthContext(ctx)

	// logging and instrumentation of simple per-request metrics
	log.With(log.Fields{"context": ctx}).Debug(name)
	metrics.Inc(name+".requested", 1)
	defer metrics.TimeElapsed(name+".response_time", time.Now())

	// call into the provided handler
	err := handler(ctx)

	// map API errors from service layer back to GRPC codes
	return grpcError(err)
}

func grpcError(err error) error {
	return grpcmapper.Convert(err)
}
