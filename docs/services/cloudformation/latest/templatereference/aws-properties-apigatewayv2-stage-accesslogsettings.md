This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Stage AccessLogSettings

Settings for logging access in a stage.

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

The ARN of the CloudWatch Logs log group to receive access logs. This parameter is required to enable access logging.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId. This parameter is required to enable access logging.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Stages](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-stages.html) in the _Amazon API Gateway_
_Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::Stage

RouteSettings
