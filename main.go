package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Secret Sender")

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(mainWidget)

	rxKeyWidget := widgets.NewQWidget(nil, 0)
	rxKeyWidget.SetLayout(widgets.NewQHBoxLayout())
	rxKeyWidget.SetContentsMargins(0, 0, 0, 0)

	rxKeyLabel := widgets.NewQLabel2("Your Key:", nil, 0)
	rxKeyWidget.Layout().AddWidget(rxKeyLabel)

	rxKeyInput := widgets.NewQLineEdit(nil)
	rxKeyInput.SetPlaceholderText("Please generate a key")
	rxKeyWidget.Layout().AddWidget(rxKeyInput)

	rxKeyGenButton := widgets.NewQPushButton2("Generate", nil)
	rxKeyWidget.Layout().AddWidget(rxKeyGenButton)

	txKeyWidget := widgets.NewQWidget(nil, 0)
	txKeyWidget.SetLayout(widgets.NewQHBoxLayout())

	txKeyLabel := widgets.NewQLabel2("Recipiant Key:", nil, 0)
	txKeyWidget.Layout().AddWidget(txKeyLabel)

	txKeyInput := widgets.NewQLineEdit(nil)
	txKeyInput.SetPlaceholderText("Paste Recipiant Key")
	txKeyWidget.Layout().AddWidget(txKeyInput)

	//messageTextWidget := widgets.NewQWidget(nil, 0)
	//messageTextWidget.SetLayout(widgets.NewQVBoxLayout())

	messageTextLabel := widgets.NewQLabel2("Text to encrypt/decrypt:", nil, 0)
	//messageTextWidget.Layout().AddWidget(messageTextLabel)

	messageTextInput := widgets.NewQPlainTextEdit(nil)
	//messageTextWidget.Layout().AddWidget(messageTextInput)

	mainWidget.Layout().AddWidget(rxKeyWidget)
	mainWidget.Layout().AddWidget(txKeyWidget)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)

	//mainWidget.Layout().AddWidget(messageTextWidget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Key...")
	mainWidget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("Generate", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	mainWidget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
