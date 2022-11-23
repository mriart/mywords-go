FROM alpine:latest

# Alpine needs some libs to run go executable. So this link
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY . ./

ENV CLOUDANT_URL https://0634c5f6-50c4-47aa-81a6-f5a4dcce30ed-bluemix.cloudantnosqldb.appdomain.cloud
ENV CLOUDANT_APIKEY cTcWGHggP9fzy7qixNQeX-MKd7UrVI6Jn0iZn6DvP8jP

EXPOSE 8090
CMD ./main
