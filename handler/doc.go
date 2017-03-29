package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/mccjul/Graph-AutoDoc-Server/model"
)

//GetDoc from id
func (h *Handler) GetDoc(c echo.Context) (err error) {
	docs := []*model.Doc{}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("autodoc").C("docs").
		Find(bson.M{"id": bson.ObjectIdHex(c.Param("id")).String()}).
		All(&docs); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, docs)
}

//PatchDoc with id
func (h *Handler) PatchDoc(c echo.Context) (err error) {
	file, e := ioutil.ReadFile("../file.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	db := h.DB.Clone()
	defer db.Close()
	if _, err = db.DB("autodoc").C("docs").RemoveAll(bson.M{"ID": c.Param("id")}); err != nil {
		return
	}

	//Insert for loop
	docs := []model.Doc{}
	if err = json.Unmarshal(file, &docs); err != nil {
		return
	}
	for i := 0; i < len(docs); i++ {
		if err = db.DB("autodoc").C("docs").Insert(bson.M{"ID": c.Param("id")}, docs[i]); err != nil {
			return
		}
	}

	return c.NoContent(http.StatusOK)
}
