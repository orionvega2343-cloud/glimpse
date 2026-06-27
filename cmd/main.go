package main

import (
	"flag"
	"fmt"
	"glimpse/internal/api"
	"glimpse/internal/capture"
	"glimpse/internal/flags"
	"log"

	"github.com/joho/godotenv"
)

var (
	briefSum = flag.Bool("b", false, "brief sum")
	fullSum  = flag.Bool("f", false, "full sum")
	test     = flag.Bool("t", false, "test")
)

func main() {

	//Установка флагов
	log.SetFlags(0)
	log.SetPrefix("Prompt:")

	//Парсинг флагов
	flag.Usage = flags.Usage
	flag.Parse()

	var prompt string
	//Проверяем вызванный флаг
	if *briefSum {
		prompt = flags.GetPrompt("b")
	} else if *fullSum {
		prompt = flags.GetPrompt("f")
	} else if *test {
		prompt = flags.GetPrompt("t")
	}

	//Загрузка .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Захват скриншота
	base64, err := capture.Capture()
	if err != nil {
		log.Fatal(err)
	}

	send, err := api.Send(base64, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(send, err)

}
