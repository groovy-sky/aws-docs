This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config UplinkSpectrumConfig

Defines a uplink spectrum.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CenterFrequency" : Frequency,
  "Polarization" : String
}

```

### YAML

```yaml

  CenterFrequency:
    Frequency
  Polarization: String

```

## Properties

`CenterFrequency`

The center frequency of the spectrum. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.

_Required_: No

_Type_: [Frequency](aws-properties-groundstation-config-frequency.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Polarization`

The polarization of the spectrum. Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"`.

_Required_: No

_Type_: String

_Allowed values_: `LEFT_HAND | RIGHT_HAND | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an uplink SpectrumConfig

The following example creates a Ground Station `SpectrumConfig` for an uplink config

#### JSON

```json

{
  "SpectrumConfig": {

    "CenterFrequency": {
      "Value": 2100,
      "Units": "MHz"
    },
    "Polarization": "RIGHT_HAND"
  }
}
```

#### YAML

```yaml

SpectrumConfig:
  CenterFrequency:
    Value: 2100
    Units: MHz
  Polarization: RIGHT_HAND
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UplinkEchoConfig

AWS::GroundStation::DataflowEndpointGroup

All content copied from https://docs.aws.amazon.com/.
