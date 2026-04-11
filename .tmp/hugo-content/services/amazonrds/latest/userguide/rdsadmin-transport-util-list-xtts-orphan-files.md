# Listing orphaned files after a tablespace import

Use the `rdsadmin.rdsadmin_transport_util.list_xtts_orphan_files` procedure to
list data files that were orphaned after a tablespace import. After you identify the data
files, you can delete them by calling
`rdsadmin.rdsadmin_transport_util.cleanup_incomplete_xtts_import`.

## Syntax

```sql

FUNCTION list_xtts_orphan_files RETURN xtts_orphan_files_list_t PIPELINED;
```

## Examples

The following example runs the procedure
`rdsadmin.rdsadmin_transport_util.list_xtts_orphan_files`. The output
shows two data files that are orphaned.

```sql

SQL> SELECT * FROM TABLE(rdsadmin.rdsadmin_transport_util.list_xtts_orphan_files);

FILENAME       FILESIZE
-------------- ---------
datafile_7.dbf 104865792
datafile_8.dbf 104865792
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Importing transportable tablespace metadata

Deleting orphaned data files

All content copied from https://docs.aws.amazon.com/.
