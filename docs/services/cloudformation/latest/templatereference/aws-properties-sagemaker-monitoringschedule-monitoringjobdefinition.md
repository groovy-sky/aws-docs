This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule MonitoringJobDefinition

Defines the monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaselineConfig" : BaselineConfig,
  "Environment" : {Key: Value, ...},
  "MonitoringAppSpecification" : MonitoringAppSpecification,
  "MonitoringInputs" : [ MonitoringInput, ... ],
  "MonitoringOutputConfig" : MonitoringOutputConfig,
  "MonitoringResources" : MonitoringResources,
  "NetworkConfig" : NetworkConfig,
  "RoleArn" : String,
  "StoppingCondition" : StoppingCondition
}

```

### YAML

```yaml

  BaselineConfig:
    BaselineConfig
  Environment:
    Key: Value
  MonitoringAppSpecification:
    MonitoringAppSpecification
  MonitoringInputs:
    - MonitoringInput
  MonitoringOutputConfig:
    MonitoringOutputConfig
  MonitoringResources:
    MonitoringResources
  NetworkConfig:
    NetworkConfig
  RoleArn: String
  StoppingCondition:
    StoppingCondition

```

## Properties

`BaselineConfig`

Baseline configuration used to validate that the data conforms to the specified
constraints and statistics

_Required_: No

_Type_: [BaselineConfig](aws-properties-sagemaker-monitoringschedule-baselineconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

Sets the environment variables in the Docker container.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringAppSpecification`

Configures the monitoring job to run a specified Docker container image.

_Required_: Yes

_Type_: [MonitoringAppSpecification](aws-properties-sagemaker-monitoringschedule-monitoringappspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringInputs`

The array of inputs for the monitoring job. Currently we support monitoring an Amazon SageMaker AI Endpoint.

_Required_: Yes

_Type_: Array of [MonitoringInput](aws-properties-sagemaker-monitoringschedule-monitoringinput.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringOutputConfig`

The array of outputs from the monitoring job to be uploaded to Amazon S3.

_Required_: Yes

_Type_: [MonitoringOutputConfig](aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringResources`

Identifies the resources, ML compute instances, and ML storage volumes to deploy for a
monitoring job. In distributed processing, you specify more than one instance.

_Required_: Yes

_Type_: [MonitoringResources](aws-properties-sagemaker-monitoringschedule-monitoringresources.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfig`

Specifies networking options for an monitoring job.

_Required_: No

_Type_: [NetworkConfig](aws-properties-sagemaker-monitoringschedule-networkconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can
assume to perform tasks on your behalf.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StoppingCondition`

Specifies a time limit for how long the monitoring job is allowed to run.

_Required_: No

_Type_: [StoppingCondition](aws-properties-sagemaker-monitoringschedule-stoppingcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringInput

MonitoringOutput

All content copied from https://docs.aws.amazon.com/.
