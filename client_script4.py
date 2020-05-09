import requests

d = {}

def updateCookies(cookieString):
    cookies = cookieString.split('=')
    d[cookies[0]] = cookies[1]

    return d


print("Create Replica...")
response =  requests.put('http://localhost:8080/replica_pdas/1001', \
    data='{\
    "group_members": [\
        "pdas/100",\
        "pdas/101",\
        "pdas/102"\
    ],\
    "pda_code": {\
        "name": "0n1n",\
        "states": ["q1", "q2", "q3", "q4"],\
        "input_alphabet": [ "0", "1" ],\
        "stack_alphabet" : [ "0", "1" ],\
        "accepting_states": ["q1", "q4"],\
        "start_state": "q1",\
        "transitions": [\
            ["q1", "null", "null", "q2", "$"],\
            ["q2", "0", "null", "q2", "0"],\
            ["q2", "0", "0", "q2", "0"],\
            ["q2", "1", "0", "q3", "null"],\
            ["q3", "1", "0", "q3", "null"],\
            ["q3", "null", "$", "q4", "null"]\
        ],\
        "eos": "$"\
    }\
}')

print(response.json())
print('---------------------------')

print("Get Replicas...")
response =  requests.get('http://localhost:8080/replica_pdas')
print(response.json())
print('---------------------------')

print("Get Members...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/members')
print(response.json())
print('---------------------------')

id = response.json()[0]
id = id.split('/')[1]
print("Get Code for PDA " + id + " ...")
url = 'http://localhost:8080/pdas/' + id + '/code'
response =  requests.get(url)
print(response.json())
print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())

id = response.json()
id = id.split('/')[1]
print("Sending put request to PDA " + id)
response =  requests.put('http://localhost:8080/pdas/' + id + '/tokens/1', data='{"token": "0"}')
d = updateCookies(response.headers['Set-Cookie'])
print(response.json())

print("Printing current state information")
response =  requests.get('http://localhost:8080/pdas/' + id + '/snapshot/5', cookies=d)
print(response.json())
print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())

id = response.json()
id = id.split('/')[1]
print("Sending put request to PDA " + id)
response =  requests.put('http://localhost:8080/pdas/' + id + '/tokens/2', data='{"token": "1"}', cookies=d)
d = updateCookies(response.headers['Set-Cookie'])
print(response.json())

print("Printing current state information")
response =  requests.get('http://localhost:8080/pdas/' + id + '/snapshot/5', cookies=d)
print(response.json())
print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())

id = response.json()
id = id.split('/')[1]
print("Sending put request to PDA " + id)
response =  requests.put('http://localhost:8080/pdas/' + id + '/tokens/3', data='{"token": "1"}', cookies=d)
d = updateCookies(response.headers['Set-Cookie'])
print(response.json())

print("Printing current state information")
response =  requests.get('http://localhost:8080/pdas/' + id + '/snapshot/5', cookies=d)
print(response.json())

print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())

id = response.json()
id = id.split('/')[1]
print("Sending put request to PDA " + id)
response =  requests.put('http://localhost:8080/pdas/' + id + '/tokens/0', data='{"token": "0"}', cookies=d)
d = updateCookies(response.headers['Set-Cookie'])
print(response.json())


print("Printing current state information")
response =  requests.get('http://localhost:8080/pdas/' + id + '/snapshot/5', cookies=d)
print(response.json())

print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())

id = response.json()
id = id.split('/')[1]
print("Caling eos...")
response = requests.put('http://localhost:8080/pdas/' + id + '/eos/4', cookies=d)
d = updateCookies(response.headers['Set-Cookie'])
print(response)

print("Printing current state information")
response =  requests.get('http://localhost:8080/pdas/' + id + '/snapshot/5', cookies=d)
print(response.json())
print('---------------------------')

print("Connecting...")
response =  requests.get('http://localhost:8080/replica_pdas/1001/connect')
print("Connected to " + response.json())
print('---------------------------')

id = response.json()
id = id.split('/')[1]
print("Caling is_accepted...")
response = requests.get('http://localhost:8080/pdas/' + id + '/is_accepted', cookies=d)
print(response.json())
print('---------------------------')

print("Deleting the replica group")
response = requests.delete('http://localhost:8080/replica_pdas/1001/delete')
print(response.json())