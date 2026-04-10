This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginRequestPolicy CookiesConfig

An object that determines whether any cookies in viewer requests (and if so, which
cookies) are included in requests that CloudFront sends to the origin.

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

Determines whether cookies in viewer requests are included in requests that CloudFront sends
to the origin. Valid values are:

- `none` – No cookies in viewer requests are included in requests that CloudFront sends
to the origin. Even when this field is set to `none`, any cookies
that are listed in a `CachePolicy` _are_ included
in origin requests.

- `whitelist` – Only the cookies in viewer requests that are listed in the
`CookieNames` type are included in requests that CloudFront sends to the
origin.

- `all` – All cookies in viewer requests are included in requests
that CloudFront sends to the origin.

- `allExcept` – All cookies in viewer requests are included in
requests that CloudFront sends to the origin, _**except**_ for those listed in the `CookieNames`
type, which are not included.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist|all|allExcept)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cookies`

Contains a list of cookie names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::OriginRequestPolicy

HeadersConfig

All content copied from https://docs.aws.amazon.com/.
