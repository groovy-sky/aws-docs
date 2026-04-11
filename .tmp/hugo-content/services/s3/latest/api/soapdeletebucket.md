# DeleteBucket (SOAP API)

###### Note

SOAP APIs for Amazon S3 are not available for new customers, and are approaching End of Life (EOL) on October 31, 2025. We recommend that you use either the REST API or the AWS SDKs.

The `DeleteBucket` operation deletes a bucket. All objects in the bucket must be deleted before the bucket itself can be deleted.

###### Example

This example deletes the "quotes" bucket.

`Sample Request`

```

<DeleteBucket xmlns="http://doc.s3.amazonaws.com/2006-03-01">
  <Bucket>quotes</Bucket>
  <AWSAccessKeyId> AWS_ACCESS_KEY_ID_REDACTED</AWSAccessKeyId>
  <Timestamp>2006-03-01T12:00:00.183Z</Timestamp>
  <Signature>Iuyz3d3P0aTou39dzbqaEXAMPLE=</Signature>
</DeleteBucket>
```

`Sample Response`

```

<DeleteBucketResponse xmlns="http://s3.amazonaws.com/doc/2006-03-01">
  <DeleteBucketResponse>
    <Code>204</Code>
    <Description>No Content</Description>
  </DeleteBucketResponse>
</DeleteBucketResponse>
```

## Elements

- `Bucket:` The name of the bucket you want to delete.

## Access Control

Only the owner of a bucket is allowed to delete it, regardless the access control policy on the bucket.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateBucket (SOAP API)

ListBucket (SOAP API)

All content copied from https://docs.aws.amazon.com/.
