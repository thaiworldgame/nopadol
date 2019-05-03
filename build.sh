#!/bin/bash

# ลบไฟล์ binary เก่าถ้ามี
rm host-linux-amd64

#หากมีการแก้ไข Dockerfile ให้รัน build.sh ก่อนเพื่อสร้าง image ใหม่
docker build -t mrtomyum/makevending .

# รันอิมเมจให้เป็น Container ขึ้นมาโดยที่ map volume ของจุดนี้ให้ตรงกับตำแหน่ง /go/src ใน container ที่ตั้งชื่อว่า maker
echo RUN container and map GOPATH folder...
docker run -dt \
 --name maker \
 -v /Users/tom/go:/go \
 mrtomyum/makevending

# เมื่อสั่ง go build ก็จะได้ไฟล์ host วางไว้ตรงนี้เลยไม่ต้อง copy เข้าๆ ออกๆ
echo Build go project...
docker exec maker \
  go build \
  -v -ldflags "-s -w \
  -X main.Version=1.0.4 \
  -X main.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
  -X main.GitHash=`git rev-parse HEAD`" \
  -o host-linux-amd64

# สั่งหยุด container และลบทิ้งซะ
docker stop maker && docker rm maker
# scp host-linux-amd64 pb:/app/host
scp host-linux-amd64 remote:~/