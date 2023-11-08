package main

func goated(name string) error {

	const theme = "blog-1"

	FetchFrontend(theme, name)
	DeployToNetlify(theme, name)
	deployFly(name)

	return nil
}

func main() {

	goated("freaking-finally")

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
