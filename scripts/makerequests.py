import requests
import time
from random import randrange

while True:
    a = randrange(1, 21)
    b = randrange(1, 21)
    add = requests.get("http://arithmeticgrpc.com/add/{}/{}".format(a, b))
    mult = requests.get("http://arithmeticgrpc.com/mult/{}/{}".format(a, b))
    print(add.content)
    print(mult.content)
    time.sleep(0.5)