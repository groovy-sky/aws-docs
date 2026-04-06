AWS Application Discovery Service is no longer open to new customers. Alternatively, use AWS Transform which provides similar capabilities. For more information, see [AWS Application Discovery Service availability change](application-discovery-service-availability-change.md).

# Data collected by the Agentless Collector VMware vCenter data collection module

The following information describes the data that's collected by the
Application Discovery Service Agentless Collector (Agentless Collector) VMware vCenter data collection
module. For information about setting up data collection, see [Setting up the Agentless Collector data collection module for VMware vCenter](agentless-collector-gs-vcenter.md).

###### Table legend for Agentless Collector VMware vCenter collected data:

- Collected data is in measurements of kilobytes (KB) unless stated otherwise.

- Equivalent data in the Migration Hub console is reported in megabytes (MB).

- Data fields denoted with an asterisk (\*) are available only in the .csv files that are
produced from the Application Discovery Service API export function.

The Agentless Collector supports data export using the AWS CLI. To export
collected data using the AWS CLI, follow the instructions described under **Export System Performance Data for All Servers** on the page [Export\
Collected Data](https://docs.aws.amazon.com/application-discovery/latest/userguide/export-data.html) in the _Application Discovery Service User Guide_.

- The polling period is in intervals of approximately 60 minutes.

- Data fields denoted with a double asterisk (\*\*) currently return a _null_ value.

Data fieldDescriptionapplicationConfigurationId\*ID of the migration application the VM is grouped under.avgCpuUsagePctAverage percentage of CPU usage over polling period.avgDiskBytesReadPerSecondAverage number of bytes read from disk over polling period.avgDiskBytesWrittenPerSecondAverage number of bytes written to disk over polling period.avgDiskReadOpsPerSecond\*\*Average number of read I/O operations per second null.avgDiskWriteOpsPerSecond\*\*Average number of write I/O operations per second.avgFreeRAMAverage free RAM expressed in MB.avgNetworkBytesReadPerSecondAverage amount of throughput of bytes read per second.avgNetworkBytesWrittenPerSecondAverage amount of throughput of bytes written per second.computerManufacturerVendor reported by the ESXi host.computerModelComputer model reported by the ESXi host.configIdID assigned by Application Discovery Service to the discovered VM.configTypeType of resource discovered.connectorIdID of the virtual appliance.cpuTypevCPU for a VM, actual model for a host.datacenterIdID of the vCenter.hostId\*ID of the VM host.hostNameName of host running the virtualization software.hypervisorType of hypervisor.idID of server.lastModifiedTimeStamp\*Latest date and time of data collection before data export.macAddressMAC address of the VM.manufacturerMaker of the virtualization software.maxCpuUsagePct Max. percentage of CPU usage during polling period.maxDiskBytesReadPerSecondMax. number of bytes read from disk over polling period.maxDiskBytesWrittenPerSecondMax. number of bytes written to disk over polling period.maxDiskReadOpsPerSecond\*\*Max. number of read I/O operations per second.maxDiskWriteOpsPerSecond\*\*Max. number of write I/O operations per second.maxNetworkBytesReadPerSecondMax. amount of throughput of bytes read per second.maxNetworkBytesWrittenPerSecondMax. amount of throughput of bytes written per second.memoryReservation\*Limit to avoid overcommitment of memory on VM.moRefIdUnique vCenter Managed Object Reference ID.name\*Name of VM or network (user specified).numCoresNumber of CPU cores assigned to VM.numCpusNumber of CPU sockets on the ESXi host.numDisks\*\*Number of disks on VM.numNetworkCards\*\*Number of network cards on VM.osNameOperating system name on VM.osVersionOperating system version on VM.portGroupId\*ID of group of member ports of VLAN.portGroupName\*Name of group of member ports of VLAN.powerState\*Status of power.serverIdApplication Discovery Service assigned ID to the discovered VM.smBiosId\*ID/version of the system management BIOS.state\*Status of the virtual appliance.toolsStatusOperational state of VMware toolstotalDiskFreeSizeFree disk space expressed in MB. Available for vCenter Server 7.0 and later
versions.totalDiskSizeTotal capacity of disk expressed in MB.totalRAMTotal amount of RAM available on VM in MB.typeType of host.vCenterIdUnique ID number of a VM.vCenterName\*Name of the vCenter host.virtualSwitchName\*Name of the virtual switch.vmFolderPathDirectory path of VM files.vmNameName of the virtual machine.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Controlling data collection scope

Using the
database and analytics data collection module
