This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Snapshot

A snapshot object that contains databases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RedshiftServerless::Snapshot",
  "Properties" : {
      "NamespaceName" : String,
      "RetentionPeriod" : Integer,
      "SnapshotName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RedshiftServerless::Snapshot
Properties:
  NamespaceName: String
  RetentionPeriod: Integer
  SnapshotName: String
  Tags:
    - Tag

```

## Properties

`NamespaceName`

The name of the namepsace.

_Required_: No

_Type_: String

_Pattern_: `^(?=^[a-z0-9-]+$).{3,64}$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RetentionPeriod`

The retention period of the snapshot created by the scheduled action.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotName`

The name of the snapshot.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^[a-z0-9-]+$).{3,64}$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of [Tag objects](https://docs.aws.amazon.com/redshift-serverless/latest/APIReference/API_Tag.html) to associate with the snapshot.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshiftserverless-snapshot-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`OwnerAccount`

The owner AWS; account of the snapshot.

`Snapshot.AdminUsername`

The username of the database within a snapshot.

`Snapshot.KmsKeyId`

The unique identifier of the KMS key used to encrypt the snapshot.

`Snapshot.NamespaceArn`

The Amazon Resource Name (ARN) of the namespace the snapshot was created from.

`Snapshot.NamespaceName`

The name of the namepsace.

`Snapshot.OwnerAccount`

The owner AWS; account of the snapshot.

`Snapshot.RetentionPeriod`

The retention period of the snapshot created by the scheduled action.

`Snapshot.SnapshotArn`

The Amazon Resource Name (ARN) of the snapshot.

`Snapshot.SnapshotCreateTime`

The timestamp of when the snapshot was created.

`Snapshot.SnapshotName`

The name of the snapshot.

`Snapshot.Status`

The status of the snapshot.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Snapshot
