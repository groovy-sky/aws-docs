This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::Application InputProcessingConfiguration

Provides a description of a processor that is used to preprocess the records in the
stream before being processed by your application code. Currently, the only input
processor available is [AWS Lambda](../../../lambda/index.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputLambdaProcessor" : InputLambdaProcessor
}

```

### YAML

```yaml

  InputLambdaProcessor:
    InputLambdaProcessor

```

## Properties

`InputLambdaProcessor`

The [InputLambdaProcessor](../userguide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.md) that is used to preprocess the records
in the stream before being processed by your application code.

_Required_: No

_Type_: [InputLambdaProcessor](aws-properties-kinesisanalytics-application-inputlambdaprocessor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputParallelism

InputSchema

All content copied from https://docs.aws.amazon.com/.
