# DeleteAnomalyMonitor

Deletes a cost anomaly monitor.

## Request Syntax

```nohighlight

{
   "MonitorArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MonitorArn](#API_DeleteAnomalyMonitor_RequestSyntax)**

The unique identifier of the cost anomaly monitor that you want to delete.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ce-2017-10-25/DeleteAnomalyMonitor)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ce-2017-10-25/DeleteAnomalyMonitor)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateCostCategoryDefinition

DeleteAnomalySubscription
