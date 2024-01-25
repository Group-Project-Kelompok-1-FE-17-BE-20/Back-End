package main

import (
	config "Laptop/app/configs"
	"Laptop/app/database"
	router "Laptop/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// logging := helpers.NewLogger()
	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)

	//call migration
	database.InitialMigration(dbMysql)

	//create a new echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	//remove pre trailingslash
	e.Pre(middleware.RemoveTrailingSlash())

	//e.Use middleware logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMysql, e)

	//start server and port
	e.Logger.Fatal(e.Start(":8083"))
}

// func customTLSConfig() (*tls.Config, error) {
// 	caCert, err := ioutil.ReadFile("server-ca.pem")
// 	if err != nil {
// 		return nil, err
// 	}

// 	certPool := x509.NewCertPool()
// 	certPool.AppendCertsFromPEM(caCert)

// 	return &tls.Config{
// 		RootCAs: certPool,
// 	}, nil
// }
