This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiGatewayManagedOverrides IntegrationOverrides

The `IntegrationOverrides` property overrides the integration settings for
an API Gateway-managed integration. If you remove this property, API Gateway restores the default values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "IntegrationMethod" : String,
  "PayloadFormatVersion" : String,
  "TimeoutInMillis" : Integer
}

```

### YAML

```yaml

  Description: String
  IntegrationMethod: String
  PayloadFormatVersion: String
  TimeoutInMillis: Integer

```

## Properties

`Description`

The description of the integration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationMethod`

Specifies the integration's HTTP method type. For WebSocket APIs, if you use a Lambda integration, you must set the integration method to `POST`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadFormatVersion`

Specifies the format of the payload sent to an integration. Required for HTTP
APIs. For HTTP APIs, supported values for Lambda proxy integrations are
`1.0` and `2.0`. For all other integrations,
`1.0` is the only supported value. To learn more, see
[Working with AWS Lambda proxy integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-lambda.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInMillis`

Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29
seconds for WebSocket APIs and 30 seconds for HTTP APIs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessLogSettings

RouteOverrides

All content copied from https://docs.aws.amazon.com/.
