This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiGatewayManagedOverrides StageOverrides

The `StageOverrides` property overrides the stage configuration for an API
Gateway-managed stage. If you remove this property, API Gateway restores the default values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessLogSettings" : AccessLogSettings,
  "AutoDeploy" : Boolean,
  "DefaultRouteSettings" : RouteSettings,
  "Description" : String,
  "RouteSettings" : Json,
  "StageVariables" : Json
}

```

### YAML

```yaml

  AccessLogSettings:
    AccessLogSettings
  AutoDeploy: Boolean
  DefaultRouteSettings:
    RouteSettings
  Description: String
  RouteSettings: Json
  StageVariables: Json

```

## Properties

`AccessLogSettings`

Settings for logging access in a stage.

_Required_: No

_Type_: [AccessLogSettings](aws-properties-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoDeploy`

Specifies whether updates to an API automatically trigger a new deployment. The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRouteSettings`

The default route settings for the stage.

_Required_: No

_Type_: [RouteSettings](aws-properties-apigatewayv2-apigatewaymanagedoverrides-routesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the API stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouteSettings`

Route settings for the stage.

_Required_: No

_Type_: [Json](aws-properties-apigatewayv2-apigatewaymanagedoverrides-routesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageVariables`

A map that defines the stage variables for a `Stage`. Variable names can have alphanumeric and underscore characters, and the values must match \[A-Za-z0-9-.\_~:/?#&=,\]+.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouteSettings

AWS::ApiGatewayV2::ApiMapping

All content copied from https://docs.aws.amazon.com/.
