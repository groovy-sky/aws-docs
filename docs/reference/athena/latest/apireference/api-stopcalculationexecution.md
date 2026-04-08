# StopCalculationExecution

Requests the cancellation of a calculation. A `StopCalculationExecution`
call on a calculation that is already in a terminal state (for example,
`STOPPED`, `FAILED`, or `COMPLETED`) succeeds but
has no effect.

###### Note

Cancelling a calculation is done on a best effort basis. If a calculation cannot
be cancelled, you can be charged for its completion. If you are concerned about
being charged for a calculation that cannot be cancelled, consider terminating the
session in which the calculation is running.

## Request Syntax

```nohighlight

{
   "CalculationExecutionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CalculationExecutionId](#API_StopCalculationExecution_RequestSyntax)**

The calculation execution UUID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

## Response Syntax

```nohighlight

{
   "State": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[State](#API_StopCalculationExecution_ResponseSyntax)**

`CREATING` \- The calculation is in the process of being created.

`CREATED` \- The calculation has been created and is ready to run.

`QUEUED` \- The calculation has been queued for processing.

`RUNNING` \- The calculation is running.

`CANCELING` \- A request to cancel the calculation has been received and the
system is working to stop it.

`CANCELED` \- The calculation is no longer running as the result of a cancel
request.

`COMPLETED` \- The calculation has completed without error.

`FAILED` \- The calculation failed and is no longer running.

Type: String

Valid Values: `CREATING | CREATED | QUEUED | RUNNING | CANCELING | CANCELED | COMPLETED | FAILED`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource, such as a workgroup, was not found.

**ResourceName**

The name of the Amazon resource.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/stopcalculationexecution.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/stopcalculationexecution.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StartSession

StopQueryExecution
