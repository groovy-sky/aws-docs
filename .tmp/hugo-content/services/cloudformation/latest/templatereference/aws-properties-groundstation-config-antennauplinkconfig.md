This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config AntennaUplinkConfig

Provides information about how AWS Ground Station should configure an antenna for uplink during a contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SpectrumConfig" : UplinkSpectrumConfig,
  "TargetEirp" : Eirp,
  "TransmitDisabled" : Boolean
}

```

### YAML

```yaml

  SpectrumConfig:
    UplinkSpectrumConfig
  TargetEirp:
    Eirp
  TransmitDisabled: Boolean

```

## Properties

`SpectrumConfig`

Defines the spectrum configuration.

_Required_: No

_Type_: [UplinkSpectrumConfig](aws-properties-groundstation-config-uplinkspectrumconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetEirp`

The equivalent isotropically radiated power (EIRP) to use for uplink transmissions. Valid values are between 20.0 to 50.0 dBW.

_Required_: No

_Type_: [Eirp](aws-properties-groundstation-config-eirp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransmitDisabled`

Whether or not uplink transmit is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an AntennaUplinkConfig

The following example creates a Ground Station `AntennaUplinkConfig`

#### JSON

```json

{
  "AntennaUplinkConfig": {
    "SpectrumConfig": {
    "CenterFrequency": {
      "Value": 2072.5,
      "Units": "MHz"
    },
    "Polarization": "RIGHT_HAND"
    },
    "TargetEirp": {
      "Value": 20,
      "Units": "dBW"
    }
  }
}
```

#### YAML

```yaml

AntennaUplinkConfig:
  SpectrumConfig:
    CenterFrequency:
      Value: 2072.5
      Units: MHz
    Polarization: RIGHT_HAND
  TargetEirp:
    Value: 20.0
    Units: dBW
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AntennaDownlinkDemodDecodeConfig

ConfigData

All content copied from https://docs.aws.amazon.com/.
