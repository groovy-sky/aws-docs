# UpdateIndex

Updates an Amazon Q Business index.

## Request Syntax

```nohighlight

PUT /applications/applicationId/indices/indexId HTTP/1.1
Content-type: application/json

{
   "capacityConfiguration": {
      "units": number
   },
   "description": "string",
   "displayName": "string",
   "documentAttributeConfigurations": [
      {
         "name": "string",
         "search": "string",
         "type": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateIndex_RequestSyntax)**

The identifier of the Amazon Q Business application connected to the index.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[indexId](#API_UpdateIndex_RequestSyntax)**

The identifier of the Amazon Q Business index.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[capacityConfiguration](#API_UpdateIndex_RequestSyntax)**

The storage capacity units you want to provision for your Amazon Q Business index. You can
add and remove capacity to fit your usage needs.

Type: [IndexCapacityConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_IndexCapacityConfiguration.html) object

Required: No

**[description](#API_UpdateIndex_RequestSyntax)**

The description of the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

Required: No

**[displayName](#API_UpdateIndex_RequestSyntax)**

The name of the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**[documentAttributeConfigurations](#API_UpdateIndex_RequestSyntax)**

Configuration information for document metadata or fields. Document metadata are
fields or attributes associated with your documents. For example, the company department
name associated with each document. For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes-types.html#doc-attributes).

Type: Array of [DocumentAttributeConfiguration](api-documentattributeconfiguration.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

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

**ServiceQuotaExceededException**

You have exceeded the set limits for your Amazon Q Business service.

**message**

The message describing a `ServiceQuotaExceededException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qbusiness-2023-11-27/UpdateIndex)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/UpdateIndex)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateDataSource

UpdatePlugin
