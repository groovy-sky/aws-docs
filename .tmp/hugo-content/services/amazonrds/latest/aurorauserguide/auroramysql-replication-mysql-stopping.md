# Stopping binary log replication for Aurora MySQL

To stop binary log replication with a MySQL DB instance, external MySQL database, or another Aurora DB cluster, follow
these steps, discussed in detail following in this topic.

[1\. Stop binary log replication on the replica target](#AuroraMySQL.Replication.MySQL.Stopping.StopReplication)

[2\. Turn off binary logging on the replication source](#AuroraMySQL.Replication.MySQL.Stopping.DisableBinaryLogging)

## 1\. Stop binary log replication on the replica target

Use the following instructions to stop binary log replication for your database engine.

Database engine  Instructions

Aurora MySQL

**To stop binary log replication on an Aurora MySQL DB cluster replica**
**target**

Connect to the Aurora DB cluster that is the replica
target, and call the [mysql.rds\_stop\_replication](mysql-stored-proc-replicating.md#mysql_rds_stop_replication)
procedure.

RDS for MySQL

**To stop binary log replication on an Amazon RDS DB instance**

Connect to the RDS DB instance that is the replica target and call the [mysql.rds\_stop\_replication](mysql-stored-proc-replicating.md#mysql_rds_stop_replication)
procedure.

MySQL (external)

**To stop binary log replication on an external MySQL**
**database**

Connect to the MySQL database and run the `STOP SLAVE` (version 5.7) or `STOP
                                            REPLICA` (version 8.0) command.

## 2\. Turn off binary logging on the replication source

Use the instructions in the following table to turn off binary logging on the replication source for your database
engine.

Database engineInstructions

Aurora MySQL

**To turn off binary logging on an**
**Amazon Aurora DB cluster**

1. Connect to the Aurora DB cluster that is the replication source.

2. Use the [mysql.rds\_set\_configuration](mysql-stored-proc-configuring.md#mysql_rds_set_configuration) procedure and specify the configuration
    parameter `binlog retention hours`, with the value `NULL`, as shown in
    the following example.

```sql

CALL mysql.rds_set_configuration('binlog retention hours', NULL);
```

###### Note

You can't use the value `0` for `binlog retention hours`.

3. Set the `binlog_format` parameter to `OFF` on the replication
    source. The `binlog_format` parameter is in the custom DB cluster parameter group
    associated with your DB cluster.

After you've changed the `binlog_format` parameter value, reboot your DB
    cluster for the change to take effect.

For more information, see [Amazon Aurora DB cluster and DB instance parameters](user-workingwithdbclusterparamgroups.md#Aurora.Managing.ParameterGroups) and [Modifying parameters in a DB parameter group in Amazon Aurora](user-workingwithparamgroups-modifying.md).

RDS for MySQL

**To turn off binary logging on an Amazon RDS DB instance**

You can't turn off binary logging directly for an Amazon RDS DB instance, but you can turn it off by
doing the following:

1. Turn off automated backups for the DB instance. You can turn off automated backups by
    modifying an existing DB instance and setting the **Backup Retention**
**Period** to 0. For more information, see [Modifying an Amazon RDS DB\
    instance](../userguide/overview-dbinstance-modifying.md) and [Working with backups](../userguide/user-workingwithautomatedbackups.md) in the _Amazon Relational Database Service User_
_Guide_.

2. Delete all read replicas for the DB instance. For more information, see [Working with read replicas of MariaDB, MySQL, and\
    PostgreSQL DB instances](../userguide/user-readrepl.md) in the _Amazon Relational Database Service User_
_Guide_.

MySQL (external)

**To turn off binary logging on an external MySQL database**

Connect to the MySQL database and call the `STOP REPLICATION` command.

1. From a command shell, stop the `mysqld` service,

```bash

sudo service mysqld stop
```

2. Edit the `my.cnf` file (this file is usually under `/etc`).

```bash

sudo vi /etc/my.cnf
```

Delete the `log_bin` and `server_id ` options from the
    `[mysqld]` section.

For more information, see [Setting the replication source configuration](http://dev.mysql.com/doc/refman/8.0/en/replication-howto-masterbaseconfig.html) in the MySQL
    documentation.

3. Start the mysql service.

```bash

sudo service mysqld start
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting up binlog replication

Scaling MySQL reads

All content copied from https://docs.aws.amazon.com/.
