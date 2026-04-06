# PredictQApp

Generates an Amazon Q App definition based on either a conversation or a problem statement
provided as input.The resulting app definition can be used to call `CreateQApp`.
This API doesn't create Amazon Q Apps directly.

## Request Syntax

```nohighlight

POST /apps.predictQApp HTTP/1.1
instance-id: instanceId
Content-type: application/json

{
   "options": { ... }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[instanceId](#API_qapps_PredictQApp_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[options](#API_qapps_PredictQApp_RequestSyntax)**

The input to generate the Q App definition from, either a conversation or problem
statement.

Type: [PredictQAppInputOptions](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_PredictQAppInputOptions.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "app": {
      "appDefinition": {
         "cards": [
            { ... }
         ],
         "initialPrompt": "string"
      },
      "description": "string",
      "title": "string"
   },
   "problemStatement": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[app](#API_qapps_PredictQApp_ResponseSyntax)**

The generated Q App definition.

Type: [PredictAppDefinition](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_PredictAppDefinition.html) object

**[problemStatement](#API_qapps_PredictQApp_ResponseSyntax)**

The problem statement extracted from the input conversation, if provided.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/amazonq/latest/api-reference/CommonErrors.html).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred while processing the request.

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

HTTP Status Code: 500

**ThrottlingException**

The requested operation could not be completed because too many requests were sent at
once. Wait a bit and try again later.

**quotaCode**

The code of the quota that was exceeded

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

**serviceCode**

The code for the service where the quota was exceeded

HTTP Status Code: 429

**UnauthorizedException**

The client is not authenticated or authorized to perform the requested operation.

HTTP Status Code: 401

**ValidationException**

The input failed to satisfy the constraints specified by the service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qapps-2023-11-27/PredictQApp)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qapps-2023-11-27/PredictQApp)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/PredictQApp)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qapps-2023-11-27/PredictQApp)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/PredictQApp)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qapps-2023-11-27/PredictQApp)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qapps-2023-11-27/PredictQApp)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qapps-2023-11-27/PredictQApp)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qapps-2023-11-27/PredictQApp)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/PredictQApp)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTagsForResource

StartQAppSession
