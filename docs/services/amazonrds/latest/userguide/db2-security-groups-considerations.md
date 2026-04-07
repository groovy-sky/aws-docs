# Considerations for security groups with Amazon RDS for Db2

For you to connect to your Amazon RDS for Db2 DB instance, it must be associated with a security
group that contains the necessary IP addresses and network configuration. Your RDS for Db2 DB
instance might use the default security group. If you assigned a default nonconfigured
security group when you created the RDS for Db2 DB instance, then the firewall prevents
internet connections. For information about creating a new security group, see [Controlling access with security groups](overview-rdssecuritygroups.md).

After you create the new security group, you modify your DB instance to associate it with the
security group. For more information, see [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

You can enhance security by using SSL to encrypt connections to your DB instance. For more
information, see [Using SSL/TLS with an Amazon RDS for Db2 DB instance](db2-concepts-ssl.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IBM Db2 Data Management Console

Securing Db2 connections
