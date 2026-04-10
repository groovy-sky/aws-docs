This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Gateway GatewayCapabilitySummary

Contains a summary of a gateway capability configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapabilityConfiguration" : String,
  "CapabilityNamespace" : String
}

```

### YAML

```yaml

  CapabilityConfiguration: String
  CapabilityNamespace: String

```

## Properties

`CapabilityConfiguration`

The JSON document that defines the configuration for the gateway capability. For more
information, see [Configuring data sources (CLI)](../../../iot-sitewise/latest/userguide/configure-sources.md#configure-source-cli) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapabilityNamespace`

The namespace of the capability configuration.
For example, if you configure OPC UA
sources for an MQTT-enabled gateway, your OPC-UA capability configuration has the namespace
`iotsitewise:opcuacollector:3`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::Gateway

GatewayPlatform

All content copied from https://docs.aws.amazon.com/.
