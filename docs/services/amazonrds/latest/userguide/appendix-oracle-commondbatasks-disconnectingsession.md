# Disconnecting a session

To disconnect the current session by ending the dedicated server process, use the
Amazon RDS procedure `rdsadmin.rdsadmin_util.disconnect`. The
`disconnect` procedure has the following parameters.

Parameter nameData typeDefaultRequiredDescription

`sid`

number

—

Yes

The session identifier.

`serial`

number

—

Yes

The serial number of the session.

`method`

varchar

'IMMEDIATE'

No

Valid values are `'IMMEDIATE'` or
`'POST_TRANSACTION'`.

The following example disconnects a session.

```sql

begin
    rdsadmin.rdsadmin_util.disconnect(
        sid    => sid,
        serial => serial_number);
end;
/
```

To get the session identifier and the session serial number, query the
`V$SESSION` view. The following example gets all sessions for the
user `AWSUSER`.

```sql

SELECT SID, SERIAL#, STATUS FROM V$SESSION WHERE USERNAME = 'AWSUSER';
```

The database must be open to use this method. For more information about
disconnecting a session, see [ALTER SYSTEM](http://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_2014.htm) in the Oracle documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

System
tasks

Terminating a session

All content copied from https://docs.aws.amazon.com/.
