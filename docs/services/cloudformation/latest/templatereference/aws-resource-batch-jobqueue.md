This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobQueue

The `AWS::Batch::JobQueue` resource specifies the parameters for an AWS Batch job queue definition. For more information, see [Job Queues](https://docs.aws.amazon.com/batch/latest/userguide/job_queues.html) in
the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::JobQueue",
  "Properties" : {
      "ComputeEnvironmentOrder" : [ ComputeEnvironmentOrder, ... ],
      "JobQueueName" : String,
      "JobQueueType" : String,
      "JobStateTimeLimitActions" : [ JobStateTimeLimitAction, ... ],
      "Priority" : Integer,
      "SchedulingPolicyArn" : String,
      "ServiceEnvironmentOrder" : [ ServiceEnvironmentOrder, ... ],
      "State" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Batch::JobQueue
Properties:
  ComputeEnvironmentOrder:
    - ComputeEnvironmentOrder
  JobQueueName: String
  JobQueueType: String
  JobStateTimeLimitActions:
    - JobStateTimeLimitAction
  Priority: Integer
  SchedulingPolicyArn: String
  ServiceEnvironmentOrder:
    - ServiceEnvironmentOrder
  State: String
  Tags:
    Key: Value

```

## Properties

`ComputeEnvironmentOrder`

The set of compute environments mapped to a job queue and their order relative to each
other. The job scheduler uses this parameter to determine which compute environment runs a
specific job. Compute environments must be in the `VALID` state before you can
associate them with a job queue. You can associate up to three compute environments with a job
queue. All of the compute environments must be either EC2 ( `EC2` or
`SPOT`) or Fargate ( `FARGATE` or `FARGATE_SPOT`); EC2 and
Fargate compute environments can't be mixed.

###### Note

All compute environments that are associated with a job queue must share the same
architecture. AWS Batch doesn't support mixing compute environment architecture types in a
single job queue.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobqueue-computeenvironmentorder.html) of [ComputeEnvironmentOrder](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobqueue-computeenvironmentorder.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobQueueName`

The name of the job queue. It can be up to 128 letters long. It can contain uppercase and
lowercase letters, numbers, hyphens (-), and underscores (\_).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobQueueType`

The type of job queue. For service jobs that run on SageMaker AI, this value is `SAGEMAKER_TRAINING`. For regular container jobs, this value is `EKS`, `ECS`, or `ECS_FARGATE` depending on the compute environment.

_Required_: No

_Type_: String

_Allowed values_: `EKS | ECS | ECS_FARGATE | SAGEMAKER_TRAINING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobStateTimeLimitActions`

The set of actions that AWS Batch perform on jobs that remain at the head of the job queue in the specified state longer than specified times. AWS Batch will perform each action after `maxTimeSeconds` has passed.

_Required_: No

_Type_: Array of [JobStateTimeLimitAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobqueue-jobstatetimelimitaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority of the job queue. Job queues with a higher priority (or a higher integer
value for the `priority` parameter) are evaluated first when associated with the
same compute environment. Priority is determined in descending order. For example, a job queue
with a priority value of `10` is given scheduling preference over a job queue with
a priority value of `1`. All of the compute environments must be either EC2
( `EC2` or `SPOT`) or Fargate ( `FARGATE` or
`FARGATE_SPOT`); EC2 and Fargate compute environments can't be mixed.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchedulingPolicyArn`

The Amazon Resource Name (ARN) of the scheduling policy. The format is
`aws:Partition:batch:Region:Account:scheduling-policy/Name`.
For example,
`aws:aws:batch:us-west-2:123456789012:scheduling-policy/MySchedulingPolicy`.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceEnvironmentOrder`

The order of the service environment associated with the job queue. Job queues with a higher priority are evaluated first when associated with the same service environment.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobqueue-serviceenvironmentorder.html) of [ServiceEnvironmentOrder](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobqueue-serviceenvironmentorder.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the job queue. If the job queue state is `ENABLED`, it is able to
accept jobs. If the job queue state is `DISABLED`, new jobs can't be added to the
queue, but jobs already in the queue can finish.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that are applied to the job queue. For more information, see [Tagging your AWS Batch resources](https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html) in
_AWS Batch User Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the job queue ARN, such as `arn:aws:batch:us-east-1:111122223333:job-queue/HighPriority`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`JobQueueArn`

Returns the job queue ARN, such as `arn:aws:batch:us-east-1:111122223333:job-queue/JobQueueName`.

## Examples

### Job queue with two compute environments

The following example defines a job queue called `HighPriority` that
has two compute environments mapped to it.

#### JSON

```json

{
  "JobQueue": {
    "Type": "AWS::Batch::JobQueue",
    "Properties": {
      "ComputeEnvironmentOrder": [
        {
          "Order": 1,
          "ComputeEnvironment": "C4OnDemand"
        },
        {
          "Order": 2,
          "ComputeEnvironment": "M4Spot"
        }
      ],
      "State": "ENABLED",
      "Priority": 1,
      "JobQueueName": "HighPriority"
    }
  }
}
```

#### YAML

```yaml

JobQueue:
  Type: 'AWS::Batch::JobQueue'
  Properties:
    ComputeEnvironmentOrder:
      - Order: 1
        ComputeEnvironment: C4OnDemand
      - Order: 2
        ComputeEnvironment: M4Spot
    State: ENABLED
    Priority: 1
    JobQueueName: HighPriority
```

## See also

- [Job Queue Parameters](https://docs.aws.amazon.com/batch/latest/userguide/job_queue_parameters.html) in the _AWS Batch User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Volume

ComputeEnvironmentOrder
