package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func main() {
	var (
		generate = flag.Bool("g", false, "generate new 32 byte key")
		file     = flag.String("f", "key32", "name of the file to store encrypted key")
	)
	flag.Parse()

	if *generate {
		log.Println("Generateing a new 32 byte random key")
		gen32bytes()
		return
	}

	key := gen32bytes()
	enkey := new([32]byte)
	pass := gen32bytes()
	copy(enkey[:], pass[:32])
	nonce := new([24]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		log.Fatal(err)
	}
	out := make([]byte, 24)
	copy(out, nonce[:])
	out = secretbox.Seal(out, key, nonce, enkey)

	if err = os.WriteFile(*file, out, 0644); err != nil {
		log.Fatal(err)
	}

	log.Print("Passphrase: ", base64.RawStdEncoding.EncodeToString(pass))

}

func gen32bytes() []byte {
	msg := new([32]byte)
	if _, err := rand.Read(msg[:]); err != nil {
		fmt.Println("Error: ", err)
	}
	return msg[:]
}
