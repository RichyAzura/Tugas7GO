package main

import (
	db "TugasGO7RichyAzura/echo-rest/db"
	routes "TugasGO7RichyAzura/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
