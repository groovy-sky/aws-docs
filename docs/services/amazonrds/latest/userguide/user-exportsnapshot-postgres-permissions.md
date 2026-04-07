# Troubleshooting RDS for PostgreSQL permissions errors

When exporting PostgreSQL databases to Amazon S3, you might see a
`PERMISSIONS_DO_NOT_EXIST` error stating that certain tables were
skipped. This error usually occurs when the superuser, which you specified when creating
the database, doesn't have permissions to access those tables.

To fix this error, run the following command:

```nohighlight

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA schema_name TO superuser_name
```

For more information on superuser privileges, see [Master user account privileges](usingwithrds-masteraccounts.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Failure messages

File naming conventions
