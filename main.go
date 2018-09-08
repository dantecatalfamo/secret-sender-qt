package main

import (
	"fmt"
	"os"

	"github.com/Shopify/ejson/crypto"
	"github.com/therecipe/qt/widgets"
)

var (
	txKP crypto.Keypair
	rxKP crypto.Keypair
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Secret Sender")

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(mainWidget)

	rxKeyLabel := widgets.NewQLabel2("Your Key:", nil, 0)
	rxKeyLabel.Font().SetPointSize(11)

	rxKeyInput := widgets.NewQLineEdit(nil)
	rxKeyInput.SetPlaceholderText("Please generate a key")
	rxKeyInput.SetReadOnly(true)

	rxKeyGenButton := widgets.NewQPushButton2("Generate", nil)
	rxKeyGenButton.ConnectClicked(func(bool) {
		rxKeyInput.SetText("Generating...")
		app.ProcessEvents(0)

		rxKP.Generate()
		rxKeyInput.SetText(rxKP.PublicString())
		fmt.Println(rxKP.PublicString())

	})

	txKeyLabel := widgets.NewQLabel2("Recipiant Key:", nil, 0)
	txKeyLabel.Font().SetPointSize(11)

	txKeyInput := widgets.NewQLineEdit(nil)
	txKeyInput.SetPlaceholderText("Paste Recipient Key")

	messageTextLabel := widgets.NewQLabel2("Text to encrypt/decrypt:", nil, 0)
	messageTextLabel.Font().SetPointSize(11)

	messageTextInput := widgets.NewQPlainTextEdit(nil)

	cryptWidget := widgets.NewQWidget(nil, 0)
	cryptLayout := widgets.NewQHBoxLayout()
	cryptLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	cryptWidget.SetLayout(cryptLayout)

	encryptButton := widgets.NewQPushButton2("Encrypt", nil)
	cryptLayout.AddWidget(encryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(encryptButton)

	decryptButton := widgets.NewQPushButton2("Decrypt", nil)
	cryptLayout.AddWidget(decryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(decryptButton)

	mainWidget.Layout().AddWidget(rxKeyLabel)
	mainWidget.Layout().AddWidget(rxKeyInput)
	mainWidget.Layout().AddWidget(rxKeyGenButton)
	mainWidget.Layout().AddWidget(txKeyLabel)
	mainWidget.Layout().AddWidget(txKeyInput)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)
	mainWidget.Layout().AddWidget(cryptWidget)

	// button := widgets.NewQPushButton2("Generate", nil)
	// button.ConnectClicked(func(bool) {
	// 	widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	// })
	//mainWidget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
