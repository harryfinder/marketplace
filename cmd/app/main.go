package main

import (
	"context"
	"flag"
	"log"
	"marketplace/cmd/app/http"
	"marketplace/internal/config"
	"marketplace/internal/database/pgx"
	"marketplace/internal/entity"
	"marketplace/internal/usecase"
	pkghttp "marketplace/pkg/controller/http"
	pkgpostgres "marketplace/pkg/storage/postgres/pgx"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// @title Swagger Stock API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3939
// host 192.168.194.78
// @Schemes http https
// @BasePath
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {

	ctx, cancelFun := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// 1. init config
	//config.InitConfig()

	var configPath = flag.String("config", "config.json", "path of the config file")
	flag.Parse()
	// go helpers.Worker(10, database.UpdateStatusCheckPhoneExperianTime)
	var cfg = config.InitConfig(*configPath)

	postgres, err := pkgpostgres.NewClient(ctx, cfg.Database, 8)
	if err != nil {
		panic(err)
	}
	database := pgx.New(postgres)

	entity := entity.New(database)

	usecase := usecase.New(entity)

	myhttp := pkghttp.NewServer()
	controller := http.NewController(usecase, myhttp)

	wg.Add(1)
	go func() {
		defer wg.Done()
		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		sig := <-quitCh
		log.Println(sig)
		cancelFun()

		ctx, cancelFun = context.WithTimeout(ctx, time.Second*10)
		defer cancelFun()

		if err := controller.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		log.Println("Server - ✓ finished goroutines")
	}()

	log.Println("Server - running at", cfg.App.Port)

	if err := controller.Serve(ctx, cfg.App.Port, cfg); err != nil {
		log.Println(err)
	}

	wg.Wait()

	log.Println("Server - ✓ finished main goroutine")
}
