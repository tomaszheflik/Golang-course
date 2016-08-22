package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Message struct {
	Priv, Pub string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple REST service (ssh keys, deploy.jsom). Hello %s!", r.URL.Path[1:])
}

func sshkeys(w http.ResponseWriter, r *http.Request) {
	// genetare private Key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	err = priv.Validate()
	if err != nil {
		fmt.Println("Validation priv failed", err)
	}
	priv_def := x509.MarshalPKCS1PrivateKey(priv)

	priv_blk := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priv_def,
	}
	priv_pem := string(pem.EncodeToMemory(&priv_blk))

	// generate public key
	pub := priv.PublicKey
	pub_def, err := x509.MarshalPKIXPublicKey(&pub)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	pub_blk := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pub_def,
	}
	pub_pem := string(pem.EncodeToMemory(&pub_blk))
	fmt.Println(priv_pem)
	fmt.Println(pub_pem)

	// Write to browser
	m := Message{priv_pem, pub_pem}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/deploy/", handler).Methods("GET")
	r.HandleFunc("/ssh/", handler).Methods("GET")
	r.HandleFunc("/ssh/keys/", sshkeys).Methods("GET")
	http.ListenAndServe(":8081", r)

}
