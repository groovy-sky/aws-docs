# RDS for Oracle users and privileges

When you create an Amazon RDS for Oracle DB instance, the default master user has most of the maximum
user permissions on the DB instance. Use the master user account for any administrative tasks,
such as creating additional user accounts in your database. Because RDS is a managed
service, you aren't allowed to log in as `SYS` and `SYSTEM`, and thus
don't have `SYSDBA` privileges.

###### Topics

- [Limitations for Oracle DBA privileges](#Oracle.Concepts.dba-limitations)

- [How to manage privileges on SYS objects](#Oracle.Concepts.Privileges.SYS-objects)

## Limitations for Oracle DBA privileges

In the database, a _role_ is a collection of privileges that you can
grant to or revoke from a user. An Oracle database uses roles to provide security. For more
information, see [Configuring Privilege and Role Authorization](https://docs.oracle.com/en/database/oracle/oracle-database/19/dbseg/configuring-privilege-and-role-authorization.html) in the Oracle Database
documentation.

The predefined role `DBA` normally allows all administrative privileges on an
Oracle database. When you create a DB instance, your master user account gets DBA privileges
(with some limitations). To deliver a managed experience, an RDS for Oracle database doesn't
provide the following privileges for the `DBA` role:

- `ALTER DATABASE`

- `ALTER SYSTEM`

- `CREATE ANY DIRECTORY`

- `DROP ANY DIRECTORY`

- `GRANT ANY PRIVILEGE`

- `GRANT ANY ROLE`

For more RDS for Oracle system privilege and role information, see [Master user account privileges](usingwithrds-masteraccounts.md).

## How to manage privileges on SYS objects

You can manage privileges on `SYS` objects by using the
`rdsadmin.rdsadmin_util` package. For example, if you create the database
user `myuser`, you could use the
`rdsadmin.rdsadmin_util.grant_sys_object` procedure to grant
`SELECT` privileges on `V_$SQLAREA` to `myuser`.
For more information, see the following topics:

- [Granting SELECT or EXECUTE privileges to SYS objects](appendix-oracle-commondbatasks-transferprivileges.md)

- [Revoking SELECT or EXECUTE privileges on SYS objects](appendix-oracle-commondbatasks-revokeprivileges.md)

- [Granting privileges to non-master users](appendix-oracle-commondbatasks-permissionsnonmasters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Oracle licensing

Oracle instance classes

All content copied from https://docs.aws.amazon.com/.
