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

_Type_: [PauseClusterMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-scheduledaction-pauseclustermessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResizeCluster`

An action that runs a `ResizeCluster` API operation.

_Required_: No

_Type_: [ResizeClusterMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-scheduledaction-resizeclustermessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResumeCluster`

An action that runs a `ResumeCluster` API operation.

_Required_: No

_Type_: [ResumeClusterMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-scheduledaction-resumeclustermessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResumeClusterMessage

Next
