This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign

Creates an outbound campaign.

###### Note

- For users to be able to view or edit a campaign at a later date by using the Amazon
Connect user interface, you must add the instance ID as a tag. For example, `{
                "tags": {"owner":
                "arn:aws:connect:{REGION}:{AWS_ACCOUNT_ID}:instance/{CONNECT_INSTANCE_ID}"}}`.

- After a campaign is created, you can't add/remove source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ConnectCampaignsV2::Campaign",
  "Properties" : {
      "ChannelSubtypeConfig" : ChannelSubtypeConfig,
      "CommunicationLimitsOverride" : CommunicationLimitsConfig,
      "CommunicationTimeConfig" : CommunicationTimeConfig,
      "ConnectCampaignFlowArn" : String,
      "ConnectInstanceId" : String,
      "EntryLimitsConfig" : EntryLimitsConfig,
      "Name" : String,
      "Schedule" : Schedule,
      "Source" : Source,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::ConnectCampaignsV2::Campaign
Properties:
  ChannelSubtypeConfig:
    ChannelSubtypeConfig
  CommunicationLimitsOverride:
    CommunicationLimitsConfig
  CommunicationTimeConfig:
    CommunicationTimeConfig
  ConnectCampaignFlowArn: String
  ConnectInstanceId: String
  EntryLimitsConfig:
    EntryLimitsConfig
  Name: String
  Schedule:
    Schedule
  Source:
    Source
  Tags:
    - Tag
  Type: String

```

## Properties

`ChannelSubtypeConfig`

Contains channel subtype configuration for an outbound campaign.

_Required_: No

_Type_: [ChannelSubtypeConfig](aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CommunicationLimitsOverride`

Communication limits configuration for an outbound campaign.

_Required_: No

_Type_: [CommunicationLimitsConfig](aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CommunicationTimeConfig`

Contains communication time configuration for an outbound campaign.

_Required_: No

_Type_: [CommunicationTimeConfig](aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectCampaignFlowArn`

The Amazon Resource Name (ARN) of the Amazon Connect campaign flow associated with the outbound campaign.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `20`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectInstanceId`

The identifier of the Amazon Connect instance. You can find the `instanceId` in
the ARN of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntryLimitsConfig`

Property description not available.

_Required_: No

_Type_: [EntryLimitsConfig](aws-properties-connectcampaignsv2-campaign-entrylimitsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the outbound campaign.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

Contains the schedule configuration.

_Required_: No

_Type_: [Schedule](aws-properties-connectcampaignsv2-campaign-schedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Contains source configuration.

_Required_: No

_Type_: [Source](aws-properties-connectcampaignsv2-campaign-source.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, `{ "tags": {"key1":"value1", "key2":"value2"} }`.

_Required_: No

_Type_: Array of [Tag](aws-properties-connectcampaignsv2-campaign-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of campaign.

_Required_: No

_Type_: String

_Allowed values_: `MANAGED | JOURNEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Connect Outbound Campaigns V2

AnswerMachineDetectionConfig

All content copied from https://docs.aws.amazon.com/.
