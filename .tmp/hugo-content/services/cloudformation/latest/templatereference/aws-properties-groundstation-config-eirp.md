This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config Eirp

Defines an equivalent isotropically radiated power (EIRP).

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

The units of the EIRP.

_Required_: No

_Type_: String

_Allowed values_: `dBW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the EIRP. Valid values are between 20.0 to 50.0 dBW.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an EIRP

The following example creates a Ground Station `EIRP`

#### JSON

```json

{
  "TargetEirp": {
    "Value": 20,
    "Units": "dBW"
  }
}
```

#### YAML

```yaml

TargetEirp:
  Value: 20.0
  Units: dBW
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DemodulationConfig

Frequency

All content copied from https://docs.aws.amazon.com/.
