This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Url Cors

The [Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
settings for your function URL. Use CORS to grant access to your function URL from any origin. You can also use CORS
to control access for specific HTTP headers and methods in requests to your function URL.

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

Whether you want to allow cookies or other credentials in requests to your function URL. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowHeaders`

The HTTP headers that origins can include in requests to your function URL. For example: `Date`, `Keep-Alive`,
`X-Custom-Header`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1024 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowMethods`

The HTTP methods that are allowed when calling your function URL. For example: `GET`, `POST`, `DELETE`,
or the wildcard character ( `*`).

_Required_: No

_Type_: Array of String

_Allowed values_: `GET | PUT | HEAD | POST | PATCH | DELETE | *`

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowOrigins`

The origins that can access your function URL. You can list any number of specific origins, separated by a comma. For example:
`https://www.example.com`, `http://localhost:60905`.

Alternatively, you can grant access to all origins with the wildcard character ( `*`).

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `253 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExposeHeaders`

The HTTP headers in your function response that you want to expose to origins that call your function URL. For example:
`Date`, `Keep-Alive`, `X-Custom-Header`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1024 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAge`

The maximum amount of time, in seconds, that browsers can cache results of a preflight request. By default, this is set to `0`,
which means the browser will not cache results.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::Url

AWS::Lambda::Version

All content copied from https://docs.aws.amazon.com/.
