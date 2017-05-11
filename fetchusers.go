package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"encoding/json"
)

type FetchUserResponse struct {
	UserId            string `json:"userId"`
	CombinationLength int `json:"CombinationLength"`
}

func (v Victim) createUserResponse() FetchUserResponse {
	response := FetchUserResponse{}
	response.UserId = v.UserId
	response.CombinationLength = v.CombinationLength
	return response
}

func fetchUsers(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(db_addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var results []Victim

	collection := session.DB(db).C(collection)
	err = collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}

	var fetchUserResponses []FetchUserResponse

	for _, result := range results {
		fetchUserResponses = append(fetchUserResponses, result.createUserResponse())
	}

	json.NewEncoder(w).Encode(fetchUserResponses)
}
