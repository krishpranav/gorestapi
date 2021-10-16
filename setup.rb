#!/usr/bin/env ruby 

class Install
    def installrestapi
        sleep 1
        system('go mod tidy')
        sleep 1
        system('go run main.go')
    end

end

class InstallFrontend
    # need to develop this block
end


def main
    function = Install.new
    function.installrestapi
end

main