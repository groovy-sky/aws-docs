# DeleteScheduledQuery

Deletes a scheduled query and stops all future executions. This operation also removes any
configured actions and associated resources.

## Request Syntax

```nohighlight

{
   "identifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[identifier](#API_DeleteScheduledQuery_RequestSyntax)**

The ARN or name of the scheduled query to delete.

Type: String

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**InternalServerException**

An internal server error occurred while processing the request. This exception is returned
when the service encounters an unexpected condition that prevents it from fulfilling the
request.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deletescheduledquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deletescheduledquery.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteRetentionPolicy

DeleteSubscriptionFilter
