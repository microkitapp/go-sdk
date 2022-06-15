package microkit

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"crypto/aes"
    // "encoding/hex"
	// "io"
	// "crypto/rand"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
)

type MicroKit struct{
	KitReady bool
	ConfigKit map[string]ConfigGroup
}

type ApiRes struct {
	Config string
}


var kitInstance *MicroKit = nil

func  InitKit(sdk_key string, options map[string]string) (*MicroKit) {
	keySize := 24
	if kitInstance == nil {
		kitInstance = new(MicroKit)
		client := new(HttpClient)
		var secret string
		sdk_key, secret = parseSdkKey(sdk_key, keySize, options["service"])
		
		resp := client.Get(options["url"], sdk_key)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		
		
		var data ApiRes
		json.Unmarshal(bodyBytes, &data)
		decryptBody, _ := decryptCBC(data.Config, secret)
		
		var rawStructJson map[string]map[string]map[string]string
		json.Unmarshal(decryptBody, &rawStructJson)
		// log.Fatalf("%v", rawStructJson)
		kitInstance.ConfigKit = InitConfigKit(rawStructJson)
		
		kitInstance.KitReady = true
	}
	return kitInstance
}

func decryptCBC( content, textKey string) (plaintext []byte, err error) {
    key, _ := base64.StdEncoding.DecodeString(textKey)
    ciphertext, _ := base64.StdEncoding.DecodeString(content)
	var block cipher.Block

    if block, err = aes.NewCipher(key); err != nil {
        return
    }

    if len(ciphertext) < aes.BlockSize {
        fmt.Printf("ciphertext too short")
        return
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    cbc := cipher.NewCBCDecrypter(block, iv)
    cbc.CryptBlocks(ciphertext, ciphertext)

    plaintext = PKCS5Trimming(ciphertext)

    return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func parseSdkKey (key string, keySize int, service string) (string, string) {
	var sdkKey string
	if service != "" {
		sdkKey = fmt.Sprintf("%s%s", key[:len(key)-keySize], service)
	} else {
		sdkKey = key[:len(key)-(keySize+1)]
	}
	return sdkKey, key[len(key)-keySize:]
}


