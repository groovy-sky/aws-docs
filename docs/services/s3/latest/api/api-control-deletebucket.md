# DeleteBucket

###### Note

This action deletes an Amazon S3 on Outposts bucket. To delete an S3 bucket, see [DeleteBucket](api-deletebucket.md) in the _Amazon S3 API Reference_.

Deletes the Amazon S3 on Outposts bucket. All objects (including all object versions and
delete markers) in the bucket must be deleted before the bucket itself can be deleted. For
more information, see [Using Amazon S3 on Outposts](../userguide/s3onoutposts.md) in
_Amazon S3 User Guide_.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-deletebucket.md#API_control_DeleteBucket_Examples) section.

###### Related Resources

- [CreateBucket](api-control-createbucket.md)

- [GetBucket](api-control-getbucket.md)

- [DeleteObject](api-deleteobject.md)

## Request Syntax

```nohighlight

DELETE /v20180820/bucket/name HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_DeleteBucket_RequestSyntax)**

Specifies the bucket being deleted.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_DeleteBucket_RequestSyntax)**

The account ID that owns the Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Sample request to delete an Amazon S3 on Outposts bucket

This request deletes the Outposts bucket named
`example-outpost-bucket`.

```HTTP

DELETE /v20180820/bucket/example-outpost-bucket/ HTTP/1.1
Host: s3-outposts.<Region>.amazonaws.com
x-amz-outpost-id: op-01ac5d28a6a232904
x-amz-account-id:example-account-id
Date: Wed, 01 Mar  2006 12:00:00 GMT
Authorization: authorization string

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/deletebucket.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deletebucket.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAccessPointScope

DeleteBucketLifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
