# Using Amazon S3 Storage Lens with AWS Organizations

Amazon S3 Storage Lens is a cloud-storage analytics feature that you can use to gain organization-wide
visibility into object-storage usage and activity. You can use S3 Storage Lens metrics to generate
summary insights, such as finding out how much storage you have across your entire organization
or which are the fastest-growing buckets and prefixes. You can also use Amazon S3 Storage Lens to collect
storage metrics and usage data for all AWS accounts that are part of your AWS Organizations hierarchy.
To do this, you must be using AWS Organizations, and you must enable S3 Storage Lens trusted access by using
your AWS Organizations management account.

After enabling trusted access, add delegated administrator access to accounts in your
organization. The delegated administrator accounts are used to create S3 Storage Lens configurations
and dashboards that collect organization-wide storage metrics and user data. For more
information about enabling trusted access, see [Amazon S3 Storage Lens\
and AWS Organizations](../../../organizations/latest/userguide/services-that-can-integrate-s3lens.md) in the _AWS Organizations User Guide_.

###### Topics

- [Enabling trusted access for S3 Storage Lens](storage-lens-with-organizations-enabling-trusted-access.md)

- [Disabling trusted access for S3 Storage Lens](storage-lens-with-organizations-disabling-trusted-access.md)

- [Registering a delegated administrator for S3 Storage Lens](storage-lens-with-organizations-registering-delegated-admins.md)

- [Deregistering a delegated administrator for S3 Storage Lens](storage-lens-with-organizations-deregistering-delegated-admins.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using AI assistants

Enabling trusted access

All content copied from https://docs.aws.amazon.com/.
