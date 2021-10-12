def solution(s):
    strList = list(map(str, s.split(' ')))
    
    answer = ''
    
    for st in strList:
        for i in range(len(st)):
            if i % 2 == 0: answer += st[i].upper()
            else: answer += st[i].lower()
        answer += ' '
        
    return answer[:-1]
