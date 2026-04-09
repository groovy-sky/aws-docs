# DeleteObjectTagging

###### Note

This operation is not supported for directory buckets.

Removes the entire tag set from the specified object. For more information about managing object
tags, see [Object\
Tagging](../dev/object-tagging.md).

To use this operation, you must have permission to perform the `s3:DeleteObjectTagging`
action.

To delete tags of a specific object version, add the `versionId` query parameter in the
request. You will need permission for the `s3:DeleteObjectVersionTagging` action.

The following operations are related to `DeleteObjectTagging`:

- [PutObjectTagging](api-putobjecttagging.md)

- [GetObjectTagging](api-getobjecttagging.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /{Key+}?tagging&versionId=VersionId HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteObjectTagging_RequestSyntax)**

The bucket name containing the objects from which to remove the tags.

**Access points** \- When you use this action with an access point for general purpose buckets, you must provide the alias of the access point in place of the bucket name or specify the access point ARN. When you use this action with an access point for directory buckets, you must provide the access point name in place of the bucket name. When using the access point ARN, you must direct requests to the access point hostname. The access point hostname takes the form _AccessPointName_- _AccountId_.s3-accesspoint. _Region_.amazonaws.com. When using this action with an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more information about access point ARNs, see [Using access points](../userguide/using-access-points.md) in the _Amazon S3 User Guide_.

**S3 on Outposts** \- When you use this action with S3 on Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on Outposts hostname takes the
form `
                     AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com`. When you use this action with S3 on Outposts, the destination bucket must be the Outposts access point ARN or the access point alias. For more information about S3 on Outposts, see [What is S3 on Outposts?](../userguide/s3onoutposts.md) in the _Amazon S3 User Guide_.

Required: Yes

**[Key](#API_DeleteObjectTagging_RequestSyntax)**

The key that identifies the object in the bucket from which to remove all tags.

Length Constraints: Minimum length of 1.

Required: Yes

**[versionId](#API_DeleteObjectTagging_RequestSyntax)**

The versionId of the object that the tag-set will be removed from.

**[x-amz-expected-bucket-owner](#API_DeleteObjectTagging_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 204
x-amz-version-id: VersionId

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response.

The response returns the following HTTP headers.

**[x-amz-version-id](#API_DeleteObjectTagging_ResponseSyntax)**

The versionId of the object the tag-set was removed from.

## Examples

### Sample Request

The following DELETE request deletes the tag set from the specified object.

```

             DELETE /exampleobject?tagging HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Date: Wed, 25 Nov 2016 12:00:00 GMT
            Authorization: signatureValue

```

### Sample Response

The following successful response shows Amazon S3 returning a 204 No Content response. The tag set
for the object has been removed.

```

           HTTP/1.1 204 No Content
           x-amz-version-id: VersionId
           Date: Wed, 25 Nov 2016 12:00:00 GMT
           Connection: close
           Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deleteobjecttagging.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deleteobjecttagging.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteObjects

DeletePublicAccessBlock

All content copied from https://docs.aws.amazon.com/.
