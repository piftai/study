
def solution(volumes):
    max_volume = max(volumes)
    operations = 0
    for i in range(len(volumes) - 1, -1, -1):
        if volumes[i] < max_volume:
            diff = max_volume - volumes[i]
            operations += diff
            # Добавляем diff к первым (i+1) чанам
            for j in range(i + 1):
                volumes[j] += diff
    return operations


n = int(input())
array = input().split()
for i in range(len(array)):
    array[i] = int(array[i])
print(solution(array), flush=True)


