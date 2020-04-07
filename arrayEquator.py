def find_even_index(arr):
    c = -1
    for i in range(len(arr)):
        a = sum(arr[:i])
        b = sum(arr[(i+1):])
        if a == b:
            c = i
            break
    return c
# =============================================================================
#     if c > -1:
#         return c
#     else:
#         c = -1
#         return c
# =============================================================================
