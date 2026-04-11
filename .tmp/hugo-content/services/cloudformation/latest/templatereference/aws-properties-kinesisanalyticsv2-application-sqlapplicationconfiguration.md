This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application SqlApplicationConfiguration

Describes the inputs, outputs, and reference data sources for a SQL-based Kinesis Data Analytics application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Inputs" : [ Input, ... ]
}

```

### YAML

```yaml

  Inputs:
    - Input

```

## Properties

`Inputs`

The array of [Input](../../../managed-flink/latest/apiv2/api-input.md) objects describing the
input streams used by the application.

_Required_: No

_Type_: Array of [Input](aws-properties-kinesisanalyticsv2-application-input.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SqlApplicationConfiguration](../../../managed-flink/latest/apiv2/api-sqlapplicationconfiguration.md) in the _Amazon Kinesis_
_Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ContentLocation

Tag

All content copied from https://docs.aws.amazon.com/.
