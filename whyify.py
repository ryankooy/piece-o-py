import sys

def whyify(cat):
  why = "y"
  add_why = cat + why
  return add_why

def print_cat(c):
  print(whyify(c))

if __name__ == "__main__":
  this_cat = str(sys.argv[1])
  print_cat(this_cat)