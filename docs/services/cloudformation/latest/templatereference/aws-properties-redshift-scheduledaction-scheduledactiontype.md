This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ScheduledAction ScheduledActionType

The action type that specifies an Amazon Redshift API operation that is supported by the Amazon Redshift scheduler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PauseCluster" : PauseClusterMessage,
  "ResizeCluster" : ResizeClusterMessage,
  "ResumeCluster" : ResumeClusterMessage
}

```

### YAML

```yaml

  PauseCluster:
    PauseClusterMessage
  ResizeCluster:
    ResizeClusterMessage
  ResumeCluster:
    ResumeClusterMessage

```

## Properties

`PauseCluster`

An action that runs a `PauseCluster` API operation.

_Required_: No

_Type_: [PauseClusterMessage](aws-properties-redshift-scheduledaction-pauseclustermessage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResizeCluster`

An action that runs a `ResizeCluster` API operation.

_Required_: No

_Type_: [ResizeClusterMessage](aws-properties-redshift-scheduledaction-resizeclustermessage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResumeCluster`

An action that runs a `ResumeCluster` API operation.

_Required_: No

_Type_: [ResumeClusterMessage](aws-properties-redshift-scheduledaction-resumeclustermessage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResumeClusterMessage

Next

All content copied from https://docs.aws.amazon.com/.
