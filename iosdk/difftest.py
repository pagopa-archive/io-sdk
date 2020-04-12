#!/usr/bin/env python3
import sys
tests = {}
base = sys.argv[1]
f = open(base, 'r')
lines = f.readlines()

fails=[]
gots = []
wants = []

i=0
while i < len(lines):
    line = lines[i]
    if(line.startswith("=== RUN")):
        if lines[i+1].startswith("--- FAIL:"):
            fails.append(i+1)
    i+=1

if len(sys.argv) == 2:
    n = 0
    for i in fails:
        print(n, lines[i], end='')
        n += 1
    sys.exit(len(fails))

n = int(sys.argv[2])
k = fails[n]+2
got = []
want = []
while not lines[k].startswith("want:"):
    #print(lines[k])
    got.append(lines[k])
    k += 1
k += 1
while not lines[k].startswith("=== RUN"):
    #print(lines[k])
    want.append(lines[k])
    k += 1


import tempfile
import os

f1n = "%s.got"%base

f2n = "%s.want"%base

with open(f1n, "w") as f1:
    with open(f2n, "w") as f2:
        f1.writelines(got)
        f2.writelines(want)

os.system("diff %s %s" % (f1n, f2n))

