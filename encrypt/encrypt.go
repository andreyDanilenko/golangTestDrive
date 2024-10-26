package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypt struct {
	Key string
}

func NewEnrypter() *Encrypt {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параматер KEY")
	}
	return &Encrypt{
		Key: key,
	}
}

func (enc *Encrypt) Encrypt(plainStr []byte) ([]byte, error) {
	//создание блока шифрования
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}
	// создание объекта GCM
	// GCM обеспечивает конфедециальность и целостность данных
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// создается срез байтов длиной, соответствующей размеру nonce
	// криптографический случайный генератор rand.Reader чтобы заполнить этот nonce
	// nonce - вектор инициализации для каждой операции шифрования.
	nonce := make([]byte, aesGCM.NonceSize())
	_, err1 := io.ReadFull(rand.Reader, nonce)
	if err1 != nil {
		return nil, err1
	}
	// nonce – место, куда будет записан результат. Здесь используется сам nonce, чтобы добавить его в начало зашифрованного текста.
	// nonce – вектор инициализации, который должен быть уникальным для каждого вызова шифрования.
	// plainStr - исходные данные, которые нужно зашифровать.
	// nil – дополнительные данные для аутентификации (AAD). В данном случае они не используются.
	ciphertext := aesGCM.Seal(nonce, nonce, plainStr, nil)
	return ciphertext, nil
}

func (enc *Encrypt) Decrypt(encryptedStr []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainStr, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainStr, nil
}
