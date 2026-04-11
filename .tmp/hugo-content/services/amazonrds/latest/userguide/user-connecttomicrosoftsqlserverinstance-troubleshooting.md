# Troubleshooting connections to your SQL Server DB instance

The following table shows error messages that you might encounter when you attempt to
connect to your SQL Server DB instance.

IssueTroubleshooting suggestions

**`Could not open a connection to SQL Server –**
**Microsoft SQL Server, Error: 53`**

Make sure that you specified the server name correctly. For
**Server name**, enter the DNS name and port
number of your sample DB instance, separated by a comma.

###### Important

If you have a colon between the DNS name and port number, change the colon to a comma.

Your server name should look like the following
example.

```

sample-instance.cg034itsfake.us-east-1.rds.amazonaws.com,1433
```

**`No connection could be made because the target machine**
**actively refused it – Microsoft SQL Server, Error:**
**10061`**

You were able to reach the DB instance but the connection was
refused. This issue is usually caused by specifying the user name or
password incorrectly. Verify the user name and password, then retry.

**`A network-related or instance-specific error**
**occurred while establishing a connection to SQL Server. The server**
**was not found or was not accessible... The wait operation timed out**
**– Microsoft SQL Server, Error: 258`**The access rules enforced by your local firewall and the IP addresses
authorized to access your DB instance might not match. The problem is
most likely the inbound rules in your security group. For more information, see
[Security in Amazon RDS](usingwithrds.md).

Your database instance must be publicly accessible. To connect to
it from outside of the VPC, the instance must have a public IP
address assigned.

###### Note

For more information on connection issues, see
[Can't connect to Amazon RDS DB instance](chap-troubleshooting.md#CHAP_Troubleshooting.Connecting).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Security group considerations

SQL Server Developer Edition

All content copied from https://docs.aws.amazon.com/.
