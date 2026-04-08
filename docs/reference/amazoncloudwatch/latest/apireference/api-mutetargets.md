# MuteTargets

Specifies which alarms an alarm mute rule applies to.

You can target up to 100 specific alarms by name. When a mute rule is active, the targeted alarms continue to evaluate metrics and transition between states, but their configured actions are muted.

## Contents

**AlarmNames**

The list of alarm names that this mute rule targets. You can specify up to 100 alarm names.

Each alarm name must be between 1 and 255 characters in length. The alarm names must match existing alarms in your AWS account and region.

Type: Array of strings

Array Members: Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/mutetargets.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/mutetargets.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/mutetargets.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricStreamStatisticsMetric

PartialFailure

All content copied from https://docs.aws.amazon.com/.
