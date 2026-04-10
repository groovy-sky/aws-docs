This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config AntennaDownlinkConfig

Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
Use an antenna downlink config in a mission profile to receive the downlink data in raw DigIF format.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SpectrumConfig" : SpectrumConfig
}

```

### YAML

```yaml

  SpectrumConfig:
    SpectrumConfig

```

## Properties

`SpectrumConfig`

Defines the spectrum configuration.

_Required_: No

_Type_: [SpectrumConfig](aws-properties-groundstation-config-spectrumconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an AntennaDownlinkConfig

The following example creates a Ground Station `AntennaDownlinkConfig`

#### JSON

```json

{
  "AntennaDownlinkConfig": {
    "SpectrumConfig": {
    "CenterFrequency": {
      "Value": 7812,
      "Units": "MHz"
    },
    "Bandwidth": {
      "Value": 30,
      "Units": "MHz"
    },
    "Polarization": "RIGHT_HAND"
    }
  }
}
```

#### YAML

```yaml

AntennaDownlinkConfig:
  SpectrumConfig:
    CenterFrequency:
      Value: 7812
      Units: MHz
    Bandwidth:
      Value: 30
      Units: MHz
    Polarization: RIGHT_HAND
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GroundStation::Config

AntennaDownlinkDemodDecodeConfig

All content copied from https://docs.aws.amazon.com/.
