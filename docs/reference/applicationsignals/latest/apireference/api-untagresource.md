# UntagResource

Removes one or more tags from the specified resource.

## Request Syntax

```nohighlight

POST /untag-resource HTTP/1.1
Content-type: application/json

{
   "ResourceArn": "string",
   "TagKeys": [ "string" ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[ResourceArn](#API_UntagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the CloudWatch resource that you want to delete tags from.

The ARN format of an Application Signals SLO is
`arn:aws:cloudwatch:Region:account-id:slo:slo-name
                  `

For more information about ARN format, see [Resource\
Types Defined by Amazon CloudWatch](../../../../services/iam/latest/userguide/list-amazoncloudwatch-amazoncloudwatch-resources-for-iam-policies.md) in the _Amazon Web Services General_
_Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[TagKeys](#API_UntagResource_RequestSyntax)**

The list of tag keys to remove from the resource.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

Resource not found.

**ResourceId**

Can't find the resource id.

**ResourceType**

The resource type is not valid.

HTTP Status Code: 404

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/untagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TagResource

UpdateServiceLevelObjective
