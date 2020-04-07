# -*- coding: utf-8 -*-
"""
Created on Tue Mar  3 20:21:00 2020

@author: denis
"""

def format_duration(seconds):
    ns = [[31536000,'year'], [86400,'day'], [3600,'hour'], [60,'minute'], [1,'second']]
    l = []
    s = seconds
#    rem = seconds
    for n in ns:
        rem = s % n[0]
#        print("rem = " + str(rem))
        s = s - rem
#        print("s = " + str(s))
        v =int( s / n[0])
        if v == 1:
            l.append([v,n[1]])
        else:
            l.append([v,n[1]+'s'])
#        print(l)
        s = rem
    tl = []
    for ent in l:
        if ent[0] > 0:
            tl.append(ent)
            print(tl)
    
    if len(tl) == 1:
        out = (f'{tl[0][0]} {tl[0][1]}')
    elif len(tl) == 2:
        out = (f'{tl[0][0]} {tl[0][1]} and {tl[1][0]} {tl[1][1]}')
    elif len(tl) == 3:
        out = (f'{tl[0][0]} {tl[0][1]}, {tl[1][0]} {tl[1][1]} and {tl[2][0]} {tl[2][1]}')
    elif len(tl) == 4:
        out = (f'{tl[0][0]} {tl[0][1]}, {tl[1][0]} {tl[1][1]}, {tl[2][0]} {tl[2][1]} and {tl[3][0]} {tl[3][1]}')
    elif len(tl) == 5:
        out = (f'{tl[0][0]} {tl[0][1]}, {tl[1][0]} {tl[1][1]}, {tl[2][0]} {tl[2][1]}, {tl[3][0]} {tl[3][1]} and {tl[4][0]} {tl[4][1]}')
    else:
        out = "now"
    print(out)
    
seconds = 5620
format_duration(seconds)