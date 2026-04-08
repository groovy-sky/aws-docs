# Collecting and maintaining the Global Status History

Amazon RDS provides a set of procedures that take snapshots of the values of status variables over time and write them to a table,
along with any changes since the last snapshot. This infrastructure is called Global Status History. For more information, see
[Managing the Global\
Status History](../userguide/appendix-mysql-commondbatasks.md#Appendix.MySQL.CommonDBATasks.GoSH).

The following stored procedures manage how the Global Status History is collected and maintained.

###### Topics

- [mysql.rds\_collect\_global\_status\_history](#mysql_rds_collect_global_status_history)

- [mysql.rds\_disable\_gsh\_collector](#mysql_rds_disable_gsh_collector)

- [mysql.rds\_disable\_gsh\_rotation](#mysql_rds_disable_gsh_rotation)

- [mysql.rds\_enable\_gsh\_collector](#mysql_rds_enable_gsh_collector)

- [mysql.rds\_enable\_gsh\_rotation](#mysql_rds_enable_gsh_rotation)

- [mysql.rds\_rotate\_global\_status\_history](#mysql_rds_rotate_global_status_history)

- [mysql.rds\_set\_gsh\_collector](#mysql_rds_set_gsh_collector)

- [mysql.rds\_set\_gsh\_rotation](#mysql_rds_set_gsh_rotation)

## mysql.rds\_collect\_global\_status\_history

Takes a snapshot on demand for the Global Status History.

### Syntax

```sql

CALL mysql.rds_collect_global_status_history;
```

## mysql.rds\_disable\_gsh\_collector

Turns off snapshots taken by the Global Status History.

### Syntax

```sql

CALL mysql.rds_disable_gsh_collector;
```

## mysql.rds\_disable\_gsh\_rotation

Turns off rotation of the `mysql.global_status_history` table.

### Syntax

```sql

CALL mysql.rds_disable_gsh_rotation;
```

## mysql.rds\_enable\_gsh\_collector

Turns on the Global Status History to take default snapshots at intervals specified by
`rds_set_gsh_collector`.

### Syntax

```sql

CALL mysql.rds_enable_gsh_collector;
```

## mysql.rds\_enable\_gsh\_rotation

Turns on rotation of the contents of the `mysql.global_status_history` table to
`mysql.global_status_history_old` at intervals specified by `rds_set_gsh_rotation`.

### Syntax

```sql

CALL mysql.rds_enable_gsh_rotation;
```

## mysql.rds\_rotate\_global\_status\_history

Rotates the contents of the `mysql.global_status_history` table to `mysql.global_status_history_old` on
demand.

### Syntax

```sql

CALL mysql.rds_rotate_global_status_history;
```

## mysql.rds\_set\_gsh\_collector

Specifies the interval, in minutes, between snapshots taken by the Global Status History.

### Syntax

```sql

CALL mysql.rds_set_gsh_collector(intervalPeriod);
```

### Parameters

`intervalPeriod`

The interval, in minutes, between snapshots. Default value is
`5`.

## mysql.rds\_set\_gsh\_rotation

Specifies the interval, in days, between rotations of the `mysql.global_status_history` table.

### Syntax

```sql

CALL mysql.rds_set_gsh_rotation(intervalPeriod);
```

### Parameters

`intervalPeriod`

The interval, in days, between table rotations. Default value is
`7`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Stored procedure reference

Configuring, starting, and stopping binary log (binlog) replication

All content copied from https://docs.aws.amazon.com/.
