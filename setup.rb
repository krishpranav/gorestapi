#!/usr/bin/env ruby 

class Install
    def installrestapi
        sleep 1
        system('go mod tidy')
        sleep 1
        puts "Go modules installed successfully and going to start the server"
        sleep 1
        system('go run main.go')
    end

end

def callbackinstallclass
    installcall = Install.new
    installcall.installrestapi
end


def main
    callbackinstallclass
end


main