ip=`ifconfig -a | grep inet | grep -v 127.0.0.1 | grep -v inet6 | awk '{print $2}' | tr -d "addr:"​ | tail -1`
port=1296
echo $ip:$port
CONSUL_PORT=8500 CONSUL_HOST=127.0.0.1 INSTANCE_HOST=$ip INSTANCE_PORT=$port go run cmd/management/main.go