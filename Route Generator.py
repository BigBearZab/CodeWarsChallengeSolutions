# -*- coding: utf-8 -*-
"""
Created on Mon Mar  2 18:00:37 2020

@author: denis
"""
from random import seed
from random import randint

rl = [[1,"Kinteton"], [2,"Southam"], [3,"Southam Reverse"], [4,"Rugby"], [5,"Rugby Reverse"]]

upcoming = []

seed(2)

for ride in range (20):
    value = randint(1,5)
    upcoming.append(value)
    
results = []

for value in upcoming:
    results.append(rl[value-1][1])
    
for i in range(len(results)-1):
    if results[i] == results[i+1]:
        results[i] = 0
#for x in range(len(results)-1):
#    if results[x] == 0:
#        del results[x]        
#        
    
print(results)    