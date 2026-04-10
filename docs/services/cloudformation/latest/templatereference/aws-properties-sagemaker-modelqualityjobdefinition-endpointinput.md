This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelQualityJobDefinition EndpointInput

Input object for the endpoint

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointName" : String,
  "EndTimeOffset" : String,
  "InferenceAttribute" : String,
  "LocalPath" : String,
  "ProbabilityAttribute" : String,
  "ProbabilityThresholdAttribute" : Number,
  "S3DataDistributionType" : String,
  "S3InputMode" : String,
  "StartTimeOffset" : String
}

```

### YAML

```yaml

  EndpointName: String
  EndTimeOffset: String
  InferenceAttribute: String
  LocalPath: String
  ProbabilityAttribute: String
  ProbabilityThresholdAttribute: Number
  S3DataDistributionType: String
  S3InputMode: String
  StartTimeOffset: String

```

## Properties

`EndpointName`

An endpoint in customer's account which has enabled `DataCaptureConfig`
enabled.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndTimeOffset`

If specified, monitoring jobs substract this time from the end time. For information
about using offsets for scheduling monitoring jobs, see [Schedule Model\
Quality Monitoring Jobs](../../../sagemaker/latest/dg/model-monitor-model-quality-schedule.md).

_Required_: No

_Type_: String

_Pattern_: `^.?P.*`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceAttribute`

The attribute of the input data that represents the ground truth label.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalPath`

Path to the filesystem where the endpoint data is available to the container.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProbabilityAttribute`

In a classification problem, the attribute that represents the class probability.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProbabilityThresholdAttribute`

The threshold for the class probability to be evaluated as a positive result.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3DataDistributionType`

Whether input data distributed in Amazon S3 is fully replicated or sharded by an
Amazon S3 key. Defaults to `FullyReplicated`

_Required_: No

_Type_: String

_Allowed values_: `FullyReplicated | ShardedByS3Key`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3InputMode`

Whether the `Pipe` or `File` is used as the input mode for
transferring data for the monitoring job. `Pipe` mode is recommended for large
datasets. `File` mode is useful for small files that fit in memory. Defaults to
`File`.

_Required_: No

_Type_: String

_Allowed values_: `Pipe | File`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartTimeOffset`

If specified, monitoring jobs substract this time from the start time. For information
about using offsets for scheduling monitoring jobs, see [Schedule Model\
Quality Monitoring Jobs](../../../sagemaker/latest/dg/model-monitor-model-quality-schedule.md).

_Required_: No

_Type_: String

_Pattern_: `^.?P.*`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetFormat

Json

All content copied from https://docs.aws.amazon.com/.
