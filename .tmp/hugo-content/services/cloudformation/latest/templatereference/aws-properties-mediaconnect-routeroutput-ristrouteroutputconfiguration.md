This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput RistRouterOutputConfiguration

The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationAddress" : String,
  "DestinationPort" : Integer
}

```

### YAML

```yaml

  DestinationAddress: String
  DestinationPort: Integer

```

## Properties

`DestinationAddress`

The destination IP address for the RIST protocol in the router output configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPort`

The destination port number for the RIST protocol in the router output configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PreferredDayTimeMaintenanceConfiguration

RouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
