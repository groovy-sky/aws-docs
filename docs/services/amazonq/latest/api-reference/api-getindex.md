# GetIndex

Gets information about an existing Amazon Q Business index.

## Request Syntax

```nohighlight

GET /applications/applicationId/indices/indexId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetIndex_RequestSyntax)**

The identifier of the Amazon Q Business application connected to the index.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[indexId](#API_GetIndex_RequestSyntax)**

The identifier of the Amazon Q Business index you want information on.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "applicationId": "string",
   "capacityConfiguration": {
      "units": number
   },
   "createdAt": number,
   "description": "string",
   "displayName": "string",
   "documentAttributeConfigurations": [
      {
         "name": "string",
         "search": "string",
         "type": "string"
      }
   ],
   "error": {
      "errorCode": "string",
      "errorMessage": "string"
   },
   "indexArn": "string",
   "indexId": "string",
   "indexStatistics": {
      "textDocumentStatistics": {
         "indexedTextBytes": number,
         "indexedTextDocumentCount": number
      }
   },
   "status": "string",
   "type": "string",
   "updatedAt": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[applicationId](#API_GetIndex_ResponseSyntax)**

The identifier of the Amazon Q Business application associated with the index.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[capacityConfiguration](#API_GetIndex_ResponseSyntax)**

The storage capacity units chosen for your Amazon Q Business index.

Type: [IndexCapacityConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_IndexCapacityConfiguration.html) object

**[createdAt](#API_GetIndex_ResponseSyntax)**

The Unix timestamp when the Amazon Q Business index was created.

Type: Timestamp

**[description](#API_GetIndex_ResponseSyntax)**

The description for the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

**[displayName](#API_GetIndex_ResponseSyntax)**

The name of the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

**[documentAttributeConfigurations](#API_GetIndex_ResponseSyntax)**

Configuration information for document attributes or metadata. Document metadata are
fields associated with your documents. For example, the company department name
associated with each document. For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes-types.html#doc-attributes).

Type: Array of [DocumentAttributeConfiguration](api-documentattributeconfiguration.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

**[error](#API_GetIndex_ResponseSyntax)**

When the `Status` field value is `FAILED`, the
`ErrorMessage` field contains a message that explains why.

Type: [ErrorDetail](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ErrorDetail.html) object

**[indexArn](#API_GetIndex_ResponseSyntax)**

The Amazon Resource Name (ARN) of the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[indexId](#API_GetIndex_ResponseSyntax)**

The identifier of the Amazon Q Business index.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[indexStatistics](#API_GetIndex_ResponseSyntax)**

Provides information about the number of documents indexed.

Type: [IndexStatistics](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_IndexStatistics.html) object

**[status](#API_GetIndex_ResponseSyntax)**

The current status of the index. When the value is `ACTIVE`, the index is
ready for use. If the `Status` field value is `FAILED`, the
`ErrorMessage` field contains a message that explains why.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | FAILED | UPDATING`

**[type](#API_GetIndex_ResponseSyntax)**

The type of index attached to your Amazon Q Business application.

Type: String

Valid Values: `ENTERPRISE | STARTER`

**[updatedAt](#API_GetIndex_ResponseSyntax)**

The Unix timestamp when the Amazon Q Business index was last updated.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qbusiness-2023-11-27/GetIndex)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/GetIndex)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetGroup

GetMedia
