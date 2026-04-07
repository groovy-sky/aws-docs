This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Contact

The `AWS::SSMContacts::Contact` resource specifies a contact or escalation
plan. Incident Manager contacts are a subset of actions and data types that you
can use for managing responder engagement and interaction.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMContacts::Contact",
  "Properties" : {
      "Alias" : String,
      "DisplayName" : String,
      "Plan" : [ Stage, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSMContacts::Contact
Properties:
  Alias: String
  DisplayName: String
  Plan:
    - Stage
  Tags:
    - Tag
  Type: String

```

## Properties

`Alias`

The unique and identifiable alias of the contact or escalation plan.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9_\-\.]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The full name of the contact or escalation plan.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\s]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Plan`

A list of stages. A contact has an engagement plan with stages that contact specified
contact channels. An escalation plan uses stages that contact specified contacts.

_Required_: No

_Type_: Array of [Stage](aws-properties-ssmcontacts-contact-stage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-ssmcontacts-contact-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of contact.

- `PERSONAL`: A single, individual contact.

- `ESCALATION`: An escalation plan.

- `ONCALL_SCHEDULE`: An on-call schedule.

_Required_: Yes

_Type_: String

_Allowed values_: `PERSONAL | ESCALATION | ONCALL_SCHEDULE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `Contact` resource, such as
`arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias`.

## Examples

### Create Incident Manager resources

The following example demonstrates the creation of not only contacts, but also a
contact channel, a contact plan, and on-call rotation schedules. We recommend
creating all Incident Manager `Contacts` resources using a single AWS CloudFormation
template.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MySSMContact": {
            "Type": "AWS::SSMContacts::Contact",
            "Properties": {
                "Alias": "my-ssm-contact",
                "DisplayName": "My_SSM_contact",
                "Type": "PERSONAL"
            }
        },
        "MyEscalationPlan": {
            "Type": "AWS::SSMContacts::Contact",
            "Properties": {
                "Alias": "my-escalation-plan",
                "DisplayName": "My_escalation_plan",
                "Type": "ESCALATION"
            }
        },
        "MyOncallScheduleContact": {
            "Type": "AWS::SSMContacts::Contact",
            "Properties": {
                "Alias": "my-oncall-schedule",
                "DisplayName": "My_on-call_schedule",
                "Type": "ONCALL_SCHEDULE"
            }
        },
        "MyContactChannelEmail": {
            "Type": "AWS::SSMContacts::ContactChannel",
            "Properties": {
                "ContactId": {"Ref": "MySSMContact"},
                "ChannelName": "MyEmailChannel",
                "ChannelType": "EMAIL",
                "ChannelAddress": "contact-email-channel@example.com"
            }
        },
        "MySSMContactEngagementPlan": {
            "Type": "AWS::SSMContacts::Plan",
            "Properties": {
                "ContactId": {"Ref": "MySSMContact"},
                "Stages": [
                    {
                        "DurationInMinutes": 10,
                        "Targets": [
                            {
                                "ChannelTargetInfo": {
                                    "ChannelId": {"Ref": "MyContactChannelEmail"},
                                    "RetryIntervalInMinutes": 3
                                }
                            }
                        ]
                    }
                ]
            }
        },
        "MySSMContactEscalationPlan": {
            "Type": "AWS::SSMContacts::Plan",
            "Properties": {
                "ContactId": {"Ref": "MyEscalationPlan"},
                "Stages": [
                    {
                        "DurationInMinutes": 0,
                        "Targets": [
                            {
                                "ContactTargetInfo": {
                                    "ContactId": {"Ref": "MySSMContact"},
                                    "IsEssential": true
                                }
                            }
                        ]
                    }
                ]
            }
        },
        "MyDailyRotation": {
            "Type": "AWS::SSMContacts::Rotation",
            "Properties": {
                "Name": "My_daily_rotation",
                "ContactIds": [
                    {"Ref": "MySSMContact"}
                ],
                "StartTime": "2025-03-03T20:20:39",
                "TimeZoneId": "America/Los_Angeles",
                "Recurrence": {
                    "NumberOfOnCalls": 1,
                    "RecurrenceMultiplier": 1,
                    "DailySettings": ["13:00"]
                }
            }
        },
        "MyWeeklyRotation": {
            "Type": "AWS::SSMContacts::Rotation",
            "Properties": {
                "Name": "My_weekly_rotation",
                "ContactIds": [
                    {"Ref": "MySSMContact"}
                ],
                "StartTime": "2025-03-03T20:20:39",
                "TimeZoneId": "America/Los_Angeles",
                "Recurrence": {
                    "NumberOfOnCalls": 1,
                    "RecurrenceMultiplier": 1,
                    "WeeklySettings": [
                        {
                            "DayOfWeek": "MON",
                            "HandOffTime": "13:00"
                        }
                    ]
                }
            }
        },
        "MyMonthlyRotation": {
            "Type": "AWS::SSMContacts::Rotation",
            "Properties": {
                "Name": "My_monthly_rotation",
                "ContactIds": [
                    {"Ref": "MySSMContact"}
                ],
                "StartTime": "2025-03-03T20:20:39",
                "TimeZoneId": "America/Los_Angeles",
                "Recurrence": {
                    "ShiftCoverages": [
                        {
                            "DayOfWeek": "MON",
                            "CoverageTimes": [
                                {
                                    "StartTime": "10:00",
                                    "EndTime": "18:00"
                                }
                            ]
                        }
                    ],
                    "NumberOfOnCalls": 1,
                    "RecurrenceMultiplier": 1,
                    "MonthlySettings": [
                        {
                            "DayOfMonth": 10,
                            "HandOffTime": "13:00"
                        }
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MySSMContact:
    Type: 'AWS::SSMContacts::Contact'
    Properties:
      Alias: my-ssm-contact
      DisplayName: My_SSM_contact
      Type: PERSONAL
  MyEscalationPlan:
    Type: 'AWS::SSMContacts::Contact'
    Properties:
      Alias: my-escalation-plan
      DisplayName: My_escalation_plan
      Type: ESCALATION
  MyOncallScheduleContact:
    Type: 'AWS::SSMContacts::Contact'
    Properties:
      Alias: my-oncall-schedule
      DisplayName: My_on-call_schedule
      Type: ONCALL_SCHEDULE
  MyContactChannelEmail:
    Type: 'AWS::SSMContacts::ContactChannel'
    Properties:
      ContactId:
        Ref: MySSMContact
      ChannelName: MyEmailChannel
      ChannelType: EMAIL
      ChannelAddress: contact-email-channel@example.com
  MySSMContactEngagementPlan:
    Type: 'AWS::SSMContacts::Plan'
    Properties:
      ContactId:
        Ref: MySSMContact
      Stages:
        - DurationInMinutes: 10
          Targets:
            - ChannelTargetInfo:
                ChannelId:
                  Ref: MyContactChannelEmail
                RetryIntervalInMinutes: 3
  MySSMContactEscalationPlan:
    Type: 'AWS::SSMContacts::Plan'
    Properties:
      ContactId:
        Ref: MyEscalationPlan
      Stages:
        - DurationInMinutes: 0
          Targets:
            - ContactTargetInfo:
                ContactId:
                  Ref: MySSMContact
                IsEssential: true
  MyDailyRotation:
    Type: 'AWS::SSMContacts::Rotation'
    Properties:
      Name: My_daily_rotation
      ContactIds:
        - Ref: MySSMContact
      StartTime: '2025-03-03T20:20:39'
      TimeZoneId: America/Los_Angeles
      Recurrence:
        NumberOfOnCalls: 1
        RecurrenceMultiplier: 1
        DailySettings:
          - '13:00'
  MyWeeklyRotation:
    Type: 'AWS::SSMContacts::Rotation'
    Properties:
      Name: My_weekly_rotation
      ContactIds:
        - Ref: MySSMContact
      StartTime: '2025-03-03T20:20:39'
      TimeZoneId: America/Los_Angeles
      Recurrence:
        NumberOfOnCalls: 1
        RecurrenceMultiplier: 1
        WeeklySettings:
          - DayOfWeek: MON
            HandOffTime: '13:00'
  MyMonthlyRotation:
    Type: 'AWS::SSMContacts::Rotation'
    Properties:
      Name: My_monthly_rotation
      ContactIds:
        - Ref: MySSMContact
      StartTime: '2025-03-03T20:20:39'
      TimeZoneId: America/Los_Angeles
      Recurrence:
        ShiftCoverages:
          - DayOfWeek: MON
            CoverageTimes:
              - StartTime: '10:00'
                EndTime: '18:00'
        NumberOfOnCalls: 1
        RecurrenceMultiplier: 1
        MonthlySettings:
          - DayOfMonth: 10
            HandOffTime: '13:00'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Systems Manager Incident Manager Contacts

ChannelTargetInfo
