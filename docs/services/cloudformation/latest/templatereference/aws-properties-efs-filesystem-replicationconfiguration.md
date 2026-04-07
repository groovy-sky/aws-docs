This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem ReplicationConfiguration

Describes the replication configuration for a specific file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destinations" : [ ReplicationDestination, ... ]
}

```

### YAML

```yaml

  Destinations:
    - ReplicationDestination

```

## Properties

`Destinations`

An array of destination objects. Only one destination object is supported.

_Required_: No

_Type_: Array of [ReplicationDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-replicationdestination.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LifecyclePolicy

ReplicationDestination
