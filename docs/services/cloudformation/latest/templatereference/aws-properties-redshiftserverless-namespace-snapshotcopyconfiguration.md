This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Namespace SnapshotCopyConfiguration

The object that you configure to copy snapshots from one namespace to a namespace in another AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationKmsKeyId" : String,
  "DestinationRegion" : String,
  "SnapshotRetentionPeriod" : Integer
}

```

### YAML

```yaml

  DestinationKmsKeyId: String
  DestinationRegion: String
  SnapshotRetentionPeriod: Integer

```

## Properties

`DestinationKmsKeyId`

The ID of the KMS key to use to encrypt your snapshots in the destination AWS Region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationRegion`

The destination AWS Region to copy snapshots to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotRetentionPeriod`

The retention period of snapshots that are copied to the destination AWS Region.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Namespace

Tag
