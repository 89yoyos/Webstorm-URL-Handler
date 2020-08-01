# Webstorm URL Handler
Designed with Redux DevTools in mind, Webstorm URL Handler listens to http requests to a given port and opens Webstorm to the specified file and line.

## Configuration

Because it is designed with Redux DevTools in mind it listens to port 80, but that can be configured via arguments.
```shell script
WebstormURLHandler.exe -port=[80] -host=["open"] -executable=["webstorm.cmd"] 
```

`-port` and `-host` work as expected. `-executable` is the path to the Webstorm executable. On a normal Windows machine, you can find the executables in `C:\Users\[User]\AppData\Local\JetBrains\Toolbox\apps\WebStorm\ch-0\[Version]\bin`. I would strongly recommend against this, however, because the directory changes when the IDE updates, and you would need to change the launch arguments, too. Instead, I suggest enabling [the "generate shell scripts" option](https://www.jetbrains.com/help/idea/working-with-the-ide-features-from-command-line.html#toolbox) in the Jetbrains Toolbox. This creates a folder full of shell files that launch the latest version of each IDE, and these scripts are automatically updated by the toolbox as new versions of the IDEs are released.

## Running Silently

### Windows
On Windows, the included [RunSilently.vbs](RunSilently.vbs) file works for launching the program in the background. On Linux, [RunSilently.sh](RunSilently.sh) should do the trick, but I have only tested it on Windows Subsystem for Linux.

Each file contains all the configuration settings with default values.
```
UrlHandlerPath="WebstormURLHandler.exe"
```
URLHandlerPath is the path to the compiled executable of this program. This can be relative or absolute.
```
WebstormPath="webstorm.cmd"
```
This is the path to the WebStorm executable **OR** the generated shell scripts from the toolbox
```
Host="open"
Port=80
```
And these work as expected and simply pass through to the executable.