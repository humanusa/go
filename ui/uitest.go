package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {

	go func() {

		window := new(app.Window)

		err := run(window)

		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

type C = layout.Context
type D = layout.Dimensions

func run(window *app.Window) error {

	theme := material.NewTheme()

	var ops op.Ops

	var boilDurationInput widget.Editor

	for {

		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(
				gtx,

				// The inputbox
				layout.Rigid(

					func(gtx C) D {

						ed := material.Editor(theme, &boilDurationInput, "sec")

						// Define insets ...
						margins := layout.Inset{
							Top:    unit.Dp(0),
							Right:  unit.Dp(170),
							Bottom: unit.Dp(40),
							Left:   unit.Dp(179),
						}

						// ... and borders ...
						border := widget.Border{
							Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
							CornerRadius: unit.Dp(3),
							Width:        unit.Dp(2),
						}

						// ... before laying it out, one inside the other
						return margins.Layout(gtx,
							func(gtx C) D {
								return border.Layout(gtx, ed.Layout)
							},
						)
					}))

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}

			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)

		}
	}
}
