This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ScheduledAction PauseClusterMessage

Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterIdentifier" : String
}

```

### YAML

```yaml

  ClusterIdentifier: String

```

## Properties

`ClusterIdentifier`

The identifier of the cluster to be paused.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Redshift::ScheduledAction

ResizeClusterMessage

All content copied from https://docs.aws.amazon.com/.
