import grammar
import threading
import time
 
global_num=0
def writetoTxt(txtFile):
    global global_num
    id = threading.currentThread().getName()
    for i in range (1000000):
        global_num+=1
        f=open("data.txt",'a')
        f.write(" ".join(grammar.sentence()))
        f.write('\n')
        f.close()

for i in range(200):
    myThread = threading.Thread(target=writetoTxt, args=("data.txt",))
    myThread.start()