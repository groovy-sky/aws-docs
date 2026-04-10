This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginRequestPolicy HeadersConfig

An object that determines whether any HTTP headers (and if so, which headers) are
included in requests that CloudFront sends to the origin.

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

Determines whether any HTTP headers are included in requests that CloudFront sends to the
origin. Valid values are:

- `none` – No HTTP headers in viewer requests are included in requests that CloudFront
sends to the origin. Even when this field is set to `none`, any
headers that are listed in a `CachePolicy` _are_
included in origin requests.

- `whitelist` – Only the HTTP headers that are listed in the `Headers`
type are included in requests that CloudFront sends to the origin.

- `allViewer` – All HTTP headers in viewer requests are included in
requests that CloudFront sends to the origin.

- `allViewerAndWhitelistCloudFront` – All HTTP headers in viewer
requests and the additional CloudFront headers that are listed in the
`Headers` type are included in requests that CloudFront sends to the
origin. The additional headers are added by CloudFront.

- `allExcept` – All HTTP headers in viewer requests are included in
requests that CloudFront sends to the origin, _**except**_ for those listed in the `Headers` type,
which are not included.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist|allViewer|allViewerAndWhitelistCloudFront|allExcept)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

Contains a list of HTTP header names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CookiesConfig

OriginRequestPolicyConfig

All content copied from https://docs.aws.amazon.com/.
