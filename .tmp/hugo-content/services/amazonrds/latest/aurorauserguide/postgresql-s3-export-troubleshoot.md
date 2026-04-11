# Troubleshooting access to Amazon S3

If you encounter connection problems when attempting to export data to Amazon S3, first confirm that the outbound access rules for
the VPC security group associated with your DB instance permit network connectivity. Specifically, the security group must have a rule
that allows the DB instance to send TCP traffic to port 443 and to any IPv4 addresses (0.0.0.0/0). For more information, see [Provide access to the DB cluster in the VPC by creating a security group](chap-settingup-aurora.md#CHAP_SettingUp_Aurora.SecurityGroup).

See also the following for recommendations:

- [Troubleshooting Amazon Aurora identity and access](security-iam-troubleshoot.md)

- [Troubleshooting Amazon S3](../../../s3/latest/userguide/troubleshooting.md) in the _Amazon Simple Storage Service User Guide_

- [Troubleshooting Amazon S3 and IAM](../../../iam/latest/userguide/troubleshoot-iam-s3.md) in the _IAM User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Function reference

Invoking a Lambda function from Aurora PostgreSQL

All content copied from https://docs.aws.amazon.com/.
