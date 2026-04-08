# Monitoring Amazon Aurora log files

Every RDS database engine generates logs that you can access for auditing and troubleshooting. The type of logs depends on your database
engine.

You can access database logs for DB instances using the AWS Management Console, the AWS Command Line Interface (AWS CLI), or the Amazon RDS API. You can't view, watch, or download transaction
logs.

###### Note

In some cases, logs contain hidden data. Therefore, the AWS Management Console might show content in a log file, but
the log file might be empty when you download it.

###### Topics

- [Viewing and listing database log files](user-logaccess-procedural-viewing.md)

- [Downloading a database log file](user-logaccess-procedural-downloading.md)

- [Watching a database log file](user-logaccess-procedural-watching.md)

- [Publishing database logs to Amazon CloudWatch Logs](user-logaccess-procedural-uploadtocloudwatch.md)

- [Reading log file contents using REST](downloadcompletedblogfile.md)

- [AuroraMySQL database log files](user-logaccess-concepts-mysql.md)

- [Aurora PostgreSQL database log files](user-logaccess-concepts-postgresql.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon RDS event categories and event messagesfor Aurora

Viewing and listing database log files

All content copied from https://docs.aws.amazon.com/.
