# GetSession

Gets the full details of a previously created session, including the session status
and configuration.

## Request Syntax

```nohighlight

{
   "SessionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/athena/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[SessionId](#API_GetSession_RequestSyntax)**

The session ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Syntax

```nohighlight

{
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
   "EngineVersion": "string",
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
   "SessionConfiguration": {
      "EncryptionConfiguration": {
         "EncryptionOption": "string",
         "KmsKey": "string"
      },
      "ExecutionRole": "string",
      "IdleTimeoutSeconds": number,
      "SessionIdleTimeoutInMinutes": number,
      "WorkingDirectory": "string"
   },
   "SessionId": "string",
   "Statistics": {
      "DpuExecutionInMillis": number
   },
   "Status": {
      "EndDateTime": number,
      "IdleSinceDateTime": number,
      "LastModifiedDateTime": number,
      "StartDateTime": number,
      "State": "string",
      "StateChangeReason": "string"
   },
   "WorkGroup": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Description](#API_GetSession_ResponseSyntax)**

The session description.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[EngineConfiguration](#API_GetSession_ResponseSyntax)**

Contains engine configuration information like DPU usage.

Type: [EngineConfiguration](https://docs.aws.amazon.com/athena/latest/APIReference/API_EngineConfiguration.html) object

**[EngineVersion](#API_GetSession_ResponseSyntax)**

The engine version used by the session (for example, `PySpark engine version
                3`). You can get a list of engine versions by calling [ListEngineVersions](https://docs.aws.amazon.com/athena/latest/APIReference/API_ListEngineVersions.html).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

**[MonitoringConfiguration](#API_GetSession_ResponseSyntax)**

Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets,
Amazon CloudWatch log groups etc.

Type: [MonitoringConfiguration](https://docs.aws.amazon.com/athena/latest/APIReference/API_MonitoringConfiguration.html) object

**[NotebookVersion](#API_GetSession_ResponseSyntax)**

The notebook version.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

**[SessionConfiguration](#API_GetSession_ResponseSyntax)**

Contains the workgroup configuration information used by the session.

Type: [SessionConfiguration](https://docs.aws.amazon.com/athena/latest/APIReference/API_SessionConfiguration.html) object

**[SessionId](#API_GetSession_ResponseSyntax)**

The session ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

**[Statistics](#API_GetSession_ResponseSyntax)**

Contains the DPU execution time.

Type: [SessionStatistics](https://docs.aws.amazon.com/athena/latest/APIReference/API_SessionStatistics.html) object

**[Status](#API_GetSession_ResponseSyntax)**

Contains information about the status of the session.

Type: [SessionStatus](https://docs.aws.amazon.com/athena/latest/APIReference/API_SessionStatus.html) object

**[WorkGroup](#API_GetSession_ResponseSyntax)**

The workgroup to which the session belongs.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/athena/latest/APIReference/CommonErrors.html).

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/athena-2017-05-18/GetSession)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/athena-2017-05-18/GetSession)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/GetSession)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/athena-2017-05-18/GetSession)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/GetSession)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/athena-2017-05-18/GetSession)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/athena-2017-05-18/GetSession)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/athena-2017-05-18/GetSession)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/athena-2017-05-18/GetSession)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/GetSession)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetResourceDashboard

GetSessionEndpoint
