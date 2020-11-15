import random
import string
import numpy as np
import sys

def gen_pass(n):
    num=random.randint(0,n-1)
    punc=random.randint(0,num) or random.randint(num,n-1)
    if num==0 and punc==num:
        punc+=1
    elif num==n-1 and punc==num:
        punc+=1
    elif punc==num:
        punc+=1
    result=""
    for i in range (n):
        if i==num:
            ch=random.choice(string.digits)
        elif i==punc:
            ch=random.choice(string.punctuation)
        elif i!=num and i!=punc:
            ch=random.choice(string.ascii_letters)
        result+=ch
    return result

def generatePassword(passlen,*users):
    nums=len(users)
    if nums==0:
        nums=1
    for i in range(nums):
        if len(users)!=0:
            print("Users_Name:",users[i])
        print("Generated_Password:",gen_pass(passlen))

print("Password_len:",9)
generatePassword(9,'Adam','Evan','Alison','Thomas','Lin')
print("Password_len:",2)
generatePassword(12,'Reid','Scotte','Tom')
print("Password_len:",5)
generatePassword(5,'Andy')
print("Password_len:",5)
generatePassword(5)