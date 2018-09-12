package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/ejson/crypto"
	"github.com/therecipe/qt/widgets"
)

var myKP crypto.Keypair

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 300)
	window.SetWindowTitle("Secret Sender")

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(mainWidget)

	myKeyLabel := widgets.NewQLabel2("Your Key:", nil, 0)
	myKeyLabel.Font().SetPointSize(11)

	myKeyInput := widgets.NewQLineEdit(nil)
	myKeyInput.SetPlaceholderText("Please generate a key")
	myKeyInput.SetReadOnly(true)

	myKeyGenButton := widgets.NewQPushButton2("Generate", nil)
	myKeyGenButton.ConnectClicked(func(bool) {
		myKeyInput.SetText("Generating...")
		app.ProcessEvents(0)

		myKP.Generate()
		myKeyInput.SetText(myKP.PublicString())
	})

	txKeyLabel := widgets.NewQLabel2("Recipient Key:", nil, 0)
	txKeyLabel.Font().SetPointSize(11)

	txKeyInput := widgets.NewQLineEdit(nil)
	txKeyInput.SetPlaceholderText("Paste Recipient Key")

	messageTextLabel := widgets.NewQLabel2("Text to encrypt/decrypt:", nil, 0)
	messageTextLabel.Font().SetPointSize(11)

	messageTextInput := widgets.NewQPlainTextEdit(nil)

	messageStatusLabel := widgets.NewQLabel(nil, 0)
	messageStatusLabel.Font().SetPointSize(10)
	messageStatusLabel.SetWordWrap(true)
	messageStatusLabel.SetHidden(true)

	cryptWidget := widgets.NewQWidget(nil, 0)
	cryptLayout := widgets.NewQHBoxLayout()
	cryptLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	cryptWidget.SetLayout(cryptLayout)

	encryptButton := widgets.NewQPushButton2("Encrypt", nil)
	cryptLayout.AddWidget(encryptButton, 0, 0)
	cryptWidget.Layout().AddWidget(encryptButton)
	encryptButton.ConnectClicked(func(bool) {
		var txKP crypto.Keypair

		emptyBytes := make([]byte, 32)
		if bytes.Equal(myKP.Public[:], emptyBytes) {
			widgets.QMessageBox_Warning(nil, "No Public Key", "Please generate a keypair before encrypting a message", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
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

		encrypter := myKP.Encrypter(txKP.Public)
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
		if bytes.Equal(myKP.Public[:], emptyBytes) {
			widgets.QMessageBox_Warning(nil, "No Public Key", "Please generate a keypair to decrypt the message with", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		decrypter := myKP.Decrypter()
		decrypted, err := decrypter.Decrypt([]byte(messageTextInput.ToPlainText()))
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Error Decrpying", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		pubKeyEncoded := strings.Split(messageTextInput.ToPlainText(), ":")[1]
		pubKey, _ := base64.StdEncoding.DecodeString(pubKeyEncoded)
		app.ProcessEvents(0)
		messageStatusLabel.SetHidden(false)
		messageStatusLabel.SetText(fmt.Sprintf("Encrypted by: %x", pubKey))

		app.ProcessEvents(0)
		messageTextInput.SetPlainText(string(decrypted))
		app.ProcessEvents(0)
	})

	mainWidget.Layout().AddWidget(myKeyLabel)
	mainWidget.Layout().AddWidget(myKeyInput)
	mainWidget.Layout().AddWidget(myKeyGenButton)
	mainWidget.Layout().AddWidget(txKeyLabel)
	mainWidget.Layout().AddWidget(txKeyInput)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)
	mainWidget.Layout().AddWidget(messageStatusLabel)
	mainWidget.Layout().AddWidget(cryptWidget)

	window.Show()

	app.Exec()
}
