# Working with storage for Amazon RDS DB instances

To specify how you want your data stored in Amazon RDS, choose a storage type and provide a
storage size when you create or modify a DB instance. Later, you can increase the amount or
change the type of storage by modifying the DB instance. For more information about which
storage type to use for your workload, see [Amazon RDS storage types](chap-storage.md#Concepts.Storage).

If your instances run RDS for Oracle or RDS for SQL Server, you can add up to three additional volumes to each
DB instance. You can choose either gp3 or io2 as the volume type, allowing you to optimize costs
and performance based on your data access patterns. The maximum storage capacity of a DB instance that uses additional volumes is
256 TiB.

###### Topics

- [Viewing storage volume details for your DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-storage-viewing.html)

- [Increasing DB instance storage capacity](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.ModifyingExisting.html)

- [Removing additional storage volumes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.RemovingAdditionalVolumes.html)

- [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.Autoscaling.html)

- [Upgrading the storage file system for a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.UpgradeFileSystem.html)

- [Modifying settings for Provisioned IOPS SSD storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/User_PIOPS.Increase.html)

- [I/O-intensive storage modifications](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.IOIntensive.html)

- [Modifying settings for General Purpose SSD (gp3) storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.gp3.html)

- [Using a dedicated log volume (DLV)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.dlv.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Getting an existing ARN

Viewing storage volume details for your DB instance
