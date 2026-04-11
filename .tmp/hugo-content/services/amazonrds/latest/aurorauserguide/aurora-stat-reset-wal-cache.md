# aurora\_stat\_reset\_wal\_cache

Resets the counter for logical wal cache.

## Syntax

To reset a specific slot

```nohighlight

SELECT * FROM aurora_stat_reset_wal_cache('slot_name')
```

To reset all slots

```nohighlight

SELECT * FROM aurora_stat_reset_wal_cache(NULL)
```

## Arguments

`NULL` or `slot_name`

## Return type

Status message, text string

- Reset the logical wal cache counter – Success message. This text is
returned when the function succeeds.

- Replication slot not found. Please try again. – Failure message.
This text is returned when the function doesn't succeed.

## Usage notes

This function is available for the following versions.

- Aurora PostgreSQL 14.5 and higher

- Aurora PostgreSQL version 13.8 and higher

- Aurora PostgreSQL version 12.12 and higher

- Aurora PostgreSQLversion 11.17 and higher

## Examples

The following example uses the `aurora_stat_reset_wal_cache` function
to reset a slot named `test_results`, and then tries to reset a slot that
doesn't exist.

```nohighlight

=> SELECT *
     FROM aurora_stat_reset_wal_cache('test_slot');
aurora_stat_reset_wal_cache
--------------------------------------
 Reset the logical wal cache counter.
(1 row)
=> SELECT *
     FROM aurora_stat_reset_wal_cache('slot-not-exist');
aurora_stat_reset_wal_cache
-----------------------------------------------
 Replication slot not found. Please try again.
(1 row)
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

aurora\_stat\_plans

aurora\_stat\_resource\_usage

All content copied from https://docs.aws.amazon.com/.
