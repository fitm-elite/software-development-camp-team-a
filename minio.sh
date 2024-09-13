#!/bin/sh

docker stop minio
docker rm minio

mkdir -p ./resources/minio/data

docker run \
    -p 9000:9000 \
    -p 9001:9001 \
    --name minio \
    -v ./resources/minio/data:/data \
    -e "MINIO_ROOT_USER=elebs" \
    -e "MINIO_ROOT_PASSWORD=vPN294iGyPgj26cWQfKwE7E7" \
    -e "MINIO_SERVER_URL=https://minio.example.com" \
    -e "MINIO_BROWSER_REDIRECT_URL=https://minio.example.com/minio/ui" \
    quay.io/minio/minio:RELEASE.2024-09-09T16-59-28Z-cpuv1 server /data --console-address ":9001"
