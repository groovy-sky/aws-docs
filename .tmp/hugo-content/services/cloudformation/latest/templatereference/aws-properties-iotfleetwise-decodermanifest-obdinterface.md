This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest ObdInterface

A network interface that specifies the On-board diagnostic (OBD) II network protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DtcRequestIntervalSeconds" : String,
  "HasTransmissionEcu" : String,
  "Name" : String,
  "ObdStandard" : String,
  "PidRequestIntervalSeconds" : String,
  "RequestMessageId" : String,
  "UseExtendedIds" : String
}

```

### YAML

```yaml

  DtcRequestIntervalSeconds: String
  HasTransmissionEcu: String
  Name: String
  ObdStandard: String
  PidRequestIntervalSeconds: String
  RequestMessageId: String
  UseExtendedIds: String

```

## Properties

`DtcRequestIntervalSeconds`

The maximum number message requests per diagnostic trouble code per second.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HasTransmissionEcu`

Whether the vehicle has a transmission control module (TCM).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the interface.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObdStandard`

The standard OBD II PID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PidRequestIntervalSeconds`

The maximum number message requests per second.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestMessageId`

The ID of the message requesting vehicle data.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseExtendedIds`

Whether to use extended IDs in the message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkInterfacesItems

ObdNetworkInterface

All content copied from https://docs.aws.amazon.com/.
