#coding=utf8

from fabric.api import run, env, cd
from fabric.operations import local
from fabric.context_managers import prefix
from fabric.context_managers import path

env.hosts = ['root@likangwei.com']

workDir = "/root/dev/gopath/src/memplus_service"
PATH = "/usr/lib/go-1.9/bin/" 

commitMsg = ""


def gitCommit():
    global commitMsg
    if commitMsg:
        local("cd ..; git add *; git commit -m '%s'; git push" % commitMsg)

def deploy():
    global commitMsg
    commitMsg = raw_input("请输入此次commit: ")
    gitCommit()

    with cd(workDir):
        with path(PATH):
            with prefix("export GOPATH=/root/dev/gopath"):
                run("echo $PATH; echo $GOPATH")
                run("git reset --hard")
                run("go get github.com/smartwalle/alipay")
                run("go get -u github.com/satori/go.uuid")
                run("git pull")
                run("go build")
                run("cp /root/dev/gopath/src/memplus.conf %s/conf/app.conf" % workDir)
                run("supervisorctl restart memoryplus")
