# Starting and stopping mail queue

Use the following instructions to start and stop the DB mail queue:

###### Topics

- [Starting the mail queue](#SQLServer.DBMail.Start)

- [Stopping the mail queue](#SQLServer.DBMail.Stop)

## Starting the mail queue

You use the `rds_sysmail_control` stored procedure to start the Database Mail process.

###### Note

Enabling Database Mail automatically starts the mail queue.

###### To start the mail queue

- Use the following SQL statement.

```

EXECUTE msdb.dbo.rds_sysmail_control start;
GO
```

## Stopping the mail queue

You use the `rds_sysmail_control` stored procedure to stop the Database Mail process.

###### To stop the mail queue

- Use the following SQL statement.

```

EXECUTE msdb.dbo.rds_sysmail_control stop;
GO
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deleting messages

Instance store support for tempdb

All content copied from https://docs.aws.amazon.com/.
