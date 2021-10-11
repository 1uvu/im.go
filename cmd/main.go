package main

import (
	"flag"
	"im/internal/api"
	"im/internal/connect"
	"im/internal/logic"
	"im/internal/task"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("start")
	var layer string

	flag.StringVar(&layer, "layer", "", "assign to run layer of IM")
	flag.Parse()

	log.Printf("start to run %s layer", layer)

	switch layer {
	case "api":
		api.NewAPI().Run()
	case "connect_ws":
		connect.NewConnect().RunWS()
	case "connect_tcp":
		log.Println("tcp connect layer has not impl now, please use ws to instead it")
	case "logic":
		logic.NewLogic().Run()
	case "task":
		task.NewTask().Run()
	default:
		log.Println("unknown layer")
	}

	log.Printf("success to run %s layer", layer)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	log.Panicln("quit the im")

}
