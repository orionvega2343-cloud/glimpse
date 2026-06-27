package overlay

import (
	"glimpse/internal/actions"
	"glimpse/internal/api"
	"glimpse/internal/capture"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Overlay() {
	a := app.New()
	w := a.NewWindow("Glimpse")

	w.Resize(fyne.NewSize(800, 600))

	brief := widget.NewLabel("Click to the button,for get summary and answers of tests")
	brief.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(brief)
	scroll.SetMinSize(fyne.NewSize(780, 450))
	w.SetContent(container.NewVBox(
		scroll,
		widget.NewButton("Get brief summary", func() {

			prompt := actions.Brief()

			base64, err := capture.Capture()
			if err != nil {
				log.Fatal(err)
			}

			result, err := api.Send(base64, prompt)
			if err != nil {
				log.Fatal(err)
			}

			brief.SetText(result)
		}),

		widget.NewButton("Get full summary", func() {
			prompt := actions.Full()

			base64, err := capture.Capture()
			if err != nil {
				log.Fatal(err)
			}

			result, err := api.Send(base64, prompt)
			if err != nil {
				log.Fatal(err)
			}
			brief.SetText(result)

		}),
		widget.NewButton("Get test summary", func() {
			prompt := actions.Test()
			base64, err := capture.Capture()
			if err != nil {
				log.Fatal(err)
			}

			result, err := api.Send(base64, prompt)
			if err != nil {
				log.Fatal(err)
			}
			brief.SetText(result)

		}),
	))

	w.ShowAndRun()
}
