package main

import (
	"frontend/gen/restapi"
	"frontend/gen/restapi/operations"
	"frontend/handler"
	"github.com/go-openapi/loads"
	"log"

	"go.uber.org/zap"
)

//go:generate swagger generate server --target ./gen --name Gremium --spec ./swagger.yaml --exclude-main

const (
	serviceName = "quiz"
	port        = 8000
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("can't create logger")
	}

	slog := logger.Sugar()

	slog.Infof("starting %s", serviceName)

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		slog.Fatal(err)
	}

	gremiumAPI := createAPI(swaggerSpec, slog)

	if err := gremiumAPI.Validate(); err != nil {
		slog.Fatal(err)
	}

	server := restapi.NewServer(gremiumAPI)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalf("stutting down the server: %v", err)
		}
	}()

	// Set the port this service will listen on
	server.Port = port
	server.EnabledListeners = []string{"http"}

	params := &handler.Params{
		Slog: slog,
	}
	registerAPIs(params, gremiumAPI)

	server.SetHandler(gremiumAPI.Serve(nil))

	// Serve the API
	if err := server.Serve(); err != nil {
		slog.Fatal(err)
	}
}

func registerAPIs(p *handler.Params, gremiumAPI *operations.GremiumAPI) {
	handler.Register(p, gremiumAPI)
}

func createAPI(swaggerSpec *loads.Document,
	slog *zap.SugaredLogger,
) *operations.GremiumAPI {
	// Create new service API
	gremiumAPI := operations.NewGremiumAPI(swaggerSpec)
	gremiumAPI.Logger = func(s string, i ...interface{}) {
		slog.Infof(s, i)
	}
	gremiumAPI.ServerShutdown = shutdownHandler(slog)
	return gremiumAPI
}

func shutdownHandler(slog *zap.SugaredLogger) func() {
	return func() {
		slog.Info("shutdown handler")
	}
}