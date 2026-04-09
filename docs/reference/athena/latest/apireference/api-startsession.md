# StartSession

Creates a session for running calculations within a workgroup. The session is ready
when it reaches an `IDLE` state.

## Request Syntax

```nohighlight

{
   "ClientRequestToken": "string",
   "CopyWorkGroupTags": boolean,
   "Description": "string",
   "EngineConfiguration": {
      "AdditionalConfigs": {
         "string" : "string"
      },
      "Classifications": [
         {
            "Name": "string",
            "Properties": {
               "string" : "string"
            }
         }
      ],
      "CoordinatorDpuSize": number,
      "DefaultExecutorDpuSize": number,
      "MaxConcurrentDpus": number,
      "SparkProperties": {
         "string" : "string"
      }
   },
   "ExecutionRole": "string",
   "MonitoringConfiguration": {
      "CloudWatchLoggingConfiguration": {
         "Enabled": boolean,
         "LogGroup": "string",
         "LogStreamNamePrefix": "string",
         "LogTypes": {
            "string" : [ "string" ]
         }
      },
      "ManagedLoggingConfiguration": {
         "Enabled": boolean,
         "KmsKey": "string"
      },
      "S3LoggingConfiguration": {
         "Enabled": boolean,
         "KmsKey": "string",
         "LogLocation": "string"
      }
   },
   "NotebookVersion": "string",
   "SessionIdleTimeoutInMinutes": number,
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ],
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ClientRequestToken](#API_StartSession_RequestSyntax)**

A unique case-sensitive string used to ensure the request to create the session is
idempotent (executes only once). If another `StartSessionRequest` is
received, the same response is returned and another session is not created. If a
parameter has changed, an error is returned.

###### Important

This token is listed as not required because AWS SDKs (for example
the AWS SDK for Java) auto-generate the token for users. If you are
not using the AWS SDK or the AWS CLI, you must provide
this token or the action will fail.

Type: String

Length Constraints: Minimum length of 32. Maximum length of 128.

Required: No

**[CopyWorkGroupTags](#API_StartSession_RequestSyntax)**

Copies the tags from the Workgroup to the Session when.

Type: Boolean

Required: No

**[Description](#API_StartSession_RequestSyntax)**

The session description.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[EngineConfiguration](#API_StartSession_RequestSyntax)**

Contains engine data processing unit (DPU) configuration settings and parameter
mappings.

Type: [EngineConfiguration](api-engineconfiguration.md) object

Required: Yes

**[ExecutionRole](#API_StartSession_RequestSyntax)**

The ARN of the execution role used to access user resources for Spark sessions and
Identity Center enabled workgroups. This property applies only to Spark enabled
workgroups and Identity Center enabled workgroups.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

Required: No

**[MonitoringConfiguration](#API_StartSession_RequestSyntax)**

Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets,
Amazon CloudWatch log groups etc.

Type: [MonitoringConfiguration](api-monitoringconfiguration.md) object

Required: No

**[NotebookVersion](#API_StartSession_RequestSyntax)**

The notebook version. This value is supplied automatically for notebook sessions in
the Athena console and is not required for programmatic session access. The
only valid notebook version is `Athena notebook version 1`. If
you specify a value for `NotebookVersion`, you must also specify a value for
`NotebookId`. See [EngineConfiguration:AdditionalConfigs](api-engineconfiguration.md#athena-Type-EngineConfiguration-AdditionalConfigs).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**[SessionIdleTimeoutInMinutes](#API_StartSession_RequestSyntax)**

The idle timeout in minutes for the session.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 480.

Required: No

**[Tags](#API_StartSession_RequestSyntax)**

A list of comma separated tags to add to the session that is created.

Type: Array of [Tag](api-tag.md) objects

Required: No

**[WorkGroup](#API_StartSession_RequestSyntax)**

The workgroup to which the session belongs.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

## Response Syntax

```nohighlight

{
   "SessionId": "string",
   "State": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[SessionId](#API_StartSession_ResponseSyntax)**

The session ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

**[State](#API_StartSession_ResponseSyntax)**

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

**SessionAlreadyExistsException**

The specified session already exists.

HTTP Status Code: 400

**TooManyRequestsException**

Indicates that the request was throttled.

**Reason**

The reason for the query throttling, for example, when it exceeds the concurrent query
limit.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/startsession.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/startsession.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/startsession.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/startsession.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/startsession.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/startsession.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/startsession.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/startsession.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/startsession.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/startsession.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartQueryExecution

StopCalculationExecution

All content copied from https://docs.aws.amazon.com/.
