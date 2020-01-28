import sys

def sockMerchant(n, ar):
  h = {}
  for i in ar:
    if i in h:
      h[i] += 1
    else:
      h[i] = 1
  c = 0
  for j in h:
    c += h[j] // 2
  return c

if __name__ == '__main__':
  sock_number = int(sys.argv[1])
  colors = ''
  for i in range(sock_number):
    colors += (sys.argv[2+i] + ' ')
  color_list = list(map(int, colors.rstrip().split()))
  result = sockMerchant(sock_number, color_list)
  print(result)