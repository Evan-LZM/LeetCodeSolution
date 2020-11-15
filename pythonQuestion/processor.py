import multiprocessing
import grammar
import time

def setcallback(x):
    with open('data.txt', 'a+') as f:
        line = " ".join(grammar.sentence()) + "\n"
        f.write(line)

def multiplication(num):
    return num*(num+1)

if __name__ == '__main__':
    e1 = time.time()
    pool = multiprocessing.Pool(6)
    # for i in range(100000000):
    for i in range(1000000):
        pool.apply_async(func=multiplication, args=(i,), callback=setcallback)
    pool.close()
    pool.join()
    e2=time.time()
    print(float(e1-e2))