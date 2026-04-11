This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy CustomHeader

An HTTP response header name and its value. CloudFront includes this header in HTTP
responses that it sends for requests that match a cache behavior that's associated with
this response headers policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : String,
  "Override" : Boolean,
  "Value" : String
}

```

### YAML

```yaml

  Header: String
  Override: Boolean
  Value: String

```

## Properties

`Header`

The HTTP response header name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

A Boolean that determines whether CloudFront overrides a response header with the same name
received from the origin with the header specified here.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the HTTP response header.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CorsConfig

CustomHeadersConfig

All content copied from https://docs.aws.amazon.com/.
