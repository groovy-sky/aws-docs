# Windows launch agents on Amazon EC2 Windows instances

Each AWS Windows AMI includes a Windows launch agent that's pre-configured with default
settings. Launch agents perform tasks during instance startup and run if an instance is
stopped and later started, or restarted. For information about a specific agent, see the
detail pages in the following list.

For more information about AWS Windows AMIs, see the [AWS \
Windows AMI reference](https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/windows-amis.html).

- [Use the EC2Launch v2 agent to perform tasks during EC2 Windows instance launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2launch-v2.html)

- [Use the EC2Launch v1 agent to perform tasks during EC2 Windows instance launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2launch.html)

- [Use the EC2Config service to perform tasks during EC2 legacy Windows operating system instance launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2config-service.html)

###### Content

- [Compare Amazon EC2 launch agents](#ec2launch-agent-compare)

- [Configure DNS Suffix for EC2 Windows launch agents](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-agents-set-dns.html)

- [Subscribe to EC2 Windows launch agent notifications](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-agents-subscribe-notifications.html)

- [Windows Service administration for EC2Launch v2 and EC2Config agents](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-agents-service-admin.html)

## Compare Amazon EC2 launch agents

The following table shows the major functional differences between EC2Config,
EC2Launch v1, and EC2Launch v2.

FeatureEC2ConfigEC2Launch v1EC2Launch v2Run asWindows ServicePowerShell ScriptsWindows ServiceSupportsLegacy OS only

Windows Server versions:

- 2016

- 2019 (LTSC and SAC)

Windows Server versions:

- 2016

- 2019 (LTSC and SAC)

- 2022

- 2025

Configuration fileXMLJSONJSON/YAMLSet Administrator usernameNoNoYesCompressed user dataNoNoYesLocal user data baked on AMINoNoYes, configurableTask configuration in user dataNoNoYesConfigurable wallpaperNoNoYesCustomize task run orderNoNoYesConfigurable tasks15920 at launchSupports Windows Event ViewerYesNoYesNumber of Event Viewer event types2030

###### Note

EC2Config documentation is provided for historical reference only. The operating system
versions it runs on are no longer supported by Microsoft. We strongly recommend that
you upgrade to the latest launch service.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configure Windows instances

Configure DNS Suffix
