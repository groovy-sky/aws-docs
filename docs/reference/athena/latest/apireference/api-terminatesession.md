# TerminateSession

Terminates an active session. A `TerminateSession` call on a session that
is already inactive (for example, in a `FAILED`, `TERMINATED` or
`TERMINATING` state) succeeds but has no effect. Calculations running in
the session when `TerminateSession` is called are forcefully stopped, but may
display as `FAILED` instead of `STOPPED`.

## Request Syntax

```nohighlight

{
   "SessionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[SessionId](#API_TerminateSession_RequestSyntax)**

The session ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

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

**[State](#API_TerminateSession_ResponseSyntax)**

The state of the session. A description of each state follows.

`CREATING` \- The session is being started, including acquiring
resources.

`CREATED` \- The session has been started.

`IDLE` \- The session is able to accept a calculation.

`BUSY` \- The session is processing another task and is unable to accept a
calculation.

`TERMINATING` \- The session is in the process of shutting down.

`TERMINATED` \- The session and its resources are no longer running.

`DEGRADED` \- The session has no healthy coordinators.

`FAILED` \- Due to a failure, the session and its resources are no longer
running.

Type: String

Valid Values: `CREATING | CREATED | IDLE | BUSY | TERMINATING | TERMINATED | DEGRADED | FAILED`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/terminatesession.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/terminatesession.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/terminatesession.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/terminatesession.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/terminatesession.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/terminatesession.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/terminatesession.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/terminatesession.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/terminatesession.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/terminatesession.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UntagResource

All content copied from https://docs.aws.amazon.com/.
