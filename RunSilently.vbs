Dim WinScriptHost
Set WinScriptHost = CreateObject("WScript.Shell")
' Path to the compiled executable:
UrlHandlerPath="WebstormURLHandler.exe"

' Path to the webstorm executable:
WebstormPath="webstorm.cmd"
' See the documentation for a good way of getting this

' Hostname to listen for in the URL
Host="open"

' Port to listen on
Port=80

' Combine it all together
WinScriptHost.Run UrlHandlerPath & " -executable=" & WebstormPath & " -host=" & Host & " -port=" & Port, 0