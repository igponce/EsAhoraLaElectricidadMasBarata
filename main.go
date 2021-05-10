package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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

func tarifaActual () int {

    var tarifa int
    var tarifa_pen[] = []int{ 1, 1, 1, 1, 1, 1, 1, 1,  // 0:00 - 7:59   - Valle
	   		      2, 2,                    // 8:00 - 9:59   - Llana
			      3, 3, 3, 3,              // 10:00 - 13:59 - Punta
			      2, 2, 2, 2,              // 14:00 - 17:59 - Llana
			      3, 3, 3, 3,              // 18:00 - 21:59 - Punta
			      2, 2,                    // 22:00 - 23:59 - Llana 
			    }
    var tarifas []string = []string{"undef","valle","llana","punta"}
    ahora := time.Now()
    dd, mm, yy := ahora.Date()
    hh, mm, ss := ahora.Clock()
    weekday := ahora.Weekday()
	
    tarifa := tarifa_pen[hh]
    fmt.Printf("%d%0d%0dT%0d:%0d:%0d - Tarifa %s", yy,mm, dd, hh, mm, ss)

    return tarifa
}
