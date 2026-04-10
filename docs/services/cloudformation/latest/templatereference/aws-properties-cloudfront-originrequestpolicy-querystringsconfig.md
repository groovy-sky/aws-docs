This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginRequestPolicy QueryStringsConfig

An object that determines whether any URL query strings in viewer requests (and if so,
which query strings) are included in requests that CloudFront sends to the origin.

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

Determines whether any URL query strings in viewer requests are included in requests
that CloudFront sends to the origin. Valid values are:

- `none` – No query strings in viewer requests are included in requests that CloudFront
sends to the origin. Even when this field is set to `none`, any query
strings that are listed in a `CachePolicy` _are_
included in origin requests.

- `whitelist` – Only the query strings in viewer requests that are listed in the
`QueryStringNames` type are included in requests that CloudFront sends to
the origin.

- `all` – All query strings in viewer requests are included in requests that CloudFront
sends to the origin.

- `allExcept` – All query strings in viewer requests are included in
requests that CloudFront sends to the origin, _**except**_ for those listed in the
`QueryStringNames` type, which are not included.

_Required_: Yes

_Type_: String

_Pattern_: `^(none|whitelist|all|allExcept)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStrings`

Contains a list of query string names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginRequestPolicyConfig

AWS::CloudFront::PublicKey

All content copied from https://docs.aws.amazon.com/.
