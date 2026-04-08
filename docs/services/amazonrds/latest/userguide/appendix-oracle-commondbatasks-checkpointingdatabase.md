# Checkpointing a database

To checkpoint the database, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.checkpoint`. The `checkpoint`
procedure has no parameters.

The following example checkpoints the database.

```sql

EXEC rdsadmin.rdsadmin_util.checkpoint;
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Working with Oracle external tables

Setting distributed recovery

All content copied from https://docs.aws.amazon.com/.
