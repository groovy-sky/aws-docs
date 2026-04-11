# EC2Config version history

The following table describes the released versions of EC2Config. For information about
the updates for SSM Agent, see [Systems Manager SSM Agent Release Notes](https://github.com/aws/amazon-ssm-agent/blob/master/RELEASENOTES.md).

###### Important

EC2Config has reached the end of support. Only the latest version of the EC2Config agent is
available for download. Prior versions are marked as private.

VersionDetailsRelease date4.9.5777

- Fixed issue where RSS configuration was set incorrectly for some
instance types.

- New version of SSM Agent `3.3.484.0`.

17 June 20244.9.5554

- Limit domain name devolution based on registry entry:
`HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\Dnscache\Parameters\DomainNameDevolutionLevel`.

- New version of SSM Agent `3.2.1630.0`.

4 October 20234.9.5467

- Added retry capability for discovering console port.

- New version of SSM Agent `3.1.2282.0`.

1 August 20234.9.5288

- Updated AWS Core SDK to version `3.7.103.23`.

- Fixed issue where the `AWS-UpdateEC2Config` SSM document fails
to update `EC2Config` on instances enabled with only IMDSv2.

- New version of SSM Agent `3.1.2144.0`.

8 March 20234.9.5231

- New version of SSM Agent 3.1.1927.0.

14 February 20234.9.5103

- Fixed issue where ephemeral volumes are incorrectly identified on r5d and i4i instance
families.

- New version of SSM Agent 3.1.1856.0.

5 December 20224.9.5064

- Updated to use PCI segment information to select the console port.

- Signed PowerShell scripts and added copyright headers.

- Fixed primary network adapter selection logic.

- New version of SSM Agent 3.1.1732.0.

16 November 20224.9.4588

- Updated IMDS wait logic to make only IMDSv2 requests.

- Added libec2launch.dll launch-agent shared library.

- New version of SSM Agent 3.1.1188.0.

31 May 20224.9.4556

- Added wait logic to ensure full initialization of NIC before use.

- New version of Log4Net 2.0.14.0 picks up security
patch.

- New version of SSM Agent 3.1.1045.0 picks up security patch.

1 March 20224.9.4536

- Fixed issue where userdata crashes when the Temp folder is missing.

- New version of SSM Agent 3.1.804.0.

31 January 20224.9.4508

- Fixed issue to correctly compute diskpart script path.

- New version of SSM Agent 3.1.338.0.

6 October 20214.9.4500

- Updated `Install-EgpuManagerConfig` with IMDS v2 support.

- Updated web links to use https.

- New version of SSM Agent 3.1.282.0

7 September 20214.9.4419

- Fixed IMDS version 1 fallback logic

- Updated all usage of Windows temp directory to EC2Config temp directory

- New version of SSM Agent 3.0.1124.0

2 June 20214.9.4381

- Added support for SSM document schema version 2.2 in EC2ConfigUpdater

- Added AWS Nitro Enclaves package version to console log

- New version of SSM Agent 3.0.529.0

4 May 20214.9.4326

- Removed all links in the settings UI

- This is the last EC2Config version that supports Windows
Server 2008.

3 March 20214.9.4279

- Fixed security issue related to `Ec2ConfigMonitor` scheduled task

- Fixed drive letter mapping issue and incorrect ephemeral disk count

- Added `OsCurrentBuild` and `OsReleaseId` to console output

- New version of SSM Agent 2.3.871.0

11 December 20204.9.4222

- Fixed IMDS version 1 fallback logic

- New version of SSM Agent 2.3.842.0

7 April 20204.9.4122

- Added support for IMDS v2

- New version of SSM Agent 2.3.814.0

4 March 20204.9.3865

- Fixed issue detecting COM port for Windows Server 2008 R2 on metal instances

- New version of SSM Agent 2.3.722.0

31 October 20194.9.3519

- New version of SSM Agent 2.3.634.0

18 June 20194.9.3429

- New version of SSM Agent 2.3.542.0

25 April 20194.9.3289

- New version of SSM Agent 2.3.444.0

11 February 20194.9.3270

- Added plugin for setting the monitor to never turn off to fix ACPI issues

- SQL Server edition and version written to console

- New version of SSM Agent 2.3.415.0

22 January 20194.9.3230

- Drive Letter Mapping description updated to better align to functionality

- New version of SSM Agent 2.3.372.0

10 January 20194.9.3160

- Increased wait time for primary NIC

- Added default configuration for RSS and Receive Queue settings for ENA devices

- Disabled hibernation during Sysprep

- New version of SSM Agent 2.3.344.0

- Upgraded AWS SDK to 3.3.29.13

15 December 20184.9.3067

- Improvements made to instance hibernation

- New version of SSM Agent 2.3.235.0

8 November 20184.9.3034

- Added route 169.254.169.253/32 for DNS server

- New version of SSM Agent 2.3.193.0

24 October 20184.9.2986

- Added signing for all EC2Config related binaries

- New version of SSM Agent 2.3.136.0

11 October 20184.9.2953

New version of SSM Agent (2.3.117.0)

2 October 20184.9.2926

New version of SSM Agent (2.3.68.0)

18 September 20184.9.2905

- New version of SSM Agent (2.3.50.0)

- Added route 169.254.169.123/32 to AMZN time service

- Added route 169.254.169.249/32 to GRID license service

- Fixed an issue causing EBS NVMe volumes to be marked as ephemeral

17 September 20184.9.2854

New version of SSM Agent (2.3.13.0)

17 August 20184.9.2831

New version of SSM Agent (2.2.916.0)

7 August 20184.9.2818

New version of SSM Agent (2.2.902.0)

31 July 20184.9.2756

New version of SSM Agent (2.2.800.0)

27 June 20184.9.2688

New version of SSM Agent (2.2.607.0)

25 May 20184.9.2660

New version of SSM Agent (2.2.546.0)

11 May 20184.9.2644

New version of SSM Agent (2.2.493.0)

26 April 20184.9.2586

New version of SSM Agent (2.2.392.0)

28 March 20184.9.2565

- New version of SSM Agent (2.2.355.0)

- Fixed an issue on M5 and C5 instances (unable to find PV drivers)

- Add console logging for instance type, newest PV drivers, and NVMe drivers

13 March 20184.9.2549

New version of SSM Agent (2.2.325.0)

8 March 20184.9.2461

New version of SSM Agent (2.2.257.0)

15 February 20184.9.2439

New version of SSM Agent (2.2.191.0)

6 February 20184.9.2400

New version of SSM Agent (2.2.160.0)

16 January 20184.9.2327

- New version of SSM Agent (2.2.120.0)

- Added COM port discovery on Amazon EC2 bare metal instances

- Added Hyper-V status logging on Amazon EC2 bare metal
instances

2 January 20184.9.2294

New version of SSM Agent (2.2.103.0)

4 December 20174.9.2262

New version of SSM Agent (2.2.93.0)

15 November 20174.9.2246

New version of SSM Agent (2.2.82.0)

11 November 20174.9.2218

New version of SSM Agent (2.2.64.0)

29 October 20174.9.2212

New version of SSM Agent (2.2.58.0)

23 October 20174.9.2203

New version of SSM Agent (2.2.45.0)

19 October 20174.9.2188

New version of SSM Agent (2.2.30.0)

10 October 20174.9.2180

- New version of SSM Agent (2.2.24.0)

- Added the Elastic GPU plugin for GPU instances

5 October 20174.9.2143

New version of SSM Agent (2.2.16.0)

1 October 20174.9.2140

New version of SSM Agent (2.1.10.0)

4.9.2130

New version of SSM Agent (2.1.4.0)

4.9.2106

New version of SSM Agent (2.0.952.0)

4.9.2061

New version of SSM Agent (2.0.922.0)

4.9.2047

New version of SSM Agent (2.0.913.0)

4.9.2031

New version of SSM Agent (2.0.902.0)

4.9.2016

- New version of SSM Agent (2.0.879.0)

- Fixed the CloudWatch Logs directory path for Windows Server 2003

4.9.1981

- New version of SSM Agent (2.0.847.0)

- Fixed the issue with `important.txt` being
generated in EBS volumes.

4.9.1964New version of SSM Agent (2.0.842.0)4.9.1951

- New version of SSM Agent (2.0.834.0)

- Fixed the issue with drive letter not being mapped from Z:
for ephemeral drives.

4.9.1925

- New version of SSM Agent (2.0.822.0)

- \[Bug\] This version is not a valid update target from
SSM Agent v4.9.1775.

4.9.1900New version of SSM Agent (2.0.805.0)4.9.1876

- New version of SSM Agent (2.0.796.0)

- Fixed an issue with output/error redirection for admin
userdata execution.

4.9.1863

- New version of SSM Agent (2.0.790.0)

- Fixed problems with attaching multiple EBS volumes to an
Amazon EC2 instance.

- Improved CloudWatch to take a configuration path, keeping the
backwards compatibility.

4.9.1791New version of SSM Agent (2.0.767.0)

4.9.1775

New version of SSM Agent (2.0.761.0)

4.9.1752

New version of SSM Agent (2.0.755.0)

4.9.1711

New version of SSM Agent (2.0.730.0)

4.8.1676

New version of SSM Agent (2.0.716.0)

4.7.1631

New version of SSM Agent (2.0.682.0)

4.6.1579

- New version of SSM Agent (2.0.672.0)

- Fixed agent update issue with v4.3, v4.4, and v4.5

4.5.1534

New version of SSM Agent (2.0.645.1)

4.4.1503

New version of SSM Agent (2.0.633.0)

4.3.1472

New version of SSM Agent (2.0.617.1)

4.2.1442

New version of SSM Agent (2.0.599.0)

4.1.1378

New version of SSM Agent (2.0.558.0)

4.0.1343

- Run Command, State Manager, the CloudWatch agent, and domain join
support have been moved into another agent called SSM Agent.
SSM Agent will be installed as part of the EC2Config
upgrade. For more information, see [EC2Config and AWS Systems Manager](ec2config-service.md#ec2config-ssm).

- If you have a proxy set up in EC2Config, you will need to
update your proxy settings for SSM Agent before upgrading.
If you do not update the proxy settings, you will not be
able to use Run Command to manage your instances. To avoid
this, see the following information before updating to the
newer version: [Installing \
and Configuring SSM Agent on Windows Instances](../../../systems-manager/latest/userguide/ssm-agent-windows.md) in the
_AWS Systems Manager User Guide_.

- If you previously enabled CloudWatch integration on your
instances by using a local configuration file
( `AWS.EC2.Windows.CloudWatch.json`), you will need to
configure the file to work with SSM Agent.

3.19.1153

- Re-enabled activation plugin for instances with old AWS KMS configuration. Skip
activation for BYOL users.

- Change default TRIM behavior to be disabled during disk
format operation and added FormatWithTRIM for overriding
InitializeDisks plugin with userdata.

3.18.1118

- Fix to reliably add routes to the primary network
adapter.

- Updates to improve support for AWS services.

3.17.1032

- Fixes duplicate system logs appearing when filters set to
same category.

- Fixes to prevent from hanging during disk
initialization.

3.16.930

Added support to log "Window is Ready to use" event to Windows
Event Log on start.

3.15.880

Fix to allow uploading Systems Manager Run Command output to S3 bucket names with
'.' character.

3.14.786

Added support to override InitializeDisks plugin settings. For
example: To speed up SSD disk initialize, you can temporarily
disable TRIM by specifying this in userdata:

<InitializeDrivesSettings><SettingsGroup>FormatWithoutTRIM</SettingsGroup></InitializeDrivesSettings

3.13.727

Systems Manager Run Command - Fixes to process commands reliably after windows
reboot.

3.12.649

- Fix to gracefully handle reboot when running
commands/scripts.

- Fix to reliably cancel running commands.

- Add support for (optionally) uploading MSI logs to S3 when
installing applications via Systems Manager Run Command.

3.11.521

- Fixes to enable RDP thumbprint generation for Windows Server
2003.

- Fixes to include timezone and UTC offset in the EC2Config
log lines.

- Systems Manager support to run Run Command commands in parallel.

- Roll back previous change to bring partitioned disks
online.

3.10.442

- Fix Systems Manager configuration failures
when installing MSI applications.

- Fix to reliably bring storage disks online.

- Updates to improve support for AWS services.

3.9.359

- Fix in post Sysprep script to leave the configuration of
windows update in a default state.

- Fix the password generation plugin to improve the
reliability in getting GPO password policy settings.

- Restrict EC2Config/SSM log folder permissions to the local
Administrators group.

- Updates to improve support for AWS services.

3.8.294

- Fixed an issue with CloudWatch that prevented logs from
getting uploaded when not on primary drive.

- Improved the disk initialization process by adding retry
logic.

- Added improved error handling when the SetPassword plugin
occasionally failed during AMI creation.

- Updates to improve support for AWS services.

3.7.308

- Improvements to the ec2config-cli utility for config
testing and troubleshooting within instance.

- Avoid adding static routes for AWS KMS and meta-data service
on an OpenVPN adapter.

- Fixed an issue where user-data execution was not honoring
the "persist" tag.

- Improved error handling when logging to the EC2 console is
not available.

- Updates to improve support for AWS services.

3.6.269

- Windows activation reliability fix to first use link local
address 169.254.0.250/251 for activating windows via
AWS KMS

- Improved proxy handling for Systems Manager, Windows Activation and
Domain Join scenarios

- Fixed an issue where duplicate lines of user accounts were
added to the Sysprep answer file

3.5.228

- Addressed a scenario where the CloudWatch plugin may
consume excessive CPU and memory reading Windows Event
Logs

- Added a link to the CloudWatch configuration documentation
in the EC2Config Settings UI

3.4.212

- Fixes to EC2Config when used in combination with
VM-Import.

- Fixed service naming issue in the WiX installer.

3.3.174

- Improved exception handling for Systems Manager and domain join
failures.

- Change to support Systems Manager SSM schema versioning.

- Fixed formatting ephemeral disks on Win2K3.

- Change to support configuring disk size greater than
2TB.

- Reduced virtual memory usage by setting GC mode to
default.

- Support for downloading artifacts from UNC path in
`aws:psModule` and `aws:application` plugin.

- Improved logging for Windows activation plugin.

3.2.97

- Performance improvements by delay loading Systems Manager SSM
assemblies.

- Improved exception handling for malformed
sysprep2008.xml.

- Command line support for Systems Manager "Apply" configuration.

- Change to support domain join when there is a pending
computer rename.

- Support for optional parameters in the `aws:applications`
plugin.

- Support for command array in `aws:psModule` plugin.

3.0.54

- Enable support for Systems Manager.

- Automatically domain join EC2 Windows instances to an AWS
directory via Systems Manager.

- Configure and upload CloudWatch logs/metrics via
Systems Manager.

- Install PowerShell modules via Systems Manager.

- Install MSI applications via Systems Manager.

2.4.233

- Added scheduled task to recover EC2Config from service
startup failures.

- Improvements to the Console log error messages.

- Updates to improve support for AWS services.

2.3.313

- Fixed an issue with large memory consumption in some cases
when the CloudWatch Logs feature is enabled.

- Fixed an upgrade bug so that ec2config versions lower than
2.1.19 can now upgrade to latest.

- Updated COM port opening exception to be more friendly and
useful in logs.

- Ec2configServiceSettings UI disabled resizing and fixed
the attribution and version display placement in UI.

2.2.12

- Handled NullPointerException while querying a registry key
for determining Windows Sysprep state which returned null
occasionally.

- Freed up unmanaged resources in finally block.

2.2.11

Fixed a issue in CloudWatch plugin for handling empty log lines.

2.2.10

- Removed configuring CloudWatch Logs settings through UI.

- Enable users to define CloudWatch Logs settings in
`%ProgramFiles%\Amazon\Ec2ConfigService\Settings\AWS.EC2.Windows.CloudWatch.json`
file to allow future enhancements.

2.2.9

Fixed unhandled exception and added logging.

2.2.8

- Fixes Windows OS version check in EC2Config Installer to
support Windows Server 2003 SP1 and later.

- Fixes null value handling when reading registry keys
related to updating Sysprep config files.

2.2.7

- Added support for EC2Config to run during Sysprep
execution for Windows 2008 and greater.

- Improved exception handling and logging for better
diagnostics

2.2.6

- Reduced the load on the instance and on CloudWatch Logs when
uploading log events.

- Addressed an upgrade issue where the CloudWatch Logs plug-in did
not always stay enabled

2.2.5

- Added support to upload logs to CloudWatch Log Service.

- Fixed a race condition issue in Ec2OutputRDPCert
plug-in

- Changed EC2Config Service recovery option to Restart from
TakeNoAction

- Added more exception information when EC2Config
Crashes

2.2.4

- Fixed a typo in PostSysprep.cmd

- Fixed the bug which EC2Config does not pin itself onto
start menu for OS2012+

2.2.3

- Added option to install EC2Config without service starting
immediately upon install. To use, run 'Ec2Install.exe
start=false' from the command prompt

- Added parameter in wallpaper plugin to control
adding/removing wallpaper. To use, run 'Ec2WallpaperInfo.exe
set' or 'Ec2WallpaperInfo.exe revert' from the command
prompt

- Added checking for RealTimeIsUniversal key, output
incorrect settings of the RealTimeIsUniveral registry key to
the Console

- Removed EC2Config dependency on Windows temp folder

- Removed UserData execution dependency on .Net 3.5

2.2.2

- Added check to service stop behavior to check that
resources are being released

- Fixed issue with long execution times when joined to
domain

2.2.1

- Updated Installer to allow upgrades from older
versions

- Fixed Ec2WallpaperInfo bug in .Net4.5 only
environment

- Fixed intermittent driver detection bug

- Added silent install option. Execute Ec2Install.exe with
the '-q' option. eg: 'Ec2Install.exe -q'

2.2.0

- Added support for .Net4 and .Net4.5 only
environments

- Updated Installer

2.1.19

- Added ephemeral disk labeling support when using Intel
network driver (eg. C3 instance Type). For more information,
see [Enhanced networking on Amazon EC2 instances](enhanced-networking.md).

- Added AMI Origin Version and AMI Origin Name support to
the console output

- Made changes to the Console Output for consistent
formatting/parsing

- Updated Help File

2.1.18

- Added EC2Config WMI Object for Completion notification
(-Namespace root\\Amazon -Class EC2\_ConfigService)

- Improved Performance of Startup WMI query with large Event
Logs; could cause prolonged high CPU during initial
execution

2.1.17

- Fixed UserData execution issue with Standard Output and
Standard Error buffer filling

- Fixed incorrect RDP thumbprint sometimes appearing in
Console Output for >= w2k8 OS

- Console Output now contains 'RDPCERTIFICATE-SubjectName:'
for Windows 2008+, which contains the machine name
value

- Added D:\ to Drive Letter Mapping dropdown

- Moved Help button to top right and changed
look/feel

- Added Feedback survey link to top right

2.1.16

- General Tab includes link to EC2Config download page for
new Versions

- Desktop Wallpaper overlay now stored in Users Local
Appdata folder instead of My Documents to support MyDoc
redirection

- MSSQLServer name sync'd with system in Post-Sysprep script
(2008+)

- Reordered Application Folder (moved files to Plugin
directory and removed duplicate files)

- Changed System Log Output (Console):

- \*Moved to a date, name, value format for easier parsing
(Please start migrating dependencies to new format)

- \*Added 'Ec2SetPassword' plugin status

- \*Added Sysprep Start and End times

- Fixed issue of Ephemeral Disks not being labeled as
'Temporary Storage' for non-english Operating Systems

- Fixed EC2Config Uninstall failure after running
Sysprep

2.1.15

- Optimized requests to the Metadata service

- Metadata now bypass Proxy Settings

- Ephemeral Disks labeled as 'Temporary Storage' and
Important.txt placed on volume when found (Citrix PV drivers
only). For more information, see [Upgrade PV drivers on EC2 Windows instances](upgrading-pv-drivers.md).

- Ephemeral Disks assigned drive letters from Z to A (Citrix
PV drivers only) - assignment can be overwritten using Drive
Letter Mapping plugin with Volume labels 'Temporary Storage
X' where x is a number 0-25)

- UserData now runs immediately following 'Windows is
Ready'

2.1.14

Desktop wallpaper fixes

2.1.13

- Desktop wallpaper will display hostname by default

- Removed dependency on Windows Time service

- Route added in cases where multiple IPs are assigned to a
single interface

2.1.11

- Changes made to Ec2Activation Plugin

- -Verifies Activation status every 30 days

- -If Grace Period has 90 days remaining (out of 180),
reattempts activation

2.1.10

- Desktop wallpaper overlay no longer persists with Sysprep
or Shutdown without Sysprep

- Userdata option to run on every service start with
<persist>true</persist>

- Changed location and name of /DisableWinUpdate.cmd to
/Scripts/PostSysprep.cmd

- Administrator password set to not expire by default in
/Scripts/PostSysprep.cmd

- Uninstall will remove EC2Config PostSysprep script from
c:\\windows\\setup\\script\\CommandComplete.cmd

- Add Route supports custom interface metrics

2.1.9

UserData Execution no longer limited to 3851 Characters

2.1.7

- OS Version and language identifier written to
console

- EC2Config version written to console

- PV driver version written to console

- Detection of Bug Check and output to the console on next
boot when found

- Option added to config.xml to persist Sysprep
credentials

- Add Route Retry logic in cases of ENI being unavailable at
start

- User Data execution PID written to console

- Minimum generated password length retrieved from
GPO

- Set service start to retry 3 attempts

- Added S3\_DownloadFile.ps1 and S3\_Upload file.ps1 examples
to /Scripts folder

2.1.6

- Version information added to General tab

- Renamed the Bundle tab to Image

- Simplified the process of specifying passwords and moved
the password-related UI from the General tab to the Image
tab

- Renamed the Disk Settings tab to Storage

- Added a Support tab with common tools for
troubleshooting

- Windows Server 2003 `sysprep.ini` set to extend OS partition by
default

- Added the private IP address to the wallpaper

- Private IP address displayed on wallpaper

- Added retry logic for Console output

- Fixed Com port exception for metadata accessibility --
caused EC2Config to terminate before console output is
displayed

- Checks for activation status on every boot -- activates as
necessary

- Fixed issue of relative paths -- caused when manually
executing wallpaper shortcut from startup folder; pointing
to Administrator/logs

- Fixed default background color for Windows Server 2003 user
(other than Administrator)

2.1.2

- Console timestamps in UTC (Zulu)

- Removed appearance of hyperlink on Sysprep tab

- Addition of feature to dynamically expand Root Volume on
first boot for Windows 2008+

- When Set-Password is enabled, now automatically enables
EC2Config to set the password

- EC2Config checks activation status prior to running
Sysprep (presents warning if not activated)

- Windows Server 2003 `Sysprep.xml` now defaults to UTC timezone
instead of Pacific

- Randomized Activation Servers

- Renamed Drive Mapping tab to Disk Settings

- Moved Initialize Drives UI items from General to the Disk
Settings tab

- Help button now points to HTML help file

- Updated HTML help file with changes

- Updated 'Note' text for Drive Letter Mappings

- Added InstallUpdates.ps1 to /Scripts folder for automating
Patches and cleanup prior to Sysprep

2.1.0

- Desktop wallpaper displays instance information by default
upon first logon (not disconnect/reconnect)

- PowerShell can be run from the userdata by
surrounding the code with
<powershell></powershell>

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Troubleshoot EC2Config

EC2 Fast Launch for Windows
