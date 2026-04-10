This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::MissionProfile

Mission profiles specify parameters and provide references to config objects to define how Ground Station lists and executes contacts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GroundStation::MissionProfile",
  "Properties" : {
      "ContactPostPassDurationSeconds" : Integer,
      "ContactPrePassDurationSeconds" : Integer,
      "DataflowEdges" : [ DataflowEdge, ... ],
      "MinimumViableContactDurationSeconds" : Integer,
      "Name" : String,
      "StreamsKmsKey" : StreamsKmsKey,
      "StreamsKmsRole" : String,
      "Tags" : [ Tag, ... ],
      "TelemetrySinkConfigArn" : String,
      "TrackingConfigArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::GroundStation::MissionProfile
Properties:
  ContactPostPassDurationSeconds: Integer
  ContactPrePassDurationSeconds: Integer
  DataflowEdges:
    - DataflowEdge
  MinimumViableContactDurationSeconds: Integer
  Name: String
  StreamsKmsKey:
    StreamsKmsKey
  StreamsKmsRole: String
  Tags:
    - Tag
  TelemetrySinkConfigArn: String
  TrackingConfigArn: String

```

## Properties

`ContactPostPassDurationSeconds`

Amount of time in seconds after a contact ends that you’d like to receive a Ground Station Contact State Change indicating the pass has finished.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `21600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactPrePassDurationSeconds`

Amount of time in seconds prior to contact start that you'd like to receive a Ground Station Contact State Change Event indicating an upcoming pass.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `21600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataflowEdges`

A list containing lists of config ARNs. Each list of config ARNs is an edge, with a "from" config and a "to" config.

_Required_: Yes

_Type_: Array of [DataflowEdge](aws-properties-groundstation-missionprofile-dataflowedge.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumViableContactDurationSeconds`

Minimum length of a contact in seconds that Ground Station will return when listing contacts.
Ground Station will not return contacts shorter than this duration.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `21600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the mission profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamsKmsKey`

KMS key to use for encrypting streams.

_Required_: No

_Type_: [StreamsKmsKey](aws-properties-groundstation-missionprofile-streamskmskey.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamsKmsRole`

Role to use for encrypting streams with KMS key.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9-.]{1,63}:iam::[0-9]{12}:role/[\w+=,.@-]{1,64}`

_Minimum_: `30`

_Maximum_: `165`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags assigned to the mission profile.

_Required_: No

_Type_: Array of [Tag](aws-properties-groundstation-missionprofile-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TelemetrySinkConfigArn`

The ARN of a TelemetrySinkConfig object that defines how to deliver telemetry data during a contact.

When specified, telemetry data will be delivered to the configured sink (such as Amazon Kinesis Data Streams) according to the settings defined in the referenced TelemetrySinkConfig.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackingConfigArn`

The ARN of a tracking config objects that defines how to track the satellite through the sky during a contact.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the mission profile. For example:

`{ "Ref": "MissionProfile" }`

`Ref` returns the ARN of the mission profile.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the mission profile, such as
`arn:aws:groundstation:us-east-2:012345678910:mission-profile/9940bf3b-d2ba-427e-9906-842b5e5d2296`.

`Id`

The ID of the mission profile, such as
`9940bf3b-d2ba-427e-9906-842b5e5d2296`.

`Region`

The region of the mission profile.

## Examples

### Mission Profile With Downlink And Uplink

The following example shows how to create a mission profile that utilizes both downlink and uplink during a contact.

The resulting mission profile will have a name of `My Mission Profile`.
CloudWatch Events will be sent `120` seconds before the contact starts and `180` seconds after the contact completes.
No contacts less than `300` seconds in duration will be returned when listing available contacts with this mission profile selected.

The satellite will be tracked as it moves through the sky using the method specified by the referenced tracking config.

The first dataflow edge in the list shows how to specify dataflow for a downlink contact.
Data is delivered from the antenna to a dataflow endpoint.
More specifically, data is flowed:

- From the antenna using the parameters defined in this antenna downlink config:

- `arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111`

- To the dataflow endpoint defined in this dataflow endpoint config:

- `arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222`

The second dataflow edge in the list shows how to specify dataflow for an uplink contact.
Data is delivered from a dataflow endpoint to the antenna for transmission to a satellite.
More specifically, data is flowed:

- From the dataflow endpoint defined in this dataflow endpoint config:

- `arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333`

- To the antenna using the parameters defined in this antenna uplink config:

- `arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444`

#### JSON

```json

{
  "Resources": {
    "MyMissionProfile": {
      "Type": "AWS::GroundStation::MissionProfile",
      "Properties": {
        "Name": "My Mission Profile",
        "ContactPrePassDurationSeconds": 120,
        "ContactPostPassDurationSeconds": 180,
        "MinimumViableContactDurationSeconds": 300,
        "TrackingConfigArn": "arn:aws:groundstation:us-east-2:012345678910:config/tracking/00000000-0000-0000-0000-000000000000",
        "TelemetrySinkConfigArn": "arn:aws:groundstation:us-east-2:012345678910:config/telemetry-sink/55555555-5555-5555-5555-555555555555",
        "DataflowEdges": [
          {
            "Source": "arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111",
            "Destination": "arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222"
          },
          {
            "Source": "arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333",
            "Destination": "arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444"
          }
        ]
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  MyMissionProfile:
    Type: AWS::GroundStation::MissionProfile
    Properties:
      Name: "Mission Profile"
      ContactPrePassDurationSeconds: 120
      ContactPostPassDurationSeconds: 180
      MinimumViableContactDurationSeconds: 300
      TrackingConfigArn: arn:aws:groundstation:us-east-2:012345678910:config/tracking/00000000-0000-0000-0000-000000000000
      TelemetrySinkConfigArn: arn:aws:groundstation:us-east-2:012345678910:config/telemetry-sink/55555555-5555-5555-5555-555555555555
      DataflowEdges:
        - Source: arn:aws:groundstation:us-east-2:012345678910:config/antenna-downlink/11111111-1111-1111-1111-111111111111
          Destination: arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/22222222-2222-2222-2222-222222222222
        - Source: arn:aws:groundstation:us-east-2:012345678910:config/dataflow-endpoint/33333333-3333-3333-3333-333333333333
          Destination: arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/44444444-4444-4444-4444-444444444444

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UplinkDataflowDetails

DataflowEdge

All content copied from https://docs.aws.amazon.com/.
