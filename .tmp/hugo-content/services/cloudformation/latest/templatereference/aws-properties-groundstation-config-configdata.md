This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config ConfigData

Config objects provide information to Ground Station about how to configure the antenna and how data flows during a contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AntennaDownlinkConfig" : AntennaDownlinkConfig,
  "AntennaDownlinkDemodDecodeConfig" : AntennaDownlinkDemodDecodeConfig,
  "AntennaUplinkConfig" : AntennaUplinkConfig,
  "DataflowEndpointConfig" : DataflowEndpointConfig,
  "S3RecordingConfig" : S3RecordingConfig,
  "TelemetrySinkConfig" : TelemetrySinkConfig,
  "TrackingConfig" : TrackingConfig,
  "UplinkEchoConfig" : UplinkEchoConfig
}

```

### YAML

```yaml

  AntennaDownlinkConfig:
    AntennaDownlinkConfig
  AntennaDownlinkDemodDecodeConfig:
    AntennaDownlinkDemodDecodeConfig
  AntennaUplinkConfig:
    AntennaUplinkConfig
  DataflowEndpointConfig:
    DataflowEndpointConfig
  S3RecordingConfig:
    S3RecordingConfig
  TelemetrySinkConfig:
    TelemetrySinkConfig
  TrackingConfig:
    TrackingConfig
  UplinkEchoConfig:
    UplinkEchoConfig

```

## Properties

`AntennaDownlinkConfig`

Provides information for an antenna downlink config object.
Antenna downlink config objects are used to provide parameters for downlinks where no demodulation or decoding is performed by Ground Station (RF over IP downlinks).

_Required_: No

_Type_: [AntennaDownlinkConfig](aws-properties-groundstation-config-antennadownlinkconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AntennaDownlinkDemodDecodeConfig`

Provides information for a downlink demod decode config object.
Downlink demod decode config objects are used to provide parameters for downlinks where the Ground Station service will demodulate and decode the downlinked data.

_Required_: No

_Type_: [AntennaDownlinkDemodDecodeConfig](aws-properties-groundstation-config-antennadownlinkdemoddecodeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AntennaUplinkConfig`

Provides information for an uplink config object.
Uplink config objects are used to provide parameters for uplink contacts.

_Required_: No

_Type_: [AntennaUplinkConfig](aws-properties-groundstation-config-antennauplinkconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataflowEndpointConfig`

Provides information for a dataflow endpoint config object.
Dataflow endpoint config objects are used to provide parameters about which IP endpoint(s) to use during a contact.
Dataflow endpoints are where Ground Station sends data during a downlink contact and where Ground Station receives data to send to the satellite during an uplink contact.

_Required_: No

_Type_: [DataflowEndpointConfig](aws-properties-groundstation-config-dataflowendpointconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3RecordingConfig`

Provides information for an S3 recording config object.
S3 recording config objects are used to provide parameters for S3 recording during downlink contacts.

_Required_: No

_Type_: [S3RecordingConfig](aws-properties-groundstation-config-s3recordingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TelemetrySinkConfig`

Provides information for a telemetry sink config object.
Telemetry sink config objects are used to provide parameters for telemetry delivery during contacts.

_Required_: No

_Type_: [TelemetrySinkConfig](aws-properties-groundstation-config-telemetrysinkconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackingConfig`

Provides information for a tracking config object.
Tracking config objects are used to provide parameters about how to track the satellite through the sky during a contact.

_Required_: No

_Type_: [TrackingConfig](aws-properties-groundstation-config-trackingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UplinkEchoConfig`

Provides information for an uplink echo config object.
Uplink echo config objects are used to provide parameters for uplink echo during uplink contacts.

_Required_: No

_Type_: [UplinkEchoConfig](aws-properties-groundstation-config-uplinkechoconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a ConfigData

A `ConfigData` can only contain a single config subtype.
See the config subtype's documentation for more information on its properties.
An example of a `TrackingConfig` is provided below.

#### JSON

```json

{
  "TrackingConfig": {
    "Autotrack": "PREFERRED"
  }
}
```

#### YAML

```yaml

TrackingConfig:
  Autotrack: "PREFERRED"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AntennaUplinkConfig

DataflowEndpointConfig

All content copied from https://docs.aws.amazon.com/.
