AWS Application Discovery Service is no longer open to new customers. Alternatively, use AWS Transform which provides similar capabilities. For more information, see [AWS Application Discovery Service availability change](application-discovery-service-availability-change.md).

# AWS Application Discovery Agent

The AWS Application Discovery Agent (Discovery Agent) is software that you install on on-premises servers and VMs
targeted for discovery and migration. Agents capture system configuration, system
performance, running processes, and details of the network connections between systems.
Agents support most Linux and Windows operating systems, and you can deploy them on physical
on-premises servers, Amazon EC2 instances, and virtual machines.

###### Note

Before you deploy the Discovery Agent, you must choose a [Migration Hub home Region](../../../migrationhub/latest/ug/home-region.md). You must
register your agent in your home Region.

The Discovery Agent runs in your local environment and requires root privileges. When you start
the Discovery Agent, it connects securely with your home region and registers with Application Discovery Service.

- For example, if `eu-central-1` is your home Region, it registers
`arsenal-discovery.eu-central-1.amazonaws.com`
with Application Discovery Service.

- Or substitute your home Region as needed for all other Regions except
us-west-2.

- If `us-west-2` is your home Region, it registers
`arsenal.us-west-2.amazonaws.com` with Application Discovery Service.

## How it works

After registration, the agent starts collecting data for the host or VM where it
resides. The agent pings the Application Discovery Service at 15-minute intervals for configuration
information.

The collected data includes system specifications, times series utilization or
performance data, network connections, and process data. You can use this information to
map your IT assets and their network dependencies. All of these data points can help you
determine the cost of running these servers in AWS and also plan for migration.

Data is transmitted securely by the Discovery Agents to Application Discovery Service using Transport Layer
Security (TLS) encryption. Agents are configured to upgrade automatically when new
versions become available. You can change this configuration setting if desired.

###### Tip

Before downloading and beginning Discovery Agent installation, be sure to read through
all of the required prerequisites in [Prerequisites for Discovery Agent](gen-prep-agents.md)

## Data collected by Discovery Agent

AWS Application Discovery Agent (Discovery Agent) is software that you install on on-premises servers and VMs.
Discovery Agent collects system configuration, times series utilization or performance data,
process data, and Transmission Control Protocol (TCP) network connections. This section
describes the data that's collected.

###### Table legend for Discovery Agent collected data:

- The term host refers to either a physical server or a VM.

- Collected data is in measurements of kilobytes (KB) unless stated
otherwise.

- Equivalent data in the Migration Hub console is reported in megabytes (MB).

- The polling period is in intervals of approximately 15 seconds and is sent to
AWS every 15 minutes.

- Data fields denoted with an asterisk (\*) are only available in the
`.csv` files that are produced from the agent's API export
function.

Data fieldDescriptionagentAssignedProcessId\*Process ID of processes discovered by the agentagentIdUnique ID of agentagentProvidedTimeStamp\*Date and time of agent observation _(mm/dd/yyyy hh:mm:ss am/pm)_cmdLine\*Process entered at the command linecpuType Type of CPU (central processing unit) used in hostdestinationIp\*IP address of device to which packet is being sentdestinationPort\*Port number to which the data/request is to be sentfamily\*Protocol of routing familyfreeRAM (MB) Free RAM and cached RAM that can be made immediately available to
applications, measured in MBgateway\*Node address of networkhostNameName of host data was collected onhypervisorType of hypervisoripAddressIP address of the hostipVersion\*IP version numberisSystem\*Boolean attribute to indicate if a process is owned by the OSmacAddress MAC address of the hostname\*Name of the host, network, metrics, etc. data is being collected
fornetMask\*IP address prefix that a network host belongs toosName Operating system name on hostosVersionOperating system version on hostpathPath of the command sourced from the command linesourceIp\*IP address of the device sending the IP packet sourcePort\*Port number from which the data/request originates fromtimestamp\*Date and time of reported attribute logged by agenttotalCpuUsagePct Percentage of CPU usage on host during polling periodtotalDiskBytesReadPerSecond (Kbps)Total kilobits read per second across all diskstotalDiskBytesWrittenPerSecond (Kbps)Total kilobits written per second across all disks totalDiskFreeSize (GB)Free disk space expressed in GBtotalDiskReadOpsPerSecondTotal number of read I/O operations per secondtotalDiskSize (GB)Total capacity of disk expressed in GBtotalDiskWriteOpsPerSecondTotal number of write I/O operations per secondtotalNetworkBytesReadPerSecond (Kbps)Total amount of throughput of bytes read per secondtotalNetworkBytesWrittenPerSecond (Kbps)Total amount of throughput of bytes written per secondtotalNumCoresTotal number of independent processing units within CPUtotalNumCpusTotal number of central processing unitstotalNumDisksThe number of physical hard disks on a hosttotalNumLogicalProcessors\*Total number of physical cores times the number of threads that can
run on each coretotalNumNetworkCardsTotal count of network cards on servertotalRAM (MB)Total amount of RAM available on hosttransportProtocol\*Type of transport protocol used

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting up

Prerequisites

All content copied from https://docs.aws.amazon.com/.
