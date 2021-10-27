#!/usr/bin/env python3

import os
import time

Y = set(['YES', 'Y', 'yes', 'y'])
N = set(['NO', 'N', 'no', 'n'])

def setupbackend():
    setuprest = input('Want to setup backend Y / N >> ')

    if setuprest in Y:
        os.system('go mod tidy')
        time.sleep(1)
        print('https://github.com/krishpranav/gorestapi\n ⭐ Give it a star if you like it')
        time.sleep(1)
        print('Running gorestapi')
        os.system('go run main.go')
    elif setuprest in N:
        print('OK BYE')

def setupfrontned():
    setupfront = input('Want to setup frontend Y / N >> ')

    if setupfront in Y:
        print('Setting up frontend')
        os.system('cd views/')
        time.sleep(1)
        os.system('yarn install && yarn build')

    elif setupfront in N:
        print('OK BYE')

def main():
    setupbackend()
    setupfrontned()

if __name__ == "__main__":
    print("SETUP SCRIPT")
    main()