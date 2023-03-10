/*
Create time at 2023年3月10日0010下午 14:26:42
Create User at Administrator
*/

package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
