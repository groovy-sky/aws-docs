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

- [Viewing storage volume details for your DB instance](rds-storage-viewing.md)

- [Increasing DB instance storage capacity](user-piops-modifyingexisting.md)

- [Removing additional storage volumes](user-piops-removingadditionalvolumes.md)

- [Managing capacity automatically with Amazon RDS storage autoscaling](user-piops-autoscaling.md)

- [Upgrading the storage file system for a DB instance](user-piops-upgradefilesystem.md)

- [Modifying settings for Provisioned IOPS SSD storage](user-piops-increase.md)

- [I/O-intensive storage modifications](user-piops-iointensive.md)

- [Modifying settings for General Purpose SSD (gp3) storage](user-piops-gp3.md)

- [Using a dedicated log volume (DLV)](user-piops-dlv.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Getting an existing ARN

Viewing storage volume details for your DB instance

All content copied from https://docs.aws.amazon.com/.
