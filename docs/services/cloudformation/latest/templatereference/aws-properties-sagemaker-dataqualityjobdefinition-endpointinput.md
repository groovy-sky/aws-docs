This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::DataQualityJobDefinition EndpointInput

Input object for the endpoint

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointName" : String,
  "ExcludeFeaturesAttribute" : String,
  "LocalPath" : String,
  "S3DataDistributionType" : String,
  "S3InputMode" : String
}

```

### YAML

```yaml

  EndpointName: String
  ExcludeFeaturesAttribute: String
  LocalPath: String
  S3DataDistributionType: String
  S3InputMode: String

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

`ExcludeFeaturesAttribute`

The attributes of the input data to exclude from the analysis.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalPath`

Path to the filesystem where the endpoint data is available to the container.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetFormat

Json

All content copied from https://docs.aws.amazon.com/.
