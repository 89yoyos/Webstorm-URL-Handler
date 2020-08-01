#!/bin/bash

# Path to the compiled executable:
UrlHandlerPath="./WebstormURLHandler"

# Path to the webstorm executable:
WebstormPath="./webstorm.cmd"
# See the documentation for a good way of getting this

# Hostname to listen for in the URL
Host="open"

# Port to listen on
Port="80"

# Combine it all together
cmd="$UrlHandlerPath -port=$Port -host=$Host -executable=$WebstormPath"

$cmd &
disown