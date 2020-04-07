# -*- coding: utf-8 -*-
"""
Created on Sun Mar  1 22:20:32 2020

@author: denis
"""

def pig_it(text):
    text = text.split()
    for i in range(len(text)):
        if '!' in text[i] or '?' in text[i]:
            text[i] = text[i]
        else:
            end = text[i][0]
            print(end)
            word = text[i][1:] + end + "ay"
            print(word)
            text[i] = word
    text = ' '.join(text)
    return(text)