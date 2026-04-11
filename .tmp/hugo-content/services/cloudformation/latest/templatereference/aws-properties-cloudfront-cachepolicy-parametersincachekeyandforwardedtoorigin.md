This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy ParametersInCacheKeyAndForwardedToOrigin

This object determines the values that CloudFront includes in the cache key. These values
can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to
find an object in its cache that it can return to the viewer.

The headers, cookies, and query strings that are included in the cache key are also included
in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find an
object in its cache that matches the request's cache key. If you want to send values to
the origin but _not_ include them in the cache key, use
`OriginRequestPolicy`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CookiesConfig" : CookiesConfig,
  "EnableAcceptEncodingBrotli" : Boolean,
  "EnableAcceptEncodingGzip" : Boolean,
  "HeadersConfig" : HeadersConfig,
  "QueryStringsConfig" : QueryStringsConfig
}

```

### YAML

```yaml

  CookiesConfig:
    CookiesConfig
  EnableAcceptEncodingBrotli: Boolean
  EnableAcceptEncodingGzip: Boolean
  HeadersConfig:
    HeadersConfig
  QueryStringsConfig:
    QueryStringsConfig

```

## Properties

`CookiesConfig`

An object that determines whether any cookies in viewer requests (and if so, which cookies)
are included in the cache key and in requests that CloudFront sends to the origin.

_Required_: Yes

_Type_: [CookiesConfig](aws-properties-cloudfront-cachepolicy-cookiesconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAcceptEncodingBrotli`

A flag that can affect whether the `Accept-Encoding` HTTP header is
included in the cache key and included in requests that CloudFront sends to the origin.

This field is related to the `EnableAcceptEncodingGzip` field. If one or
both of these fields is `true` _and_ the viewer request includes the `Accept-Encoding`
header, then CloudFront does the following:

- Normalizes the value of the viewer's `Accept-Encoding`
header

- Includes the normalized header in the cache key

- Includes the normalized header in the request to the origin, if a request is
necessary

For more information, see [Compression support](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-policy-compressed-objects) in the
_Amazon CloudFront Developer Guide_.

If you set this value to `true`, and this cache behavior also has an origin
request policy attached, do not include the `Accept-Encoding` header in the
origin request policy. CloudFront always includes the `Accept-Encoding` header in
origin requests when the value of this field is `true`, so including this
header in an origin request policy has no effect.

If both of these fields are `false`, then CloudFront treats the
`Accept-Encoding` header the same as any other HTTP header in the viewer
request. By default, it's not included in the cache key and it's not included in origin
requests. In this case, you can manually add `Accept-Encoding` to the headers
whitelist like any other HTTP header.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAcceptEncodingGzip`

A flag that can affect whether the `Accept-Encoding` HTTP header is
included in the cache key and included in requests that CloudFront sends to the origin.

This field is related to the `EnableAcceptEncodingBrotli` field. If one or
both of these fields is `true` _and_ the viewer request includes the `Accept-Encoding`
header, then CloudFront does the following:

- Normalizes the value of the viewer's `Accept-Encoding`
header

- Includes the normalized header in the cache key

- Includes the normalized header in the request to the origin, if a request is
necessary

For more information, see [Compression support](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-policy-compressed-objects) in the
_Amazon CloudFront Developer Guide_.

If you set this value to `true`, and this cache behavior also has an origin
request policy attached, do not include the `Accept-Encoding` header in the
origin request policy. CloudFront always includes the `Accept-Encoding` header in
origin requests when the value of this field is `true`, so including this
header in an origin request policy has no effect.

If both of these fields are `false`, then CloudFront treats the
`Accept-Encoding` header the same as any other HTTP header in the viewer
request. By default, it's not included in the cache key and it's not included in origin
requests. In this case, you can manually add `Accept-Encoding` to the headers
whitelist like any other HTTP header.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeadersConfig`

An object that determines whether any HTTP headers (and if so, which headers) are included
in the cache key and in requests that CloudFront sends to the origin.

_Required_: Yes

_Type_: [HeadersConfig](aws-properties-cloudfront-cachepolicy-headersconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringsConfig`

An object that determines whether any URL query strings in viewer requests (and if so, which
query strings) are included in the cache key and in requests that CloudFront sends to the
origin.

_Required_: Yes

_Type_: [QueryStringsConfig](aws-properties-cloudfront-cachepolicy-querystringsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeadersConfig

QueryStringsConfig

All content copied from https://docs.aws.amazon.com/.
