# Amazon EC2 macOS AMIs release notes

The following information provides details about the packages included by default in the EC2 macOS AMIs and summarizes the changes for each EC2 macOS AMI release.

For information about how to subscribe to macOS AMI notifications, see [Subscribe to macOS AMI notifications](macos-subscribe-notifications.md).

Mac instances can run one of the following operating systems:

- macOS Mojave (version 10.14) (x86 Mac instances only)

- macOS Catalina (version 10.15) (x86 Mac instances only)

- macOS Big Sur (version 11) (x86 and M1 Mac instances)

- macOS Monterey (version 12) (x86 and M1 Mac instances)

- macOS Ventura (version 13) (all Mac instances, M2 and M2 Pro Mac
instances support macOS Ventura version 13.2 or later)

- macOS Sonoma (version 14) (all Mac instances)

- macOS Sequoia (version 15) (all Mac instances)

###### Note

M4 and M4 Pro Mac instances support macOS Sequoia version 15.6 or later.

## Approve Local Network Privacy policies for macOS Sequoia

macOS Sequoia (version 15) has a new Local Network Privacy feature that impacts users of local IP-based services, including
Amazon EC2 Instance Metadata Service (IMDS).

###### Important

To make sure that you have uninterrupted access to local IP-based services, use the following steps to approve the Local Network Privacy policies.

###### To approve Local Network Privacy policies

