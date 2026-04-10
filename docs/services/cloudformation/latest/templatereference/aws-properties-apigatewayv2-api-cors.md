This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Api Cors

The `Cors` property specifies a CORS configuration for an API.
Supported only for HTTP APIs. See [Configuring CORS](../../../apigateway/latest/developerguide/http-api-cors.md) for more information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowCredentials" : Boolean,
  "AllowHeaders" : [ String, ... ],
  "AllowMethods" : [ String, ... ],
  "AllowOrigins" : [ String, ... ],
  "ExposeHeaders" : [ String, ... ],
  "MaxAge" : Integer
}

```

### YAML

```yaml

  AllowCredentials: Boolean
  AllowHeaders:
    - String
  AllowMethods:
    - String
  AllowOrigins:
    - String
  ExposeHeaders:
    - String
  MaxAge: Integer

```

## Properties

`AllowCredentials`

Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowHeaders`

Represents a collection of allowed headers. Supported only for HTTP APIs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowMethods`

Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowOrigins`

Represents a collection of allowed origins. Supported only for HTTP APIs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExposeHeaders`

Represents a collection of exposed headers. Supported only for HTTP APIs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAge`

The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodyS3Location

AWS::ApiGatewayV2::ApiGatewayManagedOverrides

All content copied from https://docs.aws.amazon.com/.
