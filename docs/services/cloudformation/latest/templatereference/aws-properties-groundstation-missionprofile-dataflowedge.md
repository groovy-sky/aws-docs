This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::MissionProfile DataflowEdge

A dataflow edge defines from where and to where data will flow during a contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "Source" : String
}

```

### YAML

```yaml

  Destination: String
  Source: String

```

## Properties

`Destination`

The ARN of the destination for this dataflow edge.
For example, specify the ARN of a dataflow endpoint config for a downlink edge or an antenna uplink config for an uplink edge.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The ARN of the source for this dataflow edge.
For example, specify the ARN of an antenna downlink config for a downlink edge or a dataflow endpoint config for an uplink edge.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Downlink Dataflow Edge](#aws-properties-groundstation-missionprofile-dataflowedge--examples--Downlink_Dataflow_Edge)

- [Uplink Dataflow Edge](#aws-properties-groundstation-missionprofile-dataflowedge--examples--Uplink_Dataflow_Edge)

### Downlink Dataflow Edge

The dataflow edge below shows how to specify dataflow for a downlink contact.
Data is delivered from the antenna to a dataflow endpoint.
More specifically, data is flowed:

- From the antenna using the parameters defined in this antenna downlink config:

- `arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111`

- To the dataflow endpoint defined in this dataflow endpoint config:

- `arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222`

#### JSON

```json

"DataflowEdges": [
  {
    "Source": "arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111",
    "Destination": "arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222"
  }
]

```

#### YAML

```yaml

DataflowEdges:
  - Source: arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111
    Destination: arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222

```

### Uplink Dataflow Edge

The dataflow edge below shows how to specify dataflow for an uplink contact.
Data is delivered from a dataflow endpoint to the antenna for transmission to a satellite.
More specifically, data is flowed:

- From the dataflow endpoint defined in this dataflow endpoint config:

- `arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333`

- To the antenna using the parameters defined in this antenna uplink config:

- `arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444`

#### JSON

```json

"DataflowEdges": [
  [
    "arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333",
    "arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444"
  ]
]

```

#### YAML

```yaml

DataflowEdges:
  - - arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333
    - arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GroundStation::MissionProfile

StreamsKmsKey

All content copied from https://docs.aws.amazon.com/.
