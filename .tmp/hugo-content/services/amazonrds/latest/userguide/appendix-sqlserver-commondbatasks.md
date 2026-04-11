# Common DBA tasks for Amazon RDS for Microsoft SQL Server

This section describes the Amazon RDS-specific implementations of some common DBA tasks for DB
instances that are running the Microsoft SQL Server database engine. In order to deliver a
managed service experience, Amazon RDS does not provide shell access to DB instances, and it
restricts access to certain system procedures and tables that require advanced privileges.

###### Note

When working with a SQL Server DB instance, you can run scripts to modify a newly
created database, but you cannot modify the \[model\] database, the database used as the
model for new databases.

###### Topics

- [Accessing the tempdb database on Microsoft SQL Server DB instances on Amazon RDS](sqlserver-tempdb.md)

- [Analyzing your database workload on an Amazon RDS for SQL Server DB instance with Database Engine Tuning Advisor](appendix-sqlserver-commondbatasks-workload.md)

- [Changing the db\_owner to the rdsa account for your Amazon RDS for SQL Server database](appendix-sqlserver-commondbatasks-changedbowner.md)

- [Managing collations and character sets for Amazon RDS for Microsoft SQL Server](appendix-sqlserver-commondbatasks-collation.md)

- [Creating a database user for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-createuser.md)

- [Determining a recovery model for your Amazon RDS for SQL Server database](appendix-sqlserver-commondbatasks-databaserecovery.md)

- [Determining the last failover time for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-lastfailover.md)

- [Troubleshooting point-in-time-recovery failures due to a log sequence number gap](appendix-sqlserver-commondbatasks-pitr-lsn-gaps.md)

- [Deny or allow viewing database names for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-manageview.md)

- [Disabling fast inserts during bulk loading for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-disablefastinserts.md)

- [Dropping a database in an Amazon RDS for Microsoft SQL Server DB instance](appendix-sqlserver-commondbatasks-dropmirrordb.md)

- [Renaming a Amazon RDS for Microsoft SQL Server database in a Multi-AZ deployment](appendix-sqlserver-commondbatasks-renamingdb.md)

- [Resetting the db\_owner role membership for master user for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-resetpassword.md)

- [Restoring license-terminated DB instances for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-restorelti.md)

- [Transitioning a Amazon RDS for SQL Server database from OFFLINE to ONLINE](appendix-sqlserver-commondbatasks-transitiononline.md)

- [Using change data capture for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-cdc.md)

- [Using SQL Server Agent for Amazon RDS](appendix-sqlserver-commondbatasks-agent.md)

- [Working with Amazon RDS for Microsoft SQL Server logs](appendix-sqlserver-commondbatasks-logs.md)

- [Working with trace and dump files for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-tracefiles.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Best practices

Accessing the tempdb database

All content copied from https://docs.aws.amazon.com/.
