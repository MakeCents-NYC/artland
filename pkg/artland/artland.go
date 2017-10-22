package artland

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Artland struct {
	m *MasterCard
}

func NewArtLand() *Artland {
	return &Artland{}
}

func (al *Artland) Serve() error {
	http.HandleFunc("/post", postFunc)

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

}
