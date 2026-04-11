# DeleteBucketOwnershipControls

###### Note

This operation is not supported for directory buckets.

Removes `OwnershipControls` for an Amazon S3 bucket. To use this operation, you must have the
`s3:PutBucketOwnershipControls` permission. For more information about Amazon S3 permissions,
see [Specifying\
Permissions in a Policy](../dev/using-with-s3-actions.md).

For information about Amazon S3 Object Ownership, see [Using Object Ownership](../dev/about-object-ownership.md).

The following operations are related to `DeleteBucketOwnershipControls`:

- [GetBucketOwnershipControls](api-getbucketownershipcontrols.md)

- [PutBucketOwnershipControls](api-putbucketownershipcontrols.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /?ownershipControls HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteBucketOwnershipControls_RequestSyntax)**

The Amazon S3 bucket whose `OwnershipControls` you want to delete.

Required: Yes

**[x-amz-expected-bucket-owner](#API_DeleteBucketOwnershipControls_RequestSyntax)**

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

### Sample DeleteBucketOwnershipControls Request

This example illustrates one usage of DeleteBucketOwnershipControls.

```

          DELETE /example-bucket?/ownershipControls HTTP/1.1
          Host: examplebucket.s3.<Region>.amazonaws.com
          Date: Thu, 18 Jun 2017 00:17:22 GMT
          Authorization: signatureValue;

```

### Sample DeleteBucketOwnershipControls Response

This example illustrates one usage of DeleteBucketOwnershipControls.

```

          HTTP/1.1 204 No Content
          x-amz-id-2: dVrxJD3XHDcjZHFtd7eSB+ovpY8hQ6kSe9jPzyRVkWp27cij05qV1pTIvz/hjlsrupiy9gEkSdw=
          x-amz-request-id: 4BFC0B777B448C97
          Date: Thu, 18 Jun 2020 22:54:03 GMT
          Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deletebucketownershipcontrols.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletebucketownershipcontrols.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucketMetricsConfiguration

DeleteBucketPolicy

All content copied from https://docs.aws.amazon.com/.
