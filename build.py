import subprocess
import os
if os.name == 'nt':
    import pyuac

def patcher(undo = False):
    patches = [] # we don't need any patches for now

    for patch in patches:
        filePath = ""
        if os.name == 'nt':
            filePath = patch['file_path_win']
        else:
            filePath = patch['file_path_lin']

        with open(filePath, 'r', encoding="utf8") as file:
            filedata = file.read()

        if undo == False:
            filedata = filedata.replace(patch['find'], patch['replace'])
        else:
            filedata = filedata.replace(patch['replace'], patch['find'])

        with open(filePath, 'w', encoding="utf8") as file:
            file.write(filedata)

def cleanGoBuildCache():
    subprocess.check_output("go clean -cache", shell=True)

def build():
    sys_env = os.environ.copy()

    print("Building for Linux amd64...")
    sys_env['GOARCH'] = "amd64"
    sys_env['GOOS'] = "linux"
    subprocess.check_output("go build -o bin/simpledns_linux_amd64 -ldflags=\"-w -s\" -trimpath -v", env=sys_env, shell=True)

    print("Building for Linux arm64...")
    sys_env['GOARCH'] = "arm64"
    sys_env['GOOS'] = "linux"
    subprocess.check_output("go build -o bin/simpledns_linux_arm64 -ldflags=\"-w -s\" -trimpath -v", env=sys_env, shell=True)

    print("Building for Windows amd64...")
    sys_env['GOARCH'] = "amd64"
    sys_env['GOOS'] = "windows"
    subprocess.check_output("go build -o bin/simpledns_windows_amd64.exe -ldflags=\"-w -s\" -trimpath -v", env=sys_env, shell=True)

def main():
    patcher()
    cleanGoBuildCache()
    build()
    patcher(True)

    input("Done. Press any key to exit.")

if __name__ == "__main__":
    if os.name == 'nt':
        if not pyuac.isUserAdmin():
            # https://stackoverflow.com/a/19719292/8524395
            print("Re-launching as admin...")
            pyuac.runAsAdmin()
        else:        
            main()
    else:
        main()