package main

import (
	"bufio"
	"fmt"
	"fs/initial"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		th := os.Args[2]
		t := th + ".txt"
		out := os.Args[3]
		if initial.Check_for_banner(out) {
			w := strings.SplitAfter(out, "=")
			if initial.Check_for_hash(t) {
				file, err := os.Open(t)
				if err != nil {
					log.Fatal(err)
				}
				r := bufio.NewScanner(file)
				var array []string
				for r.Scan() {
					array = append(array, r.Text())
				}
				arg := os.Args[1]
				d := initial.Check_for_letters(th, arg)
				if d != nil {
					fmt.Println(d.Error())
					return
				}
				fl, _ := os.Create(w[1])

				m := initial.Create_the_map(array, arg)
				h := bufio.NewWriter(fl)
				h.WriteString(m)
				h.Flush()
			} else {
				fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --output=<fileName.txt>")
			}
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --output=<fileName.txt>")
		}
	} else if len(os.Args) == 2 {
		f := "standard"
		arg := os.Args[1]
		d := initial.Check_for_letters(f, arg)
		if d != nil {
			fmt.Println(d.Error())
			return
		}
		file, err := os.Open(f + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		r := bufio.NewScanner(file)
		var array []string
		for r.Scan() {
			array = append(array, r.Text())
		}
		fmt.Print(initial.Create_the_map(array, arg))
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --output=<fileName.txt>")
	}
}
