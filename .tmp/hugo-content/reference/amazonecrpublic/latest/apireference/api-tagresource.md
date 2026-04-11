# TagResource

Associates the specified tags to a resource with the specified `resourceArn`.
If existing tags on a resource aren't specified in the request parameters, they aren't
changed. When a resource is deleted, the tags associated with that resource are also
deleted.

## Request Syntax

```nohighlight

{
   "resourceArn": "string",
   "tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource to add tags to. Currently, the supported
resource is an Amazon ECR Public repository.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[tags](#API_TagResource_RequestSyntax)**

The tags to add to the resource. A tag is an array of key-value pairs.
Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

HTTP Status Code: 400

**InvalidTagParameterException**

An invalid parameter has been specified. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.

HTTP Status Code: 400

**RepositoryNotFoundException**

The specified repository can't be found. Check the spelling of the specified repository
and ensure that you're performing operations on the correct registry.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

HTTP Status Code: 500

**TooManyTagsException**

The list of tags on the repository is over the limit. The maximum number of tags that
can be applied to a repository is 50.

HTTP Status Code: 400

**UnsupportedCommandException**

The action isn't supported in this Region.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-public-2020-10-30/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-public-2020-10-30/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SetRepositoryPolicy

UntagResource

All content copied from https://docs.aws.amazon.com/.
