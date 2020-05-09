package main

import (
	"net/http"
	"strings"
	"log"
	"fmt"
)

func handleCookies(w http.ResponseWriter, r *http.Request, id string) {
	cookie := readSetCookie(r,id)
	if cookie != "" {
		lastUpdatedPdaId := cookie
		log.Print("Latest updated Pid...............", lastUpdatedPdaId)
		updateStateInfo(id, lastUpdatedPdaId)
	} else {
		log.Print("No PDA was updated before")
	}

	setCookie(w, id)
}

func setCookie(w http.ResponseWriter, pdaId string) {
	pda := cache[pdaId]
	gid := pda.Gid
	
	c := http.Cookie{
		Name: gid,
		Value: pdaId,
	}
	http.SetCookie(w, &c)
	
}

func readSetCookie(r *http.Request, newPdaId string) string {

	// Loop over header names
	latestPdaId := ""
	for name, values := range r.Header {
		if name == "Cookie" {
			for _, value := range values {
				pairs := strings.Split(value, ";")
				pda := cache[newPdaId]
				gid := pda.Gid

				fmt.Println(pairs)
				for _, pair := range pairs {
					ids := strings.Split(pair, "=")
					ids[0] = strings.TrimSpace(ids[0])
					ids[1] = strings.TrimSpace(ids[1])
					if(gid==ids[0]) {
						latestPdaId = ids[1]
					}
				}
			}
		}
	}
	return latestPdaId
}

