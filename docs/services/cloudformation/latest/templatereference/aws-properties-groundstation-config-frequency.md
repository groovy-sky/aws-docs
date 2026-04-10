This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config Frequency

Defines a frequency.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Units" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Units: String
  Value: Number

```

## Properties

`Units`

The units of the frequency.

_Required_: No

_Type_: String

_Allowed values_: `GHz | MHz | kHz`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the frequency. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a Frequency

The following example creates a Ground Station `Frequency`

#### JSON

```json

{
  "CenterFrequency": {
    "Value": 7812,
    "Units": "MHz"
  },
}
```

#### YAML

```yaml

CenterFrequency:
  Value: 7812
  Units: MHz
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Eirp

FrequencyBandwidth

All content copied from https://docs.aws.amazon.com/.
