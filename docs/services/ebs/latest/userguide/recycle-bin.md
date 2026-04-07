# Recover deleted EBS volumes, EBS snapshots, and EBS-backed AMIs with Recycle Bin

Recycle Bin is a data recovery feature that enables you to restore accidentally
deleted EBS volumes, EBS snapshots, and EBS-backed AMIs. When using Recycle Bin, if your resources are deleted, they are
retained in the Recycle Bin for a time period that you specify before being permanently deleted.

You can restore a resource from the Recycle Bin at any time before its retention period expires.
After you restore a resource from the Recycle Bin, the resource is removed from the Recycle Bin and
you can use it in the same way that you use any other resource of that type in your account. If the retention
period expires and the resource is not restored, the resource is permanently deleted from the Recycle Bin
and it is no longer available for recovery.

Using Recycle Bin helps to ensure business continuity by protecting your business-critical
data against accidental deletion.

Recycle Bin is assessed as a service capability of Amazon Elastic Block Store (Amazon EBS). Any [AWS services in Scope by Compliance](https://aws.amazon.com/compliance/services-in-scope)
Program (FedRAMP, HIPAA BAA, SOC, etc) which lists Amazon EBS will also apply to Recycle Bin.

###### Topics

- [Supported resources](#supported-resources)

- [How does it work?](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-concepts.html)

- [Considerations](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-factors.html)

- [Quotas](#recycle-bin-quotas)

- [Related services](#recycle-bin-integrations)

- [Pricing](#recycle-bin-pricing)

- [Control access](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-perms.html)

- [Create retention rule](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-create-rule.html)

- [Update retention rule](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-update-rule.html)

- [Lock retention rule](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-lock.html)

- [Unlock retention rule](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-unlock.html)

- [Tag retention rules](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-tag-resource.html)

- [Delete retention rules](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-delete-rule.html)

- [Recover deleted snapshots](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-working-with-snaps.html)

- [Recover deleted volumes](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-working-with-volumes.html)

- [Recover deleted AMIs](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-working-with-amis.html)

- [Monitor using EventBridge](https://docs.aws.amazon.com/ebs/latest/userguide/rbin-eventbridge.html)

- [Monitor using CloudTrail](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-ct.html)

- [Service endpoints](https://docs.aws.amazon.com/ebs/latest/userguide/rbin-service-endpoints.html)

- [Use interface VPC endpoints](https://docs.aws.amazon.com/ebs/latest/userguide/rbin-vpcendpoints.html)

## Supported resources

Recycle Bin supports the following resource types:

- Amazon EBS volumes

- Amazon EBS snapshots

###### Important

Recycle Bin retention rules also apply to archived snapshots in the archive storage
tier. If you delete an archived snapshot that matches a retention rule, that snapshot is
retained in the Recycle Bin for the period defined in the retention rule. Archived
snapshots are billed at the rate for archived snapshots while they are in the Recycle Bin.

- Amazon EBS-backed Amazon Machine Images (AMIs)

###### Note

Retention rules also apply to disabled AMIs.

## Quotas

The following quotas apply to Recycle Bin.

QuotaDefault quota

Retention rules per Region

250

Tag key and value pairs per retention rule

50

## Related services

Recycle Bin works with the following services:

- **AWS CloudTrail** — Enables you to record events
that occur in Recycle Bin. For more information, see [Monitor Recycle Bin using AWS CloudTrail](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-ct.html).

## Pricing

There are no additional charges for using Recycle Bin and retention rules. For more
information, see [Amazon EBS pricing](https://aws.amazon.com/ebs/pricing).

- **Amazon EBS volumes** — Volumes
in the Recycle Bin are billed at the same rate as regular volumes
in your account.

- **Amazon EBS snapshots** — Snapshots
in the Recycle Bin are billed at the same rate as regular snapshots
in your account.

- **EBS-backed AMIs** — AMIs in the
Recycle Bin do not incur any additional charges.

###### Note

Some resources might still appear in the Recycle Bin console or in the AWS CLI and
API output for a short period after their retention periods have expired and they
have been permanently deleted. You are not billed for these resources. Billing stops
as soon as the retention period expires.

You can use the following AWS generated cost allocation tags for cost tracking and
allocation purposes when using AWS Billing and Cost Management.

- Key: `aws:recycle-bin:resource-in-bin`

- Value: `true`

For more information, see [AWS-generated cost\
allocation tags](../../../awsaccountbilling/latest/aboutv2/aws-tags.md) in the _AWS Billing and Cost Management User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FAQs

How does it work?
