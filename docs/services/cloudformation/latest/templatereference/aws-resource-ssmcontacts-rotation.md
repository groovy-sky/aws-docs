This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Rotation

Specifies a rotation in an on-call schedule.

###### Note

**Template example**: We recommend creating all Incident Manager `Contacts` resources using a single AWS CloudFormation template. For a
demonstration, see the examples for [AWS::SSMContacts::Contacts](../userguide/aws-resource-ssmcontacts-contact.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMContacts::Rotation",
  "Properties" : {
      "ContactIds" : [ String, ... ],
      "Name" : String,
      "Recurrence" : RecurrenceSettings,
      "StartTime" : String,
      "Tags" : [ Tag, ... ],
      "TimeZoneId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSMContacts::Rotation
Properties:
  ContactIds:
    - String
  Name: String
  Recurrence:
    RecurrenceSettings
  StartTime: String
  Tags:
    - Tag
  TimeZoneId: String

```

## Properties

`ContactIds`

The Amazon Resource Names (ARNs) of the contacts to add to the rotation.

###### Note

Only the `PERSONAL` contact type is supported. The contact types
`ESCALATION` and `ONCALL_SCHEDULE` are not supported for this
operation.

The order in which you list the contacts is their shift order in the rotation
schedule.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the rotation.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recurrence`

Information about the rule that specifies when shift team members rotate.

_Required_: Yes

_Type_: [RecurrenceSettings](aws-properties-ssmcontacts-rotation-recurrencesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The date and time the rotation goes into effect.

_Required_: Yes

_Type_: String

_Pattern_: `^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata to assign to the rotation. Tags enable you to categorize a resource in
different ways, such as by purpose, owner, or environment. For more information, see [Tagging\
Incident Manager resources](../../../incident-manager/latest/userguide/tagging.md) in the _Incident Manager User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-ssmcontacts-rotation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZoneId`

The time zone to base the rotation’s activity on, in Internet Assigned Numbers Authority
(IANA) format. For example: "America/Los\_Angeles", "UTC", or "Asia/Seoul". For more
information, see the [Time Zone\
Database](https://www.iana.org/time-zones) on the IANA website.

###### Note

Designators for time zones that don’t support Daylight Savings Time rules, such as
Pacific Standard Time (PST), are not supported.

_Required_: Yes

_Type_: String

_Pattern_: `^[:a-zA-Z0-9_\-\s\.\\/]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `Rotation` resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Targets

CoverageTime

All content copied from https://docs.aws.amazon.com/.
