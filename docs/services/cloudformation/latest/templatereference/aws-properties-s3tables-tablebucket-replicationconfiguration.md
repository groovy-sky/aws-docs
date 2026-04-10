This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TableBucket ReplicationConfiguration

The `ReplicationConfiguration` property type specifies Property description not available. for an [AWS::S3Tables::TableBucket](aws-resource-s3tables-tablebucket.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Role" : String,
  "Rules" : [ ReplicationRule, ... ]
}

```

### YAML

```yaml

  Role: String
  Rules:
    - ReplicationRule

```

## Properties

`Role`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

Property description not available.

_Required_: Yes

_Type_: Array of [ReplicationRule](aws-properties-s3tables-tablebucket-replicationrule.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricsConfiguration

ReplicationDestination

All content copied from https://docs.aws.amazon.com/.
