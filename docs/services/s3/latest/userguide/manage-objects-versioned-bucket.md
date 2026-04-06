# Working with objects in a versioning-enabled bucket

Objects that are stored in an Amazon S3 bucket before you set the versioning state have a
version ID of `null`. When you enable versioning, existing objects in your bucket
do not change. What changes is how Amazon S3 handles the objects in future requests.

###### Transitioning object versions

You can define lifecycle configuration rules for objects that have a well-defined lifecycle
to transition object versions to the S3 Glacier Flexible Retrieval storage class at a
specific time in the object's lifetime. For more information, see [Managing the lifecycle of objects](object-lifecycle-mgmt.md).

The topics in this section explain various object operations in a versioning-enabled
bucket. For more information about versioning, see [Retaining multiple versions of objects with S3 Versioning](versioning.md).

###### Topics

- [Adding objects to versioning-enabled buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/AddingObjectstoVersioningEnabledBuckets.html)

- [Listing objects in a versioning-enabled bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/list-obj-version-enabled-bucket.html)

- [Retrieving object versions from a versioning-enabled bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RetrievingObjectVersions.html)

- [Deleting object versions from a versioning-enabled bucket](deletingobjectversions.md)

- [Configuring versioned object permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VersionedObjectPermissionsandACLs.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuring MFA delete

Adding objects
