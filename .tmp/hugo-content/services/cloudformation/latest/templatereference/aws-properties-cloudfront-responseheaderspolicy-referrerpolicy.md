This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy ReferrerPolicy

Determines whether CloudFront includes the `Referrer-Policy` HTTP response header
and the header's value.

For more information about the `Referrer-Policy` HTTP response header, see
[Referrer-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Override" : Boolean,
  "ReferrerPolicy" : String
}

```

### YAML

```yaml

  Override: Boolean
  ReferrerPolicy: String

```

## Properties

`Override`

A Boolean that determines whether CloudFront overrides the `Referrer-Policy` HTTP
response header received from the origin with the one specified in this response headers
policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferrerPolicy`

The value of the `Referrer-Policy` HTTP response header. Valid values
are:

- `no-referrer`

- `no-referrer-when-downgrade`

- `origin`

- `origin-when-cross-origin`

- `same-origin`

- `strict-origin`

- `strict-origin-when-cross-origin`

- `unsafe-url`

For more information about these values, see [Referrer-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.

_Required_: Yes

_Type_: [String](aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.md)

_Pattern_: `^(no-referrer|no-referrer-when-downgrade|origin|origin-when-cross-origin|same-origin|strict-origin|strict-origin-when-cross-origin|unsafe-url)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameOptions

RemoveHeader

All content copied from https://docs.aws.amazon.com/.
