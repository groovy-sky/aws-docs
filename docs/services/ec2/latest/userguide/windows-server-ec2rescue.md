# Troubleshoot impaired Amazon EC2 Windows instance using EC2Rescue

EC2Rescue for Windows Server is an easy-to-use tool that you run on an Amazon EC2 Windows Server
instance to diagnose and troubleshoot possible problems. It is valuable for collecting log
files and troubleshooting issues and also proactively searching for possible areas of
concern. It can even examine Amazon EBS root volumes from other instances and collect relevant
logs for troubleshooting Windows Server instances using that volume. The following are some
common issues that EC2Rescue can address:

- Instance connectivity issues due to firewall, Remote Desktop Protocol (RDP), or
network interface configuration

- Operating system boot issues due to a stop error, boot loop, or corrupted
registry

- Issues that might need advanced log analysis and troubleshooting

EC2Rescue for Windows Server has two different modules:

- A **data collector module** that collects data from all
different sources

- An **analyzer module** that parses the data collected
against a series of predefined rules to identify issues and provide suggestions

The EC2Rescue for Windows Server tool only runs on Amazon EC2 instances running Windows Server 2012
and later. When the tool starts, it checks whether it is running on an Amazon EC2 instance.

###### Note

The `AWSSupport-ExecuteEC2Rescue` AWS Systems Manager Automation runbook uses the
EC2Rescue tool to troubleshoot and, where possible, fix common connectivity issues with
the specified EC2 instance. For more information, and to run this automation, see
[>AWSSupport-ExecuteEC2Rescue](https://docs.aws.amazon.com/systems-manager-automation-runbooks/latest/userguide/automation-awssupport-executeec2rescue.html).

If you are using a Linux instance, see [Troubleshoot impaired Amazon EC2 Linux instance using EC2Rescue](linux-server-ec2rescue.md).

###### Topics

- [Troubleshoot using EC2Rescue GUI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rw-gui.html)

- [Troubleshoot using EC2Rescue CLI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rw-cli.html)

- [Troubleshoot using EC2Rescue and Systems Manager](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rw-ssm.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Develop EC2Rescue modules

Troubleshoot using EC2Rescue GUI
