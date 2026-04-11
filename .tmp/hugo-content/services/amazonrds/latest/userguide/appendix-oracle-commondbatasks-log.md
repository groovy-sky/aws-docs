# Performing common log-related tasks for Oracle DB instances

Following, you can find how to perform certain common DBA tasks related to logging on
your Amazon RDS DB instances running Oracle. To deliver a managed service experience, Amazon RDS
doesn't provide shell access to DB instances, and restricts access to certain system
procedures and tables that require advanced privileges.

For more information, see [Amazon RDS for Oracle database log files](user-logaccess-concepts-oracle.md).

###### Topics

- [Setting force logging](#Appendix.Oracle.CommonDBATasks.SettingForceLogging)

- [Setting supplemental logging](#Appendix.Oracle.CommonDBATasks.AddingSupplementalLogging)

- [Switching online log files](#Appendix.Oracle.CommonDBATasks.SwitchingLogfiles)

- [Adding online redo logs](#Appendix.Oracle.CommonDBATasks.RedoLogs)

- [Dropping online redo logs](#Appendix.Oracle.CommonDBATasks.DroppingRedoLogs)

- [Resizing online redo logs](appendix-oracle-commondbatasks-resizingredologs.md)

- [Retaining archived redo logs](appendix-oracle-commondbatasks-retainredologs.md)

- [Accessing online and archived redo logs](appendix-oracle-commondbatasks-log-download.md)

- [Downloading archived redo logs from Amazon S3](appendix-oracle-commondbatasks-download-redo-logs.md)

## Setting force logging

In force logging mode, Oracle logs all changes to the database except changes in
temporary tablespaces and temporary segments ( `NOLOGGING` clauses are
ignored). For more information, see [Specifying FORCE LOGGING mode](https://docs.oracle.com/cd/E11882_01/server.112/e25494/create.htm) in the Oracle documentation.

To set force logging, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.force_logging`. The
`force_logging` procedure has the following parameters.

Parameter nameData typeDefaultYesDescription

`p_enable`

boolean

true

No

Set to `true` to put the database in force logging
mode, `false` to remove the database from force
logging mode.

The following example puts the database in force logging mode.

```sql

EXEC rdsadmin.rdsadmin_util.force_logging(p_enable => true);
```

## Setting supplemental logging

If you enable supplemental logging, LogMiner has the necessary information to
support chained rows and clustered tables. For more information, see [Supplemental logging](https://docs.oracle.com/cd/E11882_01/server.112/e22490/logminer.htm) in the Oracle documentation.

Oracle Database doesn't enable supplemental logging by default. To enable and
disable supplemental logging, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.alter_supplemental_logging`. For more
information about how Amazon RDS manages the retention of archived redo logs for Oracle
DB instances, see [Retaining archived redo logs](appendix-oracle-commondbatasks-retainredologs.md).

The `alter_supplemental_logging` procedure has the following
parameters.

Parameter nameData typeDefaultRequiredDescription

`p_action`

varchar2

—

Yes

`'ADD'` to add supplemental logging,
`'DROP'` to drop supplemental logging.

`p_type`

varchar2

null

No

The type of supplemental logging. Valid values are
`'ALL'`, `'FOREIGN KEY'`,
`'PRIMARY KEY'`, `'UNIQUE'`, or
`PROCEDURAL`.

The following example enables supplemental logging.

```sql

begin
    rdsadmin.rdsadmin_util.alter_supplemental_logging(
        p_action => 'ADD');
end;
/
```

The following example enables supplemental logging for all fixed-length maximum
size columns.

```sql

begin
    rdsadmin.rdsadmin_util.alter_supplemental_logging(
        p_action => 'ADD',
        p_type   => 'ALL');
end;
/
```

The following example enables supplemental logging for primary key columns.

```sql

begin
    rdsadmin.rdsadmin_util.alter_supplemental_logging(
        p_action => 'ADD',
        p_type   => 'PRIMARY KEY');
end;
/
```

## Switching online log files

To switch log files, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.switch_logfile`. The
`switch_logfile` procedure has no parameters.

The following example switches log files.

```sql

EXEC rdsadmin.rdsadmin_util.switch_logfile;
```

## Adding online redo logs

An Amazon RDS DB instance running Oracle starts with four online redo logs, 128 MB
each. To add additional redo logs, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.add_logfile`.

The `add_logfile` procedure has the following parameters.

###### Note

The parameters are mutually exclusive.

Parameter nameData typeDefaultRequiredDescription

`bytes`

positive

null

No

The size of the log file in bytes.

Use this parameter only if the size of the log is under
2147483648 bytes (2 GiB). Otherwise, RDS issues an error. For
log sizes above this byte value, use the `p_size`
parameter instead.

`p_size`

varchar2

—

Yes

The size of the log file in kilobytes (K), megabytes (M), or
gigabytes (G).

The following command adds a 100 MB log file.

```sql

EXEC rdsadmin.rdsadmin_util.add_logfile(p_size => '100M');
```

## Dropping online redo logs

To drop redo logs, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.drop_logfile`. The `drop_logfile`
procedure has the following parameters.

Parameter nameData typeDefaultRequiredDescription

`grp`

positive

—

Yes

The group number of the log.

The following example drops the log with group number 3.

```sql

EXEC rdsadmin.rdsadmin_util.drop_logfile(grp => 3);
```

You can only drop logs that have a status of unused or inactive. The following
example gets the statuses of the logs.

```sql

SELECT GROUP#, STATUS FROM V$LOG;

GROUP#     STATUS
---------- ----------------
1          CURRENT
2          INACTIVE
3          INACTIVE
4          UNUSED
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting
the default displayed values for full redaction

Resizing online redo logs

All content copied from https://docs.aws.amazon.com/.
