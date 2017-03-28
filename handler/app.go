package handler

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/mccjul/Graph-AutoDoc-Server/model"
)

//GetApps associated with maintainer
func (h *Handler) GetApps(c echo.Context) (err error) {
	apps := []*model.App{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("autodoc").C("apps").
		Find(bson.M{"maintainer": c.Param("id")}).
		All(&apps); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, apps)
}

// CreateApp for maintainer
func (h *Handler) CreateApp(c echo.Context) (err error) {
	token, err := newUUID()
	if err != nil {
		return err
	}

	a := &model.App{
		ID:    bson.NewObjectId(),
		Token: token,
		Name:  "DefaultName",
	}

	if err = c.Bind(a); err != nil {
		return err
	}

	if a.Name == "" || a.Maintainer == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid name or maintainer fields"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("autodoc").C("apps").Insert(a); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a)
}

//PatchApp for maintainer
func (h *Handler) PatchApp(c echo.Context) (err error) {
	a := &model.App{}

	if err = c.Bind(a); err != nil {
		return err
	}

	if a.Name == "" || a.Token == "" || a.Maintainer == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid fields"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("autodoc").C("apps").Update(bson.M{"ID": c.Param("id")}, a); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a)
}

//RemoveApp for maintainer
func (h *Handler) RemoveApp(c echo.Context) (err error) {
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("autodoc").C("apps").Remove(bson.M{"ID": c.Param("id")}); err != nil {
		return err
	}

	// Remove the Doc

	return c.NoContent(http.StatusOK)
}
