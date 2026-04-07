# GetResourceDashboard

Gets the Live UI/Persistence UI for a session.

## Request Syntax

```nohighlight

{
   "ResourceARN": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceARN](#API_GetResourceDashboard_RequestSyntax)**

The The Amazon Resource Name (ARN) for a session.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Required: Yes

## Response Syntax

```nohighlight

{
   "Url": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Url](#API_GetResourceDashboard_ResponseSyntax)**

The Live UI/Persistence UI url for a session.

Type: String

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/athena-2017-05-18/GetResourceDashboard)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/GetResourceDashboard)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetQueryRuntimeStatistics

GetSession
