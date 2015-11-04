package main

import (
	"log"

	"github.com/smreed/kubecon-2015/pkg/util"
)

func main() {
	images, err := util.FindImagesInCwd()
	if err != nil {
		log.Fatal("error finding images:", err)
	}
	for _, img := range images {
		log.Println("Building", img)
		if err = img.Build(); err != nil {
			log.Println("error building", img)
			log.Fatal(err)
		}
		log.Println("Pushing", img)
		if err = img.Push(); err != nil {
			log.Println("error ppushing", img)
			log.Fatal(err)
		}
	}

	log.Println("Done")
}
