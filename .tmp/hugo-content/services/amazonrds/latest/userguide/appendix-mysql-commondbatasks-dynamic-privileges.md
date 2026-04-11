# Dynamic privileges for RDS for MySQL

Dynamic privileges are MySQL privileges that you can explicitly grant by using the
`GRANT` statement. Depending on your version of RDS for MySQL, RDS allows
you to grant only specific dynamic privileges. RDS disallows some of these privileges
because they can interfere with the specific database operations, such as replication
and backup.

The following table shows which of these privileges you can grant for different MySQL
versions. If you are upgrading from a MySQL version lower than 8.0.36 to version 8.0.36
or higher, you might have to update your application code if granting a particular
privilege is no longer allowed.

PrivilegeMySQL 8.0.35 and lowerMySQL 8.0.36 and higher minor versionsMySQL 8.4.3 and higher

[ALLOW\_NONEXISTENT\_DEFINER](https://dev.mysql.com/doc/refman/8.4/en/privileges-provided.html)

Not available

Not available

Disallowed

[APPLICATION\_PASSWORD\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[AUDIT\_ABORT\_EXEMPT](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[AUDIT\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[AUTHENTICATION\_POLICY\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[BACKUP\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[BINLOG\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[BINLOG\_ENCRYPTION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[CLONE\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[CONNECTION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[ENCRYPTION\_KEY\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[FIREWALL\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[FIREWALL\_EXEMPT](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[FIREWALL\_USER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[FLUSH\_OPTIMIZER\_COSTS](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[FLUSH\_PRIVILEGES](https://dev.mysql.com/doc/refman/8.4/en/privileges-provided.html)

Not available

Not available

Allowed

[FLUSH\_STATUS](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[FLUSH\_TABLES](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[FLUSH\_USER\_RESOURCES](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[GROUP\_REPLICATION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[GROUP\_REPLICATION\_STREAM](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[INNODB\_REDO\_LOG\_ARCHIVE](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[INNODB\_REDO\_LOG\_ENABLE](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[MASKING\_DICTIONARIES\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[NDB\_STORED\_USER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[OPTIMIZE\_LOCAL\_TABLE](https://dev.mysql.com/doc/refman/8.4/en/privileges-provided.html)

Not available

Not available

Disallowed

[PASSWORDLESS\_USER\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[PERSIST\_RO\_VARIABLES\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[REPLICATION\_APPLIER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[REPLICATION\_SLAVE\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[RESOURCE\_GROUP\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[RESOURCE\_GROUP\_USER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[ROLE\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[SENSITIVE\_VARIABLES\_OBSERVER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[SERVICE\_CONNECTION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[SESSION\_VARIABLES\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[SET\_ANY\_DEFINER](https://dev.mysql.com/doc/refman/8.4/en/privileges-provided.html)

Not available

Not available

Allowed

[SET\_USER\_ID](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Not available

[SHOW\_ROUTINE](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[SKIP\_QUERY\_REWRITE](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[SYSTEM\_USER](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[SYSTEM\_VARIABLES\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[TABLE\_ENCRYPTION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[TELEMETRY\_LOG\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Disallowed

Disallowed

[TP\_CONNECTION\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[TRANSACTION\_GTID\_TAG](https://dev.mysql.com/doc/refman/8.4/en/privileges-provided.html)

Not available

Not available

Disallowed

[VERSION\_TOKEN\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Disallowed

Disallowed

Disallowed

[XA\_RECOVER\_ADMIN](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html)

Allowed

Allowed

Allowed

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Role-based privilege model

Ending a session or query

All content copied from https://docs.aws.amazon.com/.
