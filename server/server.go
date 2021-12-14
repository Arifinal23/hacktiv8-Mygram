package server

import (
	"fmt"
	"os"

	"github.com/arfan21/golang-mygram/config/configdb"
	"github.com/gin-gonic/gin"
)

func Start() error {
	db, err := configdb.New()
	if err != nil {
		return err
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%s", port))
	return nil
}
