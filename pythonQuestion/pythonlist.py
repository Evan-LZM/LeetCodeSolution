
import time
import itertools

def decorator(function):
    def wrapTheFunction(*args, **kwargs):
        t1 = time.time()
        function(*args, **kwargs)
        t2 = time.time()
        print("➡➡➡ Using Time: ", t2 - t1, "\n")
    return wrapTheFunction

@decorator
def func1():
    list1=[1,2,3]
    for i in range(1,100):
        list2=[4,5,6]
        list1.extend(list2)

@decorator
def func2():
    list1=[1,2,3]
    for i in range(1,100):
        list2=[4,5,6]
        list1=list1+list2

@decorator
def func3():
    list1=['1','2','3']
    result="".join(list1)
    list1.append("123")

@decorator
def func4():
    list1=['1','2','3']
    list2=['1','2','3']
    list1.append("\n")
    list1.extend(list2)
    print(list1)
    list1.append("\n")
    print(list1)
    result=[*list1,*list2]
    print("result")
    print(result)

@decorator
def func5():
    list1=['1','2','3']
    list2=['1','2','3']
    list3=['1','2','3']
    result = list(itertools.chain(list2, list2, list3))
    print(result)



# func1()
# func2()
# func3()
# func4()
func5()