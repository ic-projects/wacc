#!/usr/bin/env python3

import fileinput
import glob

for filename in glob.iglob('test_cases/valid/**/*.s', recursive=True):
    print("Cleaning ", filename, "...")
    for line in fileinput.FileInput(filename, inplace=1):
        if (len(line) > 0 and line[0].isdigit()):
            line = line.split('\t', 1)[-1]
            if (len(line) > 0 and not line[0].isdigit()):
                print(line, end='')
