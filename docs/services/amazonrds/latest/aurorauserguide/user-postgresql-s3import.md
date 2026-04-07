# Importing data from Amazon S3 into an Aurora PostgreSQL DB cluster

You can import data that's been stored using Amazon Simple Storage Service into a table on an Aurora PostgreSQL DB cluster instance. To do this, you first install the Aurora PostgreSQL `aws_s3` extension. This extension provides
the functions that you use to import data from an Amazon S3 bucket. A _bucket_ is an Amazon S3 container for objects
and files.
The data can be in a comma-separate value (CSV) file, a text file,
or a compressed (gzip) file. Following, you can learn how to install the extension and how to import data
from Amazon S3 into a table.

Your database must be running PostgreSQL version 10.7 or higher to import from Amazon S3 into

Aurora PostgreSQL.

If you don't have data stored on Amazon S3, you need to first create a bucket and store the data.
For more information, see the following topics in the _Amazon Simple Storage Service User Guide_.

- [Create a bucket](../../../s3/latest/userguide/getstartedwiths3.md#creating-bucket)

- [Add an object to a\
bucket](../../../s3/latest/userguide/getstartedwiths3.md#uploading-an-object-bucket)

Cross-account import from Amazon S3 is supported. For more information, see [Granting cross-account permissions](../../../s3/latest/userguide/example-walkthroughs-managing-access-example2.md) in the _Amazon Simple Storage Service User Guide_.

You can use the customer managed key for encryption while importing data from S3. For more information, see [KMS keys stored in AWS KMS](../../../s3/latest/userguide/usingkmsencryption.md) in the _Amazon Simple Storage Service User Guide_.

###### Note

Importing data from Amazon S3 isn't supported for Aurora Serverless v1. It is supported for Aurora Serverless v2.

###### Topics

- [Installing the aws\_s3 extension](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PostgreSQL.S3Import.InstallExtension.html)

- [Overview of importing data from Amazon S3 data](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PostgreSQL.S3Import.Overview.html)

- [Setting up access to an Amazon S3 bucket](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PostgreSQL.S3Import.AccessPermission.html)

- [Importing data from Amazon S3 to your Aurora PostgreSQL DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PostgreSQL.S3Import.FileFormats.html)

- [Function reference](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PostgreSQL.S3Import.Reference.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Integrating Aurora PostgreSQL with
AWS services

Installing the extension
