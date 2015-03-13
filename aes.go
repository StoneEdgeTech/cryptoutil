package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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

func _AESEncrypt(key, text []byte, iv []byte) ([]byte, error) {
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
