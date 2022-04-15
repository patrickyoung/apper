package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	initWeb()
}

const PORT = 8080
const ADDRESS = "127.0.0.1"

var HOME_URL = fmt.Sprintf("http://localhost:%d/health/ping", PORT)

func initWeb() {
	r := gin.Default()
	r.SetTrustedProxies([]string{ADDRESS})

	templateFunctions(r)
	routes(r)

	Connect()

	openbrowser(HOME_URL)

	r.Run(fmt.Sprintf(":%d", PORT))
}
