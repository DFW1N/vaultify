package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"os"
	"strings"
)

func deriveKey(passphrase string, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return nil, nil, fmt.Errorf("%w", err)
		}
	}
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)
	return key, salt, nil
}

func encryptContents(filestream []byte, passphrase string) ([]byte, error) {
	key, salt, err := deriveKey(passphrase, nil)
	if err != nil {
		return nil, fmt.Errorf("❌ Error generating salt: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("❌ Error creating cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("❌ Error creating GCM: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("❌ Error generating nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nonce, nonce, filestream, nil)

	if ciphertext == nil {
		return nil, fmt.Errorf("❌ Error creating output file: %w", err)
	}

	encryptStream := fmt.Sprintf("%s-%s", hex.EncodeToString(salt), hex.EncodeToString(ciphertext))

	return []byte(encryptStream), nil
}

func decryptFile(filename string, passphrase string) (string, error) {
	outfileName := strings.TrimSuffix(filename, ".enc")
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("❌ Error reading file: %w", err)
	}

	parts := strings.Split(string(data), "-")
	if len(parts) != 2 {
		return "", fmt.Errorf("❌ Error invalid ciphertext format")
	}

	salt, _ := hex.DecodeString(parts[0])
	ciphertext, _ := hex.DecodeString(parts[1])

	key, _, err := deriveKey(passphrase, salt)
	if err != nil {
		return "", fmt.Errorf("❌ Error parsing salt: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("❌ Error creating cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("❌ Error creating GCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("❌ Error decrypting: %w", err)
	}

	err = os.WriteFile(outfileName, plaintext, 0644)
	if err != nil {
		return "", fmt.Errorf("❌ Error writing decrypted file: %w", err)
	}

	return "", nil
}
