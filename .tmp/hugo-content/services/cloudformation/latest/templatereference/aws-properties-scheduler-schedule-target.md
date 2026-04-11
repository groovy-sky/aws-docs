This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule Target

The schedule's target. EventBridge Scheduler supports templated target that invoke common API operations, as well as universal targets that you can customize to
invoke over 6,000 API operations across more than 270 services. You can only specify one templated or universal target for a schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "DeadLetterConfig" : DeadLetterConfig,
  "EcsParameters" : EcsParameters,
  "EventBridgeParameters" : EventBridgeParameters,
  "Input" : String,
  "KinesisParameters" : KinesisParameters,
  "RetryPolicy" : RetryPolicy,
  "RoleArn" : String,
  "SageMakerPipelineParameters" : SageMakerPipelineParameters,
  "SqsParameters" : SqsParameters
}

```

### YAML

```yaml

  Arn: String
  DeadLetterConfig:
    DeadLetterConfig
  EcsParameters:
    EcsParameters
  EventBridgeParameters:
    EventBridgeParameters
  Input: String
  KinesisParameters:
    KinesisParameters
  RetryPolicy:
    RetryPolicy
  RoleArn: String
  SageMakerPipelineParameters:
    SageMakerPipelineParameters
  SqsParameters:
    SqsParameters

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the target.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeadLetterConfig`

An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule. If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.

_Required_: No

_Type_: [DeadLetterConfig](aws-properties-scheduler-schedule-deadletterconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsParameters`

The templated target type for the Amazon ECS [`RunTask`](../../../../reference/amazonecs/latest/apireference/api-runtask.md) API operation.

_Required_: No

_Type_: [EcsParameters](aws-properties-scheduler-schedule-ecsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeParameters`

The templated target type for the EventBridge [`PutEvents`](../../../../reference/eventbridge/latest/apireference/api-putevents.md) API operation.

_Required_: No

_Type_: [EventBridgeParameters](aws-properties-scheduler-schedule-eventbridgeparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Input`

The text, or well-formed JSON, passed to the target. If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target,
the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, Amazon EventBridge Scheduler
delivers a default notification to the target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisParameters`

The templated target type for the Amazon Kinesis [`PutRecord`](../../../../reference/kinesis/latest/apireference/api-putrecord.md) API operation.

_Required_: No

_Type_: [KinesisParameters](aws-properties-scheduler-schedule-kinesisparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryPolicy`

A `RetryPolicy` object that includes information about the retry policy settings, including the maximum age of an event, and the maximum number of times EventBridge Scheduler will try to deliver the event to a target.

_Required_: No

_Type_: [RetryPolicy](aws-properties-scheduler-schedule-retrypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that EventBridge Scheduler will use for this target when the schedule is invoked.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:iam::\d{12}:role\/[\w+=,.@\/-]+$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerPipelineParameters`

The templated target type for the Amazon SageMaker [`StartPipelineExecution`](../../../../reference/sagemaker/latest/apireference/api-startpipelineexecution.md) API operation.

_Required_: No

_Type_: [SageMakerPipelineParameters](aws-properties-scheduler-schedule-sagemakerpipelineparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsParameters`

The templated target type for the Amazon SQS [`SendMessage`](../../../../reference/awssimplequeueservice/latest/apireference/api-sendmessage.md) API operation.
Contains the message group ID to use when the target is a FIFO queue. If you specify an Amazon SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
For more information, see [Using the Amazon SQS message deduplication ID](../../../awssimplequeueservice/latest/sqsdeveloperguide/using-messagededuplicationid-property.md) in the
_Amazon SQS Developer Guide_.

_Required_: No

_Type_: [SqsParameters](aws-properties-scheduler-schedule-sqsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SqsParameters

AWS::Scheduler::ScheduleGroup

All content copied from https://docs.aws.amazon.com/.
