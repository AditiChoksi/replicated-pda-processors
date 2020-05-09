printf "\n---------------- Create Replica with id 1001 -----------------\n" 
curl -X PUT -H "Content-Type: application/json" -d '{
    "group_members": [
        "pdas/100",
        "pdas/101",
        "pdas/102"
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

printf "\n---------------- Create Replica with id 2001 -----------------\n" 
curl -X PUT -H "Content-Type: application/json" -d '{
    "group_members": [
        "pdas/200",
        "pdas/201",
        "pdas/202"
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
}' http://localhost:8080/replicas_pdas/2001

curl -c cookies.txt -X PUT -H "Content-Type: application/json" -d '{"token": "0"}' http://localhost:8080/pdas/101/tokens/0

curl -b "client-id=101" -c cookies1.txt -X PUT -H "Content-Type: application/json" -d '{"token": "0"}' http://localhost:8080/pdas/102/tokens/1

curl -b "client-id=102" -c cookies3.txt -X PUT -H "Content-Type: application/json" -d '{"token": "1"}' http://localhost:8080/pdas/100/tokens/2

curl -b "client-id=100" -c cookies3.txt -X PUT -H "Content-Type: application/json" -d '{"token": "1"}' http://localhost:8080/pdas/102/tokens/3


curl -b "client-id=102" -c cookies2.txt -X PUT -H "Content-Type: application/json" -d '{"token": "0"}' http://localhost:8080/pdas/202/tokens/0



curl -b "client-id=202" -c cookies2.txt -X PUT -H "Content-Type: application/json" -d '{"token": "0"}' http://localhost:8080/pdas/200/tokens/1

curl -b "client-id=200" -c cookies2.txt -X PUT -H "Content-Type: application/json" -d '{"token": "1"}' http://localhost:8080/pdas/201/tokens/2

curl -b "client-id=201" -c cookies2.txt -X PUT -H "Content-Type: application/json" -d '{"token": "1"}' http://localhost:8080/pdas/202/tokens/3




