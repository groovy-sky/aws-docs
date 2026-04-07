# Migrating data from a MySQL DB snapshot to a MariaDB DB instance

You can migrate an RDS for MySQL DB snapshot to a new DB instance running MariaDB using the
AWS Management Console, the AWS CLI, or Amazon RDS API. You must use a DB snapshot that was created from an Amazon RDS DB instance running MySQL 5.6 or 5.7.
To learn how to create an RDS for MySQL DB snapshot, see [Creating a DB snapshot for a Single-AZ DB instance for Amazon RDS](user-createsnapshot.md).

Migrating the snapshot doesn't affect the original DB instance from which the snapshot was taken. You can test and validate
the new DB instance before diverting traffic to it as a replacement for the original DB instance.

After you migrate from MySQL to MariaDB, the MariaDB DB instance is associated with the default
DB parameter group and option group. After you restore the DB snapshot, you can associate a custom DB parameter group
with the new DB instance. However, a MariaDB parameter group has a different set of configurable system variables.
For information about the differences between MySQL and MariaDB system variables,
see [System Variable Differences between MariaDB and MySQL](https://mariadb.com/kb/en/system-variable-differences-between-mariadb-and-mysql).
To learn about DB parameter groups, see [Parameter groups for Amazon RDS](user-workingwithparamgroups.md).
To learn about option groups, see [Working with option groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html).

## Performing the migration

You can migrate an RDS for MySQL DB snapshot to a new MariaDB DB instance using the AWS Management Console, the AWS CLI, or the
RDS API.

###### To migrate a MySQL DB snapshot to a MariaDB DB instance

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Snapshots**,
    and then select the MySQL DB snapshot you want to migrate.

3. For **Actions**, choose **Migrate**
**snapshot**. The **Migrate database** page
    appears.

4. For **Migrate to DB Engine**, choose **mariadb**.

Amazon RDS selects the **DB engine version** automatically. You can't change the DB engine version.

![The Migrate database page to migrate from MySQL to MariaDB in the Amazon RDS console.](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/MigrateMariaDB.png)

5. For the remaining sections, specify your DB instance settings. For information about each setting, see
    [Settings for DB instances](user-createdbinstance-settings.md).

6. Choose **Migrate**.

To migrate data from a MySQL DB snapshot to a MariaDB DB instance, run the AWS CLI
[`restore-db-instance-from-db-snapshot`](https://docs.aws.amazon.com/cli/latest/reference/rds/restore-db-instance-from-db-snapshot.html) command with
the following options:

- --db-instance-identifier – Name of the DB instance to create from the DB snapshot.

- --db-snapshot-identifier – The identifier for the DB snapshot to restore from.

- --engine – The database engine to use for the new instance.

###### Example

For Linux, macOS, or Unix:

```nohighlight

aws rds restore-db-instance-from-db-snapshot \
    --db-instance-identifier newmariadbinstance \
    --db-snapshot-identifier mysqlsnapshot \
    --engine mariadb
```

For Windows:

```nohighlight

aws rds restore-db-instance-from-db-snapshot ^
    --db-instance-identifier newmariadbinstance ^
    --db-snapshot-identifier mysqlsnapshot ^
    --engine mariadb
```

To migrate data from a MySQL DB snapshot to a MariaDB DB instance, call the Amazon RDS API operation
[`RestoreDBInstanceFromDBSnapshot`](../../../../reference/amazonrds/latest/apireference/api-restoredbinstancefromdbsnapshot.md).

## Incompatibilities between MariaDB and MySQL

Incompatibilities between MySQL and MariaDB include the following:

- You can't migrate a DB snapshot created with MySQL 8.0 to MariaDB.

- If the source MySQL database uses a SHA256 password hash, make sure to reset user passwords
that are SHA256 hashed before you connect to the MariaDB database. The following
code shows how to reset a password that is SHA256 hashed.

```sql

SET old_passwords = 0;
UPDATE mysql.user SET plugin = 'mysql_native_password',
Password = PASSWORD('new_password')
WHERE (User, Host) = ('master_user_name', %);
FLUSH PRIVILEGES;

```

- If your RDS master user account uses the SHA-256 password hash, make sure to reset the
password using the AWS Management Console, the [`modify-db-instance`](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-instance.html) AWS CLI command, or the [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) RDS
API operation. For information about modifying a DB instance, see [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

- MariaDB doesn't support the Memcached plugin. However, the data used by the Memcached plugin
is stored as InnoDB tables. After you migrate a MySQL DB snapshot, you can
access the data used by the Memcached plugin using SQL. For more information
about the innodb\_memcache database, see [InnoDB memcached Plugin Internals](https://dev.mysql.com/doc/refman/8.0/en/innodb-memcached-internals.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Parameters for MariaDB

MariaDB on Amazon RDS SQL reference
