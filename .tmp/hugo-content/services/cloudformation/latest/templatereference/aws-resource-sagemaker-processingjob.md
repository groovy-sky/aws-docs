This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob

An Amazon SageMaker processing job that is used to analyze data and evaluate models. For more information, see
[Process Data and Evaluate\
Models](../../../sagemaker/latest/dg/processing-job.md).

Also, note the following details specific to processing jobs created using CloudFormation stacks:

- When you delete a CloudFormation stack with a processing job resource, the processing job is stopped using the
[StopProcessingJob](../../../../reference/sagemaker/latest/apireference/api-stopprocessingjob.md) API but not deleted. Any tags associated with the processing job are deleted using the
[DeleteTags](../../../../reference/sagemaker/latest/apireference/api-deletetags.md) API.

- If any part of your CloudFormation stack deployment fails and a rollback initiates, processing jobs with a
specified `ProcessingJobName` value might cause the stack to become stuck in a failed state. This occurs
because during a rollback, CloudFormation attempts to recreate the stack resources. Processing job names must be
unique, so when CloudFormation attempts to recreate a processing job using the already defined name, this results
in an `AlreadyExists` error. To prevent this, we recommend that you don't specify the optional
`ProcessingJobName` property, thereby allowing SageMaker to auto-generate a unique name for your
processing job. This ensures successful stack rollbacks when necessary. If you must use custom job names, you have
to manually modify the `ProcessingJobName` and redeploy the stack to recover from a failed
rollback.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::ProcessingJob",
  "Properties" : {
      "AppSpecification" : AppSpecification,
      "Environment" : {Key: Value, ...},
      "ExperimentConfig" : ExperimentConfig,
      "NetworkConfig" : NetworkConfig,
      "ProcessingInputs" : [ ProcessingInputsObject, ... ],
      "ProcessingJobName" : String,
      "ProcessingOutputConfig" : ProcessingOutputConfig,
      "ProcessingResources" : ProcessingResources,
      "RoleArn" : String,
      "StoppingCondition" : StoppingCondition,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::ProcessingJob
Properties:
  AppSpecification:
    AppSpecification
  Environment:
    Key: Value
  ExperimentConfig:
    ExperimentConfig
  NetworkConfig:
    NetworkConfig
  ProcessingInputs:
    - ProcessingInputsObject
  ProcessingJobName: String
  ProcessingOutputConfig:
    ProcessingOutputConfig
  ProcessingResources:
    ProcessingResources
  RoleArn: String
  StoppingCondition:
    StoppingCondition
  Tags:
    - Tag

```

## Properties

`AppSpecification`

Configuration to run a processing job in a specified container image.

_Required_: Yes

_Type_: [AppSpecification](aws-properties-sagemaker-processingjob-appspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

Sets the environment variables in the Docker container.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExperimentConfig`

Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the [CreateProcessingJob](../../../../reference/sagemaker/latest/apireference/api-createprocessingjob.md)
API.

_Required_: No

_Type_: [ExperimentConfig](aws-properties-sagemaker-processingjob-experimentconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfig`

Networking options for a job, such as network traffic encryption between containers,
whether to allow inbound and outbound network calls to and from containers, and the VPC
subnets and security groups to use for VPC-enabled jobs.

_Required_: No

_Type_: [NetworkConfig](aws-properties-sagemaker-processingjob-networkconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProcessingInputs`

List of input configurations for the processing job.

_Required_: No

_Type_: Array of [ProcessingInputsObject](aws-properties-sagemaker-processingjob-processinginputsobject.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProcessingJobName`

The name of the processing job. If you don't provide a job name, then a unique name is automatically created for
the job.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProcessingOutputConfig`

Contains information about the output location for the compiled model and the target
device that the model runs on. `TargetDevice` and `TargetPlatform`
are mutually exclusive, so you need to choose one between the two to specify your target
device or platform. If you cannot find your device you want to use from the
`TargetDevice` list, use `TargetPlatform` to describe the
platform of your edge device and `CompilerOptions` if there are specific
settings that are required or recommended to use for particular TargetPlatform.

_Required_: No

_Type_: [ProcessingOutputConfig](aws-properties-sagemaker-processingjob-processingoutputconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProcessingResources`

Identifies the resources, ML compute instances, and ML storage volumes to deploy for a
processing job. In distributed training, you specify more than one instance.

_Required_: Yes

_Type_: [ProcessingResources](aws-properties-sagemaker-processingjob-processingresources.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the role used to create the processing job.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StoppingCondition`

Configures conditions under which the processing job should be stopped, such as how long the processing job has
been running. After the condition is met, the processing job is stopped.

_Required_: No

_Type_: [StoppingCondition](aws-properties-sagemaker-processingjob-stoppingcondition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs. For more information, see [Using Cost Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md#allocation-whatURL) in
the _AWS Billing and Cost Management User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-processingjob-tag.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutoMLJobArn`

The Amazon Resource Name (ARN) of the AutoML job associated with this processing
job.

`CreationTime`

The time the processing job was created.

`ExitMessage`

A string, up to one KB in size, that contains metadata from the processing container
when the processing job exits.

`FailureReason`

A string, up to one KB in size, that contains the reason a processing job failed, if
it failed.

`LastModifiedTime`

The time the processing job was last modified.

`MonitoringScheduleArn`

The ARN of a monitoring schedule for an endpoint associated with this processing
job.

`ProcessingEndTime`

The time that the processing job ended.

`ProcessingJobArn`

The ARN of the processing job.

`ProcessingJobStatus`

The status of the processing job.

`ProcessingStartTime`

The time that the processing job started.

`TrainingJobArn`

The ARN of the training job associated with this processing job.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AppSpecification

All content copied from https://docs.aws.amazon.com/.
