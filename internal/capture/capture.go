package capture

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/go-vgo/robotgo"
)

//Захватываем экран и сохраняем изображение

func capture() (string, error) {
	//Скрин экрана
	img1, _ := robotgo.CaptureImg()

	//Энкодим в байты
	var buf bytes.Buffer
	err := png.Encode(&buf, img1)
	if err != nil {
		return "", err
	}

	byteSlice := buf.Bytes()

	//Ответ в формате base64
	return base64.StdEncoding.EncodeToString(byteSlice), nil
}
