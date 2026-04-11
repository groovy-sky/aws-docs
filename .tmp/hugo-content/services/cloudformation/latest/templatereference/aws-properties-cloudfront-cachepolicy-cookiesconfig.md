This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy CookiesConfig

An object that determines whether any cookies in viewer requests (and if so, which cookies)
are included in the cache key and in requests that CloudFront sends to the origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CookieBehavior" : String,
  "Cookies" : [ String, ... ]
}

```

### YAML

```yaml

  CookieBehavior: String
  Cookies:
    - String

```

## Properties

`CookieBehavior`

Determines whether any cookies in viewer requests are included in the cache key and in
requests that CloudFront sends to the origin. Valid values are:

- `none` – No cookies in viewer requests are included in the cache key or in
requests that CloudFront sends to the origin. Even when this field is set to
`none`, any cookies that are listed in an
`OriginRequestPolicy` _are_ included in origin
requests.

- `whitelist` – Only the cookies in viewer requests that are listed in the
`CookieNames` type are included in the cache key and in requests that
CloudFront sends to the origin.

- `allExcept` – All cookies in viewer requests are included in the cache key and
in requests that CloudFront sends to the origin, _**except**_ for those that are listed in the
`CookieNames` type, which are not included.

- `all` – All cookies in viewer requests are included in the cache key and in
requests that CloudFront sends to the origin.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist|allExcept|all)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cookies`

Contains a list of cookie names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CachePolicyConfig

HeadersConfig

All content copied from https://docs.aws.amazon.com/.
