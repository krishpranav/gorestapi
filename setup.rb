#!/usr/bin/env ruby 

class Install
    def installrestapi
        sleep 1
        system('go mod tidy')
        sleep 1
        system('go run main.go')
    end

end

def main
    function = Install.new
    function.installrestapi
end

main