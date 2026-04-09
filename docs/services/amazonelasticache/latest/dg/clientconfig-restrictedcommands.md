# Restricted commands

To deliver a managed service experience, ElastiCache restricts access to certain cache engine-specific commands that require advanced privileges.
For clusters running Redis OSS, the following commands are unavailable:

- `bgrewriteaof`

- `bgsave`

- `config`

- `debug`

- `migrate`

- `replicaof`

- `save`

- `slaveof`

- `shutdown`

- `sync`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuring clients

Finding node endpoints and port numbers

All content copied from https://docs.aws.amazon.com/.
