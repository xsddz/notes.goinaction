package main

import (
	"notes.goinaction/chapter06/channel"
	"notes.goinaction/chapter06/race"
	"notes.goinaction/chapter06/tgoroutine"
)

func main() {
	tgoroutine.Run()

	race.RunRace()
	race.RunAtom()
	race.RunMutex()

	channel.RunUnbuffered()
	channel.RunBuffered()
}
