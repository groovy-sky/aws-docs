# Directory bucket API operations

To manage directory buckets, you can use Regional (bucket level, or control plane) endpoint API operations. To manage objects in
your directory buckets, you can use Zonal (object level, or data plane) endpoint API operations. For more information, see [Networking for directory buckets](s3-express-networking.md) and [Endpoints and gateway VPC endpoints](directory-bucket-high-performance.md#s3-express-overview-endpoints).

###### Regional endpoint API operations

The following Regional endpoint API operations are supported for directory buckets:

- [CreateAccessPoint](../api/api-control-createaccesspoint.md)

- [CreateBucket](../api/api-createbucket.md)

- [DeleteAccessPoint](../api/api-control-deleteaccesspoint.md)

- [DeleteAccessPointPolicy](../api/api-control-deleteaccesspointpolicy.md)

- [DeleteAccessPointScope](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointScope.html)

- [DeleteBucket](../api/api-deletebucket.md)

- [DeleteBucketLifecycle](../api/api-deletebucketlifecycle.md)

- [DeleteBucketPolicy](../api/api-deletebucketpolicy.md)

- [GetAccessPoint](../api/api-control-getaccesspoint.md)

- [GetAccessPointPolicy](../api/api-control-getaccesspointpolicy.md)

- [GetAccessPointScope](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointScope.html)

- [GetBucketLifecycleConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html)

- [GetBucketPolicy](../api/api-getbucketpolicy.md)

- [ListAccessPointsForDirectoryBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForDirectoryBuckets.html)

- [ListDirectoryBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListDirectoryBuckets.html)

- [ListTagsForResource](../api/api-control-listtagsforresource.md)

- [PutAccessPointPolicy](../api/api-control-putaccesspointpolicy.md)

- [PutAccessPointScope](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointScope.html)

- [PutBucketLifecycleConfiguration](../api/api-putbucketlifecycleconfiguration.md)

- [PutBucketPolicy](../api/api-putbucketpolicy.md)

- [DeleteBucketEncryption](../api/api-deletebucketencryption.md)

- [GetBucketEncryption](../api/api-getbucketencryption.md)

- [PutBucketEncryption](../api/api-putbucketencryption.md)

- [TagResource](../api/api-control-tagresource.md)

- [UntagResource](../api/api-control-untagresource.md)

###### Zonal endpoint API operations

The following Zonal endpoint API operations are supported for directory buckets:

- [CreateSession](../api/api-createsession.md)

- [CopyObject](../api/api-copyobject.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjects](../api/api-deleteobjects.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectAttributes](../api/api-getobjectattributes.md)

- [HeadBucket](../api/api-headbucket.md)

- [HeadObject](../api/api-headobject.md)

- [ListObjectsV2](../api/api-listobjectsv2.md)

- [PutObject](../api/api-putobject.md)

- [RenameObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_RenameObject.html)

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultiPartUpload](../api/api-completemultipartupload.md)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [ListMultipartUploads](../api/api-listmultipartuploads.md)

- [ListParts](../api/api-listparts.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Working with directory buckets by using the S3 console, AWS CLI, and AWS SDKs

Tagging directory buckets
