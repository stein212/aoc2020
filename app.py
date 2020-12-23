import heapq

t = int(input())

for i in range(t):
    ps = [int(num) for num in input().split(" ")]
    n = input()
    drawDowns = []

    peak = ps[0]
    low = ps[0]
    declining = False
    for p in ps:
        if p > peak:
            peak = p
            if declining:
                drawDowns.append(peak - low)
                low = peak
        else:
            if p < peak:
                declining = True
            if p < low:
                low = p
    
    print(drawDowns)