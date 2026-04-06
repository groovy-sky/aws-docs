# Managing access to shared datasets in directory buckets with access points

Amazon S3 Access Points simplify managing data access at scale for shared datasets in Amazon S3. Access points are unique hostnames you create to enforce distinct permissions and network controls for all requests made through an access point. You can create hundreds of access points per bucket, each with a distinct name and permissions customized for each application. Each access point works in conjunction with the bucket policy that is attached to the underlying
bucket.

In directory buckets, an access point name consists of a base name you provide, followed by the Zone ID (AWS Availability Zone or Local Zone) of your directory bucket location, and then `--xa-s3`. For example, `accesspointname--zoneID--xa-s3`. After you create an access point, you can't change the name or the Zone ID.

With access points for directory buckets, you can use the access point scope to restrict access to specific prefixes or API operations. You can specify any amount of prefixes, but the total length of characters of all prefixes must be less than 256 bytes.

You can configure any access point to accept requests only from a virtual private cloud (VPC). This restricts Amazon S3 data access to a private network.

In this section, the topics explain how to use access points for directory buckets. For information about directory buckets, see [Working with directory buckets](directory-buckets-overview.md).

###### Topics

- [Access points for directory buckets naming rules, restrictions, and limitations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-restrictions-limitations-naming-rules.html)

- [Referencing access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-naming.html)

- [Object operations for access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-service-api-support.html)

- [Configuring IAM policies for using access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-policies.html)

- [Monitoring and logging access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-monitoring-logging.html)

- [Creating access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-access-points-directory-buckets.html)

- [Managing your access points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-manage.html)

- [Using tags with S3 Access Points for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-db-tagging.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Security best practices

Naming rules, restrictions, and limitations
