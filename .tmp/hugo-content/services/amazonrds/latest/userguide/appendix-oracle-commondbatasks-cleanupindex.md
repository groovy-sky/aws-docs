# Cleaning up interrupted online index builds

To clean up failed online index builds, use the Amazon RDS procedure
`rdsadmin.rdsadmin_dbms_repair.online_index_clean`.

The `online_index_clean` procedure has the following parameters.

Parameter nameData typeDefaultRequiredDescription

`object_id`

binary\_integer

`ALL_INDEX_ID`

No

The object ID of the index. Typically, you can use the object
ID from the ORA-08104 error text.

`wait_for_lock`

binary\_integer

`rdsadmin.rdsadmin_dbms_repair.lock_wait`

No

Specify `rdsadmin.rdsadmin_dbms_repair.lock_wait`,
the default, to try to get a lock on the underlying object and
retry until an internal limit is reached if the lock
fails.

Specify `rdsadmin.rdsadmin_dbms_repair.lock_nowait`
to try to get a lock on the underlying object but not retry if
the lock fails.

The following example cleans up a failed online index build:

```sql

declare
  is_clean boolean;
begin
  is_clean := rdsadmin.rdsadmin_dbms_repair.online_index_clean(
    object_id     => 1234567890,
    wait_for_lock => rdsadmin.rdsadmin_dbms_repair.lock_nowait
  );
end;
/
```

For more information, see [ONLINE\_INDEX\_CLEAN function](https://docs.oracle.com/database/121/ARPLS/d_repair.htm) in the Oracle documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Disabling auditing for the SYS.AUD$ table

Skipping corrupt blocks

All content copied from https://docs.aws.amazon.com/.
