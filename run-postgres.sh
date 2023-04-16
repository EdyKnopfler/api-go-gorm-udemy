docker run \
    --rm \
    --name some-postgres \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -p 5432:5432 \
    -v ./data:/var/lib/postgresql/data \
    postgres
