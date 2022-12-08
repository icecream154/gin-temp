prod_ip='39.100.120.238'
prod_password='03bfC644B39fc1eQ'

set timeout -1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o my-app-scheduler cmd/scheduler/auto_event_scheduler.go
expect -c "
set timeout -1
spawn scp -p my-app-scheduler root@$prod_ip:~/youni-go-pre
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

rm my-app-scheduler

for env in dev test prod
do
  expect -c "
  set timeout -1
  spawn scp -r ./config/app/config_$env.yml root@$prod_ip:~/youni-go/my-app-scheduler/config/app
  expect \"*password*\"
  send \"$prod_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/db/gorm_v2_$env.yml root@$prod_ip:~/youni-go/my-app-scheduler/config/db
  expect \"*password*\"
  send \"$prod_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/rpc/rpc_$env.yml root@$prod_ip:~/youni-go/my-app-scheduler/config/rpc
  expect \"*password*\"
  send \"$prod_password\n\"
  expect eof"

done

for env_docker in dev_docker test_docker prod_docker
do
  expect -c "
  set timeout -1
  spawn scp -r ./config/db/gorm_v2_$env_docker.yml root@$prod_ip:~/youni-go/my-app-scheduler/config/db
  expect \"*password*\"
  send \"$prod_password\n\"
  expect eof"

  expect -c "
  set timeout -1
  spawn scp -r ./config/rpc/rpc_$env_docker.yml root@$prod_ip:~/youni-go/my-app-scheduler/config/rpc
  expect \"*password*\"
  send \"$prod_password\n\"
  expect eof"
done

expect -c "
set timeout -1
spawn scp -r ./config/white_list/internal_accounts.json root@$prod_ip:~/youni-go/my-app-scheduler/config/white_list
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/config_test.json root@$prod_ip:~/youni-go/my-app/config/oss
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/config_dev.json root@$prod_ip:~/youni-go/my-app/config/oss
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/white_list/internal_accounts.json root@$prod_ip:~/youni-go/my-app/config/white_list
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/policy/bucket_read_policy.txt root@$prod_ip:~/youni-go/my-app-scheduler/config/oss/policy
expect \"*password*\"
send \"$prod_password\n\"
expect eof"

expect -c "
set timeout -1
spawn scp -r ./config/oss/policy/bucket_write_policy.txt root@$prod_ip:~/youni-go/my-app-scheduler/config/oss/policy
expect \"*password*\"
send \"$prod_password\n\"
expect eof"