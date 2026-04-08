# UntagResource

Removes one or more tags from the specified resource.

## Request Parameters

**ResourceARN**

The ARN of the CloudWatch resource that you're removing tags from.

The ARN format of an alarm is
`arn:aws:cloudwatch:Region:account-id:alarm:alarm-name
                  `

The ARN format of a Contributor Insights rule is
`arn:aws:cloudwatch:Region:account-id:insight-rule/insight-rule-name
                  `

For more information about ARN format, see [Resource Types Defined by Amazon CloudWatch](../../../../services/iam/latest/userguide/list-amazoncloudwatch.md#amazoncloudwatch-resources-for-iam-policies) in the _Amazon Web_
_Services General Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**TagKeys**

The list of tag keys to remove from the resource.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModificationException**

More than one process tried to modify a resource at the same time.

HTTP Status Code: 429

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 409

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

Data Types

All content copied from https://docs.aws.amazon.com/.
