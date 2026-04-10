This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume ClientConfigurations

Specifies who can mount an OpenZFS file system and the options available while
mounting the file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Clients" : String,
  "Options" : [ String, ... ]
}

```

### YAML

```yaml

  Clients: String
  Options:
    - String

```

## Properties

`Clients`

A value that specifies who can mount the file system. You can provide a wildcard
character ( `*`), an IP address ( `0.0.0.0`), or a CIDR address
( `192.0.2.0/24`). By default, Amazon FSx uses the wildcard
character when specifying the client.

_Required_: Yes

_Type_: String

_Pattern_: `^[ -~]{1,128}$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

The options to use when mounting the file system. For a list of options that you can
use with Network File System (NFS), see the [exports(5) - Linux man page](https://linux.die.net/man/5/exports). When
choosing your options, consider the following:

- `crossmnt` is used by default. If you don't specify
`crossmnt` when changing the client configuration, you won't be
able to see or access snapshots in your file system's snapshot directory.

- `sync` is used by default. If you instead specify
`async`, the system acknowledges writes before writing to disk.
If the system crashes before the writes are finished, you lose the unwritten
data.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutocommitPeriod

NfsExports

All content copied from https://docs.aws.amazon.com/.
