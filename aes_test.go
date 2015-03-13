package cryptoutil

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestBigCommerce(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Authentication", func() {
		g.It("should look up authentication", func() {})
	})
	g.Describe("Cryptography", func() {
		g.It("should generate random initialization vector", func() {
			iv := makeIV()
			Expect(len(iv)).To(Equal(16))
		})
		g.It("should encrypt a string with a given key", func() {
			iv := []byte{255, 137, 197, 214, 116, 165, 208, 214, 167, 19, 82, 250, 13, 192, 62, 145}
			key := "this is a key which is exactly sixteen bytes long"[:16]
			plaintext := "this is the plaintext"
			encrypted, err := _AESEncrypt([]byte(key), []byte(plaintext), iv)
			Expect(err).NotTo(HaveOccurred())
			expected := []byte{60, 187, 164, 154, 208, 202, 12, 73, 118,
				192, 13, 14, 12, 216, 10, 161, 44, 183, 157, 30, 28}
			expected = append(iv, expected...)
			Expect(encrypted).To(Equal(expected))
		})
		g.It("should decrypt a string assuming the first 16 bytes are the initialization vector", func() {
			key := "this is a key which is exactly sixteen bytes long"[:16]
			encrypted := []byte{255, 137, 197, 214, 116, 165,
				208, 214, 167, 19, 82, 250, 13, 192, 62, 145,
				60, 187, 164, 154, 208, 202, 12, 73, 118, 192,
				13, 14, 12, 216, 10, 161, 44, 183, 157, 30, 28}
			decrypted, err := AESDecrypt([]byte(key), encrypted)
			Expect(err).NotTo(HaveOccurred())
			Expect(decrypted).To(Equal("this is the plaintext"))
		})
	})
}
