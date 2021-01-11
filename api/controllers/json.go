package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Kabupatens struct {
	WilayahID string `json:"wilayah_id"`
	Name      string `json:"nama"`
	Parent    string `json:"parent"`
}

type OutKabupatens struct {
	WilayahID string `json:"wilayah_id"`
	Name      string `json:"nama"`
	Parent    int    `json:"parent"`
	Tingkat   int    `json:"tingkat"`
	Singkatan string `json:"singkatan"`
}

type GetInputJsonParam struct {
	Id string `json:"id" form:"id" query:"id"`
}

func getKabupatens() []Kabupatens {
	kabs := make([]Kabupatens, 3)
	raw, err := ioutil.ReadFile("assets/kabupaten.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &kabs)
	return kabs
}

func findById(id string, messages chan OutKabupatens) {
	Kabupatens := getKabupatens()
	var filtered OutKabupatens
	filtered.WilayahID = id

	for _, kab := range Kabupatens {
		if kab.WilayahID == id {
			var parent, _ = strconv.Atoi(kab.Parent)
			filtered.WilayahID = kab.WilayahID
			filtered.Name = kab.Name
			filtered.Parent = parent
			filtered.Tingkat = 4
			filtered.Singkatan = "Singkatan"
		}

	}

	messages <- filtered

}

func (server *Server) Json(c echo.Context) error {
	newKab := []OutKabupatens{}
	Kabupatens := getKabupatens()
	var filtered OutKabupatens

	for _, kab := range Kabupatens {
		if kab.Parent == "1" {
			var parent, _ = strconv.Atoi(kab.Parent)
			filtered.WilayahID = kab.WilayahID
			filtered.Name = kab.Name
			filtered.Parent = parent
			filtered.Tingkat = 4
			filtered.Singkatan = "Singkatan"

			newKab = append(newKab, filtered)
		}

	}

	return c.JSON(http.StatusOK, newKab)
}

func (server *Server) Find(c echo.Context) error {

	var messages = make(chan OutKabupatens)

	body := new(GetInputJsonParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ids := strings.Split(body.Id, ",")

	// ids := []string{
	// 	"6518",
	// 	"3205",
	// }

	for _, id := range ids {
		go findById(id, messages)
	}

	response := make([]OutKabupatens, len(ids))
	for i, _ := range response {
		response[i] = <-messages
	}

	return c.JSON(http.StatusOK, response)
}
