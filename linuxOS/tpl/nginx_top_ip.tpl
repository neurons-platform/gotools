# f LogName
# 找到访问量最大的IP
f() {
	find {{ .GetNginxLogPath }} -name $1 -type f |xargs -i tail -n1000 {} |awk '{print $1}' |sort -h |uniq -c |sort -rn  |head -n 10
}

