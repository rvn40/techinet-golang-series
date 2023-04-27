# Go Labs(Ngoprek): How to Install Golang on Linux Ubuntu

Today go language is become one of popular language, especially for create an RESTAPI programs. It’s light and versatile enough for use in web development. 

## Prerequisite
- 2 cpu and 2GB RAM (Minimum)
- Ubuntu OS
- Access to a terminal/cli
- An ubuntu user with sudo privileges
- Internet connection

## Instructions

For instruction in video format:

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/PythonSeries/E2.png" width="560" height="315">](https://www.youtube.com/embed/7cng0PQeBzE)

## How to install

#### Go to Offical Website

So first of all we need to go to install go programming language is open golang.dev Website.

You can easily see easily one big download button. You can click that and it will navigate your browser to download page. On that page you can choose golang version that you willing to install and the os version for your machine. I am going to choose the latest version of golang for linux installation because my machine is ubuntu. 

```
Notes: By the times this docs created, the latest version of golang is go1.20.3.
```

#### Remove any previous Go installation

By deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

```
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
```
(You may need to run the command as root or through sudo).

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

#### Add /usr/local/go/bin to the PATH environment variable

You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```
export PATH=$PATH:/usr/local/go/bin
```

Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

#### Verify Golang Version

```
go version
```