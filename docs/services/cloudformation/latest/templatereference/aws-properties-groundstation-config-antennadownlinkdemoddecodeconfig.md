This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config AntennaDownlinkDemodDecodeConfig

Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
Use an antenna downlink demod decode config in a mission profile to receive the downlink data that has been demodulated and decoded.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DecodeConfig" : DecodeConfig,
  "DemodulationConfig" : DemodulationConfig,
  "SpectrumConfig" : SpectrumConfig
}

```

### YAML

```yaml

  DecodeConfig:
    DecodeConfig
  DemodulationConfig:
    DemodulationConfig
  SpectrumConfig:
    SpectrumConfig

```

## Properties

`DecodeConfig`

Defines how the RF signal will be decoded.

_Required_: No

_Type_: [DecodeConfig](aws-properties-groundstation-config-decodeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DemodulationConfig`

Defines how the RF signal will be demodulated.

_Required_: No

_Type_: [DemodulationConfig](aws-properties-groundstation-config-demodulationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpectrumConfig`

Defines the spectrum configuration.

_Required_: No

_Type_: [SpectrumConfig](aws-properties-groundstation-config-spectrumconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an AntennaDownlinkDemodDecodeConfig

The following example creates a Ground Station `AntennaDownlinkDemodDecodeConfig`

#### JSON

```json

{
  "AntennaDownlinkDemodDecodeConfig": {
    "SpectrumConfig": {
      "CenterFrequency": {
        "Value": 7812,
        "Units": "MHz"
      },
      "Polarization": "RIGHT_HAND",
      "Bandwidth": {
        "Value": 30,
        "Units": "MHz"
      }
    },
    "DemodulationConfig": {
      "UnvalidatedJSON": "{ \"type\":\"QPSK\", \"qpsk\":{ \"carrierFrequencyRecovery\":{ \"centerFrequency\":{ \"value\":7812, \"units\":\"MHz\" }, \"range\":{ \"value\":250, \"units\":\"kHz\" } }, \"symbolTimingRecovery\":{ \"symbolRate\":{ \"value\":15, \"units\":\"Msps\" }, \"range\":{ \"value\":0.75, \"units\":\"ksps\" }, \"matchedFilter\":{ \"type\":\"ROOT_RAISED_COSINE\", \"rolloffFactor\":0.5 } } } }"
    },
    "DecodeConfig": {
      "UnvalidatedJSON": "{ \"edges\":[ { \"from\":\"I-Ingress\", \"to\":\"IQ-Recombiner\" }, { \"from\":\"Q-Ingress\", \"to\":\"IQ-Recombiner\" }, { \"from\":\"IQ-Recombiner\", \"to\":\"CcsdsViterbiDecoder\" }, { \"from\":\"CcsdsViterbiDecoder\", \"to\":\"NrzmDecoder\" }, { \"from\":\"NrzmDecoder\", \"to\":\"UncodedFramesEgress\" } ], \"nodeConfigs\":{ \"I-Ingress\":{ \"type\":\"CODED_SYMBOLS_INGRESS\", \"codedSymbolsIngress\":{ \"source\":\"I\" } }, \"Q-Ingress\":{ \"type\":\"CODED_SYMBOLS_INGRESS\", \"codedSymbolsIngress\":{ \"source\":\"Q\" } }, \"IQ-Recombiner\":{ \"type\":\"IQ_RECOMBINER\" }, \"CcsdsViterbiDecoder\":{ \"type\":\"CCSDS_171_133_VITERBI_DECODER\", \"ccsds171133ViterbiDecoder\":{ \"codeRate\":\"ONE_HALF\" } }, \"NrzmDecoder\":{ \"type\":\"NRZ_M_DECODER\" }, \"UncodedFramesEgress\":{ \"type\":\"UNCODED_FRAMES_EGRESS\" } } }"
    }
  }
}

```

#### YAML

```yaml

AntennaDownlinkDemodDecodeConfig:
  SpectrumConfig:
    CenterFrequency:
      Value: 7812
      Units: MHz
    Polarization: RIGHT_HAND
    Bandwidth:
      Value: 30
      Units: MHz
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
  DecodeConfig:
    UnvalidatedJSON: '{
      "edges":[
        {
          "from":"I-Ingress",
          "to":"IQ-Recombiner"
        },
        {
          "from":"Q-Ingress",
          "to":"IQ-Recombiner"
        },
        {
          "from":"IQ-Recombiner",
          "to":"CcsdsViterbiDecoder"
        },
        {
          "from":"CcsdsViterbiDecoder",
          "to":"NrzmDecoder"
        },
        {
          "from":"NrzmDecoder",
          "to":"UncodedFramesEgress"
        }
      ],
      "nodeConfigs":{
        "I-Ingress":{
          "type":"CODED_SYMBOLS_INGRESS",
          "codedSymbolsIngress":{
            "source":"I"
          }
        },
        "Q-Ingress":{
          "type":"CODED_SYMBOLS_INGRESS",
          "codedSymbolsIngress":{
            "source":"Q"
          }
        },
        "IQ-Recombiner":{
          "type":"IQ_RECOMBINER"
        },
        "CcsdsViterbiDecoder":{
          "type":"CCSDS_171_133_VITERBI_DECODER",
          "ccsds171133ViterbiDecoder":{
            "codeRate":"ONE_HALF"
          }
        },
        "NrzmDecoder":{
          "type":"NRZ_M_DECODER"
        },
        "UncodedFramesEgress":{
          "type":"UNCODED_FRAMES_EGRESS"
        }
      }
    }'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AntennaDownlinkConfig

AntennaUplinkConfig

All content copied from https://docs.aws.amazon.com/.
