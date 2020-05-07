package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


var replicaCache = make(map[string]ReplicaGroup)

func returnAllReplicas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Return all Replicas")
	
	var replicalist[]ReplicaGroup
	for key, _ := range replicaCache {
		info := ReplicaGroup {
			GroupId: key,
		}
		replicalist = append(replicalist, info)
	}
	json.NewEncoder(w).Encode(replicalist)
}


func createReplica(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Endpoint Hit: Create Replica")
	var rg ReplicaGroup
	
	var vars = mux.Vars(r)
	var gid = vars["gid"]

	err := json.NewDecoder(r.Body).Decode(&rg)
	fmt.Println("rg is: ", rg)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	var replicaGroup = ReplicaGroup {
		GroupId: gid,
	}
	fmt.Println("The replica group")
	fmt.Println(replicaGroup)
	// created := open(id, p)

	// if created {
	// 	json.NewEncoder(w).Encode("PDA successfully created.")
	// } else 
	// {
	// 	json.NewEncoder(w).Encode("Cannot create PDA. A PDA with this id already exists.")
	// }
}