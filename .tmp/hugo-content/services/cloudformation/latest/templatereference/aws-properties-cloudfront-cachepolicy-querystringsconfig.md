This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy QueryStringsConfig

An object that determines whether any URL query strings in viewer requests (and if so, which
query strings) are included in the cache key and in requests that CloudFront sends to the
origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueryStringBehavior" : String,
  "QueryStrings" : [ String, ... ]
}

```

### YAML

```yaml

  QueryStringBehavior:
    String
  QueryStrings:
    - String

```

## Properties

`QueryStringBehavior`

Determines whether any URL query strings in viewer requests are included in the cache key
and in requests that CloudFront sends to the origin. Valid values are:

- `none` – No query strings in viewer requests are included in the cache key or
in requests that CloudFront sends to the origin. Even when this field is set to
`none`, any query strings that are listed in an
`OriginRequestPolicy` _are_ included in origin
requests.

- `whitelist` – Only the query strings in viewer requests that are listed in the
`QueryStringNames` type are included in the cache key and in requests
that CloudFront sends to the origin.

- `allExcept` – All query strings in viewer requests are included in the cache
key and in requests that CloudFront sends to the origin, _**except**_ those that are listed in the
`QueryStringNames` type, which are not included.

- `all` – All query strings in viewer requests are included in the cache key and
in requests that CloudFront sends to the origin.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist|allExcept|all)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStrings`

Contains a list of query string names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParametersInCacheKeyAndForwardedToOrigin

AWS::CloudFront::CloudFrontOriginAccessIdentity

All content copied from https://docs.aws.amazon.com/.
