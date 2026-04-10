This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule BaselineConfig

Baseline configuration used to validate that the data conforms to the specified constraints and
statistics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstraintsResource" : ConstraintsResource,
  "StatisticsResource" : StatisticsResource
}

```

### YAML

```yaml

  ConstraintsResource:
    ConstraintsResource
  StatisticsResource:
    StatisticsResource

```

## Properties

`ConstraintsResource`

The Amazon S3 URI for the constraints resource.

_Required_: No

_Type_: [ConstraintsResource](aws-properties-sagemaker-monitoringschedule-constraintsresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatisticsResource`

The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.

_Required_: No

_Type_: [StatisticsResource](aws-properties-sagemaker-monitoringschedule-statisticsresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::MonitoringSchedule

BatchTransformInput

All content copied from https://docs.aws.amazon.com/.
