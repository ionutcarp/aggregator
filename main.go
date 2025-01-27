package main

import (
	"fmt"
	"github.com/ionutcarp/aggregator/internal/config"
	"log"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	fmt.Println("postgres://example")
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config file: %+v\n", cfg)

	err = cfg.SetUser("ionut")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}
	fmt.Printf("Set user: %+v\n", cfg)

	/*	reader := bufio.NewScanner(os.Stdin)
		//kk
				for {
					fmt.Print("aggreGATOR> ")
					reader.Scan()
					words := cleanInput(reader.Text())
					if len(words) == 0 {
						continue
					}
					fmt.Printf("postgres://example")
				}*/
}
