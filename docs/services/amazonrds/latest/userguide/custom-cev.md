# Working with custom engine versions for Amazon RDS Custom for Oracle

###### Note

End of support notice: On March 31, 2027, AWS will end support for Amazon RDS Custom for Oracle. After March 31, 2027, you will no longer be able to access the RDS Custom for Oracle console or RDS Custom for Oracle resources. For more information, see [RDS Custom for Oracle end of support](rds-custom-for-oracle-end-of-support.md).

A _custom engine version (CEV)_ for Amazon RDS Custom for Oracle is a binary volume
snapshot of a database engine and specific Amazon Machine Image (AMI). By default, RDS Custom for Oracle uses the
latest available AMI managed by RDS Custom, but you can specify an AMI that was used in a previous CEV. You store your database
installation files in Amazon S3. RDS Custom uses the installation files and the AMI to create your
CEV for you.

###### Topics

- [Preparing to create a CEV](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.preparing.html)

- [Creating a CEV](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.create.html)

- [Modifying CEV status](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.modify.html)

- [Viewing CEV details for Amazon RDS Custom for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.view.html)

- [Deleting a CEV](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.delete.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting up your RDS Custom for Oracle environment

Preparing to create a CEV
