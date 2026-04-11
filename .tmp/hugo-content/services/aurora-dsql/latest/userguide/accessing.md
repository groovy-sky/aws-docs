# Accessing Aurora DSQL with PostgreSQL-compatible clients

Aurora DSQL uses the [PostgreSQL\
wire protocol](https://www.postgresql.org/docs/current/protocol.html). You can connect to PostgreSQL using a variety of tools and clients, such as
AWS CloudShell, psql, DBeaver, and DataGrip. The following table summarizes how Aurora DSQL maps common
PostgreSQL connection parameters:

PostgreSQLAurora DSQLNotesRole (also known as User or Group)Database RoleAurora DSQL creates a role for you named `admin`. When you create custom database
roles, you must use the `admin` role to associate them with IAM roles for
authenticating when connecting to your cluster. For more information, see
[Using database roles and IAM authentication](using-database-and-iam-roles.md).Host (also known as hostname or hostspec)Cluster EndpointAurora DSQL single-Region clusters provide a single managed endpoint and automatically
redirect traffic if there is unavailability within the Region.PortN/A – use default `5432`This is the PostgreSQL default.Database (dbname)use `postgres`Aurora DSQL creates this database for you when you create the cluster.SSL ModeSSL is always enabled server-sideIn Aurora DSQL, Aurora DSQL supports the `require` SSL Mode. Connections without SSL
are rejected by Aurora DSQL.PasswordAuthentication TokenAurora DSQL requires temporary authentication tokens instead of long-lived passwords. To
learn more, see [Generating an authentication token in Amazon Aurora DSQL](section-authentication-token.md).

When connecting, Aurora DSQL requires a signed IAM [authentication\
token](section-authentication-token.md) in place of a traditional password. These temporary tokens are generated using
AWS Signature Version 4 and are used only during connection establishment. Once connected, the
session remains active until it ends or the client disconnects.

If you attempt to open a new session with an expired token, the connection request fails and
a new token must be generated. For more information, see [Generating an authentication token in Amazon Aurora DSQL](section-authentication-token.md).

## Access Aurora DSQL using SQL clients

Aurora DSQL supports multiple PostgreSQL-compatible clients for connecting to your cluster. The
following sections describe how to connect using PostgreSQL with AWS CloudShell or your local command
line, as well as GUI-based tools like DBeaver and JetBrains DataGrip. Each client requires a
valid authentication token as described in the previous section.

###### Topics

- [Use DBeaver to access Aurora DSQL](accessing-dbeaver.md)

- [Use JetBrains DataGrip to access Aurora DSQL](accessing-datagrip.md)

- [Use the PostgreSQL interactive terminal (psql) to access Aurora DSQL](accessing-psql.md)

- [Use Aurora DSQL driver for SQLTools](accessing-vscode.md)

- [Troubleshooting](#accessing-troubleshooting)

## Troubleshooting

**Authentication credentials expiration for the SQL**
**Clients**

Established sessions remain authenticated for a maximum of 1 hour or until an explicit
disconnect or a client-side timeout takes place. If new connections need to be established, a
new authentication token must be generated and provided in the **Password** field of the connection. Trying to open a new session (for example, to list
new tables, or open a new SQL console) forces a new authentication attempt. If the authentication
token configured in the **Connection** settings is no longer valid,
that new session will fail and all previously opened sessions will become invalid. Keep this in
mind when choosing the duration of your IAM authentication token with the
`expires-in` option, which can be set to 15 minutes by default and can be set to a
maximum value of seven days.

Additionally, see the [Troubleshooting](troubleshooting.md) section of the
Aurora DSQL documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rust connector

DBeaver

All content copied from https://docs.aws.amazon.com/.
