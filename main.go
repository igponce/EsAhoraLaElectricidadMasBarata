package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/getlantern/systray"
)

func main() {
  systray.Run(onReady, onExit)
}

func onReady() {
  systray.SetIcon(Icondata)
  tooltipmsg := fmt.Sprintf("Random value: %d", rand.Int() )
  systray.SetTitle("Cuanto cuesta la electricidad?")
  systray.SetTooltip(tooltipmsg)
  mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

  // Sets the icon of a menu item. Only available on Mac and Windows.
  mQuit.SetIcon(Icondata)
}

func onExit() {
   // clean up here
   log.Print("Saliendo")
} 

