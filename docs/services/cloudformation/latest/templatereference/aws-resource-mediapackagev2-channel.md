This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::Channel

Creates a channel to receive content.

After it's created, a channel provides static input URLs. These URLs remain the same throughout the lifetime of the channel, regardless of any failures or upgrades that might occur. Use these URLs to configure the outputs of your upstream encoder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackageV2::Channel",
  "Properties" : {
      "ChannelGroupName" : String,
      "ChannelName" : String,
      "Description" : String,
      "InputSwitchConfiguration" : InputSwitchConfiguration,
      "InputType" : String,
      "OutputHeaderConfiguration" : OutputHeaderConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackageV2::Channel
Properties:
  ChannelGroupName: String
  ChannelName: String
  Description: String
  InputSwitchConfiguration:
    InputSwitchConfiguration
  InputType: String
  OutputHeaderConfiguration:
    OutputHeaderConfiguration
  Tags:
    - Tag

```

## Properties

`ChannelGroupName`

The name of the channel group associated with the channel configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelName`

The name of the channel.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the channel.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSwitchConfiguration`

The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.

_Required_: No

_Type_: [InputSwitchConfiguration](aws-properties-mediapackagev2-channel-inputswitchconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputType`

The input type will be an immutable field which will be used to define whether the channel will allow CMAF ingest or HLS ingest. If unprovided, it will default to HLS to preserve current behavior.

The allowed values are:

- `HLS` \- The HLS streaming specification (which defines M3U8 manifests and TS segments).

- `CMAF` \- The DASH-IF CMAF Ingest specification (which defines CMAF segments with optional DASH manifests).

_Required_: No

_Type_: String

_Allowed values_: `HLS | CMAF`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputHeaderConfiguration`

The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.

_Required_: No

_Type_: [OutputHeaderConfiguration](aws-properties-mediapackagev2-channel-outputheaderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackagev2-channel-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `arn:aws:mediapackagev2:region:AccountId:ChannelGroup/ChannelGroupName/Channel/ChannelName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The attributes of the channels. You can only use the `GetAtt` function for `readOnly` properties. For example, you can use the `GetAtt` function to retrieve the read-only `CreatedAt` property.

`Arn`

The Amazon Resource Name (ARN) of the channel.

`CreatedAt`

The timestamp of the creation of the channel.

`IngestEndpoints`

The ingest endpoints associated with the channel.

`IngestEndpointUrls`

The ingest domain URL where the source stream should be sent.

`ModifiedAt`

The timestamp of the modification of the channel.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaPackage V2

IngestEndpoint

All content copied from https://docs.aws.amazon.com/.
