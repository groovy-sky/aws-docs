This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ScheduledAudit

Use the `AWS::IoT::ScheduledAudit` resource to create a scheduled audit that
is run at a specified time interval. For API reference, see [CreateScheduleAudit](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateScheduledAudit.html)
and for general information, see [Audit](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-audit.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ScheduledAudit",
  "Properties" : {
      "DayOfMonth" : String,
      "DayOfWeek" : String,
      "Frequency" : String,
      "ScheduledAuditName" : String,
      "Tags" : [ Tag, ... ],
      "TargetCheckNames" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ScheduledAudit
Properties:
  DayOfMonth: String
  DayOfWeek: String
  Frequency: String
  ScheduledAuditName: String
  Tags:
    - Tag
  TargetCheckNames:
    - String

```

## Properties

`DayOfMonth`

The day of the month on which the scheduled audit is run (if the
`frequency` is "MONTHLY").
If days 29-31 are specified, and the month does not have that many
days, the audit takes place on the "LAST" day of the month.

_Required_: No

_Type_: String

_Pattern_: `^([1-9]|[12][0-9]|3[01])$|^LAST$|^UNSET_VALUE$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DayOfWeek`

The day of the week on which the scheduled audit is run (if the `frequency`
is "WEEKLY" or "BIWEEKLY").

_Required_: No

_Type_: String

_Allowed values_: `SUN | MON | TUE | WED | THU | FRI | SAT | UNSET_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Frequency`

How often the scheduled audit occurs.

_Required_: Yes

_Type_: String

_Allowed values_: `DAILY | WEEKLY | BIWEEKLY | MONTHLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledAuditName`

The name of the scheduled audit.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the scheduled audit.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-scheduledaudit-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetCheckNames`

Which checks are performed during the scheduled audit. Checks must be enabled for your
account. (Use `DescribeAccountAuditConfiguration` to see the list of all checks,
including those that are enabled or use `UpdateAccountAuditConfiguration` to
select which checks are enabled.)

The following checks are currently available:

- `AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`

- `CA_CERTIFICATE_EXPIRING_CHECK`

- `CA_CERTIFICATE_KEY_QUALITY_CHECK`

- `CONFLICTING_CLIENT_IDS_CHECK`

- `DEVICE_CERTIFICATE_EXPIRING_CHECK`

- `DEVICE_CERTIFICATE_KEY_QUALITY_CHECK`

- `DEVICE_CERTIFICATE_SHARED_CHECK`

- `IOT_POLICY_OVERLY_PERMISSIVE_CHECK`

- `IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK`

- `IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK`

- `LOGGING_DISABLED_CHECK`

- `REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK`

- `REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK`

- `UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the scheduled audit name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ScheduledAuditArn`

The ARN of the scheduled audit.

## Examples

In this ScheduledAudit example, all audit checks are enabled, the frequency of the
audit is weekly, and the audit will occur every Monday.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT ScheduledAudit Sample Template",
  "Resources": {
    "MyScheduledAudit": {
      "Type": "AWS::IoT::ScheduledAudit",
      "Properties": {
        "ScheduledAuditName": "MyScheduledAudit",
        "DayOfWeek": "MON",
        "Frequency": "WEEKLY",
        "TargetCheckNames": [
          "AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK",
          "CA_CERTIFICATE_EXPIRING_CHECK",
          "CA_CERTIFICATE_KEY_QUALITY_CHECK",
          "CONFLICTING_CLIENT_IDS_CHECK",
          "DEVICE_CERTIFICATE_EXPIRING_CHECK",
          "DEVICE_CERTIFICATE_KEY_QUALITY_CHECK",
          "DEVICE_CERTIFICATE_SHARED_CHECK",
          "IOT_POLICY_OVERLY_PERMISSIVE_CHECK",
          "IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK",
          "IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK",
          "LOGGING_DISABLED_CHECK",
          "REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK",
          "REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK",
          "UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT ScheduledAudit Sample Template
Resources:
  MyScheduledAudit:
    Type: AWS::IoT::ScheduledAudit
    Properties:
      ScheduledAuditName: MyScheduledAudit
      DayOfWeek: 'MON'
      Frequency: WEEKLY
      TargetCheckNames:
        - AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK
        - CA_CERTIFICATE_EXPIRING_CHECK
        - CA_CERTIFICATE_KEY_QUALITY_CHECK
        - CONFLICTING_CLIENT_IDS_CHECK
        - DEVICE_CERTIFICATE_EXPIRING_CHECK
        - DEVICE_CERTIFICATE_KEY_QUALITY_CHECK
        - DEVICE_CERTIFICATE_SHARED_CHECK
        - IOT_POLICY_OVERLY_PERMISSIVE_CHECK
        - IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK
        - IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK
        - LOGGING_DISABLED_CHECK
        - REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK
        - REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK
        - UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK
```

## See also

For more information on audit checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
