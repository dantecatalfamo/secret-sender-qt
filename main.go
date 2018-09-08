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

	// rxKeyGroup := widgets.NewQGroupBox2("Your Key", nil)
	// rxKeyLayout := widgets.NewQVBoxLayout()
	// rxKeyLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	// rxKeyGroup.SetLayout(rxKeyLayout)
	// rxKeyWidget := widgets.NewQWidget(nil, 0)
	// rxKeyLayout := widgets.NewQVBoxLayout()
	// rxKeyLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	// rxKeyWidget.SetLayout(rxKeyLayout)

	rxKeyLabel := widgets.NewQLabel2("Your Key:", nil, 0)
	rxKeyLabel.Font().SetPointSize(11)
	//rxKeyWidget.Layout().AddWidget(rxKeyLabel)

	rxKeyInput := widgets.NewQLineEdit(nil)
	rxKeyInput.SetPlaceholderText("Please generate a key")
	rxKeyInput.SetReadOnly(true)
	//rxKeyGroup.Layout().AddWidget(rxKeyInput)
	//rxKeyWidget.Layout().AddWidget(rxKeyInput)

	rxKeyGenButton := widgets.NewQPushButton2("Generate", nil)
	//rxKeyGroup.Layout().AddWidget(rxKeyGenButton)
	//rxKeyWidget.Layout().AddWidget(rxKeyGenButton)

	// txKeyWidget := widgets.NewQWidget(nil, 0)
	// txKeyLayout := widgets.NewQVBoxLayout()
	// txKeyLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	// txKeyWidget.SetLayout(txKeyLayout)

	// txKeyGroup := widgets.NewQGroupBox2("Recipient Key", nil)
	// txKeyGroup.SetLayout(widgets.NewQVBoxLayout())

	txKeyLabel := widgets.NewQLabel2("Recipiant Key:", nil, 0)
	txKeyLabel.Font().SetPointSize(11)
	// txKeyWidget.Layout().AddWidget(txKeyLabel)

	txKeyInput := widgets.NewQLineEdit(nil)
	txKeyInput.SetPlaceholderText("Paste Recipient Key")
	// txKeyWidget.Layout().AddWidget(txKeyInput)

	//messageTextWidget := widgets.NewQWidget(nil, 0)
	//messageTextWidget.SetLayout(widgets.NewQVBoxLayout())

	messageTextLabel := widgets.NewQLabel2("Text to encrypt/decrypt:", nil, 0)
	messageTextLabel.Font().SetPointSize(11)
	//messageTextWidget.Layout().AddWidget(messageTextLabel)

	messageTextInput := widgets.NewQPlainTextEdit(nil)
	//messageTextWidget.Layout().AddWidget(messageTextInput)

	cryptWidget := widgets.NewQWidget(nil, 0)
	cryptLayout := widgets.NewQHBoxLayout()
	cryptLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	cryptWidget.SetLayout(cryptLayout)

	encryptButton := widgets.NewQPushButton2("Encrypt", nil)
	cryptLayout.AddWidget(encryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(encryptButton)

	//cryptWidget.Layout().AddItem(widgets.NewQSpacerItem(-1, -1, -1, -1))

	decryptButton := widgets.NewQPushButton2("Decrypt", nil)
	cryptLayout.AddWidget(decryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(decryptButton)

	mainWidget.Layout().AddWidget(rxKeyLabel)
	mainWidget.Layout().AddWidget(rxKeyInput)
	mainWidget.Layout().AddWidget(rxKeyGenButton)
	mainWidget.Layout().AddWidget(txKeyLabel)
	mainWidget.Layout().AddWidget(txKeyInput)
	//mainWidget.Layout().AddWidget(rxKeyGroup)
	// mainWidget.Layout().AddWidget(rxKeyWidget)
	// mainWidget.Layout().AddWidget(txKeyWidget)
	//mainWidget.Layout().AddWidget(txKeyGroup)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)
	mainWidget.Layout().AddWidget(cryptWidget)

	//mainWidget.Layout().AddWidget(messageTextWidget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Key...")
	//mainWidget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("Generate", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	//mainWidget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
