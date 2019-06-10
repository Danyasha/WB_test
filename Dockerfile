FROM golang

MAINTAINER danyasha

ENV app_dir /opt/app/
ENV GOPATH /opt/app/
ENV go_app_port 80
WORKDIR ${app_dir}
EXPOSE ${go_app_port}

RUN mkdir -p ${app_dir}
# RUN export export GOPATH=$(pwd)=$(pwd)
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/mattn/go-sqlite3"

ADD /app/ ${app_dir}
RUN go install main
ADD /tv_storage.db ${app_dir}
RUN go build main
CMD ["/bin/bash"]
# ENTRYPOINT [ "/bin/bash" ]
#docker rm -f $(docker ps -aq)
#docker build -t go_app .
#docker run -d -p 8080:80 --name myawesomework go_app 
#http://192.168.99.100