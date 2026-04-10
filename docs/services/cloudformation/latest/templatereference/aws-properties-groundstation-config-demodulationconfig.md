This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config DemodulationConfig

Defines demodulation settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UnvalidatedJSON" : String
}

```

### YAML

```yaml

  UnvalidatedJSON: String

```

## Properties

`UnvalidatedJSON`

The demodulation settings are in JSON format and define parameters for demodulation, for example which modulation scheme (e.g. PSK, QPSK, etc.) and matched filter to use.

_Required_: No

_Type_: String

_Pattern_: `^[{}\[\]:.,"0-9A-z\-_\s]{1,8192}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a DemodulationConfig

The following example creates a Ground Station `DemodulationConfig`

#### JSON

```json

{
  "DemodulationConfig": {
    "UnvalidatedJSON": "{ \"type\":\"QPSK\", \"qpsk\":{ \"carrierFrequencyRecovery\":{ \"centerFrequency\":{ \"value\":7812, \"units\":\"MHz\" }, \"range\":{ \"value\":250, \"units\":\"kHz\" } }, \"symbolTimingRecovery\":{ \"symbolRate\":{ \"value\":15, \"units\":\"Msps\" }, \"range\":{ \"value\":0.75, \"units\":\"ksps\" }, \"matchedFilter\":{ \"type\":\"ROOT_RAISED_COSINE\", \"rolloffFactor\":0.5 } } } }"
  },
}

```

#### YAML

```yaml

DemodulationConfig:
  UnvalidatedJSON: '{
    "type":"QPSK",
    "qpsk":{
      "carrierFrequencyRecovery":{
        "centerFrequency":{
          "value":7812,
          "units":"MHz"
        },
        "range":{
          "value":250,
          "units":"kHz"
        }
      },
      "symbolTimingRecovery":{
        "symbolRate":{
          "value":15,
          "units":"Msps"
      },
      "range":{
        "value":0.75,
        "units":"ksps"
      },
      "matchedFilter":{
        "type":"ROOT_RAISED_COSINE",
        "rolloffFactor":0.5
      }
    }
  }
}'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecodeConfig

Eirp

All content copied from https://docs.aws.amazon.com/.
