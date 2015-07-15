# Build Sysbench image
FROM ubuntu:14.04
MAINTAINER Wangkirin <wangqilin2@huawei.com>
ENV DEBIAN_FRONTEND noninteractive


# Install sysbench
RUN sudo apt-get update
RUN sudo apt-get install sysbench -y
RUN mkdir  /data ; cd /data ; 

# load  data volume
VOLUME /data



