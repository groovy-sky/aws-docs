# AWS Backup quotas

The following quotas apply when working with AWS Backup.

###### Quotas

- [Backup](#aws-backup-quotas-table)

- [Backup index and search quotas](#aws-backup-search-quotas-table)

- [Policy quotas](#aws-backup-policies-quotas-table)

- [Malware scanning quotas](#backup-scanning-quotas-table)

- [Amazon Timestream resource quotas](#backup-timestream-quotas-table)

- [AWS Backup Audit Manager quotas](#backup-audit-manager-quotas-table)

- [Restore testing plan quotas](#backup-restore-testing-quotas-table)

- [AWS Backup gateway quotas](#backup-gateway-quotas-table)

- [Amazon EKS quotas](#backup-eks-quotas-table)

- [Logically air-gapped vault quotas](#lag-vault-quotas-table)

- [Related quotas](#backup-related-quotas)

## Backup

NameDefaultAdjustableTotal vaults (backup and logically air-gapped) per Region per account300YesRecovery points per backup vault1,000,000YesBackup plans per Region per account300YesVersions per backup plan2,000YesResource assignments per backup plan100NoAmazon S3 buckets per account100YesConcurrent cross-Region or in-Region copy jobs per account in destination Region1002NoAdditional cross-Region copy jobs per vault in a destination Region
after the limit in row above entry has been reached.152NoConcurrent cross-account copies that can be made of the same resource to the same destination
Region30NoConcurrent backup and copy jobs per resource1NoTags per resource selection in a cross account backup policy30No. Include additional tags using multiple resource assignments or backups plans.Hypervisors20NoLegal holds per account50NoNested backup layers of application stacks10No

1The limit for concurrent copy jobs from one Region to
another Region is 100 per account per Region. Once this limit is reached, if a specific
vault in the destination Region has fewer than 5 concurrent copy jobs, new copy jobs
can begin, up to a maximum of 5 concurrently.

2Limit only apply to resource types [fully managed by AWS Backup](backup-feature-availability.md).

## Backup index and search quotas

NameDefaultAdjustableConcurrent indexing jobs in account (most AWS Regions)40YesConcurrent indexing jobs in account in Asia Pacific (Malaysia),
Canada (Central), Asia Pacific (Thailand), and
Mexico (Central) AWS Regions.10YesConcurrent indexing jobs for each resource5NoConcurrent on-demand indexing job1NoConcurrent search jobs in account10Concurrent export jobs5Number of recovery points included in search job20Concurrent Amazon EBS file level restore jobs (most AWS Regions)25Concurrent Amazon EBS file level restore jobs in Asia Pacific (Malaysia),
Canada (Central), Asia Pacific (Thailand), and
Mexico (Central) AWS Regions.5

## Policy quotas

NameDefaultAdjustableResource assignments per backup plan100NoTags in a resource selection30NoResource selections that use tags in a plan10NoBackup plan rules in a plan10NoTags added to a recovery point10NoCopy actions per backup rule5NoConditions in a resource assignment in a backup plan30No

## Malware scanning quotas

The following table lists quotas for AWS Backup malware scanning with Amazon GuardDuty.

NameDefaultAdjustableRunning scan jobs per resource per account5YesRunning scan jobs per recovery point per account1NoRunning scan jobs per account150YesCreated scan jobs per resource per account10NoIncremental scan base constraint1No

When you hit your created scan jobs per resource per account limit, we will fail oldest queued job.

You might also encounter quotas imposed by Amazon GuardDuty. For more information, see
[Amazon GuardDuty quotas](../../../guardduty/latest/ug/guardduty-limits.md)
in the _AWS General Reference_.

## Amazon Timestream resource quotas

NameDefaultAdjustableConcurrent Timestream backup jobs per account4YesConcurrent Timestream restore jobs per account1Yes

## AWS Backup Audit Manager quotas

NameDefaultAdjustableFrameworks per account per Region15YesControls per account per Region50YesReport plans per account20YesFrameworks per report plan1,000No\[Number of accounts\] multiplied by \[number of Regions in a report plan\]
multiplied by \[number of daily jobs plus evaluations in a report plan\]100,000No

## [Restore testing](restore-testing.md) plan quotas

NameDefaultAdjustableRestore testing plans100NoTags per plan50NoSelections per plan30NoARNs per restore testing selection30NoConditions per selection (both `StringEquals` and
`StringNotEquals`)30NoVault selectors per restore testing selection30NoMaximum value (in days) of selection window365 daysNoBoundaries of start window hoursMinimum: 1 hour; Maximum: 168 hoursNoMaximum character length of restore testing plan name50 characters (alphanumeric and underscores, no white spaces)NoMaximum character length of restore testing selection name50 characters (alphanumeric and underscores, no white spaces)No

Each resource type has a limit on the number of concurrent restore jobs that can exist
at one time for restore jobs that are created through a restore testing plan. Once this
limit is reached, no new restore jobs for that resource type will be created until a job
in a state of `RUNNING` transitions to `COMPLETED`.

If a scheduled restore job did not start due to this quota, that job will result in a
`FAILED` status with the status message `"Restore job was unable to
                start within the specified start window. Try increasing your start window."`.
If you receive a failed job with this status message, the best practice is to first
increase your start window with sufficient time to allow jobs to finish. Then, retry the
jobs.

Note quotas do not apply to on-demand restore jobs, but to restore jobs created by a
[restore testing plan](restore-testing.md#restore-testing-create). For some
resource types, you may request an increase in the quota limit.

NameDefaultAdjustableAmazon Aurora40YesAmazon DocumentDB40YesAmazon DynamoDB40NoAmazon EBS100YesAmazon EC2100YesAmazon EFS30YesAmazon FSx40YesAmazon Neptune40YesAmazon RDS40YesAmazon S330Yes

## AWS Backup gateway quotas

NameDefaultAdjustableBackup or restore jobs per gateway4No. Create more gateways and connect them to your hypervisor.VMware tags to resource tags mapping per hypervisor10No

## Amazon EKS quotas

NameDefaultAdjustableNamespaces per EKS cluster backup1000YesPersistent Storage backups per EKS cluster backup1000YesRestore jobs per target EKS cluster1NoEKS Restore jobs per account5Yes

## Logically air-gapped vault quotas

Resource typeMaximum number of concurrent copiesEC2The EBS concurrent copy limit applies to any snapshots

being copied as part of an AMI copy.

EBS20Aurora20DocumentDB20Neptune20Storage Gateway5FSx5

## Related quotas

There are [quotas on a single resource assignment](assigning-resources.md#assigning-resources-quotas) in a single backup rule. You can create a backup plan
with multiple backup rules.

When you manage backups across multiple accounts using AWS Organizations, you might encounter
quotas that AWS Organizations imposes. For these quotas, see [Quotas for\
AWS Organizations](../../../organizations/latest/userguide/orgs-reference-limits.md) in the _AWS Organizations User Guide_.

You might also encounter quotas imposed by a AWS Backup-supported service,
including the following:

- [Amazon Elastic File System](../../../efs/latest/ug/limits.md)

- [Amazon Elastic Block Store](../../../ebs/latest/userguide/ebs-resource-quotas.md)

- [Amazon RDS](../../../amazonrds/latest/userguide/chap-limits.md)

- [Amazon Aurora](../../../amazonrds/latest/aurorauserguide/chap-limits.md)

- [Amazon EC2](../../../ec2/latest/userguide/ec2-resource-limits.md)

- [AWS Storage Gateway](../../../storagegateway/latest/userguide/resource-gateway-limits.md)

- [Amazon DynamoDB](../../../dynamodb/latest/developerguide/servicequotas.md)

- [Amazon FSx for Lustre](../../../fsx/latest/lustreguide/limits.md)

- [Amazon FSx for Windows File Server](../../../fsx/latest/windowsguide/managing-user-quotas.md)

- [Amazon DocumentDB](../../../documentdb/latest/developerguide/limits.md)

- [Amazon Neptune](../../../neptune/latest/userguide/limits.md)

- [Amazon Simple Storage Service](../../../../general/latest/gr/s3.md#limits_s3)

- [Amazon Timestream](../../../timestream/latest/developerguide/backups-limits.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Resilience

Monitoring AWS Backup

All content copied from https://docs.aws.amazon.com/.
