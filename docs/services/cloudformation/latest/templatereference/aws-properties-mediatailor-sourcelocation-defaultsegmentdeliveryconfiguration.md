This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::SourceLocation DefaultSegmentDeliveryConfiguration

The optional configuration for a server that serves segments. Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseUrl" : String
}

```

### YAML

```yaml

  BaseUrl: String

```

## Properties

`BaseUrl`

The hostname of the server that will be used to serve segments. This string must include the protocol, such as **https://**.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessConfiguration

HttpConfiguration

All content copied from https://docs.aws.amazon.com/.
