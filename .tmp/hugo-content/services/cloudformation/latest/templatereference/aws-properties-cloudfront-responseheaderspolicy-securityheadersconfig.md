This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy SecurityHeadersConfig

A configuration for a set of security-related HTTP response headers. CloudFront adds these
headers to HTTP responses that it sends for requests that match a cache behavior
associated with this response headers policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentSecurityPolicy" : ContentSecurityPolicy,
  "ContentTypeOptions" : ContentTypeOptions,
  "FrameOptions" : FrameOptions,
  "ReferrerPolicy" : ReferrerPolicy,
  "StrictTransportSecurity" : StrictTransportSecurity,
  "XSSProtection" : XSSProtection
}

```

### YAML

```yaml

  ContentSecurityPolicy:
    ContentSecurityPolicy
  ContentTypeOptions:
    ContentTypeOptions
  FrameOptions:
    FrameOptions
  ReferrerPolicy:
    ReferrerPolicy
  StrictTransportSecurity:
    StrictTransportSecurity
  XSSProtection:
    XSSProtection

```

## Properties

`ContentSecurityPolicy`

The policy directives and their values that CloudFront includes as values for the
`Content-Security-Policy` HTTP response header.

For more information about the `Content-Security-Policy` HTTP response
header, see [Content-Security-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.

_Required_: No

_Type_: [ContentSecurityPolicy](aws-properties-cloudfront-responseheaderspolicy-contentsecuritypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentTypeOptions`

Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response
header with its value set to `nosniff`.

For more information about the `X-Content-Type-Options` HTTP response
header, see [X-Content-Type-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.

_Required_: No

_Type_: [ContentTypeOptions](aws-properties-cloudfront-responseheaderspolicy-contenttypeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameOptions`

Determines whether CloudFront includes the `X-Frame-Options` HTTP response header
and the header's value.

For more information about the `X-Frame-Options` HTTP response header, see
[X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.

_Required_: No

_Type_: [FrameOptions](aws-properties-cloudfront-responseheaderspolicy-frameoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferrerPolicy`

Determines whether CloudFront includes the `Referrer-Policy` HTTP response header
and the header's value.

For more information about the `Referrer-Policy` HTTP response header, see
[Referrer-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.

_Required_: No

_Type_: [ReferrerPolicy](aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrictTransportSecurity`

Determines whether CloudFront includes the `Strict-Transport-Security` HTTP
response header and the header's value.

For more information about the `Strict-Transport-Security` HTTP response
header, see [Security headers](../../../amazoncloudfront/latest/developerguide/understanding-response-headers-policies.md#understanding-response-headers-policies-security) in the _Amazon CloudFront Developer Guide_ and [Strict-Transport-Security](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.

_Required_: No

_Type_: [StrictTransportSecurity](aws-properties-cloudfront-responseheaderspolicy-stricttransportsecurity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XSSProtection`

Determines whether CloudFront includes the `X-XSS-Protection` HTTP response
header and the header's value.

For more information about the `X-XSS-Protection` HTTP response header, see
[X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.

_Required_: No

_Type_: [XSSProtection](aws-properties-cloudfront-responseheaderspolicy-xssprotection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseHeadersPolicyConfig

ServerTimingHeadersConfig

All content copied from https://docs.aws.amazon.com/.
