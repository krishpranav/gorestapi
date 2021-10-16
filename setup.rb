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

class InstallFrontend
    def installfrontendfunc
        sleep 1
        system('cd views')
        sleep 1
        system('yarn install && yarn build')
        sleep 1
        puts "frontend website has been builded successfully :)"
    end
end

def callbackinstallclass
    installcall = Install.new
    installcall.installrestapi
end

def callbackfrontendclass
    installfrontend = InstallFrontend.new
    installfrontend.installfrontendfunc
end




def main
    callbackfrontendclass
    sleep 1
    callbackinstallclass
end


main