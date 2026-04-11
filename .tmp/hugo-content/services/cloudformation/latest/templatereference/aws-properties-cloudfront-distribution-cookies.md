This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution Cookies

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

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Forward" : String,
  "WhitelistedNames" : [ String, ... ]
}

```

### YAML

```yaml

  Forward: String
  WhitelistedNames:
    - String

```

## Properties

`Forward`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include cookies in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send cookies to the origin but not include them in the cache key, use
origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

Specifies which cookies to forward to the origin for this cache behavior: all, none,
or the list of cookies specified in the `WhitelistedNames` complex
type.

Amazon S3 doesn't process cookies. When the cache behavior is forwarding requests to an
Amazon S3 origin, specify none for the `Forward` element.

_Required_: Yes

_Type_: String

_Allowed values_: `none | whitelist | all`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhitelistedNames`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field.

If you want to include cookies in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) in the _Amazon CloudFront Developer Guide_.

If you want to send cookies to the origin but not include them in the cache key, use
an origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) in the
_Amazon CloudFront Developer Guide_.

Required if you specify `whitelist` for the value of `Forward`.
A complex type that specifies how many different cookies you want CloudFront to forward to the
origin for this cache behavior and, if you want to forward selected cookies, the names
of those cookies.

If you specify `all` or `none` for the value of
`Forward`, omit `WhitelistedNames`. If you change the value of
`Forward` from `whitelist` to `all` or
`none` and you don't delete the `WhitelistedNames` element and
its child elements, CloudFront deletes them automatically.

For the current limit on the number of cookie names that you can whitelist for each
cache behavior, see [CloudFront\
Limits](../../../../general/latest/gr/xrefaws-service-limits.md#limits_cloudfront) in the _AWS General Reference_.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CookiePreference](../../../../reference/cloudfront/latest/apireference/api-cookiepreference.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionFunctionAssociation

CustomErrorResponse

All content copied from https://docs.aws.amazon.com/.
