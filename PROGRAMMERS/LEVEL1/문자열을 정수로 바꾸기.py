def solution(s):
    answer = 0
    for ss in s:
        if ss == '+':
            answer = int(s[1:])
        elif ss == '-':
            answer = int(s[1:]) * -1
        else:
            answer = int(s)
    return answer

def solution(s):
    return int(s)
