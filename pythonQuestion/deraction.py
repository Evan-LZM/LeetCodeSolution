###编写函数：求一千万以内的是3的倍数或者5的倍数的所有数的和。

###–要求：通过装饰器实现该函数运行时间的计算。

###装饰器示例####

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
    sum=0
    for i in range ( 1 , 10000000 ):
     if i % 3 == 0 or i % 5 == 0 :
         sum += i
    print(sum)
    return sum

def gaa_w(n):
    return (n+1)*n/2

@decorator
def fun2(n):
    sum=0
    sum=3*gaa_w((n - 1) / 3) + 5 * gaa_w((n - 1) / 5) - 15 * gaa_w((n - 1) / 15)
    print(sum)
    return sum


fun1(MAX)
fun2(MAX)