package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/newrelic/go-agent/v3/newrelic"

	graph "github.com/sheelendar196/go-projects/graphql-project/graph"
	generated "github.com/sheelendar196/go-projects/graphql-project/graph/generated"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/service"
	"github.com/sheelendar196/go-projects/graphql-project/internal/handlers/api"
	"github.com/sheelendar196/go-projects/graphql-project/internal/infrastructure/repository"
	logger "golang.org/x/exp/slog"
)

var appName = "graphql-project"

const defaultPort = "8080"

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigLicense("1232323432323454323456765432345676543234"),
		newrelic.ConfigEnabled(true),
	)
	if err != nil {
		logger.Error("error starting new relic err: %v", err)
	}

	server := startServer(app)
	shutdownServer(server)

}

func startServer(app *newrelic.Application) *http.Server {

	empRepo := repository.NewEmployeeRepo(nil, app)
	roleInteractor := service.NewEmployeeInteractor(empRepo)
	if roleInteractor == nil {
		panic("error while creating repo")
	}
	empProcessor := api.NewEmployeeProceeor(roleInteractor)

	graphConfig := generated.NewExecutableSchema(
		generated.Config{Resolvers: &graph.Resolver{EmployeeProcessor: empProcessor}})
	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graphConfig)

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", defaultPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	logger.Info("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("error when start server")
	}

	return server
}

func shutdownServer(server *http.Server) {
	// need to implement
}
