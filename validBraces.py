# -*- coding: utf-8 -*-
"""
Created on Sat Feb 29 23:11:38 2020

@author: denis
"""

# =============================================================================
# Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.
# 
# This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!
# 
# All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.
# 
# What is considered Valid?
# A string of braces is considered valid if all braces are matched with the correct brace.
# =============================================================================

def validBraces(string):
  for i in range(len(string)):
      if string[i] == "(":
          temp = string[i:]
          for j in range(len(temp)):
              if temp[j] == ")":
                  string = string.replace(string[i],'0')
                  string = string.replace(string[i+j],'0',1)
                  break
      if string[i] == "[":
          temp = string[i:]
          for j in range(len(temp)):
              if temp[j] == "]":
                  string = string.replace(string[i],'0')
                  string = string.replace(string[i+j],'0',1)
                  break
      if string[i] == "{":
          temp = string[i:]
          for j in range(len(temp)):
              if temp[j] == "}":
                  string = string.replace(string[i],'0')
                  string = string.replace(string[i+j],'0',1)
                  break  
#  string = string.replace("0",'')    
  return string
              
                
# =============================================================================
#       elif string[i] == "[": 
#       
#       elif string[i] == "{":
# =============================================================================
          
      