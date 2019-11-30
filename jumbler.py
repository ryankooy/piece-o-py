import sys
import random

def new_name(word):
  l = list(word)
  random.shuffle(l)
  new = ""
  for x in l:
    new += x
  return new.capitalize()

if __name__ == "__main__":
  name = sys.argv[1]
  print(new_name(name))