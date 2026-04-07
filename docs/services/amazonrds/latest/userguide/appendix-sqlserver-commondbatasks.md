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

- [Accessing the tempdb database on Microsoft SQL Server DB instances on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.TempDB.html)

- [Analyzing your database workload on an Amazon RDS for SQL Server DB instance with Database Engine Tuning Advisor](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Workload.html)

- [Changing the db\_owner to the rdsa account for your Amazon RDS for SQL Server database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.ChangeDBowner.html)

- [Managing collations and character sets for Amazon RDS for Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Collation.html)

- [Creating a database user for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CreateUser.html)

- [Determining a recovery model for your Amazon RDS for SQL Server database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.DatabaseRecovery.html)

- [Determining the last failover time for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.LastFailover.html)

- [Troubleshooting point-in-time-recovery failures due to a log sequence number gap](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.PITR-LSN-Gaps.html)

- [Deny or allow viewing database names for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.ManageView.html)

- [Disabling fast inserts during bulk loading for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.DisableFastInserts.html)

- [Dropping a database in an Amazon RDS for Microsoft SQL Server DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.DropMirrorDB.html)

- [Renaming a Amazon RDS for Microsoft SQL Server database in a Multi-AZ deployment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.RenamingDB.html)

- [Resetting the db\_owner role membership for master user for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.ResetPassword.html)

- [Restoring license-terminated DB instances for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.RestoreLTI.html)

- [Transitioning a Amazon RDS for SQL Server database from OFFLINE to ONLINE](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.TransitionOnline.html)

- [Using change data capture for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CDC.html)

- [Using SQL Server Agent for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Agent.html)

- [Working with Amazon RDS for Microsoft SQL Server logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Logs.html)

- [Working with trace and dump files for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.TraceFiles.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Best practices

Accessing the tempdb database
