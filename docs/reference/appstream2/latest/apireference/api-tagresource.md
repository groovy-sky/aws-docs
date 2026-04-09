# TagResource

Adds or overwrites one or more tags for the specified WorkSpaces Applications resource. You can tag WorkSpaces Applications image builders, images, fleets, and stacks.

Each tag consists of a key and an optional value. If a resource already has a tag with the same key,
this operation updates its value.

To list the current tags for your resources, use [ListTagsForResource](api-listtagsforresource.md).
To disassociate tags from your resources, use [UntagResource](api-untagresource.md).

For more information about tags, see [Tagging Your Resources](../../../../services/appstream2/latest/developerguide/tagging-basic.md) in the _Amazon WorkSpaces Applications Administration Guide_.

## Request Syntax

```nohighlight

{
   "ResourceArn": "string",
   "Tags": {
      "string" : "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceArn](#API_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: Yes

**[Tags](#API_TagResource_RequestSyntax)**

The tags to associate. A tag is a key-value pair, and the value is optional. For example, Environment=Test. If you do not specify a value, Environment=.

If you do not specify a value, the value is set to an empty string.

Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following special characters:

\_ . : / = + \ - @

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(^(?!aws:).[\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopImageBuilder

UntagResource

All content copied from https://docs.aws.amazon.com/.
