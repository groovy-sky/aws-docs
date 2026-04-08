# DeleteAnomalySubscription

Deletes a cost anomaly subscription.

## Request Syntax

```nohighlight

{
   "SubscriptionArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[SubscriptionArn](#API_DeleteAnomalySubscription_RequestSyntax)**

The unique identifier of the cost anomaly subscription that you want to delete.

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

**UnknownSubscriptionException**

The cost anomaly subscription does not exist for the account.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ce-2017-10-25/deleteanomalysubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/deleteanomalysubscription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteAnomalyMonitor

DeleteCostCategoryDefinition
