# f keyWord topN
f() {
               Pid=""
               while read pid
               do
                  if $(cat /proc/$pid/environ |strings  |grep $1 2>&1 >/dev/null)
                  then
                       Pid=$pid
                  fi
               done <  <(ps -ef |grep java |grep -v grep  |awk '{print $2}')

               top -b -n1  -H -p $Pid |  head -n 20

               if [ ! -z "$Pid" ]
               then
                   n=$2
                   sp=$(top -b -H -p $Pid -n 1 |grep java  |sed 's/ *\([0-9]\+\)/\1/g' |awk '{print $1}' |sed -n "1,${n}p"|sed 's/[^0-9|\]//g' |xargs -i  printf "0x%x\n" {} |sed  ':a;N;s/\n/\\\|/g;ta')
                   {{ .GetJstackPath }}  $Pid |sed  -n  "/$sp/,/^$/p"
                   echo $sp
                   #printf "%d\n" $sp
               else
                    echo "not find process"
               fi
}

