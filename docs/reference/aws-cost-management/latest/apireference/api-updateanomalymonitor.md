# UpdateAnomalyMonitor

Updates an existing cost anomaly monitor. The changes made are applied going forward, and
doesn't change anomalies detected in the past.

## Request Syntax

```nohighlight

{
   "MonitorArn": "string",
   "MonitorName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MonitorArn](#API_UpdateAnomalyMonitor_RequestSyntax)**

Cost anomaly monitor Amazon Resource Names (ARNs).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: Yes

**[MonitorName](#API_UpdateAnomalyMonitor_RequestSyntax)**

The new name for the cost anomaly monitor.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

## Response Syntax

```nohighlight

{
   "MonitorArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[MonitorArn](#API_UpdateAnomalyMonitor_ResponseSyntax)**

A cost anomaly monitor ARN.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**LimitExceededException**

You made too many calls in a short period of time. Try again later.

HTTP Status Code: 400

**UnknownMonitorException**

The cost anomaly monitor does not exist for the account.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ce-2017-10-25/updateanomalymonitor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/updateanomalymonitor.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateAnomalySubscription
