This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy StrictTransportSecurity

Determines whether CloudFront includes the `Strict-Transport-Security` HTTP
response header and the header's value.

For more information about the `Strict-Transport-Security` HTTP response
header, see [Strict-Transport-Security](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessControlMaxAgeSec" : Integer,
  "IncludeSubdomains" : Boolean,
  "Override" : Boolean,
  "Preload" : Boolean
}

```

### YAML

```yaml

  AccessControlMaxAgeSec: Integer
  IncludeSubdomains: Boolean
  Override: Boolean
  Preload: Boolean

```

## Properties

`AccessControlMaxAgeSec`

A number that CloudFront uses as the value for the `max-age` directive in the
`Strict-Transport-Security` HTTP response header.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSubdomains`

A Boolean that determines whether CloudFront includes the `includeSubDomains`
directive in the `Strict-Transport-Security` HTTP response header.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

A Boolean that determines whether CloudFront overrides the
`Strict-Transport-Security` HTTP response header received from the origin
with the one specified in this response headers policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Preload`

A Boolean that determines whether CloudFront includes the `preload` directive in
the `Strict-Transport-Security` HTTP response header.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerTimingHeadersConfig

XSSProtection

All content copied from https://docs.aws.amazon.com/.
