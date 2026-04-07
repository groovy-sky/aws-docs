This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::NetworkAnalyzerConfiguration TraceContent

Trace content for your wireless gateway and wireless device resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogLevel" : String,
  "WirelessDeviceFrameInfo" : String
}

```

### YAML

```yaml

  LogLevel: String
  WirelessDeviceFrameInfo: String

```

## Properties

`LogLevel`

The log level for a log message. The log levels can be disabled, or set to
`ERROR` to display less verbose logs containing only error information, or to
`INFO` for more detailed logs

_Required_: No

_Type_: String

_Allowed values_: `INFO | ERROR | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WirelessDeviceFrameInfo`

`FrameInfo` of your wireless device resources for the trace content. Use
FrameInfo to debug the communication between your LoRaWAN end devices and the network
server.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IoTWireless::PartnerAccount
