
# Heaps algorithm method is too slow
#  Generating permutation using Heap Algorithm
def heapPermutation(a, size, ls):
    # a = [int(char) for char in str(a)]
    # if size becomes 1 then prints the obtained
    # permutation
    if size == 1:
        ls.append(list_nums_to_int(a))
        return
 
    for i in range(size):
        heapPermutation(a, size-1, ls)
 
        # if size is odd, swap 0th i.e (first)
        # and (size-1)th i.e (last) element
        # else If size is even, swap ith
        # and (size-1)th i.e (last) element
        if size & 1:
            a[0], a[size-1] = a[size-1], a[0]
        else:
            a[i], a[size-1] = a[size-1], a[i]

    return ls
 
 
def list_nums_to_int(list_nums):
    return int(''.join([str(num) for num in list_nums]))

def next_bigger(start_num):
    # Driver code
    a = [int(num) for num in str(start_num)]
    n = len(a)
    ls = []
    heapPermutation(a,n, ls)
    larger_list = [num for num in ls if num > start_num]
    if len(larger_list) == 0:
        return -1
    else:
        return min(larger_list)

#######################################################################################################

# Method 2

class Test_Next_Bigger:
    def test_1(self):
        assert next_bigger(12) == 21
    def test_2(self):
        assert next_bigger(1) == -1
    def test_3(self):
        assert next_bigger(513) == 531
    def test_4(self):
        assert next_bigger(2017) == 2071
    def test_5(self):
        assert next_bigger(414) == 441
    def test_6(self):
        assert next_bigger(441) == -1
