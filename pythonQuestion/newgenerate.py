import grammar
import threading
import time

e1=time.time() 
f=open("data.txt",'a')

for i in range (1000000):  
    f.write(" ".join(grammar.sentence()))
    f.write('\n')

f.close()
e2=time.time()
print(float(e2-e1))