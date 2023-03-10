/*
Create time at 2023年3月10日0010下午 14:27:32
Create User at Administrator
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Skysolderone/MongoDbRestfulAPI-Golang/pkg/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserControllers struct {
	Session *mgo.Session
}

func (c UserControllers) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := c.Session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (c *UserControllers) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	c.Session.DB("mogo-golang").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (c *UserControllers) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	if err := c.Session.DB("mogo-golang").C("users").Remove(oid); err != nil {
		w.WriteHeader(404)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Delete user ", oid, "\n")
}

func UserController(s *mgo.Session) *UserControllers {
	return &UserControllers{s}
}
