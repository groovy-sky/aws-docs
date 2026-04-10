This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume NfsExports

The configuration object for mounting a Network File System (NFS) file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientConfigurations" : [ ClientConfigurations, ... ]
}

```

### YAML

```yaml

  ClientConfigurations:
    - ClientConfigurations

```

## Properties

`ClientConfigurations`

A list of configuration objects that contain the client and options for mounting the
OpenZFS file system.

_Required_: Yes

_Type_: [Array](aws-properties-fsx-volume-clientconfigurations.md) of [ClientConfigurations](aws-properties-fsx-volume-clientconfigurations.md)

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientConfigurations

OntapConfiguration

All content copied from https://docs.aws.amazon.com/.
