# UntagResource

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Remove tags from an App Runner resource.

## Request Syntax

```nohighlight

{
   "ResourceArn": "string",
   "TagKeys": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceArn](#API_UntagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource that you want to remove tags from.

It must be the ARN of an App Runner resource.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

**[TagKeys](#API_UntagResource_RequestSyntax)**

A list of tag keys that you want to remove.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^(?!aws:).+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**InvalidStateException**

You can't perform this action when the resource is in its current state.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Remove tags from an App Runner service

This example illustrates how to remove two tags from an App Runner service.

#### Sample Request

```json

$ aws apprunner untag-resource --cli-input-json "`cat`"
{
  "ResourceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa",
  "TagKeys": [
    "Department",
    "CustomerId"
  ]
}
```

#### Sample Response

```json

{
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/untagresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UpdateDefaultAutoScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
