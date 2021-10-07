def solution(n):
    arr = list(str(n))
    arr.sort()
    arr.reverse()
    return int(''.join(arr))
