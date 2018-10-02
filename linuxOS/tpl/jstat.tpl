# f keyWord
# 使用jstat 查看java进程gc
f() {
     Pid=""
     while read pid
     do
        if $(cat /proc/$pid/environ |strings  |grep $1 2>&1 >/dev/null)
        then
             Pid=$pid
        fi
     done <  <(ps -ef |grep java |grep -v grep  |awk '{print $2}')
     if [ ! -z "$Pid" ]
     then
          {{ .GetJstatPath }}  -gcutil -h10 $Pid  1s 5
     else
          echo "not find process"
     fi
}


