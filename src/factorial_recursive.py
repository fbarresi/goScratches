import time

def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

start = time.time()
print(factorial(500))
print(time.time()-start)