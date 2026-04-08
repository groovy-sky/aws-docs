# aurora\_stat\_bgwriter

`aurora_stat_bgwriter` is a statistics view showing information about
Optimized Reads cache writes.

## Syntax

```nohighlight

aurora_stat_bgwriter()
```

## Arguments

None

## Return type

SETOF record with all `pg_stat_bgwriter` columns and the following
additional columns. For more information on `pg_stat_bgwriter` columns,
see [`pg_stat_bgwriter`](https://www.postgresql.org/docs/current/monitoring-stats.html).

You can reset stats for this function using
`pg_stat_reset_shared("bgwriter")`.

- `orcache_blks_written` – Total number of optimized reads
cache data blocks written.

- `orcache_blk_write_time` – If
`track_io_timing` is enabled, it tracks the total time spent
writing optimized reads cache data file blocks, in milliseconds. For more
information, see [track\_io\_timing](https://www.postgresql.org/docs/current/runtime-config-statistics.html).

## Usage notes

This function is available in the following Aurora PostgreSQL versions:

- 15.4 and higher 15 versions

- 14.9 and higher 14 versions

## Examples

```nohighlight

=> select * from aurora_stat_bgwriter();
-[ RECORD 1 ]-----------------+-----------
orcache_blks_written          | 246522
orcache_blk_write_time        | 339276.404

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

aurora\_stat\_backend\_waits

aurora\_stat\_database

All content copied from https://docs.aws.amazon.com/.
