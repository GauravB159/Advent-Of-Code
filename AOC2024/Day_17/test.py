import math
check = int("1_000_000_000_000_000_000_000_000_000_000_000_010_000_110_100", base=2)
answer = "2413751503435530"
num = 2

def process(A):
    val = ''
    while A > 0:
        val += str((((A%8) ^ 3) ^ 5 ^ int(A / (math.pow(2, ((A%8) ^ 3)))) % 8))
        A //= 8
    return val

def main():
    for A in range(check, check + 10000000 * int("1" + "_000" * num, base=2), int("1" + "_000" * num, base=2)):
        val = process(A)
        if val[:num+1] == answer[:num+1]:
            print(A, ' ', "".join([x + "_" if (len(bin(A)[2:]) - i - 1) % 3 == 0 else x for i, x in enumerate(bin(A)[2:])]), ': ', val, len(val))
            
main()