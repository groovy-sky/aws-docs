# DeleteAlarmMuteRule

Deletes a specific alarm mute rule.

When you delete a mute rule, any alarms that are currently being muted by that rule are immediately unmuted. If those alarms are in an ALARM state, their configured actions will trigger.

This operation is idempotent. If you delete a mute rule that does not exist, the operation succeeds without returning an error.

**Permissions**

To delete a mute rule, you need the `cloudwatch:DeleteAlarmMuteRule` permission on the alarm mute rule resource.

## Request Parameters

**AlarmMuteRuleName**

The name of the alarm mute rule to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Delete a mute rule

Delete a specific alarm mute rule. This operation does not return any output on success.

#### Sample Request

```bash

aws cloudwatch delete-alarm-mute-rule \
	--alarm-mute-rule-name "DailyMaintenanceWindow"

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/deletealarmmuterule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/deletealarmmuterule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

DeleteAlarms

All content copied from https://docs.aws.amazon.com/.
