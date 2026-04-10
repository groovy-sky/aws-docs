This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::UsagePlan ApiStage

API stage name of the associated API stage in a usage plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiId" : String,
  "Stage" : String,
  "Throttle" : {Key: Value, ...}
}

```

### YAML

```yaml

  ApiId: String
  Stage: String
  Throttle:
    Key: Value

```

## Properties

`ApiId`

API Id of the associated API stage in a usage plan.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

API stage name of the associated API stage in a usage plan.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Throttle`

Map containing method level throttling information for API stage in a usage plan.

_Required_: No

_Type_: Object of [ThrottleSettings](aws-properties-apigateway-usageplan-throttlesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [UsagePlan](../../../apigateway/latest/api/api-usageplan.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::UsagePlan

QuotaSettings

All content copied from https://docs.aws.amazon.com/.
