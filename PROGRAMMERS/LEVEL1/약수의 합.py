def solution1(n):
    answer = 0
    
    for i in range(1, int(n ** 0.5) + 1):
        if (n % i == 0):
            answer += i
            if (n // i != i):
                answer += (n // i)
            
    return answer

def solution2(n):
    return sum([i for i in range(1, (n // 2) + 1) if n % i == 0]) + n
