This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice FPorts

List of FPorts assigned for different LoRaWAN application packages to use.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Applications" : [ Application, ... ]
}

```

### YAML

```yaml

  Applications:
    - Application

```

## Properties

`Applications`

LoRaWAN application configuration, which can be used to perform geolocation.

_Required_: No

_Type_: Array of [Application](aws-properties-iotwireless-wirelessdevice-application.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Application

LoRaWANDevice

All content copied from https://docs.aws.amazon.com/.
