# WorkSpaces Applications Base Image and Managed Image Update Release Notes

Amazon WorkSpaces Applications provides base images to help you create images that include your own
applications. Base images are Amazon Machine Images (AMIs) that contain software
configurations specific to the operating system. For WorkSpaces Applications, each base image includes the
WorkSpaces Applications agent and the latest version of one of the following operating systems:

###### Important

Operating system versions that are no longer supported by the vendor are not
guaranteed to work and are not supported by AWS Support.

- Windows Server 2025 Base — Available on the following image types: Base,
Graphics G4dn, Graphics G5, and Graphics G6

- Windows Server 2022 Base — Available on the following image types: Base,
Graphics G4dn, Graphics G5, and Graphics G6

- Windows Server 2019 Base — Available on the following image types: Base,
Graphics G4dn and Graphics G5

- Windows Server 2016 Base — Available on the following image types: Base,
Graphics G4dn and Graphics G5

- Amazon Linux 2 – Available on the following image types: Base, Graphics G4dn,
and Graphics G5

- Red Hat Enterprise Linux 8 – Available on the following image types: Base,
Graphics G4dn, Graphics G5, and Graphics G6

- Rocky Linux 8 – Available on the following image types: Base, Graphics G4dn,
Graphics G5, and Graphics G6

After you create your own image that includes your own applications, you are responsible
for installing and maintaining the updates for the operating system, your applications, and
their dependencies. WorkSpaces Applications provides an automated way to update your image using managed
WorkSpaces Applications image updates. With managed image updates, you select the image that you want to
update. WorkSpaces Applications creates an image builder in the same AWS account and Region to install the
updates and create the new image. After the new image is created, you can test it on a
pre-production fleet before updating your production fleets or sharing the image with other
AWS accounts. For more information, see "Keep Your WorkSpaces Applications Image Up-to-Date" in [Administer Your Amazon WorkSpaces Applications Images](administer-images.md).

For information about the latest WorkSpaces Applications agent, see [WorkSpaces Applications Agent Release Notes](agent-software-versions.md).

The following table lists the latest released images.

###### Note

Public base images for Graphics Pro instances are no longer available from AWS
after 10/31/2025 due to End of Life of hardware supporting Graphics Pro instance types.

Public base images for Graphics Design instances are no longer available from
AWS after 12/31/2025 due to End of Life of hardware supporting Graphics Design
instance types.

Image typeImage nameBase

- AppStream-WinServer2025-12-18-2025

- AppStream-WinServer2022-11-10-2025

- AppStream-WinServer2019-11-10-2025

- AppStream-WinServer2016-11-10-2025

- AppStream-AmazonLinux2-02-11-2025

- AppStream-RHEL8-02-18-2026

- AppStream-RockyLinux8-02-18-2026

Graphics G4dn

- AppStream-Graphics-G4dn-WinServer2025-12-18-2025

- AppStream-Graphics-G4dn-WinServer2022-11-10-2025

- AppStream-Graphics-G4dn-WinServer2019-11-10-2025

- AppStream-Graphics-G4dn-WinServer2016-11-10-2025

- AppStream-Graphics-G4dn-AmazonLinux2-02-11-2025

- AppStream-Graphics-G4dn-RHEL8-02-18-2026

- AppStream-Graphics-G4dn-RockyLinux8-02-18-2026

Graphics G5

- AppStream-Graphics-G5-WinServer2025-12-18-2025

- AppStream-Graphics-G5-WinServer2022-11-10-2025

- AppStream-Graphics-G5-WinServer2019-11-10-2025

- AppStream-Graphics-G5-WinServer2016-11-10-2025

- AppStream-Graphics-G5-AmazonLinux2-02-11-2025

- AppStream-Graphics-G5-RHEL8-02-18-2026

- AppStream-Graphics-G5-RockyLinux8-02-18-2026

Graphics G6

- AppStream-Graphics-G6-WinServer2025-12-18-2025

- AppStream-Graphics-G6-WinServer2022-11-10-2025

- AppStream-Graphics-G6-RHEL8-02-18-2026

- AppStream-Graphics-G6-RockyLinux8-02-18-2026

Sample apps

Amazon-AppStream2-Sample-Image-06-17-2024

For information about how to access this base image, see [Get Started with Amazon WorkSpaces Applications: Set Up With Sample Applications](getting-started.md).

