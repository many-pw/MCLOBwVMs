package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tjarratt/babble"
	"jjaa.me/models"
	"jjaa.me/persist"
)

var babbler = babble.NewBabbler()

func main() {
	fmt.Println("worker")
	rand.Seed(time.Now().UnixNano())
	Db := persist.Connection()
	babbler.Count = 4
	name := babbler.Babble()
	for {
		video, _ := models.SelectVideoForWorker(Db, name)
		fmt.Println("got", video)

		time.Sleep(time.Second * 60)
	}
}
