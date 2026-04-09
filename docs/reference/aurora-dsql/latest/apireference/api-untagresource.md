# UntagResource

Removes a tag from a resource.

## Request Syntax

```nohighlight

DELETE /tags/resourceArn?tagKeys=tagKeys HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_UntagResource_RequestSyntax)**

The ARN of the resource from which to remove tags.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**[tagKeys](#API_UntagResource_RequestSyntax)**

The array of keys of the tags that you want to remove.

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9_.:/=+\-@ ]*`

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

The request processing has failed because of an unknown error, exception or
failure.

**retryAfterSeconds**

Retry after seconds.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource could not be found.

**resourceId**

The resource ID could not be found.

**resourceType**

The resource type could not be found.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to request throttling.

**message**

The message that the request was denied due to request throttling.

**quotaCode**

The request exceeds a request rate quota.

**retryAfterSeconds**

The request exceeds a request rate quota. Retry after seconds.

**serviceCode**

The request exceeds a service quota.

HTTP Status Code: 429

**ValidationException**

The input failed to satisfy the constraints specified by an AWS service.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the validation exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dsql-2018-05-10/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dsql-2018-05-10/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dsql-2018-05-10/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dsql-2018-05-10/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dsql-2018-05-10/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dsql-2018-05-10/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dsql-2018-05-10/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dsql-2018-05-10/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dsql-2018-05-10/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dsql-2018-05-10/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UpdateCluster

All content copied from https://docs.aws.amazon.com/.
