import grammar
import threading
import time
import itertools

MAX = 10000000
VALUE = 1000000


def decorator(function):
    def wrapTheFunction(*args, **kwargs):
        t1 = time.time()
        function(*args, **kwargs)
        t2 = time.time()
        print("➡➡➡ Using Time: ", t2 - t1, "\n")
    return wrapTheFunction


@decorator
def fun1(n):
    content = ""
    for x in range(1, VALUE):
        content += " ".join(grammar.sentence())
    f = open('data2.txt', 'a+')
    f.write(content)
    f.close()


@decorator
def fun2(n):
    content = []
    for x in range(1, VALUE):
        content.append(" ".join(grammar.sentence()))
    result = "".join(content)
    f = open('data2.txt', 'a+')
    f.write(result)
    f.close()


@decorator
def fun3(n):
    content = []
    for x in range(1, VALUE):
        content.extend(grammar.sentence())
    result = ' '.join(content)
    f = open('data2.txt', 'a+')
    f.write(result)
    f.close()

# @decorator
# def fun4(n):
#     content = []
#     for x in range(1,VALUE):
#         lista=[]
#         lista.extend(grammar.sentence())
#         lista.append("\n")
#         content=[*content,*lista]
#     result=''.join(content)
#     f = open('data2.txt', 'a+')
#     f.write(result)
#     f.close()


@decorator
def fun5(n):
    content = []
    for x in range(1, VALUE):
        lista = []
        lista.extend(grammar.sentence())
        content = list(itertools.chain(content, lista))
    result = ' '.join(content)
    f = open('data2.txt', 'a+')
    f.write(result)
    f.close()


fun1(MAX)
fun2(MAX)
fun3(MAX)
# fun5(MAX)
