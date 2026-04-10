This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::Channel HlsIngest

HLS ingest configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ingestEndpoints" : [ IngestEndpoint, ... ]
}

```

### YAML

```yaml

  ingestEndpoints:
    - IngestEndpoint

```

## Properties

`ingestEndpoints`

The input URL where the source stream should be sent.

_Required_: No

_Type_: Array of [IngestEndpoint](aws-properties-mediapackage-channel-ingestendpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaPackage::Channel

IngestEndpoint

All content copied from https://docs.aws.amazon.com/.
