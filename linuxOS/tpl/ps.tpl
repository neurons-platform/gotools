# ps 过滤进程
f() {
   ps -ef | grep $1 | grep -v grep
}

