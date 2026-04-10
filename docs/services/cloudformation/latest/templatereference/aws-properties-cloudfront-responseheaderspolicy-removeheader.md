This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy RemoveHeader

The name of an HTTP header that CloudFront removes from HTTP responses to requests that match the
cache behavior that this response headers policy is attached to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : String
}

```

### YAML

```yaml

  Header: String

```

## Properties

`Header`

The HTTP header name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferrerPolicy

RemoveHeadersConfig

All content copied from https://docs.aws.amazon.com/.
