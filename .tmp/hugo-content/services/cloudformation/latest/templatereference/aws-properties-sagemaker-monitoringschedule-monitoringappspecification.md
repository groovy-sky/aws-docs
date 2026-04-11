This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule MonitoringAppSpecification

Container image configuration object for the monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerArguments" : [ String, ... ],
  "ContainerEntrypoint" : [ String, ... ],
  "ImageUri" : String,
  "PostAnalyticsProcessorSourceUri" : String,
  "RecordPreprocessorSourceUri" : String
}

```

### YAML

```yaml

  ContainerArguments:
    - String
  ContainerEntrypoint:
    - String
  ImageUri: String
  PostAnalyticsProcessorSourceUri: String
  RecordPreprocessorSourceUri: String

```

## Properties

`ContainerArguments`

An array of arguments for the container used to run the monitoring job.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerEntrypoint`

Specifies the entrypoint for a container used to run the monitoring job.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUri`

The container image to be run by the monitoring job.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostAnalyticsProcessorSourceUri`

An Amazon S3 URI to a script that is called after analysis has been performed. Applicable
only for the built-in (first party) containers.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordPreprocessorSourceUri`

An Amazon S3 URI to a script that is called per row prior to running analysis. It can
base64 decode the payload and convert it into a flattened JSON so that the built-in container can use
the converted data. Applicable only for the built-in (first party) containers.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Json

MonitoringExecutionSummary

All content copied from https://docs.aws.amazon.com/.
