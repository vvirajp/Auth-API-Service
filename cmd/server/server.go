package server

import (
	"API/pkg/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Serve() {

	appCfg := config.AppCfg()

	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	LoadRoutes(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Info().Msg("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	log.Info().Msgf("Server is running on port: %v", appCfg.Port)
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Err(err).Msg("Oops... server is not running! error")
	}

}
