# Viewing audit logs

Your audit logs are stored in `D:\rdsdbdata\SQLAudit`.

After SQL Server finishes writing to an audit log file—when the file reaches its
size limit—Amazon RDS uploads the file to your S3 bucket. If retention is enabled,
Amazon RDS moves the file into the retention folder:
`D:\rdsdbdata\SQLAudit\transmitted`.

For information about configuring retention, see [Adding SQL Server Audit to the DB instance options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.Audit.Adding.html).

Audit records are kept on the DB instance until the audit log file is uploaded. You can
view the audit records by running the following command.

```sql

SELECT   *
	FROM     msdb.dbo.rds_fn_get_audit_file
	             ('D:\rdsdbdata\SQLAudit\*.sqlaudit'
	             , default
	             , default )
```

You
can use the same command to view audit records in your retention folder by changing the filter
to
`D:\rdsdbdata\SQLAudit\transmitted\*.sqlaudit`.

```sql

SELECT   *
	FROM     msdb.dbo.rds_fn_get_audit_file
	             ('D:\rdsdbdata\SQLAudit\transmitted\*.sqlaudit'
	             , default
	             , default )
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using SQL Server Audit

Configuring an S3 bucket
