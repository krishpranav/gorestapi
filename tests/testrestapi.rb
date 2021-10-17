#!/usr/bin/env ruby

class TestRestApi
    def testrestapi
        sleep 1
        system('go mod tidy')
        sleep 1
        system('go run test')
    end
end

def main
    test = TestRestApi.new
    test.testrestapi
end

main