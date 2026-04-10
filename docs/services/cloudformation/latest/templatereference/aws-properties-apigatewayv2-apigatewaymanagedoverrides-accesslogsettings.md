This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiGatewayManagedOverrides AccessLogSettings

The `AccessLogSettings` property overrides the access log settings for an API Gateway-managed stage.

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

The ARN of the CloudWatch Logs log group to receive access logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::ApiGatewayManagedOverrides

IntegrationOverrides

All content copied from https://docs.aws.amazon.com/.
