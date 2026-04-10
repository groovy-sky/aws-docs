This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config FrequencyBandwidth

Defines a bandwidth.

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

The units of the bandwidth.

_Required_: No

_Type_: String

_Allowed values_: `GHz | MHz | kHz`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the bandwidth. AWS Ground Station currently has the following bandwidth limitations:

- For `AntennaDownlinkDemodDecodeconfig`, valid values are between 125 kHz to 650 MHz.

- For `AntennaDownlinkconfig`, valid values are between 10 kHz to 54 MHz.

- For `AntennaUplinkConfig`, valid values are between 10 kHz to 54 MHz.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a Bandwidth

The following example creates a Ground Station `Bandwidth`

#### JSON

```json

{
  "Bandwidth": {
    "Value": 30,
    "Units": "MHz"
  },
}
```

#### YAML

```yaml

Bandwidth:
  Value: 30
  Units: MHz
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Frequency

KinesisDataStreamData

All content copied from https://docs.aws.amazon.com/.
