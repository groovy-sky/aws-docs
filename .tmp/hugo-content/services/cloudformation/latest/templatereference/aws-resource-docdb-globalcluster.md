This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::GlobalCluster

A data type representing an Amazon DocumentDB global cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDB::GlobalCluster",
  "Properties" : {
      "DeletionProtection" : Boolean,
      "Engine" : String,
      "EngineVersion" : String,
      "GlobalClusterIdentifier" : String,
      "SourceDBClusterIdentifier" : String,
      "StorageEncrypted" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DocDB::GlobalCluster
Properties:
  DeletionProtection: Boolean
  Engine: String
  EngineVersion: String
  GlobalClusterIdentifier: String
  SourceDBClusterIdentifier: String
  StorageEncrypted: Boolean
  Tags:
    - Tag

```

## Properties

`DeletionProtection`

The deletion protection setting for the new global cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The Amazon DocumentDB database engine used by the global cluster.

_Required_: No

_Type_: String

_Allowed values_: `docdb`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

Indicates the database engine version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GlobalClusterIdentifier`

Contains a user-supplied global cluster identifier. This identifier is the unique key that identifies a global cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDBClusterIdentifier`

The Amazon Resource Name (ARN) to use as the primary cluster of the global cluster. This parameter is optional.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

The storage encryption setting for the global cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to be assigned to the Amazon DocumentDB resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-docdb-globalcluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`GlobalClusterArn`

The Amazon Resource Name (ARN) for the global cluster.

`GlobalClusterResourceId`

The AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS customer master key (CMK) for the cluster is accessed.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DocDB::EventSubscription

Tag

All content copied from https://docs.aws.amazon.com/.
