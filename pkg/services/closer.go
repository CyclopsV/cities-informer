package services

import (
	"log"
	"os"
	"os/signal"
)

type Exit interface {
	Exit()
}

func Closer(fExit ...Exit) {
	chanSignal := make(chan os.Signal)
	signal.Notify(chanSignal)
	<-chanSignal
	for _, e := range fExit {
		e.Exit()
	}
	close(chanSignal)
	log.Println("Программа завершена")
}
