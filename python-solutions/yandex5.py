def compare_ters(coord1, coord2):
    coord1 = coord1.split()
    coord2 = coord2.split()
    if (int(coord1[2]) > int(coord2[0])
            and int(coord1[3]) > int(coord2[1])
            and int(coord1[0]) != int(coord2[2])
            and int(coord1[1]) != int(coord2[3])):
        return True
    return False


n = int(input())
ters = []
for i in range(n):
    s = input()
    ters.append(s)
output = []
cnt = 0
for ter in ters:
    cnt_ter = 0
    for anotherTer in ters:
        if compare_ters(ter, anotherTer) and ter != anotherTer:
            cnt_ter += 1
            print("compare ", ter, " ", anotherTer)
        else:
            print("not compare", ter, " ", anotherTer)
    output.append(cnt_ter)
print(output)

#  сделать еще один-два теста