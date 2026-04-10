This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge EgressGatewayBridge

Create a bridge with the egress bridge type. An egress bridge is a cloud-to-ground bridge. The content comes from an existing MediaConnect flow and is delivered to your premises.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxBitrate" : Integer
}

```

### YAML

```yaml

  MaxBitrate: Integer

```

## Properties

`MaxBitrate`

The maximum expected bitrate (in bps) of the egress bridge.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BridgeSource

FailoverConfig

All content copied from https://docs.aws.amazon.com/.
