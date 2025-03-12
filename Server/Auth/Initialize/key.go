package Initialize

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/Rinai-R/Gocument/Server/Auth/handle"
	"log"
	"os"
)

func InitKey() {
	privatePEM, err := os.ReadFile("/home/rinai/PROJECTS/Gocument/keys/private.pem")
	if err != nil {
		log.Fatal("open error ", err)
	}
	block, _ := pem.Decode(privatePEM)
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	publicPEM, _ := os.ReadFile("/home/rinai/PROJECTS/Gocument/keys/public.pem")
	if privateKey == nil {
		log.Fatal("Private Key is nil")
	}
	handle.Authsrv = &handle.AuthService{
		PrivateKey: privateKey.(*rsa.PrivateKey),
		PublicKey:  string(publicPEM),
	}
}
