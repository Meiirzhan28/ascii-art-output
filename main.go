package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"output/src"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		str := os.Args[1]
		banner := os.Args[2]
		flag := os.Args[3]

		if src.Checkflag(flag) == nil {
			q := strings.Split(flag, "=")

			if src.Validstr(str, banner) == nil {
				if src.Banner(banner) == nil {
					file, err := os.Open(banner + ".txt")
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)
					ascii := []string{}
					for scanner.Scan() {
						s := strings.ReplaceAll(scanner.Text(), "\n", "")
						ascii = append(ascii, s)
					}
					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}
					d := src.ReadOut(ascii, str)

					f, err := os.Create(q[1])
					if err != nil {
						log.Fatal(err)
					}
					h := bufio.NewWriter(f)
					h.WriteString(d)
					h.Flush()
				} else {
					fmt.Println(src.Banner(banner))
				}
			} else {
				fmt.Println(src.Validstr(str, banner))
			}
		} else {
			fmt.Println(src.Checkflag(flag))
		}
	} else if len(os.Args) == 3 {
		str := os.Args[1]
		banner := "standard"
		flag := os.Args[2]

		if src.Checkflag(flag) == nil {
			q := strings.Split(flag, "=")

			if src.Validstr(str, banner) == nil {
				if src.Banner(banner) == nil {
					file, err := os.Open(banner + ".txt")
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)
					ascii := []string{}
					for scanner.Scan() {
						s := strings.ReplaceAll(scanner.Text(), "\n", "")
						ascii = append(ascii, s)
					}
					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}
					d := src.ReadOut(ascii, str)

					f, err := os.Create(q[1])
					if err != nil {
						log.Fatal(err)
					}
					h := bufio.NewWriter(f)
					h.WriteString(d)
					h.Flush()
				} else {
					fmt.Println(src.Banner(banner))
				}
			} else {
				fmt.Println(src.Validstr(str, banner))
			}
		} else {
			fmt.Println(src.Checkflag(flag))
		}
	}
}
