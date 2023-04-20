curl -X POST -H "Content-Type: application/json" \
    -d '{"phone": "fakePhone", "password": "fakePassword", "code": "TGKS"}' \
    http://localhost:30000/yi/common/register

curl -X POST -H "Content-Type: application/json" \
    -d '{"phone": "fakePhone", "password": "fakePassword"}' \
    http://localhost:30000/yi/common/login

curl -X POST -H "Content-Type: application/json" \
    -d '{"phone": "fakePhone", "content": "Tell me ten fruits"}' \
    http://localhost:30000/yi/qa/submitInput
