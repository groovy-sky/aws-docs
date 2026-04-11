# EC2Launch v2 version histories

###### Version histories

- [EC2Launch v2 version history](#ec2launchv2-version-history)

## EC2Launch v2 version history

To ensure that you have the latest launch agent installed, see [Install the latest version of EC2Launch v2](ec2launch-v2-install.md). You can
receive notifications when new versions of the EC2Launch v2 agent are released.
For more information, see [Subscribe to EC2 Windows launch agent notifications](launch-agents-subscribe-notifications.md).

The following table describes the released versions of EC2Launch v2.

VersionDetailsRelease date2.4.0

- Simplified agent logging by removing redundant `console.log`, `err.log` files.

- Enhanced agent logging capabilities by adding log rotation, limiting log size to 1 MB and improving time precision to milliseconds.

- Reduced time to Windows is Ready by removing Hyper-V status query.

March 5, 20262.3.237

- Fixed an agent crash issue caused by the incompatibility between certain Visual C++ runtime versions and agent telemetry functions.

February 25, 20262.3.108

- Added IPv6 address display to wallpaper. For Memory and Network attributes, consult
[AWS Documentation](https://aws.amazon.com/ec2/instance-types).

December 29, 20252.3.56

- Fixed an issue where preconfigured persistent network routes
were not handled properly by EC2Launch v2.

November 4, 20252.3.5

- Improved primary network interface detection using
IMDS when adding routes.

September 15, 20252.2.63

- Updated launch agent to improve [Telemetry](ec2launch-v2.md#ec2launch-v2-telemetry).

July 29, 20252.1.1

- Added full support for IPv6.

May 14, 20252.0.2107

- Enhanced route addition to handle scenarios when IPv4 or
IPv6 addresses are not available on the interface.

March 27, 20252.0.2081

- Fixed an issue where RDP certificate information was not
properly retrieved or validated. Added functionality to
automatically start the Remote Desktop Services if
needed.

- Adjusted EC2Launch v2 service permissions to fix an issue that
occurs when querying the service status.

February 4, 20252.0.2046

- Updated the wallpaper path in the
`agent-config.yml` file to use the
default operating system wallpaper path.

- Added telemetry to monitor the locations where agent
errors occurred.

- Updated agent log messaging.

October 3, 20242.0.1981

- Updated `EC2Launch.exe` CLI command
error messages for non-Administrator users.

August 6, 20242.0.1948

- Added telemetry to monitor usage of admin password
options.

- Modified EC2Launch permissions.

July 1, 20242.0.1924

- Updated the EC2Launch Settings UI.

- Updated the wallpaper CLI command.

- Updated the EC2Launch installer.

June 10, 20242.0.1914

- Add routes with unspecified gateway addresses
( `0.0.0.0` for IPv4 or `::` for
IPv6).

- Always add both IPv4 and IPv6 routes.

- Fixed an issue where the `Administrator`
username was added to the
`agent-config.yml` file when it
wasn't specified.

- Modified EC2Launch v2 permissions.

June 5, 20242.0.1881

- Added an encrypted password option to
`setAdminAccount` task.

- Added CLI command to encrypt static password in
agent-config.yml.

- Fixed an issue where XML user data doesn't add PowerShell
arguments when it runs with Administrator permissions. For
more details, see [How Amazon EC2 handles user data for Windows instances](user-data.md#ec2-windows-user-data).

- Adjusted PowerShell arguments for the
`executeScript` task and user data scripts
when they run with `LocalSystem` permissions.
When arguments are empty, the agent uses the following
default value: `-ExecutionPolicy
  										Unrestricted`.

- Prevented printing duplicate driver versions to the
console log.

May 8, 20242.0.1815

- Adjusted error handling to fail on critical setup issues
before sysprep.

- Fixed an issue where wallpaper and hostname tasks could
use an incorrect IP address on instances with multiple IP
addresses assigned to the primary network interface.

- Wallpaper and hostname tasks changed to get private IP
from IMDS first, then fail back to WMI if IMDS is
disabled.

- Fixed an issue with the `initializeVolume` task
where `sc1` volumes failed to initialize due to a
transient error.

March 6, 20242.0.1739

- Fixed an issue that prevented exit codes from being
captured by `executeScript` tasks that were run
as the Windows Administrator user.

January 17, 20242.0.1702

- Restricted `Telemetry.log` permissions
to `read-execute` only for standard users.

- Configured the EC2Launch Windows service to restart on
start-up failure.

- Made `add-routes` failures actionable by
logging `route.exe` `stderr` output.

- Fixed an issue that occurs when route metrics are outside
of the range \[1, 9999\].

- Added wallpaper support to several new instance
types.

- Fixed an issue caused by user data scripts that run as the
Windows Administrator user and send output to
`stderr`.

January 4, 20242.0.1643

- Updated the `ebsnvme-id.exe` tool to
version 1.1.0.7.

- Fixed an issue with receive side scaling (RSS) and receive
queue depth settings on metal instance types that begin with
'metal-\*', such as metal-48x1.

- Removed telemetry event that reports on XML userdata
commands that block the agent.

- Updated `setDnsSuffix` task to limit domain
name devolution based on registry entry:
`HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\Dnscache\Parameters\DomainNameDevolutionLevel`.

- Added a public task and CLI that adds network
routes.

- _Note_ – This is the
last version to officially support Windows Server
2012.

- _Note_ – This is the
last version to officially support 32-bit operating
systems.

October 4, 20232.0.1580

- Changed the way that the launch agent handles errors when
you modify log file permissions.

- Added a timeout for connecting to the serial port. The
timeout allows the launch agent to continue running if the
serial port is in use.

September 5, 20232.0.1521

- Deprecated the `—block` flag of the
`EC2Launch.exe` **reset** and **sysprep**
commands.

- Updated `EC2Launch.exe` to detect and
handle the **reset** and
**sysprep** commands that are used in
inline `executeScript` tasks. Those commands
cause the agent to stop running after the
`executeScript` task runs them.

- Updated XML userdata scripts to run inline by
default.

- Enable XML userdata scripts to run detached with the new
`detach` tag. For more details, see [User data scripts](user-data.md#user-data-scripts).

- Made the following changes to the agent log.

- Updated agent log messages.

- Removed `executeScript` content and
output from the agent log.

- Removed `executeProgram` arguments and
output from the agent log.

- Made the following changes to the console log.

- Added
`EnableSCSIPersistentReservations`
value to the console log.

July 3, 20232.0.1303

- Added additional error handling and log lines when adding
network routes.

- Allowed `executeScript` and
`executeProgram` tasks in the PreReady
stage.

- Updated `executeProgram` task to generate
output files similar to the output from the executeScript
task. For more information, see [executeProgram](ec2launch-v2-task-definitions.md#ec2launch-v2-executeprogram).

- Added telemetry to monitor usage of blocking agent
commands in XML user data.

May 3, 20232.0.1245

- Improved visibility into crashes by logging crash call
stacks in clear text.

- Added the EventLog service as a startup dependency to fix
a crash when the Amazon EC2Launch service starts up faster
than the EventLog service.

- Made XML user data run before PostReady stage from the
agent config file (like EC2Launch v1 and EC2Config).

- Added YAML user data version 1.1 to make user data run
before PostReady stage from the agent config file (YAML user
data version 1.0 runs after PostReady stage from the agent
config file).

March 8, 20232.0.1173

- Adds an optional feature to display instance tags on
wallpaper. For more information, see [setWallpaper](ec2launch-v2-task-definitions.md#ec2launch-v2-setwallpaper).

- Adds error handling when the security group for Elastic
Graphics is not properly set up.

- Fixes a timeout when the Instance Metadata Service is not
enabled.

February 6, 20232.0.1121

- Fixes an issue where a 404 error is printed to the
wallpaper when no public IPv4 address is assigned.

- Fixes an issue where the volume's file system is formatted
as `RAW` instead of `NTFS` when its
device's drive letter is set to `D`.

- Fixes an issue where NVMe SSD volumes are incorrectly
identified as EBS volumes.

- Fixes an error when activating Windows when IMDS is
disabled.

January 4, 20232.0.1082

- Fixes an issue where the `setWallpaper`:
`privateIpAddress` field is blank when IMDS
is disabled.

- Fixes an issue with setting the hostname to the private
IPv4 address when IMDS is disabled.

- Fixes an issue with initializing volumes on Windows Server
2012.

- Fixes an issue with setting jumbo frames.

- Fixes an error when no SSH key is specified at instance
launch.

- Fixes an error on Windows Server 2012 when Windows does
not have a 'ReleaseId' registry key.

December 7, 20222.0.1011

- Fixes logic for finding network adapter when PnPDeviceID
is empty.

November 11, 20222.0.1009

- Uses PCI segment information to select the console
port.

November 8, 20222.0.982

- Adds retry logic to get RDP information.

- Fixes errors during volume initialization on
`d2.8xlarge` instances.

- Fixes issue where an incorrect network adapter can be
selected after a reboot.

- Removes false alarm error message when ACPI SPCR is
unavailable.

October 31, 20222.0.863

- Updates IMDS wait logic to make only IMDSv2
requests.

- Adds logic to assign drive letter to volumes that are
already initialized but not mounted.

- Prints a more specific error message when key pair type is
not supported.

- Fixes 3010 reboot code bug.

- Adds check for invalid base64-encoded user data.

July 6, 20222.0.698

- Fixes typo in log output when executing scripts.

January 30, 20222.0.674

- Telemetry uploads the enabled/disabled privacy
control.

- Fixes `index out of bounds` bug.

- Removes wallpaper shortcuts during
`sysprep`.

November 15, 20212.0.651

- Adds logic to uninstall legacy agents during EC2Launch v2
installation.

- Fixes `list-volume` CLI issue when root volume
is not listed as volume 0.

October 7, 20212.0.592

- Fixes bug to correctly report stage status.

- Removes false alarm error messages when log files are
closed.

- Adds telemetry.

August 31, 20212.0.548

- Adds leading zeros for hex IP hostname.

- Fixes file permissions for `enableOpenSsh`
task.

- Fixes sysprep command crash.

August 4, 20212.0.470

- Fixes bug in network stage to wait for DHCP to assign an
IP to the instance.

- Fixes bug with `setDnsSuffix` when
`SearchList` registry key does not
exist.

- Fixes bug in DNS devolution logic in
`setDnsSuffix`.

- Adds network routes after intermediate reboots.

- Allows `initializeVolume` to re-letter existing
volumes.

- Removes extra information from version subcommand.

July 20, 20212.0.285

- Adds option to run user scripts in a detached
process.

- Legacy userdata (XML userdata) now runs in a detached
process, which is similar behavior to the prior launch
agent.

- Adds CLI flag to the `sysprep` and
`reset` commands, which allows them to block
until the service stops.

- Restricts the config folder permissions.

March 8, 20212.0.207

- Adds optional `hostName` field to
`setHostName` task.

- Fixes reboot bug. Reboot tasks `executeScript`
and `executeProgram` will be marked as running.

- Adds more return codes to the status command.

- Adds bootstrap service to fix startup issue when running
on `t2.nano` instance type.

- Fixes clean installation mode to remove files not tracked
by installer.

February 2, 20212.0.160

- Fixes `validate` command to detect invalid
stage name.

- Adds `w32tm resync` command in
`addroutes` task.

- Fixes issue with changing DNS suffix search order.

- Adds check conditions to better report invalid user
data.

December 4, 20202.0.153Adds Sysprep functionality in UserData.November 3, 20202.0.146

- Fixes issue with RootExtend on non-English AMIs.

- Grants users group write permission to log files.

- Creates MS Reserved partition for GPT volumes.

- Adds list-volumes command and volume dropdown in Amazon
EC2Launch settings.

- Adds get-agent-config command for printing
agent-config.yml file in yaml or json format.

- Erases static password if no public key detected.

October 6, 20202.0.124

- Adds option to display OS version on wallpaper.

- Initializes encrypted EBS volumes.

- Adds routes for VPCs with no local DNS name.

September 10, 20202.0.104

- Creates DNS suffix search list if it does not exist.

- Skips Hibernation if not requested.

August 12, 20202.0.0Initial release.June 30, 2020

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Troubleshoot EC2Launch v2

EC2Launch
