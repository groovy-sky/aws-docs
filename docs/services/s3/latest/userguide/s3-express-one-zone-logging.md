# Logging with AWS CloudTrail for directory buckets

Amazon S3 is integrated with AWS CloudTrail, a service that provides a record of actions
taken by a user, role, or an AWS service. CloudTrail captures all API calls for Amazon S3
as events. Using the information collected by CloudTrail, you can determine the request that
was made to Amazon S3, the IP address from which the request was made, when it was made, and
additional details. When a supported event activity occurs in Amazon S3, that activity is
recorded in a CloudTrail event. You can use AWS CloudTrail trail to log management events
and data events for directory buckets. For more information, see [Amazon S3 CloudTrail events](cloudtrail-logging-s3-info.md) and [What is AWS CloudTrail?](../../../awscloudtrail/latest/userguide/cloudtrail-user-guide.md) in the _AWS CloudTrail User_
_Guide_.

## CloudTrail management events for directory buckets

By default, CloudTrail logs bucket-level actions for directory buckets as management
events. The `eventsource` for CloudTrail management events for directory buckets is `s3express.amazonaws.com`.
When you set up your AWS account,
CloudTrail management events are enabled by default. The following Regional endpoint API
operations (bucket-level, or control plane, API operations) are logged to CloudTrail.

- [`CreateBucket`](../api/api-createbucket.md)

- [`DeleteBucket`](../api/api-deletebucket.md)

- [`DeleteBucketPolicy`](../api/api-deletebucketpolicy.md)

- [`PutBucketPolicy`](../api/api-putbucketpolicy.md)

- [`GetBucketPolicy`](../api/api-getbucketpolicy.md)

- [`ListDirectoryBuckets`](../api/api-listdirectorybuckets.md)

- [`ListMultipartUploads`](../api/api-listmultipartuploads.md)

- [`GetBucketEncryption`](../api/api-getbucketencryption.md)

- [`PutBucketEncryption`](../api/api-putbucketencryption.md)

- [`DeleteBucketEncryption`](../api/api-deletebucketencryption.md)

###### Note

`ListMultipartUploads` is a Zonal endpoint API operation. However, this
API operation is logged to CloudTrail as a management event. For more information,
see [`ListMultipartUploads`](../api/api-listmultipartuploads.md) in the _Amazon Simple_
_Storage Service API Reference_.

For more information on CloudTrail management events, see [Logging management events](../../../awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.md) in the _AWS CloudTrail User_
_Guide._

## CloudTrail data events for directory buckets

Data events provide information about the resource operations performed on or in a
resource (for example, reading or writing to an Amazon S3 object). These are also known
as data plane operations. Data events are often high-volume activities. By default,
CloudTrail trails don't log data events, but you can configure trails to log data events
for objects stored in general purpose buckets and directory buckets. For more
information, see [Enable logging for objects in a bucket using the console](enable-cloudtrail-logging-for-s3.md#enable-cloudtrail-events).

When you log data events for a trail in CloudTrail, you can choose to use advanced
event selectors or basic event selectors. To log data events for objects stored in
directory buckets, you must use advanced event selectors. When configuring advanced
resource selectors, you will choose or specify the resource type
which is `AWS::S3Express::Object`.

The following Zonal endpoint API operations (object-level , or. data plane, API
operations) are logged to CloudTrail.

- [`AbortMultipartUpload`](../api/api-abortmultipartupload.md)

- [`CompleteMultipartUpload`](../api/api-completemultipartupload.md)

- [`CreateSession`](../api/api-createsession.md)

- [`CopyObject`](../api/api-copyobject.md)

- [`CreateMultipartUpload`](../api/api-createmultipartupload.md)

- [`DeleteObject`](../api/api-deleteobject.md)

- [`DeleteObjects`](../api/api-deleteobjects.md)

- [`GetObject`](../api/api-getobject.md)

- [`GetObjectAttributes`](../api/api-getobjectattributes.md)

- [`HeadBucket`](../api/api-headbucket.md)

- [`HeadObject`](../api/api-headobject.md)

- [`ListObjectsV2`](../api/api-listobjectsv2.md)

- [`ListParts`](../api/api-listparts.md)

- [`PutObject`](../api/api-putobject.md)

- [`RenameObject`](../api/api-renameobject.md)

- [`UploadPart`](../api/api-uploadpart.md)

- [`UploadPartCopy`](../api/api-uploadpartcopy.md)

For more information on CloudTrail data events, see [Logging data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) in the _AWS CloudTrail User_
_Guide_.

For additional information about CloudTrail events for directory buckets, see the
following topics:

###### Topics

- [CloudTrail log file examples for directory buckets](s3-express-log-files.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deleting a tag from an access point for directory buckets

Next

All content copied from https://docs.aws.amazon.com/.
