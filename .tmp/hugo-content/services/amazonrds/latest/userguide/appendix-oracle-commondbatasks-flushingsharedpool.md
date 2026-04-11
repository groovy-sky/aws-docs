# Flushing the shared pool

To flush the shared pool, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.flush_shared_pool`. The
`flush_shared_pool` procedure has no parameters.

The following example flushes the shared pool.

```sql

EXEC rdsadmin.rdsadmin_util.flush_shared_pool;
```

## Flushing the buffer cache

To flush the buffer cache, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.flush_buffer_cache`. The
`flush_buffer_cache` procedure has no parameters.

The following example flushes the buffer cache.

```sql

EXEC rdsadmin.rdsadmin_util.flush_buffer_cache;
```

## Flushing the database smart flash cache

To flush the database smart flash cache, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.flush_flash_cache`. The `flush_flash_cache` procedure has
no parameters. The following example flushes the database smart flash cache.

```

EXEC rdsadmin.rdsadmin_util.flush_flash_cache;
```

For more information about using the database smart flash cache with RDS for Oracle, see [Storing temporary data in an RDS for Oracle instance store](chap-oracle-advanced-features-instance-store.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Enabling and disabling restricted sessions

Granting SELECT or EXECUTE privileges to SYS objects

All content copied from https://docs.aws.amazon.com/.
