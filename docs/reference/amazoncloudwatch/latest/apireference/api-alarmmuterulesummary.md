# AlarmMuteRuleSummary

Summary information about an alarm mute rule, including its name, status, and configuration details.

## Contents

**AlarmMuteRuleArn**

The Amazon Resource Name (ARN) of the alarm mute rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**ExpireDate**

The date and time when the mute rule expires and is no longer evaluated. This field is only present if an expiration date was configured.

Type: Timestamp

Required: No

**LastUpdatedTimestamp**

The date and time when the mute rule was last updated.

Type: Timestamp

Required: No

**MuteType**

Indicates whether the mute rule is one-time or recurring. Valid values are `ONE_TIME` or `RECURRING`.

Type: String

Required: No

**Status**

The current status of the alarm mute rule. Valid values are `SCHEDULED`, `ACTIVE`, or `EXPIRED`.

Type: String

Valid Values: `SCHEDULED | ACTIVE | EXPIRED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/alarmmuterulesummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/alarmmuterulesummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/alarmmuterulesummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmHistoryItem

AlarmPromQLCriteria

All content copied from https://docs.aws.amazon.com/.
