This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy CorsConfig

A configuration for a set of HTTP response headers that are used for cross-origin
resource sharing (CORS). CloudFront adds these headers to HTTP responses that it sends for
CORS requests that match a cache behavior associated with this response headers
policy.

For more information about CORS, see [Cross-Origin Resource\
Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessControlAllowCredentials" : Boolean,
  "AccessControlAllowHeaders" : AccessControlAllowHeaders,
  "AccessControlAllowMethods" : AccessControlAllowMethods,
  "AccessControlAllowOrigins" : AccessControlAllowOrigins,
  "AccessControlExposeHeaders" : AccessControlExposeHeaders,
  "AccessControlMaxAgeSec" : Integer,
  "OriginOverride" : Boolean
}

```

### YAML

```yaml

  AccessControlAllowCredentials: Boolean
  AccessControlAllowHeaders:
    AccessControlAllowHeaders
  AccessControlAllowMethods:
    AccessControlAllowMethods
  AccessControlAllowOrigins:
    AccessControlAllowOrigins
  AccessControlExposeHeaders:
    AccessControlExposeHeaders
  AccessControlMaxAgeSec: Integer
  OriginOverride: Boolean

```

## Properties

`AccessControlAllowCredentials`

A Boolean that CloudFront uses as the value for the
`Access-Control-Allow-Credentials` HTTP response header.

For more information about the `Access-Control-Allow-Credentials` HTTP
response header, see [Access-Control-Allow-Credentials](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) in the MDN Web Docs.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControlAllowHeaders`

A list of HTTP header names that CloudFront includes as values for the
`Access-Control-Allow-Headers` HTTP response header.

For more information about the `Access-Control-Allow-Headers` HTTP response
header, see [Access-Control-Allow-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.

_Required_: Yes

_Type_: [AccessControlAllowHeaders](aws-properties-cloudfront-responseheaderspolicy-accesscontrolallowheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControlAllowMethods`

A list of HTTP methods that CloudFront includes as values for the
`Access-Control-Allow-Methods` HTTP response header.

For more information about the `Access-Control-Allow-Methods` HTTP response
header, see [Access-Control-Allow-Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.

_Required_: Yes

_Type_: [AccessControlAllowMethods](aws-properties-cloudfront-responseheaderspolicy-accesscontrolallowmethods.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControlAllowOrigins`

A list of origins (domain names) that CloudFront can use as the value for the
`Access-Control-Allow-Origin` HTTP response header.

For more information about the `Access-Control-Allow-Origin` HTTP response
header, see [Access-Control-Allow-Origin](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.

_Required_: Yes

_Type_: [AccessControlAllowOrigins](aws-properties-cloudfront-responseheaderspolicy-accesscontrolalloworigins.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControlExposeHeaders`

A list of HTTP headers that CloudFront includes as values for the
`Access-Control-Expose-Headers` HTTP response header.

For more information about the `Access-Control-Expose-Headers` HTTP
response header, see [Access-Control-Expose-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.

_Required_: No

_Type_: [AccessControlExposeHeaders](aws-properties-cloudfront-responseheaderspolicy-accesscontrolexposeheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControlMaxAgeSec`

A number that CloudFront uses as the value for the `Access-Control-Max-Age` HTTP
response header.

For more information about the `Access-Control-Max-Age` HTTP response
header, see [Access-Control-Max-Age](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in the MDN Web Docs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginOverride`

A Boolean that determines whether CloudFront overrides HTTP response headers received from
the origin with the ones specified in this response headers policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentTypeOptions

CustomHeader

All content copied from https://docs.aws.amazon.com/.
