# Data residency workloads

AWS Dedicated Local Zones (Dedicated Local Zones) are a type of AWS Infrastructure that are fully managed by AWS,
built for exclusive use by you or your community, and placed in a location or data center
specified by you to help comply with regulatory requirements. Dedicated Local Zones are a type of AWS Local Zones
(Local Zones) offering. For more information, see [AWS Dedicated Local Zones](https://aws.amazon.com/dedicatedlocalzones).

In Dedicated Local Zones, you can create S3 directory buckets to store data in a specific data
perimeter, which helps support data residency and isolation use cases.

Directory buckets in Dedicated Local Zones can support the S3 Express One Zone and S3 One Zone-Infrequent Access (S3 One Zone-IA; Z-IA) storage classes.
Directory buckets are not currently available in other [AWS Local Zones locations](https://aws.amazon.com/about-aws/global-infrastructure/localzones/locations).

You can use the
AWS Management Console, REST API, AWS Command Line Interface (AWS CLI), and AWS SDKs in Dedicated Local Zones.

For more information about working with the directory buckets in Local Zones, see the following topics:

###### Topics

- [Concepts for directory buckets in Local Zones](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html)

- [Enable accounts for Local Zones](https://docs.aws.amazon.com/AmazonS3/latest/userguide/opt-in-directory-bucket-lz.html)

- [Private connectivity from your VPC](https://docs.aws.amazon.com/AmazonS3/latest/userguide/connectivity-lz-directory-buckets.html)

- [Creating a directory bucket in a Local Zone](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-directory-bucket-LZ.html)

- [Authenticating and authorizing for directory buckets in Local Zones](https://docs.aws.amazon.com/AmazonS3/latest/userguide/iam-directory-bucket-LZ.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Optimizing S3 Express One Zone performance

Concepts for directory buckets in Local Zones
