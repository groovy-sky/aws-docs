# Automate backups with Amazon Data Lifecycle Manager

You can use Amazon Data Lifecycle Manager to automate the creation, retention, and deletion of EBS snapshots and
EBS-backed AMIs. When you automate snapshot and AMI management, it helps you to:

- Protect valuable data by enforcing a regular backup schedule.

- Create standardized AMIs that can be refreshed at regular intervals.

- Retain backups as required by auditors or internal compliance.

- Reduce storage costs by deleting outdated backups.

- Create disaster recovery backup policies that back up data to isolated Regions
or accounts.

When combined with the monitoring features of Amazon EventBridge and AWS CloudTrail, Amazon Data Lifecycle Manager provides
a complete backup solution for Amazon EC2 instances and individual EBS volumes at no additional
cost.

###### Important

- Amazon Data Lifecycle Manager can't manage snapshots or AMIs created by any other means.

- Amazon Data Lifecycle Manager can't automate the creation, retention, and deletion of instance
store-backed AMIs.

Amazon Data Lifecycle Manager is assessed as a service capability of Amazon Elastic Block Store (Amazon EBS). Any [AWS Services in Scope by Compliance Program](https://aws.amazon.com/compliance/services-in-scope) (FedRAMP,
HIPAA BAA, SOC, etc) which lists Amazon EBS will also apply to Amazon Data Lifecycle Manager.

###### Contents

- [Quotas](#dlm-quotas)

- [How it works](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-elements.html)

- [Default vs custom policies](https://docs.aws.amazon.com/ebs/latest/userguide/policy-differences.html)

- [Create default policies](https://docs.aws.amazon.com/ebs/latest/userguide/default-policies.html)

- [Create custom policy for snapshots](https://docs.aws.amazon.com/ebs/latest/userguide/snapshot-ami-policy.html)

- [Create custom policy for AMIs](https://docs.aws.amazon.com/ebs/latest/userguide/ami-policy.html)

- [Automate cross-account snapshot copies](https://docs.aws.amazon.com/ebs/latest/userguide/event-policy.html)

- [Modify policies](https://docs.aws.amazon.com/ebs/latest/userguide/modify.html)

- [Delete policies](https://docs.aws.amazon.com/ebs/latest/userguide/delete.html)

- [Control access](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-prerequisites.html)

- [Monitor policies](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-monitor-lifecycle.html)

- [Service endpoints](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-service-endpoints.html)

- [Interface VPC endpoints](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-vpc-endpoints.html)

- [Troubleshoot](https://docs.aws.amazon.com/ebs/latest/userguide/dlm-troubleshooting.html)

## Quotas

Your AWS account has the following quotas related to Amazon Data Lifecycle Manager:

DescriptionQuotaCustom lifecycle policies per Region100Default policies for EBS snapshots per Region1Default policies for EBS-backed AMIs per Region1Tags per resource45

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Benchmark EBS volumes

How it works
