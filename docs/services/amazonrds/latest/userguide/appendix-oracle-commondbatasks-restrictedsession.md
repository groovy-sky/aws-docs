# Enabling and disabling restricted sessions

To enable and disable restricted sessions, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.restricted_session`. The
`restricted_session` procedure has the following parameters.

Parameter nameData typeDefaultYesDescription

`p_enable`

boolean

true

No

Set to `true` to enable restricted sessions,
`false` to disable restricted sessions.

The following example shows how to enable and disable restricted sessions.

```sql

/* Verify that the database is currently unrestricted. */

SELECT LOGINS FROM V$INSTANCE;

LOGINS
-------
ALLOWED

/* Enable restricted sessions */

EXEC rdsadmin.rdsadmin_util.restricted_session(p_enable => true);

/* Verify that the database is now restricted. */

SELECT LOGINS FROM V$INSTANCE;

LOGINS
----------
RESTRICTED

/* Disable restricted sessions */

EXEC rdsadmin.rdsadmin_util.restricted_session(p_enable => false);

/* Verify that the database is now unrestricted again. */

SELECT LOGINS FROM V$INSTANCE;

LOGINS
-------
ALLOWED
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Canceling
a SQL statement in a session

Flushing the shared pool

All content copied from https://docs.aws.amazon.com/.
