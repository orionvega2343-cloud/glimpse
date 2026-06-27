package capture

import (
	"encoding/base64"
	"os"
	"os/exec"
)

//Захватываем экран и сохраняем изображение

func Capture() (string, error) {
	//Скрин экрана
	err := exec.Command("spectacle", "-b", "-n", "-o", "/tmp/screenshot.png").Run()
	if err != nil {
		return "", err
	}
	//Читаем файл
	r, err := os.ReadFile("/tmp/screenshot.png")
	if err != nil {
		return "", err
	}

	//Ответ в формате base64
	return base64.StdEncoding.EncodeToString(r), nil
}
