# Amazon S3 on Outposts API operations

This topic lists the Amazon S3, Amazon S3 Control, and Amazon S3 on Outposts API operations that you can
use with Amazon S3 on Outposts.

###### Topics

- [Amazon S3 API operations for managing objects](#S3OutpostsAPIsObject)

- [Amazon S3 Control API operations for managing buckets](#S3OutpostsAPIsBucket)

- [S3 on Outposts API operations for managing Outposts](#S3OutpostsAPIs)

## Amazon S3 API operations for managing objects

S3 on Outposts is designed to use the same object API operations as Amazon S3. You must use
access points to access any object in an Outpost bucket. When you use an object API operation
with S3 on Outposts, you provide either the Outposts access point Amazon Resource Name
(ARN) or the access point alias. For more information about access point aliases, see [Using a bucket-style alias for your S3 on Outposts bucket access point](s3-outposts-access-points-alias.md).

Amazon S3 on Outposts supports the following Amazon S3 API operations:

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CopyObject](../api/api-copyobject.md)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjects](../api/api-deleteobjects.md)

- [DeleteObjectTagging](../api/api-deleteobjecttagging.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectTagging](../api/api-getobjecttagging.md)

- [HeadBucket](../api/api-headbucket.md)

- [HeadObject](../api/api-headobject.md)

- [ListMultipartUploads](../api/api-listmultipartuploads.md)

- [ListObjects](../api/api-listobjects.md)

- [ListObjectsV2](../api/api-listobjectsv2.md)

- [ListObjectVersions](../api/api-listobjectversions.md)

- [ListParts](../api/api-listparts.md)

- [PutObject](../api/api-putobject.md)

- [PutObjectTagging](../api/api-putobjecttagging.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)

## Amazon S3 Control API operations for managing buckets

S3 on Outposts supports the following Amazon S3 Control API operations for working with
buckets.

- [CreateAccessPoint](../api/api-control-createaccesspoint.md)

- [CreateBucket](../api/api-control-createbucket.md)

- [DeleteAccessPoint](../api/api-control-deleteaccesspoint.md)

- [DeleteAccessPointPolicy](../api/api-control-deleteaccesspointpolicy.md)

- [DeleteBucket](../api/api-control-deletebucket.md)

- [DeleteBucketLifecycleConfiguration](../api/api-control-deletebucketlifecycleconfiguration.md)

- [DeleteBucketPolicy](../api/api-control-deletebucketpolicy.md)

- [DeleteBucketReplication](../api/api-control-deletebucketreplication.md)

- [DeleteBucketTagging](../api/api-control-deletebuckettagging.md)

- [GetAccessPoint](../api/api-control-getaccesspoint.md)

- [GetAccessPointPolicy](../api/api-control-getaccesspointpolicy.md)

- [GetBucket](../api/api-control-getbucket.md)

- [GetBucketLifecycleConfiguration](../api/api-control-getbucketlifecycleconfiguration.md)

- [GetBucketPolicy](../api/api-control-getbucketpolicy.md)

- [GetBucketReplication](../api/api-control-getbucketreplication.md)

- [GetBucketTagging](../api/api-control-getbuckettagging.md)

- [GetBucketVersioning](../api/api-control-getbucketversioning.md)

- [ListAccessPoints](../api/api-control-listaccesspoints.md)

- [ListRegionalBuckets](../api/api-control-listregionalbuckets.md)

- [PutAccessPointPolicy](../api/api-control-putaccesspointpolicy.md)

- [PutBucketLifecycleConfiguration](../api/api-control-putbucketlifecycleconfiguration.md)

- [PutBucketPolicy](../api/api-control-putbucketpolicy.md)

- [PutBucketReplication](../api/api-control-putbucketreplication.md)

- [PutBucketTagging](../api/api-control-putbuckettagging.md)

- [PutBucketVersioning](../api/api-control-putbucketversioning.md)

## S3 on Outposts API operations for managing Outposts

S3 on Outposts supports the following Amazon S3 on Outposts API operations for managing
endpoints.

- [CreateEndpoint](../api/api-s3outposts-createendpoint.md)

- [DeleteEndpoint](../api/api-s3outposts-deleteendpoint.md)

- [ListEndpoints](../api/api-s3outposts-listendpoints.md)

- [ListOutpostsWithS3](../api/api-s3outposts-listoutpostswiths3.md)

- [ListSharedEndpoints](../api/api-s3outposts-listsharedendpoints.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Supported regions

Configuring S3 control
client
