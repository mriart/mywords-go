FROM alpine:3.17

# Alpine needs some libs to run go executable. So this link
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add gcompat

# "COPY . ./" copies everything, but I do not want to copy .git files. For directories, one by one
# COPY . ./
COPY Dockerfile README.md env_auth.sh main main.go ./
COPY public/ ./public/
COPY templates/ ./templates/

ENV CLOUDANT_URL https://0634c5f6-50c4-47aa-81a6-f5a4dcce30ed-bluemix.cloudantnosqldb.appdomain.cloud
ENV CLOUDANT_APIKEY cTcWGHggP9fzy7qixNQeX-MKd7UrVI6Jn0iZn6DvP8jP

EXPOSE 8080
CMD ["./main"]
