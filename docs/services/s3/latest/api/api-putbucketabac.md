# PutBucketAbac

Sets the attribute-based access control (ABAC) property of the general purpose bucket. You must have `s3:PutBucketABAC` permission to perform this action. When you enable ABAC, you can use tags for access control on your buckets. Additionally, when ABAC is enabled, you must use the [TagResource](api-control-tagresource.md) and [UntagResource](api-control-untagresource.md) actions to manage tags on your buckets. You can nolonger use the [PutBucketTagging](api-putbuckettagging.md) and [DeleteBucketTagging](api-deletebuckettagging.md) actions to tag your bucket. For more information, see [Enabling ABAC in general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html).

## Request Syntax

```nohighlight

PUT /?abac HTTP/1.1
Host: Bucket.s3.amazonaws.com
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<AbacStatus xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Status>string</Status>
</AbacStatus>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutBucketAbac_RequestSyntax)**

The name of the general purpose bucket.

Required: Yes

**[Content-MD5](#API_PutBucketAbac_RequestSyntax)**

The MD5 hash of the `PutBucketAbac` request body.

For requests made using the AWS Command Line Interface (CLI) or AWS SDKs, this field is calculated automatically.

**[x-amz-expected-bucket-owner](#API_PutBucketAbac_RequestSyntax)**

The AWS account ID of the general purpose bucket's owner.

**[x-amz-sdk-checksum-algorithm](#API_PutBucketAbac_RequestSyntax)**

Indicates the algorithm that you want Amazon S3 to use to create the checksum. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in the _Amazon S3 User Guide_.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[AbacStatus](#API_PutBucketAbac_RequestSyntax)**

Root level tag for the AbacStatus parameters.

Required: Yes

**[Status](#API_PutBucketAbac_RequestSyntax)**

The ABAC status of the general purpose bucket.

Type: String

Valid Values: `Enabled | Disabled`

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/PutBucketAbac)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/PutBucketAbac)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListParts

PutBucketAccelerateConfiguration
