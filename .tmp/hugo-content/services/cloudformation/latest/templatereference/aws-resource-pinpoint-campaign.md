This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign

Specifies the settings for a campaign. A _campaign_ is a messaging
initiative that engages a specific segment of users for an Amazon Pinpoint
application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::Campaign",
  "Properties" : {
      "AdditionalTreatments" : [ WriteTreatmentResource, ... ],
      "ApplicationId" : String,
      "CampaignHook" : CampaignHook,
      "CustomDeliveryConfiguration" : CustomDeliveryConfiguration,
      "Description" : String,
      "HoldoutPercent" : Integer,
      "IsPaused" : Boolean,
      "Limits" : Limits,
      "MessageConfiguration" : MessageConfiguration,
      "Name" : String,
      "Priority" : Integer,
      "Schedule" : Schedule,
      "SegmentId" : String,
      "SegmentVersion" : Integer,
      "Tags" : [ Tag, ... ],
      "TemplateConfiguration" : TemplateConfiguration,
      "TreatmentDescription" : String,
      "TreatmentName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::Campaign
Properties:
  AdditionalTreatments:
    - WriteTreatmentResource
  ApplicationId: String
  CampaignHook:
    CampaignHook
  CustomDeliveryConfiguration:
    CustomDeliveryConfiguration
  Description: String
  HoldoutPercent: Integer
  IsPaused: Boolean
  Limits:
    Limits
  MessageConfiguration:
    MessageConfiguration
  Name: String
  Priority: Integer
  Schedule:
    Schedule
  SegmentId: String
  SegmentVersion: Integer
  Tags:
    - Tag
  TemplateConfiguration:
    TemplateConfiguration
  TreatmentDescription: String
  TreatmentName: String

```

## Properties

`AdditionalTreatments`

An array of requests that defines additional treatments for the campaign, in
addition to the default treatment for the campaign.

_Required_: No

_Type_: Array of [WriteTreatmentResource](aws-properties-pinpoint-campaign-writetreatmentresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the campaign is
associated with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CampaignHook`

Specifies the Lambda function to use as a code hook for a campaign.

_Required_: No

_Type_: [CampaignHook](aws-properties-pinpoint-campaign-campaignhook.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDeliveryConfiguration`

The delivery configuration settings for sending the treatment through a custom
channel. This object is required if the `MessageConfiguration` object
for the treatment specifies a `CustomMessage` object.

_Required_: No

_Type_: [CustomDeliveryConfiguration](aws-properties-pinpoint-campaign-customdeliveryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A custom description of the campaign.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoldoutPercent`

The allocated percentage of users (segment members) who shouldn't receive
messages from the campaign.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsPaused`

Specifies whether to pause the campaign. A paused campaign doesn't run unless
you resume it by changing this value to `false`. If you restart a
campaign, the campaign restarts from the beginning and not at the point you
paused it. If a campaign is running it will complete and then pause. Pause only
pauses or skips the next run for a recurring future scheduled campaign. A campaign
scheduled for immediate can't be paused.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limits`

The messaging limits for the campaign.

_Required_: No

_Type_: [Limits](aws-properties-pinpoint-campaign-limits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageConfiguration`

The message configuration settings for the treatment.

_Required_: No

_Type_: [MessageConfiguration](aws-properties-pinpoint-campaign-messageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the campaign.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

An integer between 1 and 5, inclusive, that represents the priority of the
in-app message campaign, where 1 is the highest priority and 5 is the lowest. If
there are multiple messages scheduled to be displayed at the same time, the
priority determines the order in which those messages are displayed.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule settings for the treatment.

_Required_: Yes

_Type_: [Schedule](aws-properties-pinpoint-campaign-schedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentId`

The unique identifier for the segment to associate with the campaign.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentVersion`

The version of the segment to associate with the campaign.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

The message template to use for the treatment.

_Required_: No

_Type_: [TemplateConfiguration](aws-properties-pinpoint-campaign-templateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatmentDescription`

A custom description of the treatment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatmentName`

A custom name for the treatment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string that combines the unique identifier for the
Amazon Pinpoint application with the unique identifier for the segment that the campaign
targets.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the campaign.

`CampaignId`

The unique identifier for the campaign.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::BaiduChannel

CampaignCustomMessage

All content copied from https://docs.aws.amazon.com/.
