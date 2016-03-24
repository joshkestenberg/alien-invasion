import requests
import json

url = 'https://api.quadrigacx.com/v2/ticker'
r = requests.get(url,"last")
print r.json()

