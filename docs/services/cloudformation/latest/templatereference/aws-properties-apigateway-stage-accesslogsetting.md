This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Stage AccessLogSetting

The `AccessLogSetting` property type specifies settings for logging access in this stage.

`AccessLogSetting` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.

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

The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`. This parameter is required to enable access logging.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

A single line format of the access logs of data, as specified by selected [$context variables](../../../apigateway/latest/developerguide/api-gateway-mapping-template-reference.md#context-variable-reference). The format must include at least `$context.requestId`. This parameter is required to enable access logging.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Stage](../../../apigateway/latest/api/api-stage.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::Stage

CanarySetting
