This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy HeadersConfig

An object that determines whether any HTTP headers (and if so, which headers) are included
in the cache key and in requests that CloudFront sends to the origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderBehavior" : String,
  "Headers" : [ String, ... ]
}

```

### YAML

```yaml

  HeaderBehavior: String
  Headers:
    - String

```

## Properties

`HeaderBehavior`

Determines whether any HTTP headers are included in the cache key and in requests that CloudFront
sends to the origin. Valid values are:

- `none` – No HTTP headers are included in the cache key or in requests that CloudFront
sends to the origin. Even when this field is set to `none`, any
headers that are listed in an `OriginRequestPolicy` _are_ included in origin requests.

- `whitelist` – Only the HTTP headers that are listed in the `Headers`
type are included in the cache key and in requests that CloudFront sends to the
origin.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

Contains a list of HTTP header names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CookiesConfig

ParametersInCacheKeyAndForwardedToOrigin

All content copied from https://docs.aws.amazon.com/.
