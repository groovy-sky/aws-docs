# Troubleshooting Amazon EFS integration

Your RDS for Oracle DB instance monitors the connectivity to an Amazon EFS file system. When
monitoring detects an issue, it might try to correct the issue and publish an event in the
RDS console. For more information, see [Viewing Amazon RDS\
events](user-listevents.md).

Use the information in this section to help you diagnose and fix common issues when you
work with Amazon EFS integration.

NotificationDescriptionAction

`The EFS for RDS Oracle instance
                                    instance_name isn't available on the
                                primary host. NFS port 2049 of your EFS isn't
                            reachable.`

The DB instance can't communicate with the EFS file system.

Make sure of the following:

- The EFS file system exists.

- The security group that is attached to the EFS mount target
has an inbound rule to allow the security group or network
subnet of the RDS for Oracle DB instance on TCP/2049 (Type
NFS).

`The EFS isn't reachable.`

An error occurred during the installation of the
`EFS_INTEGRATION` option.

Make sure of the following:

- The EFS file system exists.

- The security group that is attached to the EFS mount target
has an inbound rule to allow the security group or network
subnet of the RDS for Oracle DB instance on TCP/2049 (Type
NFS).

- The `enableDnsSupport` attribute is turned on for
your VPC.

- You are using the Amazon provided DNS server in your VPC.
Amazon EFS integration doesn't work with a custom DHCP DNS.

`The associated role with your DB instance wasn't
                            found.`

An error occurred during the installation of the
`EFS_INTEGRATION` option.

Make sure that you associated an IAM role with your RDS for Oracle DB
instance.

`The associated role with your DB instance wasn't
                            found.`

An error occurred during the installation of the
`EFS_INTEGRATION` option. RDS for Oracle was restored from a
DB snapshot with the `USE_IAM_ROLE` option setting of
`TRUE`.

Make sure that you associated an IAM role with your RDS for Oracle
DB instance.

`The associated role with your DB instance wasn't
                            found.`

An error occurred during the installation of the
`EFS_INTEGRATION` option. RDS for Oracle was created from an
all-in-one CloudFormation template with the `USE_IAM_ROLE` option
setting of `TRUE`.

As a workaround, complete the following steps:

1. Create a DB instance with the IAM role and default option
    group.

2. On a subsequent stack update, add the custom option group with
    the `EFS_INTEGRATION` option.

`PLS-00302: component 'CREATE_DIRECTORY_EFS' must be
                                declared`

This error can occur when you're using a version of RDS for Oracle that
doesn't support Amazon EFS.

Make sure that you are using RDS for Oracle DB instance version
19.0.0.0.ru-2022-07.rur-2022-07.r1 or higher.

`Read access of your EFS is denied. Check your file system
                                policy.`

Your DB instance can't read the EFS file system.

Make sure that your EFS file system allows read access through the IAM
role or on the EFS file system level.

N/A

Your DB instance can't write to the EFS file system.

Take the following steps:

1. Make sure that your EFS file system is mounted on an Amazon EC2
    instance.

2. Give the `others` group write access to your RDS
    user. The simplest technique is to run the `chmod
                                           777` command on the top directory of the EFS file
    system.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Removing the option

Java virtual machine (JVM)

All content copied from https://docs.aws.amazon.com/.
