# GetAlarmMuteRule

Retrieves details for a specific alarm mute rule.

This operation returns complete information about the mute rule, including its configuration, status, targeted alarms, and metadata.

The returned status indicates the current state of the mute rule:

- **SCHEDULED**: The mute rule is configured and will become active in the future

- **ACTIVE**: The mute rule is currently muting alarm actions

- **EXPIRED**: The mute rule has passed its expiration date and will no longer become active

**Permissions**

To retrieve details for a mute rule, you need the `cloudwatch:GetAlarmMuteRule` permission on the alarm mute rule resource.

## Request Parameters

**AlarmMuteRuleName**

The name of the alarm mute rule to retrieve.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Response Elements

The following elements are returned by the service.

**AlarmMuteRuleArn**

The Amazon Resource Name (ARN) of the alarm mute rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

**Description**

The description of the alarm mute rule.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**ExpireDate**

The date and time when the mute rule expires and is no longer evaluated.

Type: Timestamp

**LastUpdatedTimestamp**

The date and time when the mute rule was last updated.

Type: Timestamp

**MuteTargets**

Specifies which alarms this rule applies to.

Type: [MuteTargets](api-mutetargets.md) object

**MuteType**

Indicates whether the mute rule is one-time or recurring. Valid values are `ONE_TIME` or `RECURRING`.

Type: String

**Name**

The name of the alarm mute rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

**Rule**

The configuration that defines when and how long alarms are muted.

Type: [Rule](api-rule.md) object

**StartDate**

The date and time when the mute rule becomes active. If not set, the rule is active immediately.

Type: Timestamp

**Status**

The current status of the alarm mute rule. Valid values are `SCHEDULED`, `ACTIVE`, or `EXPIRED`.

Type: String

Valid Values: `SCHEDULED | ACTIVE | EXPIRED`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## Examples

### Get details for a mute rule

Retrieve complete details for a specific alarm mute rule.

#### Sample Request

```bash

aws cloudwatch get-alarm-mute-rule \
	--alarm-mute-rule-name "DailyMaintenanceWindow"

```

#### Sample Response

```json

{
    "Name": "DailyMaintenanceWindow",
    "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-east-1:123456789012:alarm-mute-rule:DailyMaintenanceWindow",
    "Description": "Mute alarms during daily maintenance",
    "Rule": {
        "Schedule": {
            "Expression": "cron(0 2 * * ?)",
            "Duration": "PT2H",
            "Timezone": "UTC"
        }
    },
    "MuteTargets": {
        "AlarmNames": [
            "WebServerCPUAlarm",
            "DatabaseConnectionAlarm"
        ]
    },
    "Status": "SCHEDULED",
    "LastUpdatedTimestamp": "2026-01-15T10:30:00Z"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/getalarmmuterule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/getalarmmuterule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnableInsightRules

GetDashboard

All content copied from https://docs.aws.amazon.com/.
