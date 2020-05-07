package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func  handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/pdas", returnAllPdas)
	myRouter.HandleFunc("/pdas/{id}", createPda)
	myRouter.HandleFunc("/pdas/{id}/reset", reset)
	myRouter.HandleFunc("/pdas/{id}/tokens/{position}", put)
	myRouter.HandleFunc("/pdas/{id}/eos/{position}", eos)
	myRouter.HandleFunc("/pdas/{id}/is_accepted", is_accepted)
	myRouter.HandleFunc("/pdas/{id}/stack/top/{k}", peek)
	myRouter.HandleFunc("/pdas/{id}/stack/len", stacklen)
	myRouter.HandleFunc("/pdas/{id}/state", current_state)
	myRouter.HandleFunc("/pdas/{id}/tokens", gettokens)
	myRouter.HandleFunc("/pdas/{id}/snapshot/{k}", snapshot)
	myRouter.HandleFunc("/pdas/{id}/close", close)
	myRouter.HandleFunc("/pdas/{id}/delete", deletePda)

	// myRouter.HandleFunc("/pdas/{id}/join", joinGroup)
	// myRouter.HandleFunc("/pdas/{id}/code", getPdaSpec)
	// myRouter.HandleFunc("/pdas/{id}/c3state", c3state)

	// myRouter.HandleFunc("/replicas_pdas", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}/reset", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}/members", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}/connect", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}/close", returnAllReplicas)
	// myRouter.HandleFunc("/replicas_pdas/{gid}/delete", returnAllReplicas)


	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main(){
	fmt.Println("Server started. Listening at port 8080")

	handleRequest()
}