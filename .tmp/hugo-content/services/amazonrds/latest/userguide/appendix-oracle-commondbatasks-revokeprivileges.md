# Revoking SELECT or EXECUTE privileges on SYS objects

To revoke privileges on a single object, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.revoke_sys_object`. The procedure only
revokes privileges that the master account has already been granted through a role
or direct grant.

The `revoke_sys_object` procedure has the following parameters.

Parameter nameData typeDefaultRequiredDescription

`p_obj_name`

varchar2

—

Yes

The name of the object to revoke privileges for. The object
can be a directory, function, package, procedure, sequence,
table, or view. Object names must be spelled exactly as they
appear in `DBA_OBJECTS`. Most system objects are
defined in upper case, so we recommend you try that first.

`p_revokee`

varchar2

—

Yes

The name of the object to revoke privileges for. The object
can be a schema or a role.

`p_privilege`

varchar2

null

Yes

—

The following example revokes select privileges on an object named
`V_$SESSION` from a user named `USER1`.

```sql

begin
    rdsadmin.rdsadmin_util.revoke_sys_object(
        p_obj_name  => 'V_$SESSION',
        p_revokee   => 'USER1',
        p_privilege => 'SELECT');
end;
/
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Granting SELECT or EXECUTE privileges to SYS objects

RDS\_X$ view
tasks

All content copied from https://docs.aws.amazon.com/.
