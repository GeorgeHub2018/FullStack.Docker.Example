# start from scratch
FROM scratch

#ENV GOPATH /go/
# the test program:
#COPY --from=golang ./go/src/FullStack.Docker.Example/Service/bin FullStack.Service
# COPY bin/FullStack.Service.exe /
COPY bin/FullStack.Service.exe FullStack.Service
EXPOSE 38080
ENTRYPOINT ["./FullStack.Service"]