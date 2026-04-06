# Creating, configuring, and working with Amazon S3 general purpose buckets

To store your data in Amazon S3, you work with resources known as buckets and objects. A
_bucket_ is a container for objects. An _object_ is a file and any metadata that describes that
file.

To store an object in Amazon S3, you create a bucket and then upload the object to a bucket.
When the object is in the bucket, you can open it, download it, and move it. When you no
longer need an object or a bucket, you can clean up your resources.

The topics in this section provide an overview of working with general purpose buckets in Amazon S3. They
include information about naming, creating, accessing, and deleting general purpose buckets. For more
information about viewing or listing objects in a bucket, see [Organizing, listing, and working with your objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/organizing-objects.html).

There are several types of Amazon S3 buckets. Before creating a bucket, make sure that you choose the bucket type that best fits your application and performance requirements. For more information about the various bucket types and the appropriate use cases for each, see [Buckets](welcome.md#BasicsBucket).

###### Note

For more information about using the Amazon S3 Express One Zone storage class with directory buckets, see [S3 Express One Zone](directory-bucket-high-performance.md#s3-express-one-zone) and [Working with directory buckets](directory-buckets-overview.md).

###### Note

With Amazon S3, you pay only for what you use. For more information about Amazon S3 features and
pricing, see [Amazon S3](https://aws.amazon.com/s3). If you are a new Amazon S3
customer, you can get started with Amazon S3 for free. For more information, see [AWS Free Tier](https://aws.amazon.com/free).

###### Topics

- [General purpose buckets overview](https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingBucket.html)

- [Namespaces for general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/gpbucketnamespaces.html)

- [Common general purpose bucket patterns for building applications on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/common-bucket-patterns.html)

- [General purpose bucket naming rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)

- [General purpose bucket quotas, limitations, and restrictions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html)

- [Accessing an Amazon S3 general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-bucket-intro.html)

- [Creating a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html)

- [Viewing the properties for an S3 general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/view-bucket-properties.html)

- [Listing Amazon S3 general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/list-buckets.html)

- [Emptying a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/empty-bucket.html)

- [Deleting a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/delete-bucket.html)

- [Mount an Amazon S3 bucket as a local file system](https://docs.aws.amazon.com/AmazonS3/latest/userguide/mountpoint.html)

- [Working with Storage Browser for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-browser.html)

- [Configuring fast, secure file transfers using Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/transfer-acceleration.html)

- [Using Requester Pays general purpose buckets for storage transfers and usage](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RequesterPaysBuckets.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using Amazon S3 with the AWS CLI

General purpose buckets overview
