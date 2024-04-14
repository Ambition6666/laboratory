package main

import (
	"laboratory/internal/http/router"
	"laboratory/log"
	"laboratory/sql"
)

func main() {
	sql.InitSQL()
	log.InitLogger()
	r := router.InitRouter()
	r.Run(":9090")
}
