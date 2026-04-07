This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobQueue JobStateTimeLimitAction

Specifies an action that AWS Batch will take after the job has remained at the head of the queue in the specified state for longer than the specified time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "MaxTimeSeconds" : Integer,
  "Reason" : String,
  "State" : String
}

```

### YAML

```yaml

  Action: String
  MaxTimeSeconds: Integer
  Reason: String
  State: String

```

## Properties

`Action`

The action to take when a job is at the head of the job queue in the specified state for the specified period of
time. For job queues connected to a `ECS`, `FARGATE` or `EKS` compute environment, the only supported value is `CANCEL`, which will cancel the job.
For job queues connected to a `SAGEMAKER_TRAINING` service environment, the only supported value is `TERMINATE`, which will terminate the job.

_Required_: Yes

_Type_: String

_Allowed values_: `CANCEL | TERMINATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTimeSeconds`

The approximate amount of time, in seconds, that must pass with the job in the specified state before the action
is taken. The minimum value is 600 (10 minutes) and the maximum value is 86,400 (24 hours).

_Required_: Yes

_Type_: Integer

_Minimum_: `600`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Reason`

The reason to log for the action being taken.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the job needed to trigger the action. The only supported value is `RUNNABLE`.

_Required_: Yes

_Type_: String

_Allowed values_: `RUNNABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComputeEnvironmentOrder

ServiceEnvironmentOrder
