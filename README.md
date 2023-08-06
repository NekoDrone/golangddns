# golangddns
A lightweight cross-platform Dynamic DNS update service written in Go.
## How It Works
The program sends a GET request to ip-api in order to find the machine's current outgoing IP. It then checks against the last known outgoing IP of the machine.

If there is a difference, it sends the current outgoing IP as a HTTP Request to an API Endpoint.

This was initially built to run with the Cloudflare v4 API, but can be adapted to use any API Endpoint.
## What you need
To use this service, you will minimally need the following:
1. A DNS Provider with an API capable of updating the DNS Record of your site.
2. A valid build of this project.
3. A Go installation. The project was built and tested for Go 1.20. Earlier versions of Go may be used, but full functionality is not guaranteed.
4. ``(optional)`` A cron job scheduler if you want to continuously monitor and update any IP changes.


# Installation
## Requirements
- Go 1.20. Earlier versions of Go may be used, but support is not guaranteed.

This package relies on no other external packages, and is written fully using Go's default packages.

## Instructions
There are two main ways to install this service.
1. Build it from source.
2. Grab a pre-built package. The only pre-built package is tested on Linux/Ubuntu.

**The preferred method of installation is to build from source.**
### To Build From Source
1. Clone this repo in a folder.
2. Navigate into the ``golangddns`` folder and run ``go install ./cmd/golangddns``
3. The build will be in your GOBIN path. For default Go installs, this will be ``$HOME/go/bin/``.
4. If you cannot find your GOBIN path, or you do not know where your GOBIN path is, run ``go env GOBIN``. If this returns blank, your GOBIN path is the default path.

A compiled list of commands for Linux/Ubuntu is provided below.
```
git clone https://github.com/NekoDrone/golangddns
cd golangddns
go install ./cmd/golangddns
cd ~/go/bin
```

### After Getting a Build
1. If you have built from source, navigate to your GOBIN path and run the file ``golangddns``. Otherwise, you may run the file anywhere you want to.
2. The program will prompt you for an API Key, an API Endpoint, and the DNS Record you want to update. For the DNS Record, it follows normal rules (e.g. @ -> root, www -> www).
3. The program generates a lastIp file and a .env file to store relevant information. If you want to reset the service, delete these two files.
4. Schedule the cron job to run at ``* * * * *``. This will run the update every minute.

# Contributing
If you would like to contribute, please raise an issue for a bug, feature request, or suggestion first.

If you would like to write a pull request, please do so after submitting an issue.

Any and all contributions are welcome!
