def fib(n):
    if n <= 1:
            return n
    else:
            return fib(n-1) + fib(n-2)

num = 40
for i in range(num):
     print(fib(i))
