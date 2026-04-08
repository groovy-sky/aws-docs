# Troubleshooting access to Amazon S3

If you encounter connection problems when attempting to export data to Amazon S3, first confirm that the outbound access rules for
the VPC security group associated with your DB instance permit network connectivity. Specifically, the security group must have a rule
that allows the DB instance to send TCP traffic to port 443 and to any IPv4 addresses (0.0.0.0/0). For more information, see [Provide access to your DB instance in your VPC by creating a security group](chap-settingup.md#CHAP_SettingUp.SecurityGroup).

See also the following for recommendations:

- [Troubleshooting Amazon RDS identity and access](security-iam-troubleshoot.md)

- [Troubleshooting Amazon S3](../../../s3/latest/userguide/troubleshooting.md) in the _Amazon Simple Storage Service User Guide_

- [Troubleshooting Amazon S3 and IAM](../../../iam/latest/userguide/troubleshoot-iam-s3.md) in the _IAM User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Function reference

Invoking a Lambda function from RDS for PostgreSQL

All content copied from https://docs.aws.amazon.com/.
