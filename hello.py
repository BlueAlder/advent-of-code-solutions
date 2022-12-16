import copy

def a(b):
  g = copy.copy(b)
  g.append("d")

c = ["a"]
print(c)
a(c)
print(c)