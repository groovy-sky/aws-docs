# Troubleshoot impaired Amazon EC2 Linux instance using EC2Rescue

EC2Rescue for Linux is an easy-to-use, open-source tool that can be run on an Amazon EC2 Linux
instance to diagnose, troubleshoot, and remediate common issues using its library of over 100
_modules_. Modules are YAML files that contain either a BASH or a Python
script and the necessary metadata.

Some generalized use cases for EC2Rescue for Linux instances include:

- Gathering syslog and package manager logs

- Collecting resource utilization data

- Diagnosing and remediating known problematic kernel parameters and common OpenSSH issues

###### Note

The `AWSSupport-TroubleshootSSH` AWS Systems Manager Automation runbook installs
EC2Rescue for Linux and then uses the tool to check or attempt to fix common issues
that prevent an SSH connection to a Linux instance. For more information, see
[AWSSupport-TroubleshootSSH](https://docs.aws.amazon.com/systems-manager-automation-runbooks/latest/userguide/automation-awssupport-troubleshootssh.html).

If you are using a Windows instance, see [Troubleshoot impaired Amazon EC2 Windows instance using EC2Rescue](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Windows-Server-EC2Rescue.html).

###### Topics

- [Install EC2Rescue](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rl_install.html)

- [Run EC2Rescue commands](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rl_working.html)

- [Develop EC2Rescue modules](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2rl_moduledev.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshoot Sysprep issues

Install EC2Rescue
