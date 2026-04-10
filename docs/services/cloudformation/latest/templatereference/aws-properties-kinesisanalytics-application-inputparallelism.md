This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::Application InputParallelism

Describes the number of in-application streams to create for a given streaming source.
For information about parallelism, see [Configuring Application\
Input](../../../kinesisanalytics/latest/dev/how-it-works-input.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Count" : Integer
}

```

### YAML

```yaml

  Count: Integer

```

## Properties

`Count`

Number of in-application streams to create. For more information, see [Limits](../../../kinesisanalytics/latest/dev/limits.md).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLambdaProcessor

InputProcessingConfiguration

All content copied from https://docs.aws.amazon.com/.
