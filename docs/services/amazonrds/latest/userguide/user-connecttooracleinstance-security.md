# Considerations for security groups

For you to connect to your DB instance, it must be associated with a security group that contains the necessary IP addresses and network
configuration. Your DB instance might use the default security group. If
you assigned a default, nonconfigured security group when you created the DB instance, the firewall prevents
connections. For information about creating a new security group, see [Controlling access with security groups](overview-rdssecuritygroups.md).

After you create the new security group, you modify your DB instance to associate it
with the security group. For more information, see
[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

You can enhance security by using SSL to encrypt connections to your DB instance.
For more information, see
[Oracle Secure Sockets Layer](appendix-oracle-options-ssl.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL\*Plus

Dedicated and shared server
processes

All content copied from https://docs.aws.amazon.com/.
