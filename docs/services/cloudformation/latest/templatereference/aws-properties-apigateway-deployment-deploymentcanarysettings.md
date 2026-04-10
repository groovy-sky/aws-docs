This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Deployment DeploymentCanarySettings

The `DeploymentCanarySettings` property type specifies settings for the canary deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PercentTraffic" : Number,
  "StageVariableOverrides" : {Key: Value, ...},
  "UseStageCache" : Boolean
}

```

### YAML

```yaml

  PercentTraffic: Number
  StageVariableOverrides:
    Key: Value
  UseStageCache: Boolean

```

## Properties

`PercentTraffic`

The percentage (0.0-100.0) of traffic routed to the canary deployment.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageVariableOverrides`

A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UseStageCache`

A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [CreateDeployment](../../../apigateway/latest/api/api-createdeployment.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CanarySetting

MethodSetting

All content copied from https://docs.aws.amazon.com/.
