import concurrent.futures
import requests
import json

USERS_DICT = [
    {"name": "Alice", "age": 25, "gender": "female", "latitude": 28.704060, "longitude": 77.102493},
    {"name": "Bob", "age": 30, "gender": "male", "latitude": 28.704061, "longitude": 77.102494},
    {"name": "Charlie", "age": 35, "gender": "male", "latitude": 28.704062, "longitude": 77.102495},
    {"name": "Diana", "age": 28, "gender": "female", "latitude": 28.704063, "longitude": 77.102496},
    {"name": "Eve", "age": 22, "gender": "female", "latitude": 28.704064, "longitude": 77.102497},
    {"name": "Frank", "age": 33, "gender": "male", "latitude": 28.704065, "longitude": 77.102498},
    {"name": "Grace", "age": 27, "gender": "female", "latitude": 28.704066, "longitude": 77.102499},
    {"name": "Hank", "age": 31, "gender": "male", "latitude": 28.704067, "longitude": 77.102500},
    {"name": "Ivy", "age": 26, "gender": "female", "latitude": 28.704068, "longitude": 77.102501},
    {"name": "Jack", "age": 29, "gender": "male", "latitude": 28.704069, "longitude": 77.102502},
]

def db_insert_users():
    url = 'http://localhost:8080/dating-app/v2/users/create'
    headers = {'Content-Type': 'application/json'}

    def send_post_request(user):
        data = {
            "email": f"{user['name'].lower()}@gmail.com",
            "name": user['name'],
            "password": "password",
            "gender": user['gender'],
            "age": user['age'],
            "latitude": user['latitude'],
            "longitude": user['longitude']
        }

        response = requests.post(url, headers=headers, data=json.dumps(data))
        if response.status_code == 200:
            print(f"User {user['name']} created successfully")
        else:
            print(f"Failed to create user {user['name']} with error :- {response.text}")
            
        return response
    
    for user in USERS_DICT:
        send_post_request(user)
   
if __name__ == '__main__':
    db_insert_users()
        