This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy CachePolicyConfig

A cache policy configuration.

This configuration determines the following:

- The values that CloudFront includes in the cache key. These values can include HTTP
headers, cookies, and URL query strings. CloudFront uses the cache key to find an
object in its cache that it can return to the viewer.

- The default, minimum, and maximum time to live (TTL) values that you want
objects to stay in the CloudFront cache.

###### Important

If your minimum TTL is greater than 0, CloudFront will cache content for at least the duration specified in the cache policy's minimum TTL, even if the `Cache-Control: no-cache`, `no-store`, or `private` directives are present in the origin headers.

The headers, cookies, and query strings that are included in the cache key are also included
in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a
valid object in its cache that matches the request's cache key. If you want to send
values to the origin but _not_ include them in the cache key, use
`OriginRequestPolicy`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "DefaultTTL" : Number,
  "MaxTTL" : Number,
  "MinTTL" : Number,
  "Name" : String,
  "ParametersInCacheKeyAndForwardedToOrigin" : ParametersInCacheKeyAndForwardedToOrigin
}

```

### YAML

```yaml

  Comment: String
  DefaultTTL: Number
  MaxTTL: Number
  MinTTL: Number
  Name: String
  ParametersInCacheKeyAndForwardedToOrigin:
    ParametersInCacheKeyAndForwardedToOrigin

```

## Properties

`Comment`

A comment to describe the cache policy. The comment cannot be longer than 128
characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultTTL`

The default amount of time, in seconds, that you want objects to stay in the CloudFront
cache before CloudFront sends another request to the origin to see if the object has been
updated. CloudFront uses this value as the object's time to live (TTL) only when the origin
does _not_ send `Cache-Control` or `Expires`
headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

The default value for this field is 86400 seconds (one day). If the value of
`MinTTL` is more than 86400 seconds, then the default value for this
field is the same as the value of `MinTTL`.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTTL`

The maximum amount of time, in seconds, that objects stay in the CloudFront cache before
CloudFront sends another request to the origin to see if the object has been updated. CloudFront
uses this value only when the origin sends `Cache-Control` or
`Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

The default value for this field is 31536000 seconds (one year). If the value of
`MinTTL` or `DefaultTTL` is more than 31536000 seconds, then
the default value for this field is the same as the value of
`DefaultTTL`.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinTTL`

The minimum amount of time, in seconds, that you want objects to stay in the CloudFront
cache before CloudFront sends another request to the origin to see if the object has been
updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique name to identify the cache policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParametersInCacheKeyAndForwardedToOrigin`

The HTTP headers, cookies, and URL query strings to include in the cache key. The values
included in the cache key are also included in requests that CloudFront sends to the
origin.

_Required_: Yes

_Type_: [ParametersInCacheKeyAndForwardedToOrigin](aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::CachePolicy

CookiesConfig

All content copied from https://docs.aws.amazon.com/.
