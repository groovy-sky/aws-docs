# Backup plans

In AWS Backup, a _backup plan_ is a policy expression that defines when and
how you want to back up your AWS resources, such as Amazon DynamoDB tables or Amazon Elastic File System (Amazon EFS) file
systems. You can assign resources to backup plans, and AWS Backup automatically backs up and retains
backups for those resources according to the backup plan. You can create multiple backup plans
if you have workloads with different backup requirements. By default, backup windows are
optimized by AWS Backup. You can customize the backup window in the console or
programmatically.

AWS Backup efficiently stores your periodic backups incrementally. The first backup of an AWS
resource backs up a full copy of your data. For each successive incremental backup, only the
changes to your AWS resources are backed up. Incremental backups enable you to benefit from
the data protection of frequent backups while minimizing storage costs.

AWS Backup also seamlessly manages your backup plan's lifecycle
based on your retention settings, which allows you to restore when needed.

The following sections provide the basics of managing your backup strategy in AWS Backup.

###### Topics

- [Create a backup plan](https://docs.aws.amazon.com/aws-backup/latest/devguide/creating-a-backup-plan.html)

- [Understanding backup plan summary](https://docs.aws.amazon.com/aws-backup/latest/devguide/understanding-backup-plan-summaries.html)

- [Select AWS services to backup](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Getting started

Create a backup plan
