package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Keys struct {
	TwitterConsumerKey       string `yaml:"TwitterConsumerKey"`
	TwitterConsumerSecret    string `yaml:"TwitterConsumerSecret"`
	TwitterAccessToken       string `yaml:"TwitterAccessToken"`
	TwitterAccessTokenSecret string `yaml:"TwitterAccessTokenSecret"`
}

func main() {
	f, err := os.Open("./secret.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// m := make(map[string]string)
	k := Keys{}
	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(l, &k)
		// err = yaml.Unmarshal(l, &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	fmt.Println(k)
}