1. [Connect to your instance's graphical user interface (GUI)](connect-to-mac-instance.md#mac-instance-vnc).

2. Follow the prompts on the screen to approve the Local Network Privacy policies.

3. After you have approved the policies, create an AMI of your EC2 Mac instance. For more information,
    see [Create an Amazon EBS-backed AMI](creating-an-ami-ebs.md).

Any EC2 Mac instances that are launched from the newly created AMI will retain the Local Network Privacy permissions.

## Default packages included in Amazon EC2 macOS AMIs

The following table describes packages that are included by default in the EC2 macOS AMIs.

PackagesRelease notes

EC2 macOS Init

[https://github.com/aws/ec2-macos-init/tags](https://github.com/aws/ec2-macos-init/tags)

EC2 macOS Utils

[https://github.com/aws/ec2-macos-utils/tags](https://github.com/aws/ec2-macos-utils/tags)

Amazon SSM Agent

[https://github.com/aws/amazon-ssm-agent/releases](https://github.com/aws/amazon-ssm-agent/releases)

AWS Command Line Interface (AWS CLI) version 2

[https://raw.githubusercontent.com/aws/aws-cli/v2/CHANGELOG.rst](https://raw.githubusercontent.com/aws/aws-cli/v2/CHANGELOG.rst)

Command Line Tools for Xcode

[https://developer.apple.com/documentation/xcode-release-notes](https://developer.apple.com/documentation/xcode-release-notes)

Homebrew

[https://github.com/Homebrew/brew/releases](https://github.com/Homebrew/brew/releases)

EC2 Instance Connect

[https://github.com/aws/aws-ec2-instance-connect-config/releases](https://github.com/aws/aws-ec2-instance-connect-config/releases)

Safari

[https://developer.apple.com/documentation/safari-release-notes](https://developer.apple.com/documentation/safari-release-notes)

## Amazon EC2 macOS AMI updates

The following table describes changes included in the EC2 macOS AMI releases.
Note that some changes apply to all EC2 macOS AMIs, whereas
others apply to only a subset of these AMIs.

### EC2 macOS AMI updates

ReleaseChanges2026.03.17

###### Release of macOS Tahoe 26.3.1

- [What's new in the updates for macOS Tahoe 26](https://support.apple.com/en-us/122868)

- Updated Command Line Tools to 26.3

- Updated `awscli` to 2.34.10

- Updated Homebrew to 5.1.0

2026.03.03

###### All AMIs

- Updated `awscli` to 2.33.31

- Updated Homebrew to 5.0.15

###### Release of macOS Sonoma 14.8.4

- [Security content of macOS Sonoma 14.8.4](https://support.apple.com/en-us/126350)

- Updated Safari to 26.3

- [Security content of Safari 26.3](https://support.apple.com/en-us/126354)

###### Release of macOS Sequoia 15.7.4

- [Security content of macOS Sequoia 15.7.4](https://support.apple.com/en-us/126349)

- Updated Safari to 26.3

- [Security content of Safari 26.3](https://support.apple.com/en-us/126354)

###### Release of macOS Tahoe 26.3

- [Security content of macOS Tahoe 26.3](https://support.apple.com/en-us/126348)

2025.12.26

###### All AMIs

- Updated `awscli` to 2.32.19

- Updated `amazon-ssm-agent` to 3.3.3270.0

- Updated Homebrew to 5.0.6

###### Release of macOS Sonoma 14.8.3

- [Security content of macOS Sonoma 14.8.3](https://support.apple.com/en-us/125888)

- Updated Safari to 26.2

- [Security content of Safari 26.2](https://support.apple.com/en-us/125892)

###### Release of macOS Sequoia 15.7.3

- [Security content of macOS Sequoia 15.7.3](https://support.apple.com/en-us/125887)

- Updated Safari to 26.2

- [Security content of Safari 26.2](https://support.apple.com/en-us/125892)

- Updated Command Line Tools to 26.2

###### Release of macOS Tahoe 26.2

- [Security content of macOS Tahoe 26.2](https://support.apple.com/en-us/125886)

- Updated Command Line Tools to 26.2

2025.12.17

###### All AMIs

- Updated `awscli` to 2.32.16

- Updated Homebrew to 5.0.5

###### Release of macOS Sonoma 14.8.2

- [Security content of macOS Sonoma 14.8.2](https://support.apple.com/en-us/125636)

- Updated Safari to 26.1

- [Security content of Safari 26.1](https://support.apple.com/en-us/125640)

###### Release of macOS Sequoia 15.7.2

- [Security content of macOS Sequoia 15.7.2](https://support.apple.com/en-us/125635)

- Updated Safari to 26.1

- [Security content of Safari 26.1](https://support.apple.com/en-us/125640)

- Updated Command Line Tools to 26.1

###### Release of macOS Tahoe 26.1

- [Security content of macOS Tahoe 26.1](https://support.apple.com/en-us/125634)

- Updated Command Line Tools to 26.1

2025.11.18

###### All AMIs

- Updated `awscli` to 2.31.35

- Updated `ec2-macos-init` to 1.5.13

- Updated `ec2-macos-utils` to 1.0.7

- Updated Homebrew to 5.0.1

###### Release of macOS Sonoma 14.8.1

- [Security content of macOS Sonoma 14.8.1](https://support.apple.com/en-us/125330)

- Updated Safari to 26.0.1

- [Security content of Safari 26.0.1](https://support.apple.com/en-us/125113)

###### Release of macOS Sequoia 15.7.1

- [Security content of macOS Sequoia 15.7.1](https://support.apple.com/en-us/125329)

- Updated Safari to 26.0.1

- [Security content of Safari 18.6](https://support.apple.com/en-us/125113)

- Updated Command Line Tools to 26.0

###### Release of macOS Tahoe 26.0.1

- [Security content of macOS Tahoe 26.0.1](https://support.apple.com/en-us/125328)

- Updated Command Line Tools to 26.0

2025.09.04

###### All AMIs

- Updated `awscli` to 2.28.19

- Updated `ec2-instance-connect` to 2.0.0-5

- Updated Homebrew to 4.6.7

###### Release of macOS Ventura 13.7.8

- [Security content of macOS Ventura\
13.7.8](https://support.apple.com/en-us/124929)

###### Release of macOS Sonoma 14.7.8

- [Security content of macOS Sonoma 14.7.8](https://support.apple.com/en-us/124928)

###### Release of macOS Sequoia 15.6.1

- [Security content of macOS Sequoia 15.6.1](https://support.apple.com/en-us/124927)

2025.08.05

###### All AMIs

- Updated `awscli` to 2.28.0

- Updated Homebrew to 4.5.13

###### Release of macOS Ventura 13.7.7

- [Security content of macOS Ventura 13.7.7](https://support.apple.com/en-us/124151)

- Updated Safari to 18.6

- [Security content of Safari 18.6](https://support.apple.com/en-us/124152)

###### Release of macOS Sonoma 14.7.7

- [Security content of macOS Sonoma 14.7.7](https://support.apple.com/en-us/124150)

- Updated Safari to 18.6

- [Security content of Safari 18.6](https://support.apple.com/en-us/124152)

###### Release of macOS Sequoia 15.6

- [Security content of macOS Sequoia 15.6](https://support.apple.com/en-us/124149)

- Updated AWS ENA Ethernet to 2.0.0

2025.06.27

###### All AMIs

- Updated `awscli` to 2.27.40

- Updated Homebrew to 4.5.7

- Migrated x86\_64 images to AWS ENA Ethernet dext

###### Release of macOS Ventura 13.7.6

- [Security content of macOS Ventura 13.7.6](https://support.apple.com/en-us/122718)

- Updated Safari to 18.5

- [Security content of Safari 18.5](https://support.apple.com/en-us/122719)

###### Release of macOS Sonoma 14.7.6

- [Security content of macOS Sonoma 14.7.6](https://support.apple.com/en-us/122717)

- Updated Safari to 18.5

- [Security content of Safari 18.5](https://support.apple.com/en-us/122719)

###### Release of macOS Sequoia 15.5

- [Security content of macOS Sequoia 15.5](https://support.apple.com/en-us/122716)

- Updated Command Line Tools to 16.4

- Updated AWS ENA Ethernet to 1.0.9

2025.05.21

###### All AMIs

- Updated `awscli` to 2.27.7

- Updated `ec2-macos-init` to
1.5.11

- Updated `ec2-macos-utils` to
1.0.5

- Updated Homebrew to 4.5.1

###### Release of macOS Ventura 13.7.5

- [Security content of macOS Ventura\
13.7.5](https://support.apple.com/en-us/122375)

- Updated Safari to 18.4

- [Security content of Safari 18.4](https://support.apple.com/en-us/122379)

###### Release of macOS Sonoma 14.7.5

- [Security content of macOS Sonoma\
14.7.5](https://support.apple.com/en-us/122374)

- Updated Safari to 18.4

- [Security content of Safari 18.4](https://support.apple.com/en-us/122379)

###### Release of macOS Sequoia 15.4.1

- [Security content of macOS Sequoia\
15.4.1](https://support.apple.com/en-us/122400)

- Updated Command Line Tools to 16.3

2025.05.05

###### All AMIs

- Updated `awscli` to 2.27.1

- Updated `ec2-macos-init` to 1.5.11

- Updated `ec2-macos-system-monitor` to 1.3.1

- Updated `ec2-macos-utils` to 1.0.5

- Updated Homebrew to 4.4.32

###### Release of macOS Ventura 13.7.5

- [Security content of macOS Ventura 13.7.5](https://support.apple.com/en-us/122375)

- Updated Safari to 18.4

- [Security content of Safari 18.4](https://support.apple.com/en-us/122379)

###### Release of macOS Sonoma 14.7.5

- [Security content of macOS Sonoma 14.7.5](https://support.apple.com/en-us/122374)

- Updated Safari to 18.4

- [Security content of Safari 18.4](https://support.apple.com/en-us/122379)

###### Release of macOS Sequoia 15.4.1

- [Security content of macOS Sequoia 15.4.1](https://support.apple.com/en-us/122400)

2025.03.18

###### All AMIs

- Updated `awscli` to 2.24.2

- Updated Homebrew to 4.4.20

###### Release of macOS Sequoia 15.3.1

- [Security content of macOS Sequoia 15.3.1](https://support.apple.com/en-us/120283)

###### Release of macOS Sonoma 14.7.4

- [Security content of macOS Sonoma 14.7.4](https://support.apple.com/en-us/109035)

- Updated Safari to 18.3

###### Release of macOS Ventura 13.7.4

- [Security content of macOS Ventura 13.7.4](https://support.apple.com/en-us/106337)

- Updated Safari to 18.3

2025.01.24

###### All AMIs

- Updated `awscli` to 2.22.33

- Updated Homebrew to 4.4.15

###### Release of macOS Sequoia 15.2

- [Security content of macOS Sequoia 15.2](https://support.apple.com/en-us/121839)

- Updated Command Line Tools to 16.2

###### Release of macOS Sonoma 14.7.2

- [Security content of macOS Sonoma 14.7.2](https://support.apple.com/en-us/121840)

- Updated Safari to 18.2

- Updated Command Line Tools to 16.2

###### Release of macOS Ventura 13.7.2

- [Security content of macOS Ventura 13.7.2](https://support.apple.com/en-us/121842)

- Updated Safari to 18.2

2024.12.20

###### All AMIs

- Updated Homebrew to 4.4.8

- Updated `aws-cli` to 2.22.5

- Updated `amazon-ssm-agent` to 3.3.987.0

###### Release of macOS Sequoia 15.1.1

- [Security content of macOS Sequoia 15.1.1](https://support.apple.com/en-us/121753)

###### Release of macOS Sonoma 14.7.1

- [Security content of macOS Sonoma 14.7.1](https://support.apple.com/en-us/121570)

- Updated Safari to 18.1.1

###### Release of macOS Ventura 13.7.1

- [Security content of macOS Ventura 13.7.1](https://support.apple.com/en-us/121568)

- Updated Safari to 18.1.1

2024.10.28

###### All AMIs

- Updated Homebrew to 4.4.2

- Updated `aws-cli` to 2.18.13

- Updated `amazon-ssm-agent` to 3.3.987.0

- Updated `ec2-macos-init` to 1.5.10

- Updated `ec2-macos-utils` to 1.0.4

###### Release of macOS Sequoia 15.0

- [Security content of macOS Sequoia 15](https://support.apple.com/en-us/121238)

###### Release of macOS Sonoma 14.7

- [Security content of macOS Sonoma 14.7](https://support.apple.com/en-us/121247).

- Updated Command Line Tools to 16.0

- Updated Safari to 18.0.1

- [Security content of Safari 18](https://support.apple.com/en-us/121241)

###### Release of macOS Ventura 13.7

- [Security content of macOS Ventura 13.7](https://support.apple.com/en-us/121234)

- Updated Safari to 18.0.1

- [Security content of Safari 18](https://support.apple.com/en-us/121241)

2024.08.20

###### All AMIs

- Updated Homebrew to 4.3.14

- Updated `aws-cli` to 2.17.29

###### Release of macOS Sonoma 14.6.1

- No published CVE entries.

###### Release of macOS Ventura 13.6.9

- No published CVE entries.

- Updated Safari to 17.6

- [Security content of Safari 17.6](https://support.apple.com/en-us/120913)

###### Release of macOS Monterey 12.7.6

- [Security content of macOS Monterey 12.7.6](https://support.apple.com/en-us/120910)

- Updated Safari to 17.6

- [Security content of Safari 17.6](https://support.apple.com/en-us/120913)

2024.06.07

###### All AMIs

- Updated Homebrew to 4.3.1-1

- Updated `aws-cli` to 2.15.56

- Updated `amazon-ssm-agent` to 3.3.380.0-1

###### Release of macOS Sonoma 14.5

- [Security content of macOS Sonoma 14.5](https://support.apple.com/en-us/120903)

###### Release of macOS Ventura 13.6.7

- [Security content of macOS Ventura 13.6.7](https://support.apple.com/en-us/120900)

- Updated Safari to 17.5

- [Security content of Safari 17.5](https://support.apple.com/en-us/120896)

###### Release of macOS Monterey 12.7.5

- [Security content of macOS Monterey 12.7.5](https://support.apple.com/en-us/120896)

- Updated Safari to 17.5

- [Security content of Safari 17.5](https://support.apple.com/en-us/120896)

2024.04.12

###### All AMIs

- Updated Homebrew to 4.2.16-1

- Updated `aws-cli` to 2.15.36

###### Release of macOS Sonoma 14.4.1

- [Security content of macOS Sonoma 14.4.1](https://support.apple.com/en-us/120889)

###### Release of macOS Ventura 13.6.6

- [Security content of macOS Ventura 13.6.6](https://support.apple.com/en-us/120891)

- Updated Safari to 17.4.1

- [Security content of Safari 17.4.1](https://support.apple.com/en-us/120888)

###### For macOS Monterey

- Updated Safari to 17.4.1

- [Security content of Safari 17.4.1](https://support.apple.com/en-us/120888)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Retrieve macOS AMI IDs

EBS optimization
