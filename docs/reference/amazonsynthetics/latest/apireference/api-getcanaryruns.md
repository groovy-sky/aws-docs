# GetCanaryRuns

Retrieves a list of runs for a specified canary.

## Request Syntax

```nohighlight

POST /canary/name/runs HTTP/1.1
Content-type: application/json

{
   "DryRunId": "string",
   "MaxResults": number,
   "NextToken": "string",
   "RunType": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_GetCanaryRuns_RequestSyntax)**

The name of the canary that you want to see runs for.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[0-9a-z_\-]+$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[DryRunId](#API_GetCanaryRuns_RequestSyntax)**

The DryRunId associated with an existing canary’s dry run. You can use this DryRunId to retrieve information about the dry run.

Type: String

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**[MaxResults](#API_GetCanaryRuns_RequestSyntax)**

Specify this parameter to limit how many runs are returned each time you use
the `GetCanaryRuns` operation. If you omit this parameter, the default of 100 is used.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_GetCanaryRuns_RequestSyntax)**

A token that indicates that there is more data
available. You can use this token in a subsequent `GetCanaryRuns` operation to retrieve the next
set of results.

###### Note

When auto retry is enabled for the canary, the first subsequent retry is suffixed with \*1 indicating its the first retry and the next subsequent try is suffixed with \*2.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 252.

Required: No

**[RunType](#API_GetCanaryRuns_RequestSyntax)**

- When you provide `RunType=CANARY_RUN` and `dryRunId`, you will get an exception

- When a value is not provided for `RunType`, the default value is `CANARY_RUN`

- When `CANARY_RUN` is provided, all canary runs excluding dry runs are returned

- When `DRY_RUN` is provided, all dry runs excluding canary runs are returned

Type: String

Valid Values: `CANARY_RUN | DRY_RUN`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "CanaryRuns": [
      {
         "ArtifactS3Location": "string",
         "BrowserType": "string",
         "DryRunConfig": {
            "DryRunId": "string"
         },
         "Id": "string",
         "Name": "string",
         "RetryAttempt": number,
         "ScheduledRunId": "string",
         "Status": {
            "State": "string",
            "StateReason": "string",
            "StateReasonCode": "string",
            "TestResult": "string"
         },
         "Timeline": {
            "Completed": number,
            "MetricTimestampForRunAndRetries": number,
            "Started": number
         }
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CanaryRuns](#API_GetCanaryRuns_ResponseSyntax)**

An array of structures. Each structure contains the details of one of the
retrieved canary runs.

Type: Array of [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) objects

**[NextToken](#API_GetCanaryRuns_ResponseSyntax)**

A token that indicates that there is more data
available. You can use this token in a subsequent `GetCanaryRuns`
operation to retrieve the next
set of results.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 252.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/CommonErrors.html).

**InternalServerException**

An unknown internal error occurred.

HTTP Status Code: 500

**ResourceNotFoundException**

One of the specified resources was not found.

HTTP Status Code: 404

**ValidationException**

A parameter could not be validated.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/synthetics-2017-10-11/GetCanaryRuns)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/synthetics-2017-10-11/GetCanaryRuns)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetCanary

GetGroup
