This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::UsagePlan ThrottleSettings

`ThrottleSettings` is a property of the [AWS::ApiGateway::UsagePlan](../userguide/aws-resource-apigateway-usageplan.md) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BurstLimit" : Integer,
  "RateLimit" : Number
}

```

### YAML

```yaml

  BurstLimit: Integer
  RateLimit: Number

```

## Properties

`BurstLimit`

The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RateLimit`

The API target request rate limit.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [UsagePlan](../../../apigateway/latest/api/api-usageplan.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ApiGateway::UsagePlanKey

All content copied from https://docs.aws.amazon.com/.
