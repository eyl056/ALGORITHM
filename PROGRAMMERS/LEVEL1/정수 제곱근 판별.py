def solution(n):
    p = n ** 0.5
    if p == int(p):
        return (p + 1) ** 2
    else: return -1