The following table lists the software components for the latest released base images and
the components that are available if you update your image using managed image updates. If
the version is marked “latest”, the current stable software component available from the
vendor will be installed. If the version is marked “not included”, managed image updates is
not managing the component and the version will not be changed when you update your
image.

The following table lists the software components for the latest released Windows, Amazon
Linux, Rocky Linux, and Red Hat Enterprise Linux base images and Managed image
updates.

Windows

Software componentLatest base images (December 18, 2025)Managed image updates (February 18,
2026)Amazon AWS (AvsCamera) Driver1.0.23.01.0.23.0Amazon CloudWatch Agent1.4.379171.300063SSM Agent3.3.3050.03.3.3598.0NICE DCV Virtual Display2024.0-191432025.0-20850AMD Driver for Graphics Design instances24.20.13028.7002 24.20.13028.7002AppStream 2.0 AgentLATEST (02-09-2026)--AWS Command Line Interface (AWS CLI)

1.40.24 (Windows Server 2016/2019)

2.31.30.0 (Windows Server 2022/2025)

Not includedFirefox144 (Windows Server 2016/2019)Not includedMicrosoft Message Queuing (MSMQ)Installed with Windows ServerInstalled with Windows ServerNVIDIA Graphics Driver for G4dn, G5 and G6 instances

581.42 (Windows Server 2022/2025)

539.19 (Windows Server 2019)

512.78 (Windows Server 2016)

581.42 (Windows Server 2022/2025)

539.19 (Windows Server 2019)

512.78 (Windows Server 2016)

