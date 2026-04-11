This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config DecodeConfig

Defines decoding settings.

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

The decoding settings are in JSON format and define a set of steps to perform to decode the data.

_Required_: No

_Type_: String

_Pattern_: `^[{}\[\]:.,"0-9A-z\-_\s]{1,8192}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a DecodeConfig

The following example creates a Ground Station `DecodeConfig`

#### JSON

```json

{
  "DecodeConfig": {
    "UnvalidatedJSON": "{ \"edges\":[ { \"from\":\"I-Ingress\", \"to\":\"IQ-Recombiner\" }, { \"from\":\"Q-Ingress\", \"to\":\"IQ-Recombiner\" }, { \"from\":\"IQ-Recombiner\", \"to\":\"CcsdsViterbiDecoder\" }, { \"from\":\"CcsdsViterbiDecoder\", \"to\":\"NrzmDecoder\" }, { \"from\":\"NrzmDecoder\", \"to\":\"UncodedFramesEgress\" } ], \"nodeConfigs\":{ \"I-Ingress\":{ \"type\":\"CODED_SYMBOLS_INGRESS\", \"codedSymbolsIngress\":{ \"source\":\"I\" } }, \"Q-Ingress\":{ \"type\":\"CODED_SYMBOLS_INGRESS\", \"codedSymbolsIngress\":{ \"source\":\"Q\" } }, \"IQ-Recombiner\":{ \"type\":\"IQ_RECOMBINER\" }, \"CcsdsViterbiDecoder\":{ \"type\":\"CCSDS_171_133_VITERBI_DECODER\", \"ccsds171133ViterbiDecoder\":{ \"codeRate\":\"ONE_HALF\" } }, \"NrzmDecoder\":{ \"type\":\"NRZ_M_DECODER\" }, \"UncodedFramesEgress\":{ \"type\":\"UNCODED_FRAMES_EGRESS\" } } }"
  }
}

```

#### YAML

```yaml

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

DataflowEndpointConfig

DemodulationConfig

All content copied from https://docs.aws.amazon.com/.
