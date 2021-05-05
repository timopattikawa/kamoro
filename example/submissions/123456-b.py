T = int(input())
 
while (T > 0):
    T -= 1
    r, b, d = map(int, input().split())
    avail = min([r, b]) * (d + 1)
    if (avail >= max([r, b])):
        print("YES")
    else:
        print("NO")

# print("Anjing")