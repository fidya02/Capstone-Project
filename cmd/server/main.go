package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/fidya02/Capstone-Project/internal/builder"
	"github.com/fidya02/Capstone-Project/internal/config"
	"github.com/fidya02/Capstone-Project/internal/http/binder"
	"github.com/fidya02/Capstone-Project/internal/http/server"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	splash()

	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	midtransClient := initMidtrans(cfg)

	publicRoutes := builder.BuildPublicRoutes(cfg, db, midtransClient)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db, midtransClient)

	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	runServer(srv, cfg.Port)

	waitForShutdown(srv)
}

func initMidtrans(cfg *config.Config) snap.Client {
	snapClient := snap.Client{}

	if cfg.Env == "development" {
		snapClient.New(cfg.MidtransConfig.ServerKey, midtrans.Sandbox)
	} else {
		snapClient.New(cfg.MidtransConfig.ServerKey, midtrans.Production)
	}

	return snapClient
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func splash() {
	colorReset := "\033[0m"

	splashText := `
	________    _______    _______   ____  ____  _______   ___        __     ______   
	|"      "\  /"     "|  |   __ "\ ("  _||_ " ||   _  "\ |"  |      |" \   /" _  "\  
	(.  ___  :)(: ______)  (. |__) :)|   (  ) : |(. |_)  :)||  |      ||  | (: ( \___) 
	|: \   ) || \/    |    |:  ____/ (:  |  | . )|:     \/ |:  |      |:  |  \/ \      
	(| (___\ || // ___)_   (|  /      \\ \__/ // (|  _  \\  \  |___   |.  |  //  \ _   
	|:       :)(:      "| /|__/ \     /\\ __ //\ |: |_)  :)( \_|:  \  /\  |\(:   _) \  
	(________/  \_______)(_______)   (__________)(_______/  \_______)(__\_|_)\_______) 
																						  
																   
	`
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
