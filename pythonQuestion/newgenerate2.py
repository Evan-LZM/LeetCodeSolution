import grammar
import threading
import time

MAX=10000000

def decorator(function):
    def wrapTheFunction(*args, **kwargs):
        t1 = time.time()
        function(*args, **kwargs)
        t2 = time.time()
        print("➡➡➡ Using Time: ", t2 - t1, "\n")
    return wrapTheFunction

@decorator
def fun1(n):
    for i in range (100):
        content = ""
        for x in range(1,1000000):
            content+=" ".join(grammar.sentence())+"\n"

        f = open('data2.txt', 'a+')
        f.write(content)
        f.close()
fun1(MAX)