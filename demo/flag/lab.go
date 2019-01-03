package flag

import (
	"flag"
	"strings"
)

func Do() {
	var h = flag.Bool("h", false, "查看帮助")
	var help = flag.Bool("help", false, "查看帮助")

	var do string
	flag.StringVar(&do,"do","","do something")

	flag.Parse()

	if *h||*help{
		flag.PrintDefaults()
	}

	if len(strings.TrimSpace(do))!=0{
		println(do)
	}

}
