package main

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"friends.uacf.io/apps/friendsapi/handlers"
	"friends.uacf.io/apps/friendsapi/rpc"
	"friends.uacf.io/apps/friendsapi/serviceconf"
	"friends.uacf.io/data"
	"friends.uacf.io/services"

	"github.com/gengo/grpc-gateway/runtime"
	"go.uacf.io/capture"
	"go.uacf.io/env"
	log "go.uacf.io/logging"
	"go.uacf.io/metrics"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/crast/app.v0"
)

const appName = "friends.friendsapi"

var build string

var config struct {
	Environment   string
	LogLevel      env.LogLevel
	SentryDsn     env.Secret
	HttpServePort int
	GrpcServePort int
	Metrics       serviceconf.MetricsConfig
	MysqlUser       string
	MysqlPassword   string
	MysqlConnection string
}

func configure() error {
	// Load the environment
	err := env.Load(&config)
	if err != nil {
		return err
	}

	// Set the log level
	log.SetLevel(config.LogLevel.String())

	// Log the loaded environment
	for key, value := range env.Inspect(&config) {
		log.With(log.Fields{"var": key, "type": value.Type, "value": value.Value}).Info("environment")
	}

	// Add the Sentry log hook
	if config.SentryDsn.Value != "" {
		err = capture.UseSentry(config.SentryDsn.Value)
		if err != nil {
			return err
		}
		err = capture.LogHook(10*time.Millisecond, log.LevelPanic, log.LevelError)
		if err != nil {
			return err
		}
	}

	// Initialize metrics
	err = serviceconf.InitMetrics(config.Metrics, config.Environment, appName)
	if err != nil {
		return err
	}
	return err
}

func main() {
	log.With(log.Fields{"app": appName, "build": build}).Info("starting application 123")

	if err := configure(); err != nil {
		log.With(log.F("err", err)).Error("Configuration failed")
		os.Exit(1)
	}

	log.Debug("Thhis is my new debug log!!!!!!**!*!*!*!*!**!**")
	// Open network port for each server.
	ports := make(map[int]net.Listener)
	for _, port := range []int{config.HttpServePort, config.GrpcServePort, 82} {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.With(log.F("err", err)).Errorf("failed to listen on port %d: %v", port, err.Error())
			os.Exit(1)
		}
		ports[port] = lis
	}

	grpcServer(ports[config.GrpcServePort])
	grpcEndpoint := fmt.Sprintf("127.0.0.1:%d", config.GrpcServePort)
	grpcJsonGateway(grpcEndpoint, ports[config.HttpServePort])
	debugJsonServer(ports[82])

	// app.Main waits for all running goroutines to end and then runs closers.
	app.Main()
}

func grpcServer(listener net.Listener) {
	logger := log.With(log.Fields{"port": listener.Addr().String()})
	grpc.EnableTracing = false
	server := grpc.NewServer()
	dataLayer, dberr := data.NewFriendData(
		"root",
		"root",
		"tcp(mysql:3306)/mapmyfitness",
	)
	//config.MysqlUser, config.MysqlPassword, config.MysqlConnection)
	if dberr != nil {
		log.Debugf("******* datalayererror: %s", dberr)
	}

	log.Debugf("******* datalayer*****: %s", dataLayer)
	//if err := dataLayer.db.Ping(); err != nil {
	//	log.Debugf("DATABASE ping error: %s", err)
	//}
	service := services.NewFriendService(dataLayer)
	rpc.RegisterFriendsApiServiceServer(server, handlers.NewFriendsApiServer(service))
	app.AddCloser(server.Stop)
	app.Go(func() error {
		logger.Infof("starting grpc server")
		return server.Serve(listener)
	})
}

func grpcJsonGateway(grpcEndpoint string, listener net.Listener) {
	logger := log.With(log.Fields{"port": listener.Addr().String()})
	ctx, cancel := context.WithCancel(context.Background())
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := rpc.RegisterFriendsApiServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		logger.With(log.F("error", err)).Error("failed to start JSON-GRPC gateway")
		app.Stop()
		return
	}

	// Closer stops our listener, which stops any new requests coming in.
	// Then it cancels any ongoing requests going out to the GRPC, to speed up the shutdown.
	app.AddCloser(func() error {
		logger.Info("stopping JSON-GRPC gateway")
		err := listener.Close()
		app.Go(func() {
			time.Sleep(100 * time.Millisecond)
			cancel()
		})
		return err
	})
	app.Go(func() error {
		logger.Infof("starting JSON-GRPC gateway")
		return http.Serve(listener, mux)
	})
}

func debugJsonServer(listener net.Listener) {
	http.Handle("/debug/env", env.HttpHandler{Config: &config})
	http.HandleFunc("/debug/metrics", metrics.HttpHandler)
	server := &http.Server{Handler: http.DefaultServeMux}
	log.With(log.Fields{"port": listener.Addr()}).Info("starting debug server")
	app.Serve(listener, server)
}
