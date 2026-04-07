# Update management for Amazon EC2 instances

We recommend that you regularly patch, update, and secure the operating system and
applications on your EC2 instances. You can use [AWS Systems Manager Patch\
Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager.html) to automate the process of installing security-related updates for both the
operating system and applications.

For EC2 instances in an Auto Scaling group, you can use the [AWS-PatchAsgInstance](https://docs.aws.amazon.com/systems-manager-automation-runbooks/latest/userguide/automation-aws-patchasginstance.html) runbook to help avoid instances that are
undergoing patching from being replaced. Alternatively, you can use any automatic update
services or recommended processes for installing updates that are provided by the application
vendor.

###### Resources

- AL2023 – [Updating AL2023](https://docs.aws.amazon.com/linux/al2023/ug/updating.html) in the
_Amazon Linux 2023 User Guide_

- AL2 – [Manage software on your Amazon Linux 2 instance](https://docs.aws.amazon.com/linux/al2/ug/managing-software.html)
in the _Amazon Linux 2 User Guide_

- Windows instances – [Update management](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-windows-security-best-practices.html#ec2-windows-update-management)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Attach a role to an instance

Best practices for Windows instances
