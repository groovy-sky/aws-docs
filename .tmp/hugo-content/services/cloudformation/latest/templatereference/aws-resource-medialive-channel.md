This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel

The AWS::MediaLive::Channel resource is a MediaLive resource type that creates a channel.

A MediaLive channel ingests and transcodes (decodes and encodes) source content from the inputs that are
attached to that channel, and packages the new content into outputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Channel",
  "Properties" : {
      "AnywhereSettings" : AnywhereSettings,
      "CdiInputSpecification" : CdiInputSpecification,
      "ChannelClass" : String,
      "ChannelEngineVersion" : ChannelEngineVersionRequest,
      "ChannelSecurityGroups" : [ String, ... ],
      "Destinations" : [ OutputDestination, ... ],
      "DryRun" : Boolean,
      "EncoderSettings" : EncoderSettings,
      "InferenceSettings" : InferenceSettings,
      "InputAttachments" : [ InputAttachment, ... ],
      "InputSpecification" : InputSpecification,
      "LinkedChannelSettings" : LinkedChannelSettings,
      "LogLevel" : String,
      "Maintenance" : MaintenanceCreateSettings,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "Vpc" : VpcOutputSettings
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Channel
Properties:
  AnywhereSettings:
    AnywhereSettings
  CdiInputSpecification:
    CdiInputSpecification
  ChannelClass: String
  ChannelEngineVersion:
    ChannelEngineVersionRequest
  ChannelSecurityGroups:
    - String
  Destinations:
    - OutputDestination
  DryRun: Boolean
  EncoderSettings:
    EncoderSettings
  InferenceSettings:
    InferenceSettings
  InputAttachments:
    - InputAttachment
  InputSpecification:
    InputSpecification
  LinkedChannelSettings:
    LinkedChannelSettings
  LogLevel: String
  Maintenance:
    MaintenanceCreateSettings
  Name: String
  RoleArn: String
  Tags:
    - Tag
  Vpc:
    VpcOutputSettings

```

## Properties

`AnywhereSettings`

Property description not available.

_Required_: No

_Type_: [AnywhereSettings](aws-properties-medialive-channel-anywheresettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CdiInputSpecification`

Specification of CDI inputs for this channel.

_Required_: No

_Type_: [CdiInputSpecification](aws-properties-medialive-channel-cdiinputspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelClass`

The class for this channel. For a channel with two pipelines, the class is STANDARD. For a channel with one
pipeline, the class is SINGLE\_PIPELINE.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelEngineVersion`

Property description not available.

_Required_: No

_Type_: [ChannelEngineVersionRequest](aws-properties-medialive-channel-channelengineversionrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelSecurityGroups`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destinations`

The settings that identify the destination for the outputs in this MediaLive output package.

_Required_: No

_Type_: Array of [OutputDestination](aws-properties-medialive-channel-outputdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DryRun`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncoderSettings`

The encoding configuration for the output content.

_Required_: No

_Type_: [EncoderSettings](aws-properties-medialive-channel-encodersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceSettings`

Property description not available.

_Required_: No

_Type_: [InferenceSettings](aws-properties-medialive-channel-inferencesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputAttachments`

The list of input attachments for the channel.

_Required_: No

_Type_: Array of [InputAttachment](aws-properties-medialive-channel-inputattachment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSpecification`

The input specification for this channel. It specifies the key characteristics of the inputs for this channel:
the maximum bitrate, the resolution, and the codec.

_Required_: No

_Type_: [InputSpecification](aws-properties-medialive-channel-inputspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkedChannelSettings`

Property description not available.

_Required_: No

_Type_: [LinkedChannelSettings](aws-properties-medialive-channel-linkedchannelsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

The verbosity for logging activity for this channel. Charges for logging (which are generated through Amazon
CloudWatch Logging) are higher for higher verbosities.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Maintenance`

Maintenance settings for this channel.

_Required_: No

_Type_: [MaintenanceCreateSettings](aws-properties-medialive-channel-maintenancecreatesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role for MediaLive to assume when running this channel. The role is identified by its ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A collection of tags for this channel. Each tag is a key-value pair.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vpc`

Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.

_Required_: No

_Type_: [VpcOutputSettings](aws-properties-medialive-channel-vpcoutputsettings.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the channel.

For example: `{ "Ref": "myChannel" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the MediaLive channel. For example:
arn:aws:medialive:us-west-1:111122223333:medialive:channel:1234567

`Id`

User-specified id. This is used in an output group or an output.

`Inputs`

The inputs that are attached to this channel. The inputs are identified by their IDs (not by their names or
their ARNs).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaLive

AacSettings

All content copied from https://docs.aws.amazon.com/.
