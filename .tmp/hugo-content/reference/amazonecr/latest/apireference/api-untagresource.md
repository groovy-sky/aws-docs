# UntagResource

Deletes specified tags from a resource.

## Request Syntax

```nohighlight

{
   "resourceArn": "string",
   "tagKeys": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_UntagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource from which to remove tags. Currently, the only supported
resource is an Amazon ECR repository.

Type: String

Required: Yes

**[tagKeys](#API_UntagResource_RequestSyntax)**

The keys of the tags to be removed.

Type: Array of strings

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**InvalidTagParameterException**

An invalid parameter has been specified. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

HTTP Status Code: 400

**RepositoryNotFoundException**

The specified repository could not be found. Check the spelling of the specified
repository and ensure that you are performing operations on the correct registry.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**TooManyTagsException**

The list of tags on the repository is over the limit. The maximum number of tags that
can be applied to a repository is 50.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UpdateImageStorageClass

All content copied from https://docs.aws.amazon.com/.
