package main

import (
	"log"
	"sync"

	"github.com/smreed/kubecon-2015/pkg/util"
)

var wg sync.WaitGroup

func main() {
	images, err := util.FindImagesInCwd()
	if err != nil {
		log.Fatal("error finding images:", err)
	}

	wg.Add(len(images))
	built := make(chan util.Image)

	// Build images in goroutines
	go func() {
		for _, img := range images {
			go func(img util.Image) {
				log.Println("Building", img)
				if err = img.Build(); err != nil {
					log.Println("error building", img)
					log.Fatal(err)
				}
				built <- img
			}(img)
		}
	}()

	// Push built images in goroutines
	go func() {
		for img := range built {
			log.Println("Pushing", img)
			if err = img.Push(); err != nil {
				log.Println("error pushing", img)
				log.Fatal(err)
			}
			wg.Done()
		}
	}()

	wg.Wait()
	log.Println("Done")
}
