for i in {1..20}; do
  # ab -n 10000 -c 1000 http://localhost:5000/ping
  ab -n 10000 -c 1000 -p request-test.json -T application/json http://localhost:5000/api/v1/users/login
done
