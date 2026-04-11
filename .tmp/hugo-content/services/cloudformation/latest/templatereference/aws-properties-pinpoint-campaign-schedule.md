This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign Schedule

Specifies the schedule settings for a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "EventFilter" : CampaignEventFilter,
  "Frequency" : String,
  "IsLocalTime" : Boolean,
  "QuietTime" : QuietTime,
  "StartTime" : String,
  "TimeZone" : String
}

```

### YAML

```yaml

  EndTime: String
  EventFilter:
    CampaignEventFilter
  Frequency: String
  IsLocalTime: Boolean
  QuietTime:
    QuietTime
  StartTime: String
  TimeZone: String

```

## Properties

`EndTime`

The scheduled time, in ISO 8601 format, when the campaign ended or will
end.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventFilter`

The type of event that causes the campaign to be sent, if the value of the
`Frequency` property is `EVENT`.

_Required_: No

_Type_: [CampaignEventFilter](aws-properties-pinpoint-campaign-campaigneventfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Frequency`

Specifies how often the campaign is sent or whether the campaign is sent in
response to a specific event.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsLocalTime`

Specifies whether the start and end times for the campaign schedule use each
recipient's local time. To base the schedule on each recipient's local time, set
this value to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuietTime`

The default quiet time for the campaign. Quiet time is a specific time range
when a campaign doesn't send messages to endpoints, if all the following
conditions are met:

- The `EndpointDemographic.Timezone` property of the endpoint
is set to a valid value.

- The current time in the endpoint's time zone is later than or equal to
the time specified by the `QuietTime.Start` property for the
campaign.

- The current time in the endpoint's time zone is earlier than or equal
to the time specified by the `QuietTime.End` property for the
campaign.

If any of the preceding conditions isn't met, the endpoint will receive
messages from the campaign, even if quiet time is enabled.

_Required_: No

_Type_: [QuietTime](aws-properties-pinpoint-campaign-quiettime.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The scheduled time when the campaign began or will begin. Valid values are:
`IMMEDIATE`, to start the campaign immediately; or, a specific
time in ISO 8601 format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The starting UTC offset for the campaign schedule, if the value of the
`IsLocalTime` property is `true`. Valid values are:
`UTC, UTC+01, UTC+02, UTC+03, UTC+03:30, UTC+04, UTC+04:30, UTC+05,
                        UTC+05:30, UTC+05:45, UTC+06, UTC+06:30, UTC+07, UTC+08, UTC+09, UTC+09:30,
                        UTC+10, UTC+10:30, UTC+11, UTC+12, UTC+13, UTC-02, UTC-03, UTC-04, UTC-05,
                        UTC-06, UTC-07, UTC-08, UTC-09, UTC-10,` and
`UTC-11`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QuietTime

SetDimension

All content copied from https://docs.aws.amazon.com/.
