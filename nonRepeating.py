# -*- coding: utf-8 -*-
"""
Created on Mon Mar  2 19:00:00 2020

@author: denis
"""

def first_non_repeating_letter(string):
    l = list({let for word in string for let in word})
    
    for x in range(len(l)):
        l[x] = [string.count(l[x]), l[x], string.find(l[x])]
    fl = []
    for run in l:
        if run[0] == 1:
            fl.append(run)
    fl.sort(key = lambda x: x[2])       
    
    print(fl)
    
    
    
    
    
string = "shshsihhhmah"
first_non_repeating_letter(string)