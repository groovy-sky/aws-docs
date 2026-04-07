This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::IngestConfiguration

The `AWS::IVS::IngestConfiguration` resource specifies an ingest protocol to be used for a stage.
For more information, see [Stream Ingest](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/rt-stream-ingest.html)
in the _Amazon IVS Real-Time Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::IngestConfiguration",
  "Properties" : {
      "IngestProtocol" : String,
      "InsecureIngest" : Boolean,
      "Name" : String,
      "StageArn" : String,
      "Tags" : [ Tag, ... ],
      "UserId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IVS::IngestConfiguration
Properties:
  IngestProtocol: String
  InsecureIngest: Boolean
  Name: String
  StageArn: String
  Tags:
    - Tag
  UserId: String

```

## Properties

`IngestProtocol`

Type of ingest protocol that the user employs for broadcasting.

_Required_: No

_Type_: String

_Allowed values_: `RTMP | RTMPS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InsecureIngest`

Whether the channel allows insecure RTMP ingest. Default: `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Ingest name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageArn`

ARN of the stage with which the IngestConfiguration is associated.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:ivs:[a-z0-9-]+:[0-9]+:stage/[a-zA-Z0-9-]+$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-ingestconfiguration-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserId`

Customer-assigned name to help identify the participant using the IngestConfiguration;
this can be used to link a participant to a user in the customer’s own systems. This can be any UTF-8 encoded text.
_This field is exposed to all stage participants and should not be used for_
_personally identifying, confidential, or sensitive information._

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ingest configuration ARN. For example:

`{ "Ref": "myIngestConfiguration" }`

For the Amazon IVS ingest configuration
`myIngestConfiguration`, `Ref` returns the ingest configuration
ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ingest-configuration ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:ingest-configuration/abcdABCDefgh`

`ParticipantId`

ID of the participant within the stage. For example: `abCDEf12GHIj`

`State`

State of the ingest configuration. It is `ACTIVE` if a publisher currently is publishing to
the stage associated with the ingest configuration. Valid values: `ACTIVE` \| `INACTIVE`.

`StreamKey`

Ingest-key value for the RTMP(S) protocol. For example: `skSKABCDefgh`

## Examples

### Ingest Configuration Template Examples

The following examples specify an Amazon IVS
ingest configuration.

#### JSON

```json

{
     "AWSTemplateFormatVersion": "2010-09-09",
     "Resources": {
         "Stage": {
             "Type": "AWS::IVS::Stage",
             "Properties": {
                 "Name": "myStage",
                 "Tags": [
                     {
                         "Key": "MyKey",
                         "Value": "MyValue"
                     }
                 ]
             }
         },
         "IngestConfiguration": {
             "Type": "AWS::IVS::IngestConfiguration",
             "Properties": {
                 "Name": "myIngest",
                 "StageArn": {"Ref": "Stage"},
                 "IngestProtocol": "RTMPS",
                 "InsecureIngest": false,
                 "UserId": "myUser",
                 "Tags": [
                     {
                         "Key": "MyKey",
                         "Value": "MyValue"
                     }
                 ]
             }
         }
     },
     "Outputs": {
         "IngestConfigurationArn": {
             "Value": {"Ref": "IngestConfiguration"}
         },
         "IngestConfigurationState": {
             "Value": {
                 "Fn::GetAtt": [
                     "IngestConfiguration",
                     "State"
                 ]
             }
         },
         "IngestConfigurationStreamKey": {
             "Value": {
                 "Fn::GetAtt": [
                     "IngestConfiguration",
                     "StreamKey"
                 ]
             }
         },
         "IngestConfigurationParticipantId": {
             "Value": {
                 "Fn::GetAtt": [
                     "IngestConfiguration",
                     "ParticipantId"
                 ]
			 }
         }
     }
 }
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  Stage:
    Type: AWS::IVS::Stage
    Properties:
      Name: MyStage
      Tags:
        - Key: MyKey
          Value: MyValue

  IngestConfiguration:
    Type: AWS::IVS::IngestConfiguration
    Properties:
      Name: myIngestName
      IngestProtocol: RTMPS
      InsecureIngest: false
      UserId: myUser
      StageArn: !Ref Stage
      Tags:
        - Key: MyKey
          Value: MyValue

Outputs:
  IngestConfigurationArn:
    Value: !Ref IngestConfiguration

  IngestConfigurationState:
    Value: !GetAtt IngestConfiguration.State

  IngestConfigurationStreamKey:
    Value: !GetAtt IngestConfiguration.StreamKey

  IngestConfigurationParticipantId:
    Value: !GetAtt IngestConfiguration.ParticipantId
```

## See also

- [IVS Stream Ingest](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/rt-stream-ingest.html)

- [IngestConfiguration](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_IngestConfiguration.html) data type

- [CreateIngestConfiguration](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_CreateIngestConfiguration.html) API endpoint

- [DeleteIngestConfiguration](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_DeleteIngestConfiguration.html) API endpoint

- [GetIngestConfiguration](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_GetIngestConfiguration.html) API endpoint

- [ListIngestConfigurations](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_ListIngestConfigurations.html) API endpoint

- [UpdateIngestConfiguration](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_UpdateIngestConfiguration.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Video

Tag
