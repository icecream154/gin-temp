test_ip='120.55.170.127'
test_password='a4d14B835d7116Fa'

set timeout -1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o my-app cmd/web/main.go
expect -c "
set timeout -1
spawn scp -p my-app root@$test_ip:~/youni-go-pre
expect \"*password*\"
send \"$test_password\n\"
expect eof"

rm my-app

for env in dev test prod
do
  expect -c "
  set timeout -1
  spawn scp -r ./config/app/config_$env.yml root@$test_ip:~/youni-go/my-app/config/app
  expect \"*password*\"
  send \"$test_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/db/gorm_v2_$env.yml root@$test_ip:~/youni-go/my-app/config/db
  expect \"*password*\"
  send \"$test_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/rpc/rpc_$env.yml root@$test_ip:~/youni-go/my-app/config/rpc
  expect \"*password*\"
  send \"$test_password\n\"
  expect eof"

done

for env_docker in dev_docker test_docker prod_docker
do
  expect -c "
  set timeout -1
  spawn scp -r ./config/db/gorm_v2_$env_docker.yml root@$test_ip:~/youni-go/my-app/config/db
  expect \"*password*\"
  send \"$test_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/rpc/rpc_$env_docker.yml root@$test_ip:~/youni-go/my-app/config/rpc
  expect \"*password*\"
  send \"$test_password\n\"
  expect eof"
done

expect -c "
set timeout -1
spawn scp -r ./config/oss/config_prod.json root@$test_ip:~/youni-go/my-app/config/oss
expect \"*password*\"
send \"$test_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/config_test.json root@$test_ip:~/youni-go/my-app/config/oss
expect \"*password*\"
send \"$test_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/config_dev.json root@$test_ip:~/youni-go/my-app/config/oss
expect \"*password*\"
send \"$test_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/white_list/internal_accounts.json root@$test_ip:~/youni-go/my-app/config/white_list
expect \"*password*\"
send \"$test_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/policy/bucket_read_policy.txt root@$test_ip:~/youni-go/my-app/config/oss/policy
expect \"*password*\"
send \"$test_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/policy/bucket_write_policy.txt root@$test_ip:~/youni-go/my-app/config/oss/policy
expect \"*password*\"
send \"$test_password\n\"
expect eof"