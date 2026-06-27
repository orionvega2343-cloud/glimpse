package flags

import (
	"flag"
	"fmt"
	"os"
)

//Usage - функция справка появляется в случае невалидных данных

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: prompt <flag>")
	flag.PrintDefaults()
	os.Exit(2)
}

func GetPrompt(flag string) string {
	if flag == "" {
		return "Ошибка"
	}

	switch flag {
	case "b":
		return "Сделай мне краткую сводку в 3-4 абзаца исходя из изображения"
	case "f":
		return "Сделай мне полную сводку в 10-12 абзацев исходя из изображения"

	case "t":
		return "Проанализируй и проверь актуальную информацию по вопросам теста которые ты видишь на изображении,кратко ответь на каждый из них"
	default:
		return ""
	}

}
