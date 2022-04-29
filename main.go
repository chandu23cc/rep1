package main

import "github.com/labstack/echo/v4"
import "net/http"
import "math/rand"
import "os"


type quote struct {
	Title string
	Author string
}


//var quotes []quote=make([]quote,0)

var quotes []quote = []quote{
	{"Learn to lead","Sai Vidya Campus"},
	{"Be the change you want to see in the world","Mahatma Gandhi"},
	{""}
}

func main() {

	//rand.Seed(time.Now().Unix())
	api:=echo.New()

	api.GET("/quotes",getQuotes)
	api.GET("/quotes/random",getRandomQuote)

	port :=os.Getenv("PORT")
	if port == "" {
		port="32445"
	}//to set port externally while running,  $env:PORT= some_port_no

	/*//endpoint - handler
	api.GET("/",hello) //HANDLER
	api.POST("/",hello)
	api.PUT("/",hello)*/

	api.Start(":"+port)

}

/*func hello(c echo.Context) error {
	return c.JSON(200,"Hello World")
}

func anotherHandler(c echo.Cintext) error {
	return c.NoContent(http.StatusOK)
}*/

func getQuotes (c echo.Context) error {
	return c.JSON(http.StatusOK ,quotes)
}

func getRandomQuote(c echo.Context) error {
	
	index:= rand.Intn(len(quotes))
	return c.JSON(http.StatusOK ,quotes[index])
	return nil
}