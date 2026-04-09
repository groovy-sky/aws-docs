# PutObjectRetention

###### Note

This operation is not supported for directory buckets.

Places an Object Retention configuration on an object. For more information, see [Locking Objects](../dev/object-lock.md). Users or
accounts require the `s3:PutObjectRetention` permission in order to place an Object Retention
configuration on objects. Bypassing a Governance Retention configuration requires the
`s3:BypassGovernanceRetention` permission.

This functionality is not supported for Amazon S3 on Outposts.

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /{Key+}?retention&versionId=VersionId HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-request-payer: RequestPayer
x-amz-bypass-governance-retention: BypassGovernanceRetention
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<Retention xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Mode>string</Mode>
   <RetainUntilDate>timestamp</RetainUntilDate>
</Retention>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutObjectRetention_RequestSyntax)**

The bucket name that contains the object you want to apply this Object Retention configuration to.

**Access points** \- When you use this action with an access point for general purpose buckets, you must provide the alias of the access point in place of the bucket name or specify the access point ARN. When you use this action with an access point for directory buckets, you must provide the access point name in place of the bucket name. When using the access point ARN, you must direct requests to the access point hostname. The access point hostname takes the form _AccessPointName_- _AccountId_.s3-accesspoint. _Region_.amazonaws.com. When using this action with an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more information about access point ARNs, see [Using access points](../userguide/using-access-points.md) in the _Amazon S3 User Guide_.

Required: Yes

**[Content-MD5](#API_PutObjectRetention_RequestSyntax)**

The MD5 hash for the request body.

For requests made using the AWS Command Line Interface (CLI) or AWS SDKs, this field is calculated automatically.

**[Key](#API_PutObjectRetention_RequestSyntax)**

The key name for the object that you want to apply this Object Retention configuration to.

Length Constraints: Minimum length of 1.

Required: Yes

**[versionId](#API_PutObjectRetention_RequestSyntax)**

The version ID for the object that you want to apply this Object Retention configuration to.

**[x-amz-bypass-governance-retention](#API_PutObjectRetention_RequestSyntax)**

Indicates whether this action should bypass Governance-mode restrictions.

**[x-amz-expected-bucket-owner](#API_PutObjectRetention_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-request-payer](#API_PutObjectRetention_RequestSyntax)**

Confirms that the requester knows that they will be charged for the request. Bucket owners need not
specify this parameter in their requests. If either the source or destination S3 bucket has Requester
Pays enabled, the requester will pay for the corresponding charges. For information about
downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays\
Buckets](../dev/objectsinrequesterpaysbuckets.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

**[x-amz-sdk-checksum-algorithm](#API_PutObjectRetention_RequestSyntax)**

Indicates the algorithm used to create the checksum for the object when you use the SDK. This header will not provide any
additional functionality if you don't use the SDK. When you send this header, there must be a corresponding `x-amz-checksum` or
`x-amz-trailer` header sent. Otherwise, Amazon S3 fails the request with the HTTP status code `400 Bad Request`. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

If you provide an individual checksum, Amazon S3 ignores any provided `ChecksumAlgorithm`
parameter.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[Retention](#API_PutObjectRetention_RequestSyntax)**

Root level tag for the Retention parameters.

Required: Yes

**[Mode](#API_PutObjectRetention_RequestSyntax)**

Indicates the Retention mode for the specified object.

Type: String

Valid Values: `GOVERNANCE | COMPLIANCE`

Required: No

**[RetainUntilDate](#API_PutObjectRetention_RequestSyntax)**

The date on which this Object Lock Retention will expire.

Type: Timestamp

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
x-amz-request-charged: RequestCharged

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[x-amz-request-charged](#API_PutObjectRetention_ResponseSyntax)**

If present, indicates that the requester was successfully charged for the request. For more
information, see [Using Requester Pays buckets for storage transfers and usage](../userguide/requesterpaysbuckets.md) in the _Amazon Simple_
_Storage Service user guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/putobjectretention.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/putobjectretention.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutObjectLockConfiguration

PutObjectTagging

All content copied from https://docs.aws.amazon.com/.
