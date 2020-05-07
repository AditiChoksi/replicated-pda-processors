package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


var replicaCache = make(map[string]PDAProcessor)

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