package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strconv"
// )

// func appendToFile(fd *os.File, str string) {
// 	if _, err := fd.WriteString(str); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	w := flag.Int("w", 100, "Image width")
// 	h := flag.Int("h", 100, "Image height")
// 	d := flag.Int("d", 100, "Color deep")
// 	o := flag.String("o", "out.ppm", "Output name")
// 	flag.Parse()
// 	f, err := os.OpenFile(*o, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
// 	defer f.Close()

// 	if err != nil {
// 		panic(err)
// 	}

// 	appendToFile(f, "P2\n")
// 	appendToFile(f, fmt.Sprintf("%d %d\n", *w, *h))
// 	appendToFile(f, fmt.Sprintf("%d\n", *d))
// 	for i := 0; i < *h; i++ {
// 		for j := 0; j < *w; j++ {
// 			// num, _ := ran(16)
// 			appendToFile(f, strconv.Itoa(rand.Intn(*d)))
// 			if j != *w-1 {
// 				appendToFile(f, "  ")
// 			}
// 		}
// 		log.Println(i)
// 		appendToFile(f, "\n")
// 	}
// }
