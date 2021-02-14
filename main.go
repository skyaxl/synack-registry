package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kataras/golog"
	"github.com/skyaxl/synack-registry/api"
)

func main() {
	errs := make(chan error, 2)
	go func() {
		errs <- api.New()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	golog.Infof("[Registry Api] Has stoped. \n%s\n", <-errs)
}
