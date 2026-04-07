This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution ForwardedValues

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include values in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send values to the origin but not include them in the cache key, use an
origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

A complex type that specifies how CloudFront handles query strings, cookies, and HTTP
headers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cookies" : Cookies,
  "Headers" : [ String, ... ],
  "QueryString" : Boolean,
  "QueryStringCacheKeys" : [ String, ... ]
}

```

### YAML

```yaml

  Cookies:
    Cookies
  Headers:
    - String
  QueryString: Boolean
  QueryStringCacheKeys:
    - String

```

## Properties

`Cookies`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include cookies in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send cookies to the origin but not include them in the cache key, use
an origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

A complex type that specifies whether you want CloudFront to forward cookies to the origin
and, if so, which ones. For more information about forwarding cookies to the origin, see
[How CloudFront Forwards, Caches,\
and Logs Cookies](../../../amazoncloudfront/latest/developerguide/cookies.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: [Cookies](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distribution-cookies.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include headers in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send headers to the origin but not include them in the cache key, use
an origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

A complex type that specifies the `Headers`, if any, that you want CloudFront to
forward to the origin for this cache behavior (whitelisted headers). For the headers
that you specify, CloudFront also caches separate versions of a specified object that is based
on the header values in viewer requests.

For more information, see [Caching Content\
Based on Request Headers](../../../amazoncloudfront/latest/developerguide/header-caching.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include query strings in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send query strings to the origin but not include them in the cache key,
use an origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

Indicates whether you want CloudFront to forward query strings to the origin that is
associated with this cache behavior and cache based on the query string parameters. CloudFront
behavior depends on the value of `QueryString` and on the values that you
specify for `QueryStringCacheKeys`, if any:

If you specify true for `QueryString` and you don't specify any values for
`QueryStringCacheKeys`, CloudFront forwards all query string parameters to the
origin and caches based on all query string parameters. Depending on how many query
string parameters and values you have, this can adversely affect performance because
CloudFront must forward more requests to the origin.

If you specify true for `QueryString` and you specify one or more values
for `QueryStringCacheKeys`, CloudFront forwards all query string parameters to the
origin, but it only caches based on the query string parameters that you specify.

If you specify false for `QueryString`, CloudFront doesn't forward any query
string parameters to the origin, and doesn't cache based on query string
parameters.

For more information, see [Configuring\
CloudFront to Cache Based on Query String Parameters](../../../amazoncloudfront/latest/developerguide/querystringparameters.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringCacheKeys`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include query strings in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send query strings to the origin but not include them in the cache key,
use an origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

A complex type that contains information about the query string parameters that you
want CloudFront to use for caching for this cache behavior.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ForwardedValues](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ForwardedValues.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DistributionConfig

FunctionAssociation
