This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Deployment AccessLogSetting

The `AccessLogSetting` property type specifies settings for logging access in this stage.

`AccessLogSetting` is a property of the [StageDescription](../userguide/aws-properties-apigateway-deployment-stagedescription.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationArn" : String,
  "Format" : String
}

```

### YAML

```yaml

  DestinationArn: String
  Format: String

```

## Properties

`DestinationArn`

The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

A single line format of the access logs of data, as specified by selected $context variables. The format must include at least `$context.requestId`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [accessLogSettings](../../../apigateway/latest/api/api-stage.md#accessLogSettings) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::Deployment

CanarySetting

All content copied from https://docs.aws.amazon.com/.
