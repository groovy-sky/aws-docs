# Connecting from MySQL Workbench

###### To connect from MySQL Workbench

01. Download and install MySQL Workbench at [Download MySQL Workbench](http://dev.mysql.com/downloads/workbench).

02. Open MySQL Workbench.

    ![The Welcome screen in MySQL Workbench.](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/mysql-workbench-main.png)

03. From **Database**, choose **Manage Connections**.

04. In the **Manage Server Connections** window, choose **New**.

05. In the **Connect to Database** window, enter the following information:

- **Stored Connection** – Enter a name for the connection, such as `MyDB`.

- **Hostname** – Enter the DB instance endpoint.

- **Port** – Enter the port used by the DB instance.

- **Username** – Enter the user name of a valid
database user, such as the master user.

- **Password** – Optionally, choose **Store in Vault** and then enter and save
the password for the user.

The window looks similar to the following:

![The Manage Server Connections window in MySQL Workbench.](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/mysql-workbench-connect.png)

You can use the features of MySQL Workbench to customize connections. For example, you can use the **SSL** tab to configure SSL/TLS connections.
For information about using MySQL Workbench, see the [MySQL Workbench documentation](https://dev.mysql.com/doc/workbench/en).
Encrypting client connections to MySQL DB instances with SSL/TLS, see
[Encrypting client connections with SSL/TLS to MySQL DB instances on Amazon RDS](mysql-ssl-connections.md).

06. Optionally, choose **Test Connection** to confirm that the connection to the DB instance is successful.

07. Choose **Close**.

08. From **Database**, choose **Connect to Database**.

09. From **Stored Connection**, choose your connection.

10. Choose **OK**.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Connecting from the command-line client

Connecting with the AWS drivers
