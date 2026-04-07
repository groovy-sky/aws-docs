# Installing the MySQL command-line client

Most Linux distributions include the MariaDB client instead of the Oracle MySQL
client. To install the MySQL command-line client on Amazon Linux 2023, run the following
command:

```nohighlight

sudo dnf install mariadb105
```

To install the MySQL command-line client on Amazon Linux 2, run the following command:

```nohighlight

sudo yum install mariadb
```

To install the MySQL command-line client on most DEB-based Linux distributions, run
the following command:

```nohighlight

apt-get install mariadb-client
```

To check the version of your MySQL command-line client, run the following
command:

```nohighlight

mysql --version
```

To read the MySQL documentation for your current client version, run the following
command:

```nohighlight

man mysql
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Finding the connection information

Connecting from the command-line client
