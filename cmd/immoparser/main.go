package main

import (
	"log"

	"github.com/leangeder/immoparser/pkg/sites"
	"github.com/leangeder/immoparser/pkg/sites/domimmo"
	"github.com/leangeder/immoparser/pkg/sites/leboncoin"
)

func main() {
	log.Println("Testing")
	listAnnonceleboncoin := leboncoin.Annonce{}
	listAnnoncedomimmo := domimmo.Annonce{}

	log.Println(sites.Get(listAnnonceleboncoin))
	log.Println(sites.Get(listAnnoncedomimmo))
}
