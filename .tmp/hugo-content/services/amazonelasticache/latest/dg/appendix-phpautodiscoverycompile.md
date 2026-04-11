# Compiling the source code for the ElastiCache cluster client for PHP

This section covers how to obtain and compile the source code for the ElastiCache Cluster Client for PHP.

There are two packages you need to pull from GitHub and compile;
[aws-elasticache-cluster-client-libmemcached](https://github.com/awslabs/aws-elasticache-cluster-client-libmemcached)
and [aws-elasticache-cluster-client-memcached-for-php](https://github.com/awslabs/aws-elasticache-cluster-client-memcached-for-php).

###### Topics

- [Compiling the libmemcached library](#Appendix.PHPAutoDiscoveryCompile.Libmemcached)

- [Compiling the ElastiCache Memcached auto discovery client for PHP](#Appendix.PHPAutoDiscoveryCompile.Client)

## Compiling the libmemcached library

###### To compile the aws-elasticache-cluster-client-libmemcached library

1. Launch an Amazon EC2 instance.

2. Install the library dependencies.

- On Amazon Linux 201509 AMI

```nohighlight

sudo yum install gcc gcc-c++ autoconf libevent-devel
```

- On Ubuntu 14.04 AMI

```nohighlight

sudo apt-get update
sudo apt-get install libevent-dev gcc g++ make autoconf libsasl2-dev
```

3. Pull the repository and compile the code.

```nohighlight

Download and install  https://github.com/awslabs/aws-elasticache-cluster-client-libmemcached/archive/v1.0.18.tar.gz

```

## Compiling the ElastiCache Memcached auto discovery client for PHP

The following sections describe how to compile the ElastiCache Memcached Auto Discovery Client

###### Topics

- [Compiling the ElastiCache Memcached client for PHP 7](#Appendix.PHPAudiscoveryCompile.Client.PHP7)

- [Compiling the ElastiCache Memcached client for PHP 5](#Appendix.PHPAudiscoveryCompile.PHP5)

### Compiling the ElastiCache Memcached client for PHP 7

Run the following set of commands under the code directory.

```nohighlight

git clone https://github.com/awslabs/aws-elasticache-cluster-client-memcached-for-php.git
cd aws-elasticache-cluster-client-memcached-for-php
git checkout php7
sudo yum install php70-devel
phpize
./configure --with-libmemcached-dir=<libmemcached-install-directory> --disable-memcached-sasl
make
make install
```

###### Note

You can statically link the libmemcached library into the PHP binary so it can be ported across various Linux platforms.
To do that, run the following command before `make`:

```nohighlight

sed -i "s#-lmemcached#<libmemcached-install-directory>/lib/libmemcached.a -lcrypt -lpthread -lm -lstdc++ -lsasl2#" Makefile
```

### Compiling the ElastiCache Memcached client for PHP 5

Compile the `aws-elasticache-cluster-client-memcached-for-php` by
running the following commands under the `aws-elasticache-cluster-client-memcached-for-php/`
folder.

```nohighlight

git clone https://github.com/awslabs/aws-elasticache-cluster-client-memcached-for-php.git
cd aws-elasticache-cluster-client-memcached-for-php
sudo yum install zlib-devel
phpize
./configure --with-libmemcached-dir=<libmemcached-install-directory>
make
make install
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Removing the PHP cluster client

Configuring clients

All content copied from https://docs.aws.amazon.com/.
