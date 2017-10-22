package artland

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MakeCents-NYC/artland/pkg/mastercard"
)

var lastPost *PostRequest

type Artland struct {
	m *mastercard.MasterCard
}

func NewArtLand() *Artland {
	_m, err := mastercard.NewMasterCardAPI("keystore.p12", "protos/artland.proto", "", "keyalias", "keystorepass")
	if err != nil {
		return nil
	}

	result := &Artland{
		m: _m,
	}

	return result
}

func (al *Artland) Serve() error {
	http.HandleFunc("/post", postFunc)
	http.HandleFunc("/get", getFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

type PostRequest struct {
	Id  string
	Img string
}

func postFunc(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		w.WriteHeader(401)
		if _, err := w.Write([]byte("Bad body!")); err != nil {
			fmt.Println(err.Error())
		}
	}

	var data []byte
	if _, err := io.ReadFull(body, data); err != nil {
		w.WriteHeader(401)
		if _, err := w.Write([]byte("Bad body!")); err != nil {
			fmt.Println(err.Error())
		}
	}

	var result *PostRequest
	if err := json.Unmarshal(data, result); err != nil {
		w.WriteHeader(401)
		if _, err := w.Write([]byte("Bad body!")); err != nil {
			fmt.Println(err.Error())
		}
	}

	lastPost = result
}

func getFunc(w http.ResponseWriter, r *http.Request) {
	if lastPost != nil {
		stuff, err := json.Marshal(lastPost)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.WriteHeader(201)

		if _, err := w.Write(stuff); err != nil {
			fmt.Println(err.Error())
		}
	}
}
