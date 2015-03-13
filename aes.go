package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func lookupOauthToken() {
}

func makeIV() []byte {
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
	return iv
}

func AESEncrypt(key, text []byte) ([]byte, error) {
	return _AESEncrypt(key, text, makeIV())
}

func _AESEncrypt(key, text, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(text))
	encrypter.XORKeyStream(encrypted, text)
	encrypted = append(iv, encrypted...)
	return encrypted, nil
}

func AESDecrypt(key, ciphertext []byte) (string, error) {
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	decripter := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decripter.XORKeyStream(decrypted, ciphertext)
	return string(decrypted), nil
}

func _B64AESEncrypt(key, text, iv []byte) (string, error) {
	encryptedBytes, err := _AESEncrypt(key, text, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func B64AESEncrypt(key, text []byte) (string, error) {
	return _B64AESEncrypt(key, text, makeIV())
}

func B64AESDecrypt(key []byte, b64Ciphertext string) (string, error) {
	rawEncrypted, err := base64.StdEncoding.DecodeString(b64Ciphertext)
	if err != nil {
		return "", err
	}
	return AESDecrypt(key, rawEncrypted)
}
