This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config SpectrumConfig

Defines a spectrum.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bandwidth" : FrequencyBandwidth,
  "CenterFrequency" : Frequency,
  "Polarization" : String
}

```

### YAML

```yaml

  Bandwidth:
    FrequencyBandwidth
  CenterFrequency:
    Frequency
  Polarization: String

```

## Properties

`Bandwidth`

The bandwidth of the spectrum. AWS Ground Station currently has the following bandwidth limitations:

- For `AntennaDownlinkDemodDecodeconfig`, valid values are between 125 kHz to 650 MHz.

- For `AntennaDownlinkconfig`, valid values are between 10 kHz to 54 MHz.

- For `AntennaUplinkConfig`, valid values are between 10 kHz to 54 MHz.

_Required_: No

_Type_: [FrequencyBandwidth](aws-properties-groundstation-config-frequencybandwidth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CenterFrequency`

The center frequency of the spectrum. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.

_Required_: No

_Type_: [Frequency](aws-properties-groundstation-config-frequency.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Polarization`

The polarization of the spectrum. Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"`. Capturing both `"RIGHT_HAND"` and `"LEFT_HAND"` polarization requires two separate configs.

_Required_: No

_Type_: String

_Allowed values_: `LEFT_HAND | RIGHT_HAND | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a SpectrumConfig

The following example creates a Ground Station `SpectrumConfig`

#### JSON

```json

{
  "SpectrumConfig": {
    "Bandwidth": {
      "Value": 30,
      "Units": "MHz"
    },
    "CenterFrequency": {
      "Value": 7812,
      "Units": "MHz"
    },
    "Polarization": "RIGHT_HAND"
  }
}
```

#### YAML

```yaml

SpectrumConfig:
  Bandwidth:
    Value: 30
    Units: MHz
  CenterFrequency:
    Value: 7812
    Units: MHz
  Polarization: RIGHT_HAND
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3RecordingConfig

Tag

All content copied from https://docs.aws.amazon.com/.
