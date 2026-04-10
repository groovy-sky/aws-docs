This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy XSSProtection

Determines whether CloudFront includes the `X-XSS-Protection` HTTP response
header and the header's value.

For more information about the `X-XSS-Protection` HTTP response header, see
[X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModeBlock" : Boolean,
  "Override" : Boolean,
  "Protection" : Boolean,
  "ReportUri" : String
}

```

### YAML

```yaml

  ModeBlock: Boolean
  Override: Boolean
  Protection: Boolean
  ReportUri: String

```

## Properties

`ModeBlock`

A Boolean that determines whether CloudFront includes the `mode=block` directive
in the `X-XSS-Protection` header.

For more information about this directive, see [X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

A Boolean that determines whether CloudFront overrides the `X-XSS-Protection`
HTTP response header received from the origin with the one specified in this response
headers policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protection`

A Boolean that determines the value of the `X-XSS-Protection` HTTP response
header. When this setting is `true`, the value of the
`X-XSS-Protection` header is `1`. When this setting is
`false`, the value of the `X-XSS-Protection` header is
`0`.

For more information about these settings, see [X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportUri`

A reporting URI, which CloudFront uses as the value of the `report` directive in
the `X-XSS-Protection` header.

You cannot specify a `ReportUri` when `ModeBlock` is
`true`.

For more information about using a reporting URL, see [X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StrictTransportSecurity

AWS::CloudFront::StreamingDistribution

All content copied from https://docs.aws.amazon.com/.
