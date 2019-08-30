def bitonic_sort(up, l):
    if len(l) == 1:
        return l
    else:
        first = bitonic_sort(True, l[:len(l)//2])
        second = bitonic_sort(False, l[len(l)//2:])
        return bitonic_merge(up, first+second)

def bitonic_merge(up, l):
    if len(l) == 1:
        return l
    else:
        bitonic_compare(up, l)
        first = bitonic_merge(up, l[:len(l)//2])
        second = bitonic_merge(up, l[len(l)//2:])
        return first + second

def bitonic_merge(up, l):
    a = len(l)//2
    for i in range(a):
        if (l[i] > l[i+a]) == up:
            l[i], l[i+a] = l[i+a], l[i]
