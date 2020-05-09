package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
	"math/rand"
	"log"
)


var replicaCache = make(map[string]ReplicaGroup)

func returnAllReplicas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Return all Replicas")
	
	var replicalist[] string
	for _, member := range replicaCache {

		replicalist = append(replicalist, member.Gid)
	}
	json.NewEncoder(w).Encode(replicalist)
}


func createReplica(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Endpoint Hit: Create Replica")
	var rg ReplicaGroup
	
	var vars = mux.Vars(r)
	var gid = vars["gid"]

	err := json.NewDecoder(r.Body).Decode(&rg)

	if err != nil {
		log.Print(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	rg.Gid  = gid

	_, found := replicaCache[gid]

	if !found {

		var specs = rg.Pda_code

		for _, member := range rg.Group_members {
			s := strings.Split(member, "/")
			p:= PDAProcessor {
				Gid: gid,
				Id: s[1],
				Name: specs.Name,
				Input_alphabet: specs.Input_alphabet,
				Stack_alphabet: specs.Stack_alphabet,
				Accepting_states: specs.Accepting_states,
				Start_state: specs.Start_state,
				Transitions: specs.Transitions,
				Eos: specs.Eos,
			}
			
			open(p.Id, p)
			fmt.Println("Pda successfully created or replaced with new definition.")
			
		}

		replicaCache[gid] = rg

		log.Print("Relica Group created.")
		json.NewEncoder(w).Encode("Replica group successfully created.")
	} else 
	{
		json.NewEncoder(w).Encode("Cannot create Replica group. A Replica group with this id already exists.")
	}
}


func returnAllProcessorsInAGroup(w  http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Return all Processors in a group")
	var vars = mux.Vars(r)
	var gid = vars["gid"]
	group := replicaCache[gid]
	json.NewEncoder(w).Encode(group.Group_members)

}

func resetMembers(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Reset all members in a group")
	var vars = mux.Vars(r)
	var gid = vars["gid"]
	group := replicaCache[gid]
	Group_members := group.Group_members
	for procId, _ := range Group_members {
		baseAddress := Group_members[procId]
		s := strings.Split(baseAddress, "/")
		p := cache[s[1]]
		resetInternal(&p)
	}
}

func returnRandomPDAProcessor(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Return random ID of the PDA processor")
	var vars = mux.Vars(r)
	var gid = vars["gid"]
	group := replicaCache[gid]
	Group_members := group.Group_members
	numberOfPDAProcessors := len(Group_members)
	randomNumber := rand.Intn(numberOfPDAProcessors)
	json.NewEncoder(w).Encode(group.Group_members[randomNumber])

}

func joinGroup(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Join PDA to replica group")
	var vars = mux.Vars(r)
	var pid = vars["id"]
	var rg ReplicaGroup
	err := json.NewDecoder(r.Body).Decode(&rg)
	gid := rg.Gid
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	fmt.Println("Join ",pid,"to group ID ",gid)
	_, found := replicaCache[gid]
	if(found){
		group := replicaCache[gid]
		_, pdaExists := cache[pid]
		if(pdaExists) {
			PDAprocessor := cache[pid]
			specs := group.Pda_code
			pdaAdress :=  "pdas/"+pid
			group.Group_members = append(group.Group_members, pdaAdress)
			PDAprocessor.Gid = gid		
			PDAprocessor.Input_alphabet = specs.Input_alphabet
			PDAprocessor.Stack_alphabet = specs.Stack_alphabet
			PDAprocessor.Accepting_states = specs.Accepting_states
			PDAprocessor.Start_state = specs.Start_state
			PDAprocessor.Transitions = specs.Transitions
			PDAprocessor.Eos = specs.Eos

			replicaCache[gid] = group
			cache[pid] = PDAprocessor
		} else {
			json.NewEncoder(w).Encode("The PDA with specified id does'nt exist. First create the PDA then try joining again.")
		}
		
	} else{
		json.NewEncoder(w).Encode("Cannot join processor to Group. A Replica group with this id does not exist.")
	}
}

func deleteReplicaGroup(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Delete a replica group and all its members")
	var vars = mux.Vars(r)
	var gid = vars["gid"]
	_, found := replicaCache[gid]

	if found {
		delete(replicaCache, gid)
		json.NewEncoder(w).Encode("Replica Group deleted.")
	} else {
		json.NewEncoder(w).Encode("Replica group found.")

	}
}