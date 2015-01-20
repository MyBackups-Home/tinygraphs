package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/controllers/checkerboard"
	"github.com/taironas/tinygraphs/controllers/isogrids"
	"github.com/taironas/tinygraphs/controllers/squares"
)

var root = flag.String("root", "app", "file system path")

func main() {
	r := new(route.Router)

	r.HandleFunc("/checkerboard/?", checkerboard.Checkerboard)
	r.HandleFunc("/checkerboard/[0-8]/?", checkerboard.Color)

	r.HandleFunc("/squares/?", squares.Random)
	r.HandleFunc("/squares/random/?", squares.Random)
	r.HandleFunc("/squares/random/[0-8]/?", squares.RandomColor)
	r.HandleFunc("/squares/[a-zA-Z0-9\\.]+/?", squares.Square)      //cached
	r.HandleFunc("/squares/[0-8]/[a-zA-Z0-9\\.]+/?", squares.Color) // cached

	r.HandleFunc("/isogrids/skeleton/[a-zA-Z0-9]+/?", isogrids.Skeleton)
	r.HandleFunc("/isogrids/diagonals/[a-zA-Z0-9]+/?", isogrids.Diagonals)
	r.HandleFunc("/isogrids/halfdiagonals/[a-zA-Z0-9]+/?", isogrids.HalfDiagonals)
	r.HandleFunc("/isogrids/color/[a-zA-Z0-9]+/?", isogrids.Color)
	r.HandleFunc("/isogrids/[a-zA-Z0-9]+/?", isogrids.Isogrids)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
