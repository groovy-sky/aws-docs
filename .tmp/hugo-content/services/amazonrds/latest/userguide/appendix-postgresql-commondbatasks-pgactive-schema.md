# Understanding the pgactive schema

The pgactive schema manages active-active replication in RDS for PostgreSQL. This schema
contains tables that store replication configuration and status information.

###### Note

The pgactive schema is evolving and subject to change. Don't modify the data in these
tables directly.

The key tables in the pgactive schema include:

- `pgactive_nodes` – Stores information about nodes in the
active-active replication group.

- `pgactive_connections` – Stores connection details for each
node.

## pgactive\_nodes

The pgactive\_nodes stores information about the nodes participating in the active-active
replication group.

Column

Type

Collation

Nullable

Default

node\_sysidtext–not null–node\_timelineoid–not null–node\_dboidoid–not null–node\_statuschar–not null–node\_nametext–not null–node\_dsntext–not null–node\_init\_from\_dsntext–not null–node\_read\_onlyboolean––falsenode\_seq\_idsmallint–not null–

**node\_sysid**

Unique ID for a node, generated during `pgactive_create_group` or
`pgactive_join_group`

**node\_status**

Readiness of the node:

- **b** \- beginning setup

- **i** \- initializing

- **c** \- catchup

- **o** \- creating outbound slots

- **r** \- ready

- **k** \- killed

###### Note

This column doesn't indicate if a node is connected or disconnected.

**node\_name**

User-provided unique node name.

**node\_dsn**

Connection string or user mapping name

**node\_init\_from\_dsn**

DSN from which this node was created.

## pgactive\_connection

The pgactive\_connections stores connection details for each node.

ColumnTypeCollationNullableDefaultconn\_sysidtextnonenot nullnoneconn\_timelineoidnonenot nullnoneconn\_dboidoidnonenot nullnoneconn\_dsntextnonenot nullnoneconn\_apply\_delayintegernonenonenoneconn\_replication\_setstextnonenonenone

conn\_sysid

Node identifier for the node this entry refers to.

conn\_dsn

Same as pgactive.pgactive\_nodes `node_dsn`.

conn\_apply\_delay

If set, milliseconds to wait before applying each transaction from the remote
node. Mainly for debugging. If null, the global default applies.

## Working with replication sets

Replication sets determine which tables to include or exclude from replication
operations. By default, all tables are replicated unless you specify otherwise using the
following functions:

- `pgactive_exclude_table_replication_set()` \- Excludes specified tables
from replication

- `pgactive_include_table_replication_set()` \- Includes specified tables in
replication

###### Note

Before you configure replication sets, consider the following:

- You can configure table inclusion or exclusion only after running
`pgactive_create_group()` but before
`pgactive_join_group()`.

- After you use `pgactive_exclude_table_replication_set()`, you can't use
`pgactive_include_table_replication_set()`.

- After you use `pgactive_include_table_replication_set()`, you can't use
`pgactive_exclude_table_replication_set()`.

The system handles newly created tables differently based on your initial
configuration:

- If you excluded tables: Any new tables created after
`pgactive_join_group()` are automatically included in replication

- If you included tables: Any new tables created after
`pgactive_join_group()` are automatically excluded from replication.

To view the replication set configuration for a specific table, use the
`pgactive.pgactive_get_table_replication_sets()` function.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Understanding active-active conflicts

pgactive functions reference

All content copied from https://docs.aws.amazon.com/.
