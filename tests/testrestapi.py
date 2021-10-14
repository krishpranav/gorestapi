import os
import time

def testing():
    print('checking go version :)')
    os.system('go version')
    time.sleep(1)
    os.system('go test')


if __name__ == "__main__":
    testing()
