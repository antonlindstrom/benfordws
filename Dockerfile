FROM mischief/docker-golang

ADD ./ /root/go
WORKDIR /root/go

RUN make
RUN make test
CMD ./bin/benfordws
