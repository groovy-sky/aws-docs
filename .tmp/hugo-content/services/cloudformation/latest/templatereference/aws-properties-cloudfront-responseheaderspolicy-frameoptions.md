This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy FrameOptions

Determines whether CloudFront includes the `X-Frame-Options` HTTP response header
and the header's value.

For more information about the `X-Frame-Options` HTTP response header, see
[X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FrameOption" : String,
  "Override" : Boolean
}

```

### YAML

```yaml

  FrameOption: String
  Override: Boolean

```

## Properties

`FrameOption`

The value of the `X-Frame-Options` HTTP response header. Valid values are
`DENY` and `SAMEORIGIN`.

For more information about these values, see [X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.

_Required_: Yes

_Type_: String

_Pattern_: `^(DENY|SAMEORIGIN)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

A Boolean that determines whether CloudFront overrides the `X-Frame-Options` HTTP
response header received from the origin with the one specified in this response headers
policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomHeadersConfig

ReferrerPolicy

All content copied from https://docs.aws.amazon.com/.
