package main

import (
	"laboratory/internal/http/router"
	"laboratory/sql"
)

func main() {
	sql.InitSQL()
	r := router.InitRouter()
	r.Run(":9090")
}
