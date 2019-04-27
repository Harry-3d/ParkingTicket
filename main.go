package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	generateTicket(14)
}

func generateTicket(amount int64) {
	var hours24 int64 = 86400
	var multiplier int64 = 0
	for multiplier = 0; multiplier < amount; multiplier++ {

		// get todays date to start creating tickets from
		today := time.Now()
		now := time.Unix(today.Unix()+hours24*multiplier, 0)

		// Letters for day
		uc := strings.ToUpper(fmt.Sprintf("%s", now.Weekday()))
		day := uc[:2]

		// the day number
		dayNum := fmt.Sprintf("%d", now.Day())
		if len(dayNum) == 1 {
			dayNum = "0" + dayNum
		}

		// month number
		month := fmt.Sprintf("%d", now.Month())
		if len(month) == 1 {
			month = "0" + month
		}

		dateStr := day + " " + dayNum + "/" + month

		// Make the ticket
		importImg(dateStr)
	}
}

func importImg(date string) {
	im, err := gg.LoadImage("source.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(1645, 708)
	dc.SetRGBA(1, 1, 1, 0)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("fonts/Ubuntu-C.ttf", 192); err != nil {
		panic(err)
	}

	dc.DrawImage(im, 0, 0)

	dc.DrawStringAnchored(date, 250, 400, 0, 0)

	filename := strings.Replace(date, "/", " ", -1)

	// make output dir
	if _, err = os.Stat("output/"); os.IsNotExist(err) {
		os.Mkdir("output/", 0700)
	}

	err = dc.SavePNG("output/" + filename + ".png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Generated %s.png\n", filename)
}
