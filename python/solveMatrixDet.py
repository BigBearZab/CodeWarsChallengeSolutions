# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""
m1 = [ [1, 3], [2,5]]
m2 = [ [2,5,3], [1,-2,-1], [1, 3, 4]]

# solve the determinant of a 2x2 matrix
def solve_det2(mat2:list):
    det = mat2[0][0] * mat2[1][1] - mat2[0][1] * mat2[1][0]
    return det

# now create larger solution
def solve_det(matrix:list):
    if len(matrix) == 1:
        return matrix[[0]]
    if len(matrix) == 2:
        return solve_det2(matrix)
    if len(matrix)==3:
        return 
        
# print(solve_det([1]))
    
#%%

for i in range(3):
    print(solve_det2([m2[1][:i] + m2[1][(i+1):],
           m2[2][:i] + m2[2][(i+1):]]))