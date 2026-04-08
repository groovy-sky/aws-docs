# UntagResource

Removes a tag from the specified resource.

## Request Syntax

```nohighlight

DELETE /tags/resourceArn?tagKeys=tagKeys HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_UntagResource_RequestSyntax)**

The ARN of the resource to remove the tag from.

Required: Yes

**[tagKeys](#API_UntagResource_RequestSyntax)**

The key for the tag to add.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**InternalServerException**

Internal server error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource could not be found.

**resourceId**

The ID of the resource.

**resourceType**

The type of the resource.

HTTP Status Code: 404

**ThrottlingException**

Throttling limit exceeded error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 429

**ValidationException**

Validation exception error.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/untagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TagResource

UpdateAnalyzer
