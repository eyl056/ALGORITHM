def solution(arr):
    if len(arr) < 2: return [-1]
    arr.remove(min(arr))
    return arr
