def solution(chans):
    for i in range(len(chans) - 1):
        if chans[i] > chans[i + 1]:
            return -1
    if len(chans) == 2:
        return chans[1] - chans[0]
    maxVolume = chans[-1]
    cnt = 0
    for i in range(len(chans)-1, 0, -1):  # 0 не включается, следовательно он не смотрится
        if chans[i] == maxVolume:
            pass
        else:
            cnt += maxVolume - chans[i]
            chans[i] += maxVolume - chans[i]
            for j in range(0, len(chans), 1):
                chans[j] += maxVolume - chans[i]
    return cnt


n = int(input())
chans = input().split()

for i in range(len(chans)):
    chans[i] = int(chans[i])

print(solution(chans))


