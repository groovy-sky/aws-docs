This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiGatewayManagedOverrides

The `AWS::ApiGatewayV2::ApiGatewayManagedOverrides` resource overrides the
default properties of API Gateway-managed resources that are implicitly configured for
you when you use quick create. When you create an API by using quick create, an
`AWS::ApiGatewayV2::Route`, `AWS::ApiGatewayV2::Integration`,
and `AWS::ApiGatewayV2::Stage` are created for you and associated with your
`AWS::ApiGatewayV2::Api`. The
`AWS::ApiGatewayV2::ApiGatewayManagedOverrides` resource enables you to
set, or override the properties of these implicit resources. Supported only for HTTP
APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::ApiGatewayManagedOverrides",
  "Properties" : {
      "ApiId" : String,
      "Integration" : IntegrationOverrides,
      "Route" : RouteOverrides,
      "Stage" : StageOverrides
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::ApiGatewayManagedOverrides
Properties:
  ApiId: String
  Integration:
    IntegrationOverrides
  Route:
    RouteOverrides
  Stage:
    StageOverrides

```

## Properties

`ApiId`

The ID of the API for which to override the configuration of API Gateway-managed resources.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Integration`

Overrides the integration configuration for an API Gateway-managed integration.

_Required_: No

_Type_: [IntegrationOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-apigatewaymanagedoverrides-integrationoverrides.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Route`

Overrides the route configuration for an API Gateway-managed route.

_Required_: No

_Type_: [RouteOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-apigatewaymanagedoverrides-routeoverrides.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

Overrides the stage configuration for an API Gateway-managed stage.

_Required_: No

_Type_: [StageOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Id`

The identifier.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Cors

AccessLogSettings
