printf "\n---------------- Create Replica with id 1001 -----------------\n" 
curl -X PUT -H "Content-Type: application/json" -d '{
    "group_members": [
        "pdas/100",
        "pdas/101"
    ],
    "pda_code": {
        "name": "0n1n",
        "states": ["q1", "q2", "q3", "q4"],
        "input_alphabet": [ "0", "1" ],
        "stack_alphabet" : [ "0", "1" ],
        "accepting_states": ["q1", "q4"],
        "start_state": "q1",
        "transitions": [
            ["q1", "null", "null", "q2", "$"],
            ["q2", "0", "null", "q2", "0"],
            ["q2", "0", "0", "q2", "0"],
            ["q2", "1", "0", "q3", "null"],
            ["q3", "1", "0", "q3", "null"],
            ["q3", "null", "$", "q4", "null"]
        ],
        "eos": "$"
    }
}' http://localhost:8080/replicas_pdas/1001


curl -X GET http://localhost:8080/replicas_pdas