Process monitor4.01[Latest](https://docs.microsoft.com/en-us/sysinternals/downloads/procmon)Quality Windows Audio/Video Experience (qWAVE)Installed with Windows ServerInstalled with Windows ServerVisual C++ redistributable packagesMicrosoft Visual C++ 2013 Redistributable (x64) -
12.0.40664.0

Microsoft Visual C++ 2015-2022
Redistributable (x64) - 14.42.34438

Microsoft Visual C++ 2013 Redistributable (x64) -
12.0.30501

Microsoft Visual C++ 2015-2022
Redistributable (x64) - 14.44.35211

Windows Server updatesBase image updates as of November 2025[Latest](https://www.catalog.update.microsoft.com/home.aspx)WinSCard Filter Driver1.0.19.01.0.19.0Paravirtual (PV) driver8.6.08.6.0ENA driver2.11.02.11.0AWS NVMe driver1.7.01.7.0

Amazon Linux

Software componentLatest base images (October 22, 2024)Managed AppStream 2.0 image updates (October 22,
2024)AWS Command Line Interface (AWS CLI)1.18.147-1Not includedAmazon CloudWatch Agent1.300041.1-11.300041.1-1SSM Agent3.3.1611.0-13.3.1611.0-1NICE DCV Server AppStream2024.0.18380-12024.0.18380-1Cloud-init19.3-46Not includedAL2 Kernel4.14.355-275.570Not includedNVIDIA Graphics Driver for G4dn and G5 instances550.127.05550.127.05Cuda Version12.4Not included

Rocky Linux

Software componentLatest base images (February 18, 2026)Managed image updates (February 18, 2026)AWS Command Line Interface (AWS CLI)2.33.242.33.24Amazon CloudWatch Agent1.300064.0b1337-11.300064.0b1337-1SSM Agent3.3.3598.0-13.3.3598.0-1NICE DCV Server AppStream2024.0.17598-182024.0.17598-18Cloud-init23.4-78\_10.11.0.223.4-78\_10.11.0.2Kernel4.18.0-553.104.14.18.0-553.104.1NVIDIA Graphics Driver for G4dn, G5 and G6 instances580.95.05580.95.05Cuda Version13.013.0

Red Hat Enterprise Linux

Software componentLatest base images (February 18, 2026)Managed image updates (February 18, 2026)AWS Command Line Interface (AWS CLI)2.33.242.33.24Amazon CloudWatch Agent1.300064.0b1337-11.300064.0b1337-1SSM Agent3.3.3598.0-13.3.3598.0-1NICE DCV Server AppStream2024.0.17598-182024.0.17598-18Cloud-init23.4-78\_10.1123.4-78\_10.11Kernel4.18.0-553.105.14.18.0-553.105.1NVIDIA Graphics Driver for G4dn, G5 and G6 instances580.95.05580.95.05Cuda Version13.013.0

###### Important

The following public images are deprecated and therefore no longer available from
AWS:

- 2016/2019/2022 Windows images released before May 30, 2025

- Amazon Linux 2 images released before February 2024

- Images for the Graphics Desktop, Graphics Design, and Graphics Pro instance families

If you want to use an image for a multi-session fleet, the image must meet the
following conditions:

- The image must be created from a base image released on or after June 12,
2023\. Or, the image must be updated by using managed WorkSpaces Applications image updates
released on or after September 6, 2023. For more information, see [Update an Image by Using Managed WorkSpaces Applications Image Updates](keep-image-updated-managed-image-updates.md).

- The WorkSpaces Applications agent release version must be 09-06-2023 or later. For more
information, see [Manage WorkSpaces Applications Agent Versions](base-images-agent.md).

- If you have updated your image using Managed WorkSpaces Applications Image updates, then the
WorkSpaces Applications agent release version is not applicable. Your image must be updated using
a Managed Image Update released on or after September 6, 2023. For more
information, see [Update an Image by Using Managed WorkSpaces Applications Image Updates](keep-image-updated-managed-image-updates.md).

- Multi-session fleets are supported only for Microsoft Server 2019, 2022, and 2025.

The following table describes all released base images.

ReleasePlatformImage Changes02/18/2026Red Hat Enterprise Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- 4K resolution support for non-accelerated instance types

02/18/2026Rocky Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- 4K resolution support for non-accelerated instance types

12/18/2025Windows

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- Includes support for Windows Server 2025

11/10/2025Windows

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- General bug fixes and improvements

- Updated to latest NVIDIA drivers

- Includes new CloudWatch Agent 1.4.37917

- Include new SSM Agent 3.3.3050.0

11/10/2025Red Hat Enterprise Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- General bug fixes and improvements

- Updated to latest NVIDIA drivers

- CW Agent updated to 1.300061.0b1289-1

- SSM Agent updated to 3.3.3270.0-1

- NICE DCV Server AppStream updated to 2024.0.17598-18

11/10/2025Rocky Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- General bug fixes and improvements

- Updated to latest NVIDIA drivers

- CW Agent updated to 1.300061.0b1289-1

- SSM Agent updated to 3.3.3270.0-1

- NICE DCV Server AppStream updated to 2024.0.17598-18

09-05-2025Red Hat Enterprise Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- Updated to latest NVIDIA drivers

- Includes new CloudWatch Agent 1.300057.1b1167-1

- Include new SSM Agent 3.3.2958.0-1

- General bug fixes and improvements

09-05-2025Rocky Linux

- Base

- Graphics G4dn

- Graphics G5

- Graphics G6

- Updated to latest NVIDIA drivers

- Includes new CloudWatch Agent 1.300057.1b1167-1

- Include new SSM Agent 3.3.2958.0-1

- General bug fixes and improvements

05-30-2025Windows

- Base

- Graphics G4dn

- Graphics G5

- Updated to latest NVIDIA drivers

- Amazon DCV updated to version 2024.0-19143

- Includes new CloudWatch Agent 1.4.37911

- Include new SSM Agent 3.3.2299.0

05-30-2025Red Hat Enterprise Linux

- Base

- Graphics G4dn

- Graphics G5

- Updated to latest NVIDIA drivers

- Amazon DCV server updated to version 2024.0.17598-1

- Includes new CloudWatch Agent 1.300055.1-1

- Include new SSM Agent 3.3.2471.0-1

05-30-2025Rocky Linux

- Base

- Graphics G4dn

- Graphics G5

- Updated to latest NVIDIA drivers

- Amazon DCV server updated to version 2024.0.17598-1

- Includes new CloudWatch Agent 1.300055.1-1

- Include new SSM Agent 3.3.2471.0-1

02-11-2025Amazon Linux 2

- Base

- Graphics G4dn

- Graphics G5

- Includes latest CloudWatch Agent 1.300041.1-1

- Includes new SSM Agent 3.3.1611.0-1

02-11-2025Windows

- Base

- Graphics G4dn

- Graphics G5

- Includes new ENA and NVMe drivers

- Includes updated PV driver

- Updating drivers and software with new versions
available

12-19-2024Rocky Linux

- Base

- Graphics G4dn

- Graphics G5

- Support for Rocky Linux 8

10-22-2024Amazon Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes latest NVIDIA drivers for Linux.

- Updated Linux to version 2.0.20241014. For more information,
see [Amazon Linux 2 version 2.0.20241014 release\
notes](../../../al2/latest/relnotes/relnotes-20241014.md).

10-22-2024Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes latest NVIDIA drivers for Windows

- Includes Microsoft Visual C++ 2015-2022 Redistributable (x64)
\- 14.40.33816

- Includes Microsoft security updates up to October 9,
2024

09-12-2024Amazon Linux 2

- Base

- Graphics G4dn

- Graphics G5

- Bug fixes related to touch functionality on Android, iOS, and
Surface Pro

- Includes latest NVIDIA Graphics Driver (550.90.07) for G4dn
and G5 instances for Amazon Linux 2

- Updated Linux to version 2 2.0.20240709.1. For more
information, see [2.0.20240709.1](../../../al2/latest/relnotes/relnotes-20240709.md).

07-30-2024Red Hat Enterprise Linux

- Base

- Graphics G4dn

- Graphics G5

- Includes support for Red Hat Enterprise Linux 8

06-17-2024Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes Microsoft security updates up to June 13, 2024

- Includes CloudWatch Agent 1.4.37896

- Includes SSM Agent 3.3.484.0

- Includes AWS Command Line Interface (AWS CLI) (WinServer
2016/2019) 1.33.9

- Includes AWS Command Line Interface (AWS CLI) (WinServer 2022)
2.16.9.0

05-08-2024Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes Microsoft security updates up to May 2024

- Includes latest NVIDIA Graphics Driver (552.08) for Graphics
Pro and G4dn instances for Windows Server 2016 and Windows
Server 2019

- Includes CloudWatch Agent 1.4.37891

- Includes SSM Agent 3.3.131.0-1

- Includes AWS Command Line Interface (AWS CLI) 1.32.89

- Includes AWSVirtualSmartCardReader 1.0.0.59

05-08-2024Linux

- Base

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Updated Linux to version 2.0.20240412.0. For more information,
see [Amazon Linux 2.0.20240412.0 release notes](../../../al2/latest/relnotes/relnotes-20240419.md).

03-24-2024Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes Microsoft security updates up to March 2024

- Includes latest NVIDIA Graphics Driver (551.61) for Graphics
Pro and G4dn instances for Windows Server 2016 and Windows
Server 2019

- Includes CloudWatch Agent 1.3.50742

- Includes SSM Agent 3.2.2303.0

- Includes AWS Command Line Interface (AWS CLI) 2.15.33.0

- Includes AWSVirtualSmartCardReader 1.0.0.59

03-24-2024Linux

- Base

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Updated Linux to version 2.0.20240318.0. For more information,
see [2.0.20240318.0](../../../al2/latest/relnotes/relnotes-20240325.md).

01-26-2024Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes Microsoft security updates up to January 2024

12-11-2023Windows

- Base

- Graphics G4dn

- Graphics G5

- Add support for Windows Server 2022

11-13-2023Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Includes Microsoft security updates up to November 2023

11-13-2023

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Graphics G5

- Updated Linux to version 2.0.20231101.0. For more information,
see [Amazon Linux 2.0.20231101.0 release notes](../../../al2/latest/relnotes/relnotes-20231103.md).

06-12-2023Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft security updates up to June 2023

06-11-2023

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Updated Linux to version 2.0.20230530.0. For more information,
see [Amazon Linux 2 2.0.20230530.0 release notes](../../../al2/latest/relnotes/relnotes-20230607.md).

03-29-2023Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft security updates up to February 2023

03-15-2023

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Updated Linux to version 2.0.20220805.0. For more information,
see [Amazon Linux 2 2.0.20230221.0 release notes](../../../al2/latest/relnotes/relnotes-20230301.md).

- Improves Webcam experience

- Resolves an issue that prevents WorkSpaces Applications fleet instances from
provisioning when the system cryptography is set to use
FIPS-compliant algorithms

10-05-2022Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft security updates up to September 13,
2022

09-21-2022

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Updated Linux to version 2.0.20220805.0. For more information,
see [Amazon Linux 2.0.20220805.0 release notes](../../../al2/latest/relnotes/relnotes-20220823.md).

- Includes Image Assistant GUI

- Includes webcam support

09-14-2022

Amazon

Linux 2

- Graphics G4dn

- Graphics Pro

- Includes NVIDIA Graphics Driver (510.85.02)

09-01-2022Windows

- Graphics G4dn

- Graphics Pro

- Includes NVIDIA Graphics Driver (473.47) for Windows Server
2012 R2

- Includes NVIDIA Graphics Driver (512.78) for Windows Server
2016 and Windows Server 2019

07-12-2022Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft security updates up to June 14, 2022

- Includes latest AMD Driver (24.20.13028.7002) for Graphics
Design instances for Windows Server 2016 and Windows Server
2019

- Includes latest NVIDIA Graphics Driver (472.98) for Graphics
Pro and G4dn instances for Windows Server 2012R2

- Includes latest NVIDIA Graphics Driver (511.65) for Graphics
Pro and G4dn instances for Windows Server 2016 and Windows
Server 2019

- Includes CloudWatch Agent 1.3.50742

- Includes SSM Agent 3.1.1575.0

- Includes AWS Command Line Interface (AWS CLI) 1.23.11

06-20-2022

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Updated Linux to version 2.0.20220426.0. For more information,
see [Amazon\
Linux 2.0.20220426.0 release notes](../../../al2/latest/relnotes/relnotes-20220426.md).

03-03-2022Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft security updates up to January 11,
2022

02-18-2022

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Updated Linux to version 2.0.20211223.0. For more information,
see [Amazon\
Linux 2.0.20211223.0 release notes](../../../al2/latest/relnotes/relnotes-20211223.md).

- Latest Linux base images

11-19-2021

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Latest Linux base images, including blank screen fixes on
small instance types

11-15-2021

Amazon

Linux 2

- Base

- Graphics G4dn

- Graphics Pro

- Linux base images

10-08-2021Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Sample apps

- Includes Microsoft security updates up to September 15,
2021

- AWS Tools for PowerShell updated to version 3.15.1398

07-19-2021Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft Windows updates up to July 13, 2021

06-01-2021Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft Windows updates up to April 14, 2021

- Includes AMD driver version 24.20.13028.5012 for Graphics
Design instances

12-28-2020Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes a driver that adds support for using smart cards.
Smart cards can be used for Windows sign in, Active
Directory-joined streaming instances, and in-session
authentication for streaming applications

- Includes Microsoft Windows updates up to December 9,
2020

- Includes AWS CLI version 1.18.138

- Includes NVIDIA Graphics Driver version 451.48 for Graphics
Pro and Graphics G4dn instances

07-16-2020Windows

- Base

- Graphics Design

- Graphics G4dn

- Graphics Pro

- Includes Microsoft Windows updates up to June 9, 2020

- Includes AWS CLI version 1.18.86

- Includes NVIDIA Graphics Driver version 441.66 for Graphics
Pro instances

04-22-2020Windows

- Base (Windows Server 2019)

- Graphics Design (Windows Server 2019)

- Graphics G4dn (Windows Server 2019)

- Graphics Pro (Windows Server 2019)

- Includes Microsoft Windows updates up to March 10, 2020

- Includes AWS CLI version 1.18.21

- Includes NVIDIA Graphics Driver version 441.66 for Graphics
Pro instances

03-18-2020Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to February 11,
2020

- Includes AWS CLI version 1.17.5

- Includes NVIDIA Graphics Driver version 412.16 for Graphics
Pro instances

03-16-2020Windows

- Graphics G4dn

- Adds support for Graphics G4dn instances based on the EC2 G4dn
family (Windows Server 2012 R2)

- Includes Microsoft Windows updates up to February 11,
2020

- Includes AWS CLI version 1.17.5

03-05-2020Windows

- Graphics G4dn

- Adds support for Graphics G4dn instances based on the EC2 G4dn
family (Windows Server 2016 and Windows Server 2019)

- Includes Microsoft Windows updates up to February 11,
2020

- Includes AWS CLI version 1.17.5

01-13-2020Windows

- Graphics Design

- Adds support for Windows Server 2019, with Microsoft Windows
updates up to November 12, 2019

12-12-2019Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to November 12,
2019

- Includes AWS CLI version 1.16.284

- Includes a new version of the SSM Agent (v2.3.760.0), which
resolves an issue that prevented streaming instances from being
provisioned

09-18-2019Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to August 13, 2019 for
all Base and Graphics Pro instances and for Graphics Design
Windows Server 2012 R2. Graphics Design Windows Server 2016
instances already include this version.

- Includes AWS CLI version 1.16.222 for all Base and Graphics
Pro instances and Graphics Design Windows Server 2012 R2.
Graphics Design Windows Server 2016 instances already include
this version.

- Includes a fix to prevent Windows Defender from being enabled
by default on Windows Server 2016 and Windows Server 2019 image
builder instances. For more information, see [Windows Update and Antivirus Software on Amazon WorkSpaces Applications](windows-update-antivirus-software.md).

09-05-2019Windows

- Graphics Design

- Adds support for Windows Server 2016

- Includes Microsoft Windows updates up to August 13,
2019

- Includes AWS CLI version 1.16.222

- Includes AMD Driver version 24.20.13028.3002 for Graphics
Design instances (compatible with Windows Server 2016)

06-24-2019Windows

- Base

- Graphics Pro

- Adds support for Windows Server 2016 and Windows Server
2019

05-28-2019Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to May 14, 2019

04-29-2019Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to April 20, 2019

- Includes AWS CLI version 1.16.126

- Includes NVIDIA Graphics Driver 412.16 for Graphics Pro
instances

01-22-2019Windows

- Base

- Graphics Design

- Graphics Pro

- Includes Microsoft Windows updates up to December 10,
2018

- Includes AWS CLI version 1.16.84

- Includes NVIDIA Graphics Driver version 391.58 for Graphics
Pro instances

06-12-2018Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Includes Microsoft Windows updates up to May 9, 2018

- Includes Windows PowerShell 5.1

05-02-2018Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Includes Microsoft Windows updates up to April 10, 2018

- Adds the following language packs: Japanese, Korean,
Portuguese (Brazil), Thai, Chinese (Simplified), Chinese
(Traditional)

03-19-2018Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Includes Microsoft Windows updates up to February 23,
2018

- Includes the following language packs: German, French,
Italian, Spanish, Dutch

- Resolves intermittent issues with using Microsoft Visio and
Microsoft Project applications during streaming sessions

01-24-2018Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Includes Microsoft Windows updates up to January 5,
2018

- Includes Microsoft Windows updates for the Spectre and
Meltdown vulnerabilities

- Enables a default profile to be created on image builders and
used for the AWS Command Line Interface (CLI) during streaming
sessions

01-01-2018Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Resolves an issue with connectivity to WorkSpaces Applications instances

12-07-2017Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- Includes Microsoft Windows updates up to November 19,
2017

- Adds support for managed WorkSpaces Applications agent updates

11-13-2017Windows

- Base

- Resolves an issue with Microsoft Office 365 applications not
working during streaming sessions

- Includes Microsoft Windows updates up to October 11,
2017

09-05-2017Windows

- Base

- Graphics Design

- Graphics Desktop

- Graphics Pro

- New Graphics Design instance family

- Support for On-Demand fleets

- Updated approach for session context

- Includes Microsoft Windows updates up to August 9, 2017

- Resolves an intermittent issue with applications not coming to
the foreground

- Resolves an intermittent issue with applications not appearing
in tile view

07-25-2017Windows

- Graphics Desktop

- Graphics Pro

- New Graphics Desktop and Graphics Pro instance families

- Adds support for 2 K resolution

07-24-2017Windows

- Base

- Includes Microsoft Windows updates up to July 13, 2017

- Adds support for Microsoft Active Directory domains

06-20-2017Windows

- Base

- Sample apps

- Optimizes application launch performance

- Resolves an issue with applications not displaying in tile
view

- Resolves an issue with applications displaying in tile view
only

- Resolves an issue with applications displaying multiple times
in tile view

- Resolves an issue with recently launched application windows
not appearing in the foreground

- Resolves an issue with page margins when printing

05-18-2017Windows

- Base

- Sample apps

- Adds support for WorkSpaces Applications home folders

- Includes Microsoft Windows updates up to May 16, 2017

- Resolves an intermittent network issue that affects internet
connections from streaming instances

- Resolves an issue with application tiles not functioning
correctly

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Install AMD Driver on Graphics Design Instances

Images

All content copied from https://docs.aws.amazon.com/.
