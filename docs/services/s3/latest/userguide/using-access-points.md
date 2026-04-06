# Using Amazon S3 access points for general purpose buckets

The following examples demonstrate how to use access points for general purpose buckets with compatible operations in Amazon S3.

###### Note

S3 automatically generate access point aliases for all access points and these aliases can be used
anywhere a bucket name is used to perform object-level operations. For more
information, see [Access point aliases](access-points-naming.md#access-points-alias).

You can only use access points for general purpose buckets to perform operations on objects. You can't use access points
to perform other Amazon S3 operations, such as modifying or deleting buckets. For a
complete list of S3 operations that support access points, see [Access point compatibility](access-points-service-api-support.md).

###### Topics

- [List objects through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/list-object-ap.html)

- [Download an object through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/get-object-ap.html)

- [Configure access control lists (ACLs) through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/put-acl-permissions-ap.html)

- [Upload an object through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/put-object-ap.html)

- [Add a tag-set through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/add-tag-set-ap.html)

- [Delete an object through an access point for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/delete-object-ap.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Delete your access point for a general purpose bucket

List objects through an access point for a general purpose bucket
