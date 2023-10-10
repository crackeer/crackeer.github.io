package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	// 生成私钥和公钥
	privKeyA, pubKeyA, err := generateECDSAKeyPair()
	if err != nil {
		panic(err)
	}

	privKeyB, pubKeyB, err := generateECDSAKeyPair()
	if err != nil {
		panic(err)
	}

	// 生成共享密钥
	sharedKeyA, err := generateSharedSecret(privKeyA, pubKeyB)
	if err != nil {
		panic(err)
	}
	sharedKeyB, err := generateSharedSecret(privKeyB, pubKeyA)

	fmt.Println(string(sharedKeyA), string(sharedKeyB))
	// 加密明文
	plaintext := []byte("Hello, World!")
	ciphertext, err := encrypt(plaintext, sharedKeyA)
	if err != nil {
		panic(err)
	}

	// 解密密文
	decrypted, err := decrypt(ciphertext, append(sharedKeyB, '9'))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
}

// 生成ECDSA公私钥对
func generateECDSAKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	curve := elliptic.P256() // 使用P-256曲线
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

// 使用ECDH算法生成共享密钥
func generateSharedSecret(privKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) ([]byte, error) {
	curve := elliptic.P256() // 使用P-256曲线
	x, _ := curve.ScalarMult(pubKey.X, pubKey.Y, privKey.D.Bytes())
	sharedKey := x.Bytes()
	return sharedKey, nil
}

// 使用共享密钥加密
func encrypt(plaintext []byte, sharedKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	return append(iv, ciphertext...), nil
}

// 使用共享密钥解密
func decrypt(ciphertext []byte, sharedKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
