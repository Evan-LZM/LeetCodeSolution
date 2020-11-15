import grammar
import threading
import time

e1=time.time() 
content=[]

with open('data1.txt' ,'w') as f:
    for i in range (1000000):
        content.append(" ".join(grammar.sentence()))

    for line in content:
        f.write(line+"\n")

f.close()
e2=time.time() 
print(float(e2-e1))