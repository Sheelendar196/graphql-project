package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sheelendar196/go-projects/graphql-project/graph"
	logger "golang.org/x/exp/slog"
)

var appName = "graphql-project"

const defaultPort = "8080"

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(true),
	)
	if err != nil {
		logger.Error("error starting new relic err: %v")
	}

	server := startServer(app)

	shutdownServer(server)
	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}

func startServer(app *newrelic.Application) *http.Server {

	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//wrappPattern, WrappHandler := newrelic.WrapHandleFunc(app, "/query", srv.ServeHTTP)

	mux.Handle("/query", srv)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", defaultPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	go func() {
		logger.Info("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error when start server")
		}
	}()
	return server
}

func shutdownServer(server *http.Server) {
	// need to implement
}
