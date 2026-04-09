# DeleteBucketCors

###### Note

This operation is not supported for directory buckets.

Deletes the `cors` configuration information set for the bucket.

To use this operation, you must have permission to perform the `s3:PutBucketCORS` action.
The bucket owner has this permission by default and can grant this permission to others.

For information about `cors`, see [Enabling Cross-Origin Resource Sharing](../dev/cors.md) in the
_Amazon S3 User Guide_.

###### Related Resources

- [PutBucketCors](api-putbucketcors.md)

- [RESTOPTIONSobject](restoptionsobject.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /?cors HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteBucketCors_RequestSyntax)**

Specifies the bucket whose `cors` configuration is being deleted.

Required: Yes

**[x-amz-expected-bucket-owner](#API_DeleteBucketCors_RequestSyntax)**

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

This example illustrates one usage of `DeleteBucketCors`.

```HTTP

DELETE /?cors HTTP/1.1
Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
Date: Tue, 13 Dec 2011 19:14:42 GMT
Authorization: signatureValue

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deletebucketcors.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletebucketcors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucketAnalyticsConfiguration

DeleteBucketEncryption

All content copied from https://docs.aws.amazon.com/.
