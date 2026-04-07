# Managing an Amazon RDS Custom for Oracle DB instance

###### Note

End of support notice: On March 31, 2027, AWS will end support for Amazon RDS Custom for Oracle. After March 31, 2027, you will no longer be able to access the RDS Custom for Oracle console or RDS Custom for Oracle resources. For more information, see [RDS Custom for Oracle end of support](rds-custom-for-oracle-end-of-support.md).

Amazon RDS Custom supports a subset of the usual management tasks for Amazon RDS DB instances.
Following, you can find instructions for the supported RDS Custom for Oracle management tasks using the
AWS Management Console and the AWS CLI.

###### Topics

- [Working with container databases (CDBs) in RDS Custom for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.multitenant.html)

- [Working with high availability features for RDS Custom for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.ha.html)

- [Customizing your RDS Custom environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.customizing-env.html)

- [Modifying your RDS Custom for Oracle DB instance](custom-managing-modifying.md)

- [Changing the character set of an RDS Custom for Oracle DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.character-set.html)

- [Setting the NLS\_LANG value in RDS Custom for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.nlslang.html)

- [Support for Transparent Data Encryption](#custom-managing.tde)

- [Tagging RDS Custom for Oracle resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.tagging.html)

- [Deleting an RDS Custom for Oracle DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-managing.deleting.html)

## Support for Transparent Data Encryption

RDS Custom supports Transparent Data Encryption (TDE) for RDS Custom for Oracle DB instances.

However, you can't enable TDE using an option in a custom option group as you can in RDS for Oracle. You turn on TDE
manually. For information about using Oracle Transparent Data Encryption, see [Securing stored data using Transparent\
Data Encryption](http://docs.oracle.com/cd/E11882_01/network.112/e40393/asotrans.htm) in the Oracle documentation.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logging in as SYS

Working with container databases (CDBs) in RDS Custom for Oracle
