# GetBucket

Gets an Amazon S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts](../userguide/s3onoutposts.md) in the
_Amazon S3 User Guide_.

If you are using an identity other than the root user of the AWS account that owns the
Outposts bucket, the calling identity must have the `s3-outposts:GetBucket`
permissions on the specified Outposts bucket and belong to the Outposts bucket owner's
account in order to use this action. Only users from Outposts bucket owner account with
the right permissions can perform actions on an Outposts bucket.

If you don't have `s3-outposts:GetBucket` permissions or you're not using an
identity that belongs to the bucket owner's account, Amazon S3 returns a `403 Access
            Denied` error.

The following actions are related to `GetBucket` for Amazon S3 on Outposts:

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-getbucket.md#API_control_GetBucket_Examples) section.

- [PutObject](api-putobject.md)

- [CreateBucket](api-control-createbucket.md)

- [DeleteBucket](api-control-deletebucket.md)

## Request Syntax

```nohighlight

GET /v20180820/bucket/name HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetBucket_RequestSyntax)**

Specifies the bucket.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetBucket_RequestSyntax)**

The AWS account ID of the Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetBucketResult>
   <Bucket>string</Bucket>
   <PublicAccessBlockEnabled>boolean</PublicAccessBlockEnabled>
   <CreationDate>timestamp</CreationDate>
</GetBucketResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketResult](#API_control_GetBucket_ResponseSyntax)**

Root level tag for the GetBucketResult parameters.

Required: Yes

**[Bucket](#API_control_GetBucket_ResponseSyntax)**

The Outposts bucket requested.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

**[CreationDate](#API_control_GetBucket_ResponseSyntax)**

The creation date of the Outposts bucket.

Type: Timestamp

**[PublicAccessBlockEnabled](#API_control_GetBucket_ResponseSyntax)**

Type: Boolean

## Examples

### Sample request for getting Amazon S3 on Outposts bucket

This example illustrates one usage of GetBucket.

```HTTP

GET /v20180820/bucket/example-outpost-bucket/ HTTP/1.1
Host: s3-outposts.<Region>.amazonaws.com
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            x-amz-Date: 20200928T203757Z
            Authorization: authorization string

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getbucket.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getbucket.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getbucket.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getbucket.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getbucket.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getbucket.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getbucket.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getbucket.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getbucket.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getbucket.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccessPointScope

GetBucketLifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
