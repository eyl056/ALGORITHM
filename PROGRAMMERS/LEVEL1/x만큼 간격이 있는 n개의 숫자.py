import sys
input = sys.stdin.readline

def solution(x, n):
    answer = []
    for i in range(1, n + 1):
        answer.append(i * x)
    return answer


if __name__ == '__main__':
    x, n = map(int, input().strip().split(' '))
    solution(x, n)
