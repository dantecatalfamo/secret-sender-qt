package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

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
	window.SetMinimumSize2(250, 300)
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
	encryptButton.ConnectClicked(func(bool) {
		var tmpKP crypto.Keypair
		err := tmpKP.Generate()
		if err != nil {
			widgets.QMessageBox_Critical(nil, "No Public Key", "Please enter the recipient's public key", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		keyStr := strings.TrimSpace(txKeyInput.Text())
		if keyStr == "" {
			widgets.QMessageBox_Warning(nil, "No Public Key", "Please enter the recipient's public key", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		keyHex, err := hex.DecodeString(keyStr)
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Key Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		for i, b := range keyHex {
			txKP.Public[i] = b
		}

		encrypter := tmpKP.Encrypter(txKP.Public)
		cypherText, err := encrypter.Encrypt([]byte(messageTextInput.ToPlainText()))
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Encryption Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		app.ProcessEvents(0)
		messageTextInput.SetPlainText(string(cypherText))
		app.ProcessEvents(0)
	})

	decryptButton := widgets.NewQPushButton2("Decrypt", nil)
	cryptLayout.AddWidget(decryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(decryptButton)
	decryptButton.ConnectClicked(func(bool) {
		emptyBytes := make([]byte, 32)
		if bytes.Equal(rxKP.Public[:], emptyBytes) {
			widgets.QMessageBox_Warning(nil, "No Public Key", "Please generate a keypair to decrypt the message with", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		decrypter := rxKP.Decrypter()
		decrypted, err := decrypter.Decrypt([]byte(messageTextInput.ToPlainText()))
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Error Decrpying", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		app.ProcessEvents(0)
		messageTextInput.SetPlainText(string(decrypted))
		app.ProcessEvents(0)
	})

	mainWidget.Layout().AddWidget(rxKeyLabel)
	mainWidget.Layout().AddWidget(rxKeyInput)
	mainWidget.Layout().AddWidget(rxKeyGenButton)
	mainWidget.Layout().AddWidget(txKeyLabel)
	mainWidget.Layout().AddWidget(txKeyInput)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)
	mainWidget.Layout().AddWidget(cryptWidget)

	window.Show()

	app.Exec()
}
