This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Stage CanarySetting

Configuration settings of a canary deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentId" : String,
  "PercentTraffic" : Number,
  "StageVariableOverrides" : {Key: Value, ...},
  "UseStageCache" : Boolean
}

```

### YAML

```yaml

  DeploymentId: String
  PercentTraffic: Number
  StageVariableOverrides:
    Key: Value
  UseStageCache: Boolean

```

## Properties

`DeploymentId`

The ID of the canary deployment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentTraffic`

The percent (0-100) of traffic diverted to a canary deployment.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageVariableOverrides`

Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseStageCache`

A Boolean flag to indicate whether the canary deployment uses the stage cache or not.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Stage](../../../apigateway/latest/api/api-stage.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessLogSetting

MethodSetting
