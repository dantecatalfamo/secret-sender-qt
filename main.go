package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Shopify/ejson/crypto"
	"github.com/therecipe/qt/core"
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

	myKeyButtonWidget := widgets.NewQWidget(nil, 0)
	myKeyButtonLayout := widgets.NewQHBoxLayout()
	myKeyButtonLayout.Layout().SetContentsMargins(0, 0, 0, 0)
	myKeyButtonWidget.SetLayout(myKeyButtonLayout)

	myKeyGenButton := widgets.NewQPushButton2("Generate", nil)
	myKeyButtonWidget.Layout().AddWidget(myKeyGenButton)
	myKeyGenButton.ConnectClicked(func(bool) {
		myKeyInput.SetText("Generating...")
		app.ProcessEvents(0)

		myKP.Generate()
		myKeyInput.SetText(myKP.PublicString())
	})

	myKeySaveButton := widgets.NewQPushButton2("Save", nil)
	myKeyButtonWidget.Layout().AddWidget(myKeySaveButton)
	myKeySaveButton.ConnectClicked(func(bool) {
		fileName := widgets.QFileDialog_GetSaveFileName(nil, "Save Keypair As", core.QDir_HomePath(), "", "", 0)
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Save Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		pubString := base64.StdEncoding.EncodeToString(myKP.Public[:])
		privString := base64.StdEncoding.EncodeToString(myKP.Private[:])
		comboString := fmt.Sprintf("%s:%s", pubString, privString)
		_, err = file.Write([]byte(comboString))
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Save Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		return
	})

	myKeyLoadButton := widgets.NewQPushButton2("Load", nil)
	myKeyButtonWidget.Layout().AddWidget(myKeyLoadButton)
	myKeyLoadButton.ConnectClicked(func(bool) {
		fileName := widgets.QFileDialog_GetOpenFileName(nil, "Load Keypair", core.QDir_HomePath(), "", "", 0)
		file, err := os.Open(fileName)
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Load Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Load Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		fileString := string(fileBytes)
		pubPrivEncoded := strings.Split(fileString, ":")

		pubKey, err := base64.StdEncoding.DecodeString(pubPrivEncoded[0])
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Load Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		privKey, err := base64.StdEncoding.DecodeString(pubPrivEncoded[1])
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Load Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		for i, b := range pubKey {
			myKP.Public[i] = b
		}

		for i, b := range privKey {
			myKP.Private[i] = b
		}

		app.ProcessEvents(0)
		myKeyInput.SetText(myKP.PublicString())
		app.ProcessEvents(0)
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

		keyBytes, err := hex.DecodeString(keyStr)
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Key Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		for i, b := range keyBytes {
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
	mainWidget.Layout().AddWidget(myKeyButtonWidget)
	mainWidget.Layout().AddWidget(txKeyLabel)
	mainWidget.Layout().AddWidget(txKeyInput)
	mainWidget.Layout().AddWidget(messageTextLabel)
	mainWidget.Layout().AddWidget(messageTextInput)
	mainWidget.Layout().AddWidget(messageStatusLabel)
	mainWidget.Layout().AddWidget(cryptWidget)

	window.Show()

	app.Exec()
}
