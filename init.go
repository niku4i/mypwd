package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

var currentUser = user.Current
var key = "FnT=ba;NeSDHt0FUPZW1WAY:/Kh=WWK?"

func exitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init",
	Short:     "initialize password database",
	Long:      "initialize password database",
}

func runInit(cmd *Command, args []string) {
	fmt.Print("Enter master password: ")
	scanner := Scanner{
		Ask:        "Enter masster password: ",
		Validation: true,
	}
	passwd := scanner.Input()
	fmt.Println(passwd)

	user, userErr := currentUser()
	handleErr(userErr)
	baseDir := filepath.Join(user.HomeDir, ".mypwd")
	dbDir := filepath.Join(baseDir, "db")
	err := os.MkdirAll(dbDir, 0755)
	handleErr(err)

	services := filepath.Join(dbDir, "services.json")
	f, err := os.OpenFile(services, os.O_WRONLY, 0755)
	handleErr(err)

	s := "[]"
	_, err = f.WriteString(s)
	handleErr(err)
	f.Close()

	//passwd := filepath.Join(dbDir, "passwd.json")
	json := []byte("[]")
	ciphertext := encode(json)
	fmt.Printf("%s=>%x\n", json, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, decode(ciphertext))
}

func encode(plaintext []byte) []byte {
	// 暗号化アルゴリズムaesを生成
	var block, err = aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// 暗号化文字列
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func decode(ciphertext []byte) []byte {
	// 暗号化アルゴリズムaesを生成
	var block, err = aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 暗号化文字列
	cfb := cipher.NewCFBDecrypter(block, iv)
	//plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}
