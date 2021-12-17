# CodePay Process Transaction

## development
```bash
# build project (first start)
docker-compose up -d --build

# start project
docker-compose up -d

# stop project (don't destroy the container)
docker-compose stop

# stop project (destroy the container)
docker-compose down

# recreate container without stop
docker-compose up -d --force-recreate

# if you change the Dockerfile, you need build again
docker-compose up -d --force-recreate --build
```

## production
```bash
# create image
 docker build -t tineto/codepay-process-transaction . -f Dockerfile.prod

# test image
docker-compose -f Dockerfile.prod up -d
```