package app

import (
	"context"
	"fmt"
	"mock-server/internal/constants"
	"mock-server/internal/controller"
	"mock-server/internal/logs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Engine Engine
	Logger logs.Logger

	HealthController controller.Health
	MockController   controller.Mock
}

func (app *App) Setup() *App {
	app.Logger.InfoWithoutContext("App setup started")
	app.injectDependencies()
	app.ConfigureRoutes()

	return app
}

func (app *App) injectDependencies() {
	app.Logger.InfoWithoutContext("Injecting dependencies")

	app.HealthController = controller.NewHealth()
	app.MockController = controller.NewMock()
}

func (app *App) ConfigureRoutes() {
	app.Logger.InfoWithoutContext("Configuring routes")

	app.Engine = NewEngine()
	v1 := app.Engine.Group("/v1")
	{
		// Swagger
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET(constants.PingBasePath, app.HealthController.GetPing)
		v1.POST(constants.MockBasePath, app.MockController.SetMockResponse)
	}

	app.Engine.NoRoute(app.MockController.NotFound)
}

func (app *App) InitServer() {
	app.Logger.InfoWithoutContext("Starting server")

	srv := app.createServer()
	go func() {
		app.Logger.InfoWithoutContext("Server started")
		err := srv.ListenAndServe()
		if err != nil {
			app.Logger.ErrorWithoutContext("Error when starting server", err)
			return
		}
	}()

	app.waitForShutdownSignal(srv)
}

func (app *App) createServer() *http.Server {
	serverAddress := fmt.Sprintf(":%d", constants.ServerPort)

	server := &http.Server{
		Addr:         serverAddress,
		WriteTimeout: time.Second * 180,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      app.Engine,
	}
	return server
}

func (app *App) waitForShutdownSignal(srv *http.Server) {
	shutdownTimeout := time.Duration(100)
	channel := make(chan os.Signal, 1)
	// We will accept gracefully shutdowns when we want to quit the app via SIGTERM
	signal.Notify(channel, syscall.SIGTERM)
	// Block until we receive our signal
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	// Doesn't block if there's no connection but will wait otherwise
	// until the timeout deadline
	_ = srv.Shutdown(ctx)
}

func NewApp() *App {
	logger := logs.New("app")

	return &App{
		Logger: logger,
	}
}
