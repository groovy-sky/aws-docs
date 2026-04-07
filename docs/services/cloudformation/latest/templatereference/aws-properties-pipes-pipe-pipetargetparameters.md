This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetParameters

The parameters required to set up a target for your pipe.

For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchJobParameters" : PipeTargetBatchJobParameters,
  "CloudWatchLogsParameters" : PipeTargetCloudWatchLogsParameters,
  "EcsTaskParameters" : PipeTargetEcsTaskParameters,
  "EventBridgeEventBusParameters" : PipeTargetEventBridgeEventBusParameters,
  "HttpParameters" : PipeTargetHttpParameters,
  "InputTemplate" : String,
  "KinesisStreamParameters" : PipeTargetKinesisStreamParameters,
  "LambdaFunctionParameters" : PipeTargetLambdaFunctionParameters,
  "RedshiftDataParameters" : PipeTargetRedshiftDataParameters,
  "SageMakerPipelineParameters" : PipeTargetSageMakerPipelineParameters,
  "SqsQueueParameters" : PipeTargetSqsQueueParameters,
  "StepFunctionStateMachineParameters" : PipeTargetStateMachineParameters,
  "TimestreamParameters" : PipeTargetTimestreamParameters
}

```

### YAML

```yaml

  BatchJobParameters:
    PipeTargetBatchJobParameters
  CloudWatchLogsParameters:
    PipeTargetCloudWatchLogsParameters
  EcsTaskParameters:
    PipeTargetEcsTaskParameters
  EventBridgeEventBusParameters:
    PipeTargetEventBridgeEventBusParameters
  HttpParameters:
    PipeTargetHttpParameters
  InputTemplate: String
  KinesisStreamParameters:
    PipeTargetKinesisStreamParameters
  LambdaFunctionParameters:
    PipeTargetLambdaFunctionParameters
  RedshiftDataParameters:
    PipeTargetRedshiftDataParameters
  SageMakerPipelineParameters:
    PipeTargetSageMakerPipelineParameters
  SqsQueueParameters:
    PipeTargetSqsQueueParameters
  StepFunctionStateMachineParameters:
    PipeTargetStateMachineParameters
  TimestreamParameters:
    PipeTargetTimestreamParameters

```

## Properties

`BatchJobParameters`

The parameters for using an AWS Batch job as a target.

_Required_: No

_Type_: [PipeTargetBatchJobParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetbatchjobparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLogsParameters`

The parameters for using an CloudWatch Logs log stream as a target.

_Required_: No

_Type_: [PipeTargetCloudWatchLogsParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsTaskParameters`

The parameters for using an Amazon ECS task as a target.

_Required_: No

_Type_: [PipeTargetEcsTaskParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetecstaskparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeEventBusParameters`

The parameters for using an EventBridge event bus as a target.

_Required_: No

_Type_: [PipeTargetEventBridgeEventBusParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpParameters`

These are custom parameter to be used when the target is an API Gateway REST APIs or
EventBridge ApiDestinations.

_Required_: No

_Type_: [PipeTargetHttpParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargethttpparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputTemplate`

Valid JSON text passed to the target. In this case, nothing from the event itself is
passed to the target. For more information, see [The JavaScript Object Notation (JSON)\
Data Interchange Format](http://www.rfc-editor.org/rfc/rfc7159.txt).

To remove an input template, specify an empty string.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamParameters`

The parameters for using a Kinesis stream as a target.

_Required_: No

_Type_: [PipeTargetKinesisStreamParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetkinesisstreamparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionParameters`

The parameters for using a Lambda function as a target.

_Required_: No

_Type_: [PipeTargetLambdaFunctionParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetlambdafunctionparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftDataParameters`

These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the
Amazon Redshift Data API BatchExecuteStatement.

_Required_: No

_Type_: [PipeTargetRedshiftDataParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetredshiftdataparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerPipelineParameters`

The parameters for using a SageMaker AI pipeline as a target.

_Required_: No

_Type_: [PipeTargetSageMakerPipelineParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetsagemakerpipelineparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsQueueParameters`

The parameters for using a Amazon SQS stream as a target.

_Required_: No

_Type_: [PipeTargetSqsQueueParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepFunctionStateMachineParameters`

The parameters for using a Step Functions state machine as a target.

_Required_: No

_Type_: [PipeTargetStateMachineParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargetstatemachineparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestreamParameters`

The parameters for using a Timestream for LiveAnalytics table as a
target.

_Required_: No

_Type_: [PipeTargetTimestreamParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipetargettimestreamparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PipeTargetLambdaFunctionParameters

PipeTargetRedshiftDataParameters
