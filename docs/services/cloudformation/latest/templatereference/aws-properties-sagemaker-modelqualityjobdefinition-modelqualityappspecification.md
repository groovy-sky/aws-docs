This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelQualityJobDefinition ModelQualityAppSpecification

Container image configuration object for the monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerArguments" : [ String, ... ],
  "ContainerEntrypoint" : [ String, ... ],
  "Environment" : {Key: Value, ...},
  "ImageUri" : String,
  "PostAnalyticsProcessorSourceUri" : String,
  "ProblemType" : String,
  "RecordPreprocessorSourceUri" : String
}

```

### YAML

```yaml

  ContainerArguments:
    - String
  ContainerEntrypoint:
    - String
  Environment:
    Key: Value
  ImageUri: String
  PostAnalyticsProcessorSourceUri: String
  ProblemType: String
  RecordPreprocessorSourceUri: String

```

## Properties

`ContainerArguments`

An array of arguments for the container used to run the monitoring job.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256 | 50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerEntrypoint`

Specifies the entrypoint for a container that the monitoring job runs.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

Sets the environment variables in the container that the monitoring job runs.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageUri`

The address of the container image that the monitoring job runs.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PostAnalyticsProcessorSourceUri`

An Amazon S3 URI to a script that is called after analysis has been performed. Applicable
only for the built-in (first party) containers.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProblemType`

The machine learning problem type of the model that the monitoring job monitors.

_Required_: Yes

_Type_: String

_Allowed values_: `BinaryClassification | MulticlassClassification | Regression`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordPreprocessorSourceUri`

An Amazon S3 URI to a script that is called per row prior to running analysis. It can
base64 decode the payload and convert it into a flattened JSON so that the built-in container can use
the converted data. Applicable only for the built-in (first party) containers.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Json

ModelQualityBaselineConfig
