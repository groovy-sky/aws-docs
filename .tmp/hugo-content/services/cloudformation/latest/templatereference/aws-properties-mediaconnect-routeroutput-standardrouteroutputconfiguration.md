This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput StandardRouterOutputConfiguration

The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkInterfaceArn" : String,
  "Protocol" : String,
  "ProtocolConfiguration" : RouterOutputProtocolConfiguration
}

```

### YAML

```yaml

  NetworkInterfaceArn: String
  Protocol: String
  ProtocolConfiguration:
    RouterOutputProtocolConfiguration

```

## Properties

`NetworkInterfaceArn`

The Amazon Resource Name (ARN) of the network interface associated with the standard router output.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:routerNetworkInterface:[a-z0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol used by the standard router output.

_Required_: No

_Type_: String

_Allowed values_: `RTP | RIST | SRT_CALLER | SRT_LISTENER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolConfiguration`

The configuration settings for the protocol used by the standard router output.

_Required_: Yes

_Type_: [RouterOutputProtocolConfiguration](aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtListenerRouterOutputConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
