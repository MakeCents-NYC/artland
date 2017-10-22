package artland

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MakeCents-NYC/artland/pkg/mastercard"
)

var lastPost *PostRequest

type Artland struct {
	m *mastercard.MasterCard
}

var a *Artland

func NewArtLand() *Artland {
	_m, err := mastercard.NewMasterCardAPI("/keystore.p12", "protos/artland.proto", "Nfo8owSgU8WmSOYuAmZmMWWY8GDTWKIf78mKGkYv9690e4ab!9bd5a371243f403cb2b455a2aedd58930000000000000000", "keyalias", "keystorepass")
	if err != nil {
		return nil
	}

	a.m = _m

	return a
}

func (al *Artland) Serve() error {
	http.HandleFunc("/api/post", postFunc)
	http.HandleFunc("/api/get", getFunc)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	return nil
}

type PostRequest struct {
	Id  string
	Img string
}

func postFunc(w http.ResponseWriter, r *http.Request) {
	var result PostRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&result); err != nil {
		w.WriteHeader(300)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println(err.Error())
		}
		var stuff []byte
		if _, err := r.Body.Read(stuff); err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				fmt.Println(err.Error())
			}
		}
		if _, err := w.Write(stuff); err != nil {
			fmt.Println(err.Error())
		}
	}

	defer r.Body.Close()

	a.m.PostTransaction(PostRequest)

	lastPost = &result
}

func getFunc(w http.ResponseWriter, r *http.Request) {
	if lastPost == nil {
		lastPost = &PostRequest{
			Id:  "hello",
			Img: "world",
		}
	}

	if lastPost != nil {
		encode := json.NewEncoder(w)
		err := encode.Encode(lastPost)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.WriteHeader(200)
	}
}
