package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "errors"
    "os"
)

func EncryptPassword(password string) (string, error) {
    key := []byte(os.Getenv("ENCRYPTION_KEY"))
    plaintext := []byte(password)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonce := make([]byte, gcm.NonceSize())
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptPassword(encryptedPassword string) (string, error) {
    key := []byte(os.Getenv("ENCRYPTION_KEY"))
    ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    if len(ciphertext) < gcm.NonceSize() {
        return "", errors.New("malformed ciphertext")
    }

    nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
}