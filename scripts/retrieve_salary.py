#! /usr/bin/python3

import requests

for i in range(5000):
    response = requests.get(f"http://localhost:8080/salary/{i*1000}")
    response.raise_for_status()

    data = response.text
    print(data)