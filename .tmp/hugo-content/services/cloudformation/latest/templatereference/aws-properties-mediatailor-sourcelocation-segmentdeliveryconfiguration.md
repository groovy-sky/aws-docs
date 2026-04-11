This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::SourceLocation SegmentDeliveryConfiguration

The segment delivery configuration settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseUrl" : String,
  "Name" : String
}

```

### YAML

```yaml

  BaseUrl: String
  Name: String

```

## Properties

`BaseUrl`

The base URL of the host or path of the segment delivery server that you're using to serve segments. This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as `https://example.com/some/path`. To use a relative URL specify the relative path, such as `/some/path*`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique identifier used to distinguish between multiple segment delivery configurations in a source location.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecretsManagerAccessTokenConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
