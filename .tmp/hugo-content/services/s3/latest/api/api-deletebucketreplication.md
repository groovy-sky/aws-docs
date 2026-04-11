# DeleteBucketReplication

###### Note

This operation is not supported for directory buckets.

Deletes the replication configuration from the bucket.

To use this operation, you must have permissions to perform the
`s3:PutReplicationConfiguration` action. The bucket owner has these permissions by default
and can grant it to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](../userguide/s3-access-control.md).

###### Note

It can take a while for the deletion of a replication configuration to fully propagate.

For information about replication configuration, see [Replication](../dev/replication.md) in the
_Amazon S3 User Guide_.

The following operations are related to `DeleteBucketReplication`:

- [PutBucketReplication](api-putbucketreplication.md)

- [GetBucketReplication](api-getbucketreplication.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /?replication HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteBucketReplication_RequestSyntax)**

The bucket name.

Required: Yes

**[x-amz-expected-bucket-owner](#API_DeleteBucketReplication_RequestSyntax)**

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

The following DELETE request deletes the `replication` subresource from the specified
bucket. This removes the replication configuration that is set for the bucket.

```

           DELETE /?replication HTTP/1.1
           Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
           Date: Wed, 11 Feb 2015 05:37:16 GMT
           20150211T171320Z

           Authorization: authorization string

```

### Sample Response

When the `replication` subresource has been deleted, Amazon S3 returns a `204 No
            Content` response. It will not replicate new objects that are stored in the
`examplebucket` bucket.

```

           HTTP/1.1 204 No Content
           x-amz-id-2: Uuag1LuByRx9e6j5OnimrSAMPLEtRPfTaOAa==
           x-amz-request-id: 656c76696e672example
           Date: Wed, 11 Feb 2015 05:37:16 GMT
           Connection: keep-alive
           Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deletebucketreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletebucketreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucketPolicy

DeleteBucketTagging

All content copied from https://docs.aws.amazon.com/.
