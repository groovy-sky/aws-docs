# DeleteBucketAnalyticsConfiguration

###### Note

This operation is not supported for directory buckets.

Deletes an analytics configuration for the bucket (specified by the analytics configuration
ID).

To use this operation, you must have permissions to perform the
`s3:PutAnalyticsConfiguration` action. The bucket owner has this permission by default. The
bucket owner can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](../userguide/s3-access-control.md).

For information about the Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class\
Analysis](../dev/analytics-storage-class.md).

The following operations are related to `DeleteBucketAnalyticsConfiguration`:

- [GetBucketAnalyticsConfiguration](api-getbucketanalyticsconfiguration.md)

- [ListBucketAnalyticsConfigurations](api-listbucketanalyticsconfigurations.md)

- [PutBucketAnalyticsConfiguration](api-putbucketanalyticsconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /?analytics&id=Id HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteBucketAnalyticsConfiguration_RequestSyntax)**

The name of the bucket from which an analytics configuration is deleted.

Required: Yes

**[id](#API_DeleteBucketAnalyticsConfiguration_RequestSyntax)**

The ID that identifies the analytics configuration.

Required: Yes

**[x-amz-expected-bucket-owner](#API_DeleteBucketAnalyticsConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## Examples

### Sample Request

The following DELETE request deletes the analytics configuration with the ID `list1`.

```

            DELETE ?/analytics&id=list1 HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Date: Wed, 14 May 2014 02:11:22 GMT
            Authorization: signatureValue

```

### Sample Response

The following successful response shows Amazon S3 returning a `204 No Content` response.
The analytics configuration with the ID `list1` for the bucket has been removed.

```

            HTTP/1.1 204 No Content
            x-amz-id-2: 0FmFIWsh/PpBuzZ0JFRC55ZGVmQW4SHJ7xVDqKwhEdJmf3q63RtrvH8ZuxW1Bol5
            x-amz-request-id: 0CF038E9BCF63097
            Date: Wed, 14 May 2014 02:11:22 GMT
            Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletebucketanalyticsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucket

DeleteBucketCors

All content copied from https://docs.aws.amazon.com/.
