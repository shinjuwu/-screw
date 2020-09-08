package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

//AES ECB模式的加密解密
type AesTool struct {
	//128 192  256位的其中一个 长度 对应分别是 16 24  32字节长度
	Key []byte
}

func NewAesTool(key []byte) *AesTool {
	return &AesTool{Key: key}
}

//测试AES ECB 加密解密
func TestEncryptDecrypt() {
	key := []byte("8E003066E5FCFF03626DEBF05EDA1DB9")
	tool := NewAesTool(key)
	encrypt := tool.AESEncrypt([]byte("{\"token\":\"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxMDA2MDAzMSIsImlhdCI6MTU5OTQ2MDI1MywiZXhwIjoxNjAwMDY1MDUzfQ.5XamfusayAmydmTMX3pcmYhhFV1DjHBOhrv0qST5BgYOAGOa6YwjF3lek7msO9tsWA6onHQGE5tS8otvOD7x9g\"}"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt := tool.AesDecrypt(encrypt, key)
	fmt.Println(string(decrypt))
}

//测试AES ECB 加密解密
func TestHmacSha256() {
	key := "BAF266DA550D77A433D2503ADBA20A4D"
	oring := "App_Id=EB0B2F58C59A83A6E1&Input_Data=YNdNGwOHyiXXXl3oLVDfV8d3g4448jAZzdj54btT2EECb8WuW79JQqAYNuqpWX%2BcxtNCY8wSp0zqAaZQ6cfdViCLET%2FjTZE7eTwOeSNfcR0luRAuz0e7Ak5MRrshKRG2l5PzkLkdQGo2hcBgH%2B1S7MNjYlPgk4s8bd6x00laKfJ9pR3tTWBaEypaPDbOlZ4RrM%2FCMhKS5aWFXoAFgB6wRNfkLS1NZhq4c5KNbzTlE6C2UMTOWqNC7J%2FuOvkJP8EyvUM7PR33F97SD0nNYdE7hA%3D%3D&nonce=805496143&timestamp=1599550294"
	hmacSha256 := HmacSha256(oring, key)
	fmt.Println(hmacSha256)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

//ECB PKCS5 加密
func (this *AesTool) AESEncrypt(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	des := make([]byte, len(content))
	ecb.CryptBlocks(des, content)
	return des
}

//ECB PKCS5 解密
func (this *AesTool) AesDecrypt(crypted, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData
}

//PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func SignSHA1(orig string) []byte {
	h := sha1.New()
	h.Write([]byte(orig))
	bs := h.Sum(nil)
	return bs
}

func HmacSha256(oring string, key string) string {
	secret := []byte(key)
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(oring))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
