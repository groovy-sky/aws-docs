# ListAlarmMuteRules

Lists alarm mute rules in your AWS account and region.

You can filter the results by alarm name to find all mute rules targeting a specific alarm, or by status to find rules that are scheduled, active, or expired.

This operation supports pagination for accounts with many mute rules. Use the `MaxRecords` and `NextToken` parameters to retrieve results in multiple calls.

**Permissions**

To list mute rules, you need the `cloudwatch:ListAlarmMuteRules` permission.

## Request Parameters

**AlarmName**

Filter results to show only mute rules that target the specified alarm name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**MaxRecords**

The maximum number of mute rules to return in one call. The default is 50.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token returned from a previous call to indicate where to continue retrieving results.

Type: String

Required: No

**Statuses**

Filter results to show only mute rules with the specified statuses. Valid values are `SCHEDULED`, `ACTIVE`, or `EXPIRED`.

Type: Array of strings

Valid Values: `SCHEDULED | ACTIVE | EXPIRED`

Required: No

## Response Elements

The following elements are returned by the service.

**AlarmMuteRuleSummaries**

A list of alarm mute rule summaries.

Type: Array of [AlarmMuteRuleSummary](api-alarmmuterulesummary.md) objects

**NextToken**

The token to use when requesting the next set of results. If this field is absent, there are no more results to retrieve.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidNextToken**

The next token specified is invalid.

**message**

HTTP Status Code: 400

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## Examples

### List all mute rules

List all alarm mute rules in your account.

#### Sample Request

```bash

aws cloudwatch list-alarm-mute-rules

```

#### Sample Response

```json

{
    "AlarmMuteRuleSummaries": [
        {
            "Name": "DailyMaintenanceWindow",
            "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-east-1:123456789012:alarm-mute-rule:DailyMaintenanceWindow",
            "Status": "SCHEDULED",
            "MuteType": "RECURRING",
            "LastUpdatedTimestamp": "2026-01-15T10:30:00Z"
        },
        {
            "Name": "ProductionDeployment-2026-01-20",
            "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-east-1:123456789012:alarm-mute-rule:ProductionDeployment-2026-01-20",
            "Status": "ACTIVE",
            "MuteType": "ONE_TIME",
            "LastUpdatedTimestamp": "2026-01-20T13:00:00Z"
        },
        {
            "Name": "WeeklyBackupWindow",
            "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-west-2:123456789012:alarm-mute-rule:WeeklyBackupWindow",
            "Status": "SCHEDULED",
            "MuteType": "RECURRING",
            "ExpireDate": "2026-12-31T23:59:59Z",
            "LastUpdatedTimestamp": "2026-01-05T12:00:00Z"
        }
    ]
}

```

### List mute rules targeting a specific alarm

List all mute rules that target a specific alarm.

#### Sample Request

```bash

aws cloudwatch list-alarm-mute-rules \
	--alarm-name "WebServerCPUAlarm"

```

#### Sample Response

```json

{
    "AlarmMuteRuleSummaries": [
        {
            "Name": "DailyMaintenanceWindow",
            "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-east-1:123456789012:alarm-mute-rule:DailyMaintenanceWindow",
            "Status": "SCHEDULED",
            "MuteType": "RECURRING",
            "LastUpdatedTimestamp": "2026-01-15T10:30:00Z"
        },
        {
            "Name": "EmergencyMuteRule",
            "AlarmMuteRuleArn": "arn:aws:cloudwatch:us-east-1:123456789012:alarm-mute-rule:EmergencyMuteRule",
            "Status": "ACTIVE",
            "MuteType": "ONE_TIME",
            "LastUpdatedTimestamp": "2026-01-21T14:00:00Z"
        }
    ]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/listalarmmuterules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/listalarmmuterules.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetOTelEnrichment

ListDashboards

All content copied from https://docs.aws.amazon.com/.
