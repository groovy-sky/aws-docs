# AWS Backup feature availability

AWS Backup features are offered according to resource and AWS Region.
The following sections and tables can help you determine feature availability.

###### Contents

- [Features available for all supported resources](#features-for-all-resources)

- [Feature availability by resource](#features-by-resource)

- [Feature availability by AWS Region](#features-by-region)

- [Supported services by AWS Region](#supported-services-by-region)

## Features available for all supported resources

AWS Backup offers the following features for its supported AWS services, as well as
for supported third-party applications. Support for a feature or service should not be
assumed unless explicitly mentioned.

- [Automated backup schedules and retention management](about-backup-plans.md)

- [Centralized backup monitoring](monitoring.md)

- [Encrypted backups](encryption.md)

- [Incremental backups](creating-a-backup.md)

- [Cross-account management with AWS Organizations](manage-cross-account.md)

- [Automated backup audits\
and reports with AWS Backup Audit Manager](aws-backup-audit-manager.md)

- [Write-once, read-many (WORM) with AWS Backup Vault Lock](vault-lock.md)

## Feature availability by resource

To use AWS Backup with a supported AWS service in a particular Region, the service must be available in the
Region. To determine service availability in a Region, view the
[service endpoints](../../../../general/latest/gr/aws-service-information.md)
in the _AWS General Reference_.

For information on opt-in Regions and what resources and features are supported within,
see [Feature availability by AWS Region](#features-by-region).

AWS Backup supports[Cross-Region backup](cross-region-backup.md)[Cross-account backup](create-cross-account-backup.md)[AWS Backup Audit Manager](aws-backup-audit-manager.md)[Incremental backup](about-backup-plans.md)[Continuous backup and point-in-time restore](point-in-time-recovery.md)[Full management](whatisbackup.md#full-management)[Lifecycle to cold storage](plan-options-and-configuration.md)Item-level restore 1[Restore testing](restore-testing.md)[Logically air-gapped vault](logicallyairgappedvault.md)[Backup search](backup-search.md)[Backup tiering](backup-tiering.md)[Malware Protection](malware-protection.md)Amazon EC2✓✓✓✓✓✓✓Amazon S3✓✓✓✓✓✓✓✓✓✓✓✓Amazon EBS✓✓✓✓✓✓✓✓✓✓Amazon RDS single instance✓ 3✓ 3✓ 4✓✓✓Amazon RDS cluster✓ 3✓ 3✓ 4✓✓Amazon Aurora✓ 3✓ 3✓✓ 6✓✓✓Amazon Aurora DSQL✓✓✓✓✓Amazon EFS✓✓✓✓✓✓✓✓✓FSx for Lustre✓✓✓✓✓✓FSx for Windows File Server✓✓✓✓✓8✓FSx for ONTAP✓ 2✓✓FSx for OpenZFS✓✓✓✓AWS Storage Gateway✓✓✓✓✓Amazon DocumentDB✓ 3✓ 3✓✓Amazon Neptune✓ 3✓ 3✓✓ 9Amazon Redshift Serverless✓Amazon Timestream✓✓✓✓✓✓Windows VSS✓✓✓✓✓Virtual machines✓✓✓✓✓✓✓✓AWS CloudFormation✓ 5✓✓ 5✓Amazon DynamoDB✓✓DynamoDB with [AWS Backup advanced features](advanced-ddb-backup.md)✓✓✓✓✓✓✓SAP HANA databases on Amazon EC2 instances✓✓✓ 6✓✓✓Amazon EKS10✓✓✓✓✓11

Some resource types have both continuous backup capability and cross-Region and cross-account copy
available. When a cross-Region or cross-account copy of a continuous backup is made, the
copied recovery point (backup) becomes a snapshot (periodic) backup. PITR (Point-in-Time
Restore) is not available for these copies.

- Amazon RDS and Amazon S3 support cross-account and cross-Region copy from incremental
backups. Amazon RDS also supports simultaneous cross-Region and cross-account snapshot copying in a single action.

- Amazon Aurora and SAP HANA on Amazon EC2 instances support cross-account and
cross-Region copy from full backups. Amazon Aurora also supports simultaneous cross-Region and cross-account snapshot copying in a single action.

1 The "item" in an item-level restore varies depending on the
supported resource. For example, a file system item is a file or directory, whereas an S3
item is an S3 object. A VMware item is a disk. For more information, see the [Restore a backup by resource type](restoring-a-backup.md) section for the
supported resource.

2 AWS Backup Audit Manager supports this resource across all
controls except [cross-account copy](controls-and-remediation.md#backup-cross-account-copy) and
[cross-Region copy](controls-and-remediation.md#backup-cross-region-copy).

3 Amazon RDS, Aurora, DocumentDB, and Neptune now support cross-Region and cross-account snapshot copying in a single action. RDS multi availability
zone (Multi-AZ) database instances can be copied, but Multi-AZ clusters do not currently
support any copy operations. See [Cross-Region copy considerations with specific resources](cross-region-backup.md#cross-region-considerations) for further information.

4 See [RDS multi-availability zone\
backups](rds-multiaz-backup.md) for Regions where Backup Audit Manager support is available.

5 In [CloudFormation stack\
backups](applicationstackbackups.md), nested resources retain their source resource features. However,
resources within the stack do not retain Point-in-Time Restore (PITR) functionality (such as
Amazon S3 and Amazon RDS).

6 Snapshots are full and incremental backup is offered
through PITR.

7 Amazon FSx for OpenZFS Multi-AZ (multi-availability zone) file
systems can only be restored from the Amazon FSx console or the API request [`createFileSystemFromBackup`](../../../../reference/fsx/latest/apireference/api-createfilesystemfrombackup.md).

8 Is supported in a restore test if FSx for Windows File Server uses AWS
managed active directory

9 Is not currently available in Asia Pacific (Jakarta) Region

10 Features for Persistent Storage Backups (e.g. Amazon EBS) taken as part of Amazon EKS Backups will reflect the feature support for the respective resource types in this table.

11 Is not currently available in Middle East (Bahrain) and Middle East (UAE)

## Feature availability by AWS Region

AWS Backup is available in all the following AWS Regions.
AWS Backup features are available in all
these Regions unless otherwise noted in the following table.

Some Regions require account opt-in, as noted in the following table. Some feature
availability is determined by whether opt-in is required or not required. For more
information, see [AWS Regions your account can\
use](../../../accounts/latest/reference/manage-acct-regions.md) in the _AWS Account Management Reference Guide_.

**Considerations for opt-in Regions:**

- **Cross-account** copy is **not supported** for
Amazon DocumentDB in Regions where opt-in is required.

- **Cross-Region** copy is **not supported** for
Amazon DocumentDB in Regions where opt-in is required.

- **Cross-Region** copy of Neptune backups is currently
**supported** in Africa (Cape Town), Asia Pacific (Hong Kong),
Asia Pacific (Jakarta), Israel (Tel Aviv),
Middle East (Bahrain), and Middle East (UAE) Regions.

**Cross-Region** copy of FSx for Lustre, FSx for Windows File Server, FSx for ONTAP, and FSx for OpenZFS is
**not supported** in Regions where opt-in is required.

- **Cross-account** copy is **not supported** for CloudFormation, Neptune,
and Timestream in Regions where opt-in is required.

**Considerations and limitations for cross-account management in opt-in**
**Regions:**

- Cross-account management in AWS Regions where opt-in is required includes
cross-account monitoring and access to backup policies; delegated administrator
accounts can launch policies but do not have access to the monitoring
functions.

- Both management accounts and their child accounts can be opted into AWS Organizations. If
a child account is opted into cross-account management prior to its management account
being opted into cross-account management, there will be a delay (up to 24 hours)
before cross-account monitoring will show job statuses across the organization.

AWS Backup supportsOpt-in[Cross-Region backup copy](cross-region-backup.md)[Cross-account management](manage-cross-account.md)[Cross-account backup copy](create-cross-account-backup.md)[AWS Backup Audit Manager](aws-backup-audit-manager.md) and [Jobs dashboard](backup-dashboards.md)[Restore testing](restore-testing.md)[Backup search](backup-search.md)[Backup tiering](backup-tiering.md)[Malware Protection](malware-protection.md)US East (N. Virginia)Not required✓✓✓✓✓✓✓✓US East (Ohio)Not required✓✓✓✓✓✓✓✓US West (N. California)Not required✓✓✓✓✓✓✓✓US West (Oregon)Not required✓✓✓✓✓✓✓✓Africa (Cape Town)Required✓✓✓✓✓✓✓✓Asia Pacific (Hong Kong)Required✓✓✓✓✓✓✓✓Asia Pacific (Hyderabad)Required✓✓✓✓✓✓✓✓Asia Pacific (Jakarta)Required✓✓✓✓✓✓✓✓Asia Pacific (Malaysia)Required✓✓✓✓✓✓Asia Pacific (Melbourne)Required✓✓✓✓✓✓✓✓Asia Pacific (Mumbai)Not required✓✓✓✓✓✓✓✓Asia Pacific (New Zealand)Required✓✓✓Asia Pacific (Osaka)Not required✓✓✓✓✓✓✓✓Asia Pacific (Seoul)Not required✓✓✓✓✓✓✓✓Asia Pacific (Singapore)Not required✓✓✓✓✓✓✓✓Asia Pacific (Sydney)Not required✓✓✓✓✓✓✓✓Asia Pacific (Taipei)Required✓✓✓✓Asia Pacific (Thailand)Required✓✓✓✓✓Asia Pacific (Tokyo)Not required✓✓✓✓✓✓✓✓Canada (Central)Not required✓✓✓✓✓✓✓✓Canada West (Calgary)Required✓✓✓✓✓✓China (Beijing)[AWS in China](https://www.amazonaws.cn/en/about-aws/china)✓2✓China (Ningxia)[AWS in China](https://www.amazonaws.cn/en/about-aws/china)✓2✓Europe (Frankfurt)Not required✓✓✓✓✓✓✓✓Europe (Ireland)Not required✓✓✓✓✓✓✓✓Europe (London)Not required✓✓✓✓✓✓✓✓Europe (Milan)Required✓✓✓✓✓✓✓✓Europe (Paris)Not required✓✓✓✓✓✓✓✓Europe (Spain)Required✓✓✓✓✓✓✓✓Europe (Stockholm)Not required✓✓✓✓✓✓✓✓Europe (Zurich)Required✓✓✓✓✓✓✓✓Israel (Tel Aviv)Required✓✓✓✓✓✓Mexico (Central)Required✓✓✓✓✓Middle East (Bahrain)Required✓✓✓✓✓✓✓✓Middle East (UAE)Required✓✓✓✓✓✓✓✓South America (São Paulo)Not required✓✓✓✓✓✓✓✓AWS GovCloud (US-East)[AWS GovCloud (US)](https://aws.amazon.com/govcloud-us)✓✓✓4✓✓✓✓AWS GovCloud (US-West)[AWS GovCloud (US)](https://aws.amazon.com/govcloud-us)✓✓✓4✓✓✓✓AWS European Sovereign Cloud (Germany)[AWS European Sovereign Cloud (Germany)](https://aws.amazon.com/aws.eu)✓

1Cross-Region and cross-account copy to a logically
air-gapped vault is not currently available in Asia Pacific (Malaysia), Canada West (Calgary),
Mexico (Central), Asia Pacific (Thailand), Asia Pacific (Taipei), Asia Pacific (New Zealand),
China (Beijing), China (Ningxia), AWS GovCloud (US-East), or AWS GovCloud (US-West)
Regions.

2China (Beijing) and China (Ningxia) support cross-Region
copy from one of these two Regions to the other. Cross-Region copy is not
supported from these Regions to other Regions
or into these Regions. Cross-account copy is not supported for these
Regions.

3Jobs dashboard and AWS Backup Audit Manager organizational reporting,
and Jobs dashboard aggregation are only
available in Regions that support cross-account management and AWS Backup Audit Manager.

4AWS GovCloud (US-East) and AWS GovCloud (US-West) support
cross-Region copy from one of these two Regions to the other. Cross-Region copy is not
supported from these Regions to other Regions, or from other Regions into these Regions.

## Supported services by AWS Region

AWS Backup is available for these resource types in all Regions in which AWS Backup
and the listed resource operates:

- Aurora

- AWS CloudFormation

- DynamoDB

- DynamoDB with AWS Backup advanced features

- Amazon EBS

- Amazon EC2

- Amazon EFS

- Amazon Redshift

- Amazon RDS

###### Note

AWS Backup cannot be deployed or operated locally on AWS Outposts infrastructure. All backup operations and storage are performed in the associated AWS Region, and local backup retention on Outposts is not supported.

Region and service[Aurora DSQL](backup-aurora.md)[Amazon FSx](restoring-fsx.md)[SAP HANA on EC2 instances](backup-saphana.md)[Amazon S3](s3-backups.md)[Storage Gateway](working-with-gateways.md)[Amazon Timestream](timestream-backup.md)[VMware](backing-up-vms.md) and [Backup gateway](working-with-gateways.md)[Amazon EKS](eks-backups.md)[Amazon Neptune](creating-a-backup.md)[Amazon DocumentDB](creating-a-backup.md)US East (N. Virginia)✓✓✓✓✓✓✓✓✓✓US East (Ohio)✓✓✓✓✓✓✓✓✓✓US West (N. California)✓✓✓✓✓✓✓US West (Oregon)✓✓✓✓✓✓✓✓✓✓Africa (Cape Town)✓✓✓✓✓✓✓Asia Pacific (Hong Kong)✓✓✓✓✓✓✓Asia Pacific (Hyderabad)✓✓✓✓✓Asia Pacific (Jakarta)✓✓✓✓✓Asia Pacific (Malaysia)✓✓Asia Pacific (Melbourne)Windows; Lustre; ONTAP✓✓✓Asia Pacific (Mumbai)✓✓✓✓✓✓✓✓✓Asia Pacific (New Zealand)✓✓Asia Pacific (Osaka)✓✓✓✓✓✓✓✓Asia Pacific (Seoul)✓✓✓✓✓✓✓✓✓Asia Pacific (Singapore)✓✓✓✓✓✓✓✓Asia Pacific (Sydney)✓✓✓✓✓✓✓✓✓Asia Pacific (Taipei)✓✓Asia Pacific (Thailand)✓✓Asia Pacific (Tokyo)✓✓✓✓✓✓✓✓✓✓Canada (Central)✓✓✓✓✓✓✓✓Canada West (Calgary)✓✓China (Beijing)Windows; Lustre✓ 1✓✓✓China (Ningxia)Windows; Lustre✓ 1✓✓✓Europe (Frankfurt)✓✓✓✓✓✓✓✓✓✓Europe (Ireland)✓✓✓✓✓✓✓✓✓✓Europe (London)✓✓✓✓✓✓✓✓✓Europe (Milan)✓✓✓✓✓✓✓Europe (Paris)✓✓✓✓✓✓✓✓✓Europe (Spain)✓✓✓✓Europe (Stockholm)✓✓✓✓✓✓✓Europe (Zurich)✓✓✓✓Israel (Tel Aviv)✓✓✓✓✓Mexico (Central)✓✓Middle East (Bahrain)✓✓✓✓✓✓✓Middle East (UAE)✓✓✓✓✓✓South America (São Paulo)✓✓✓✓✓✓AWS GovCloud (US-West)Windows; Lustre; ONTAP✓✓✓✓✓✓✓AWS GovCloud (US-East)Windows; Lustre; ONTAP✓✓✓✓✓AWS European Sovereign Cloud (Germany)✓

A check under Amazon FSx indicates that FSx for Windows File Server, FSx for Lustre, FSx for ONTAP, and
FSx for OpenZFS are all supported in that Region by AWS Backup; otherwise, the supported
configurations will be listed.

1 Cross-account copy is not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

What is AWS Backup?

How it works

All content copied from https://docs.aws.amazon.com/.
