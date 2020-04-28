package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
      "os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Time(flag string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := fmt.Sprintf("%s\n%s", time.Now().Format("2006-01-02 15:04:05"), flag)
            out,err := exec.Command("./flag").Output()
            if err == nil {
            result = result + string(out)
            }
		return ctx.String(http.StatusOK, result)
	}
}


func main() {
	flag, err := ioutil.ReadFile("flag.txt")
	if err != nil {
		log.Fatalf("ReadFile() error=%v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", Time(string(flag)))
	e.Start(":3000")
}
