# Performing common database tasks for Amazon RDS for Db2 DB instances

You can perform certain common DBA tasks related to databases on your Amazon RDS for Db2 DB
instances. To deliver a managed service experience, Amazon RDS doesn't provide shell access to DB
instances. Also, the master user can't run commands or utilities requiring
`SYSADM`, `SYSMAINT`, or `SYSCTRL` authorities.

For information about common tasks for buffer pools, databases, and tablespaces, see the following topics.

###### Topics

- [Common tasks for buffer pools](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-managing-buffer-pools.html)

- [Common tasks for databases](db2-managing-databases.md)

- [Common tasks for tablespaces](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-managing-tablespaces.html)

## Managing storage

Db2 uses automatic storage to manage the physical storage for database objects such as
tables, indexes, and temporary files. Instead of manually allocating storage space and
keeping track of which storage paths are being used, automatic storage allows the Db2 system
to create and manage storage paths as needed. This can simplify administration of Db2
databases and reduce the likelihood of errors due to human mistakes. For more information,
see [Automatic storage](https://www.ibm.com/docs/en/db2/11.5?topic=overview-automatic-storage) in the IBM Db2 documentation.

With RDS for Db2, you can dynamically increase the storage size with automatic expansion
of the logical volumes and the file system. For more information, see [Working with storage for Amazon RDS DB instances](user-piops-storagetypes.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Attaching to the remote DB instance

Buffer pools
