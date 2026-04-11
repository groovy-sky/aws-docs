This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem MetadataConfiguration

The configuration that allows you to specify the performance of metadata operations for an FSx for Lustre file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iops" : Integer,
  "Mode" : String
}

```

### YAML

```yaml

  Iops: Integer
  Mode: String

```

## Properties

`Iops`

The number of Metadata IOPS provisioned for the file system.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Specifies whether the file system is using the AUTOMATIC setting of metadata IOPS or if it is using a USER\_PROVISIONED value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LustreConfiguration

NfsExports

All content copied from https://docs.aws.amazon.com/.
