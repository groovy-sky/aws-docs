# Other AWS services that use S3 on Outposts

Other AWS services that run local to your AWS Outposts can also use your Amazon S3 on Outposts
capacity. In Amazon CloudWatch the `S3Outposts` namespace shows detailed metrics for buckets within S3 on Outposts, but these
metrics don't include usage for other AWS services. To manage your S3 on Outposts capacity
that is consumed by other AWS services, see the information in the following table.

AWS serviceDescriptionLearn moreAmazon S3All direct S3 on Outposts usage has a matching account and bucket CloudWatch
metric.[See metrics](../userguide/s3outpostscapacity.md#S3OutpostsCloudWatchMetrics)Amazon Elastic Block Store (Amazon EBS)For Amazon EBS on Outposts, you can choose an AWS Outpost as your snapshot destination and store locally in your S3 on Outpost.[Learn more](../../../ec2/latest/userguide/snapshots-outposts.md)Amazon Relational Database Service (Amazon RDS)You can use Amazon RDS local backups to store your RDS backups locally on your
Outpost.[Learn more](../../../amazonrds/latest/userguide/rds-on-outposts.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sharing S3 on Outposts

Monitoring S3 on Outposts

All content copied from https://docs.aws.amazon.com/.
