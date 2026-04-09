# CreateIndex

Creates an Amazon Q Business index.

To determine if index creation has completed, check the `Status` field
returned from a call to `DescribeIndex`. The `Status` field is set
to `ACTIVE` when the index is ready to use.

Once the index is active, you can index your documents using the [`BatchPutDocument`](api-batchputdocument.md) API or the [`CreateDataSource`](api-createdatasource.md) API.

## Request Syntax

```nohighlight

POST /applications/applicationId/indices HTTP/1.1
Content-type: application/json

{
   "capacityConfiguration": {
      "units": number
   },
   "clientToken": "string",
   "description": "string",
   "displayName": "string",
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ],
   "type": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_CreateIndex_RequestSyntax)**

The identifier of the Amazon Q Business application using the index.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[capacityConfiguration](#API_CreateIndex_RequestSyntax)**

The capacity units you want to provision for your index. You can add and remove
capacity to fit your usage needs.

Type: [IndexCapacityConfiguration](api-indexcapacityconfiguration.md) object

Required: No

**[clientToken](#API_CreateIndex_RequestSyntax)**

A token that you provide to identify the request to create an index. Multiple calls to
the `CreateIndex` API with the same client token will create only one
index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**[description](#API_CreateIndex_RequestSyntax)**

A description for the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

Required: No

**[displayName](#API_CreateIndex_RequestSyntax)**

A name for the Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: Yes

**[tags](#API_CreateIndex_RequestSyntax)**

A list of key-value pairs that identify or categorize the index. You can also use tags
to help control access to the index. Tag keys and values can consist of Unicode letters,
digits, white space, and any of the following symbols: \_ . : / = + - @.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

**[type](#API_CreateIndex_RequestSyntax)**

The index type that's suitable for your needs. For more information on what's included
in each type of index, see [Amazon Q Business\
tiers](../qbusiness-ug/tiers.md#index-tiers).

Type: String

Valid Values: `ENTERPRISE | STARTER`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "indexArn": "string",
   "indexId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[indexArn](#API_CreateIndex_ResponseSyntax)**

The Amazon Resource Name (ARN) of an Amazon Q Business index.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[indexId](#API_CreateIndex_ResponseSyntax)**

The identifier for the Amazon Q Business index.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/createindex.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/createindex.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDataSource

CreatePlugin

All content copied from https://docs.aws.amazon.com/.
