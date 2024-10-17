import os
os.system("")

idxs = [[0, 6], [18, 24], [33, 39]]
string = "*hello* skdfjskfj *world* lskdfj *hello* blah"

italic  = "\x1B[3m"
inverse = "\x1B[7m"
reset   = "\x1B[0m"

#slices = []
#cursor = 0
#for _, c in enumerate(idxs):
#    start, end = c
#    before = string[cursor:start]
#    inside = italic + string[start+1:end] + reset
#    cursor = end+1
#    slices.append(before)
#    slices.append(inside)
#print(''.join(slices))


cursor = 0
for _, c in enumerate(idxs):
    start, end = c
    before = string[cursor:start] # define before.
    inside = italic + string[start+1:end] + reset # define inside.
    cursor = end+1 # update cursor.
    print(before + inside, end='') # Print before and inside.
print(string[cursor:len(string)], end='') # print after.
print('') # newline!

#def insert_italic(s, r):
#    return s[:r[0]] + italic + s[r[0]+1:r[1]] + reset + s[r[1]+1:]
#
#s = string
#r = idxs[0]
#m = insert_italic(s, r)
#
#print(r)
#print(s)
#print(m)



buffer = []
for i in range(len(string)):
    found = False
    for _, c in enumerate(idxs):
        if i == c[0]:
            print(italic, end='')
            found = True
        elif i == c[1]:
            print(reset, end='')
            found = True
    if not found:
        print(string[i], end='')
print('')
