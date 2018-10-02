# f keyWord
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
                pidstat -T ALL -drwu -p $Pid 1 5 | tail -n 17
           else
    	          echo "not find process"
           fi
}

