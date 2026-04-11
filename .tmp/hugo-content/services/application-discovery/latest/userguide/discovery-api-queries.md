AWS Application Discovery Service is no longer open to new customers. Alternatively, use AWS Transform which provides similar capabilities. For more information, see [AWS Application Discovery Service availability change](application-discovery-service-availability-change.md).

# Using the Application Discovery Service API to query discovered configuration items

A _configuration item_ is an IT asset that was discovered
in your data center by an agent or by an import. When you use
AWS Application Discovery Service (Application Discovery Service), you use the API to specify filters and query specific configuration
items for server, application, process, and connection assets. For information about the API,
see [Application Discovery Service API Reference](../../../../reference/application-discovery/latest/apireference/welcome.md).

The tables in the following sections list the available input filters and output sorting
options for two Application Discovery Service actions:

- `DescribeConfigurations`

- `ListConfigurations`

The filtering and sorting options are organized by the type of asset to which apply (server,
application, process, or connection).

###### Important

Results returned by `DescribeConfigurations`, `ListConfigurations`,
and `StartExportTask` might not contain recent updates. For more information, see
[Eventual consistency in the AWS Application Discovery Service API](#eventual-consistency).

## Using the `DescribeConfigurations` action

The `DescribeConfigurations` action retrieves attributes for a list of
configuration IDs. All the supplied IDs must be for the same asset type (server, application,
process, or connection). Output fields are specific to the asset type selected. For example,
the output for a server configuration item includes a list of attributes about the server,
such as host name, operating system, and number of network cards. For more information about
command syntax, see [DescribeConfigurations](../../../../reference/application-discovery/latest/apireference/api-describeconfigurations.md).

The `DescribeConfigurations` action does not support filtering.

###### Output fields for `DescribeConfigurations`

The following tables, organized by asset type, list the supported output fields of the
`DescribeConfigurations` action. The ones marked as mandatory are always present
in the output.

**Server assets**

FieldMandatory`server.agentId``server.applications``server.applications.hasMoreValues``server.configurationId`x`server.cpuType``server.hostName``server.hypervisor``server.networkInterfaceInfo``server.networkInterfaceInfo.hasMoreValues``server.osName``server.osVersion``server.tags``server.tags.hasMoreValues``server.timeOfCreation`x`server.type``server.performance.avgCpuUsagePct``server.performance.avgDiskReadIOPS``server.performance.avgDiskReadsPerSecondInKB``server.performance.avgDiskWriteIOPS``server.performance.avgDiskWritesPerSecondInKB``server.performance.avgFreeRAMInKB``server.performance.avgNetworkReadsPerSecondInKB``server.performance.avgNetworkWritesPerSecondInKB``server.performance.maxCpuUsagePct``server.performance.maxDiskReadIOPS``server.performance.maxDiskReadsPerSecondInKB``server.performance.maxDiskWriteIOPS``server.performance.maxDiskWritesPerSecondInKB``server.performance.maxNetworkReadsPerSecondInKB``server.performance.maxNetworkWritesPerSecondInKB``server.performance.minFreeRAMInKB``server.performance.numCores``server.performance.numCpus``server.performance.numDisks``server.performance.numNetworkCards``server.performance.totalRAMInKB`

**Process assets**

FieldMandatory`process.commandLine``process.configurationId`x`process.name``process.path``process.timeOfCreation`x

**Application assets**

FieldMandatory`application.configurationId`x`application.description``application.lastModifiedTime`x`application.name`x`application.serverCount`x`application.timeOfCreation`x

## Using the `ListConfigurations` action

The `ListConfigurations` action retrieves a list of configuration items
according to the criteria that you specify in a filter. For more information about command
syntax, see [ListConfigurations](../../../../reference/application-discovery/latest/apireference/api-listconfigurations.md).

###### Output fields for `ListConfigurations`

The following tables, organized by asset type, list the supported output fields of the
`ListConfigurations` action. The ones marked as mandatory are always present in
the output.

**Server assets**

FieldMandatory`server.configurationId`x`server.agentId``server.hostName``server.osName``server.osVersion``server.timeOfCreation`x`server.type`

**Process assets**

FieldMandatory`process.commandLine``process.configurationId`x`process.name``process.path``process.timeOfCreation`x`server.agentId``server.configurationId`x

**Application assets**

FieldMandatory`application.configurationId`x`application.description``application.name`x`application.serverCount`x`application.timeOfCreation`x`application.lastModifiedTime`x

**Connection assets**

FieldMandatory`connection.destinationIp`x`connection.destinationPort`x`connection.ipVersion`x`connection.latestTimestamp`x`connection.occurrence`x`connection.sourceIp`x`connection.transportProtocol``destinationProcess.configurationId``destinationProcess.name``destinationServer.configurationId``destinationServer.hostName``sourceProcess.configurationId``sourceProcess.name``sourceServer.configurationId``sourceServer.hostName`

###### Supported filters for `ListConfigurations`

The following tables, organized by asset type, list the supported filters for the
`ListConfigurations` action. Filters and values are in a key/value relationship
defined by one of the supported logical conditions. You can sort the output of the indicated
filters.

**Server assets**

Filter

Supported conditions

Supported values

Supported sorting

`server.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- Any valid server configuration ID

None`server.hostName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.osName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.osVersion`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.agentId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- String

None`server.connectorId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- String

None`server.type`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

String with one of the following values:

- EC2

- OTHER

- VMWARE\_VM

- VMWARE\_HOST

- VMWARE\_VM\_TEMPLATE

None`server.vmWareInfo.morefId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.vmWareInfo.vcenterId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.vmWareInfo.hostId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.networkInterfaceInfo.portGroupId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.networkInterfaceInfo.portGroupName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.networkInterfaceInfo.virtualSwitchName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.networkInterfaceInfo.ipAddress`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.networkInterfaceInfo.macAddress`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.performance.avgCpuUsagePct`

- GE

- LE

- GT

- LT

- Percentage

None`server.performance.totalDiskFreeSizeInKB`

- GE

- LE

- GT

- LT

- Double

None`server.performance.avgFreeRAMInKB`

- GE

- LE

- GT

- LT

- Double

None`server.tag.value`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.tag.key`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.application.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.application.description`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None`server.application.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- Any valid application configuration ID

None

`server.process.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ProcessId

None

`server.process.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None

`server.process.commandLine`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

None

**Application assets**

Filter

Supported conditions

Supported values

Supported sorting

`application.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ApplicationId

None`application.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`application.description`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`application.serverCount`Filtering not supported.Filtering not supported.

- ASC

- DESC

`application.timeOfCreation`Filtering not supported.Filtering not supported.

- ASC

- DESC

`application.lastModifiedTime`Filtering not supported.Filtering not supported.

- ASC

- DESC

`server.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ServerId

None

**Process assets**

Filter

Supported conditions

Supported values

Supported sorting

`process.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ProcessId

`process.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`process.commandLine`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ServerId

`server.hostName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.osName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.osVersion`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`server.agentId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

**Connection assets**

Filter

Supported conditions

Supported values

Supported sorting

`connection.sourceIp`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- IP

- ASC

- DESC

`connection.destinationIp`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- IP

- ASC

- DESC

`connection.destinationPort`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- Integer

- ASC

- DESC

`sourceServer.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ServerId

`sourceServer.hostName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`destinationServer.osName`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`destinationServer.osVersion`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`destinationServer.agentId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

`sourceProcess.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ProcessId

`sourceProcess.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`sourceProcess.commandLine`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`destinationProcess.configurationId`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- ProcessId

`destinationProcess.name`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

`destinationprocess.commandLine`

- EQUALS

- NOT\_EQUALS

- EQ

- NE

- CONTAINS

- NOT\_CONTAINS

- String

- ASC

- DESC

## Eventual consistency in the AWS Application Discovery Service API

The following update operations are eventually consistent. Updates might not be
immediately visible to the read operations [StartExportTask](../../../../reference/application-discovery/latest/apireference/api-startexporttask.md), [DescribeConfigurations](../../../../reference/application-discovery/latest/apireference/api-describeconfigurations.md), and [ListConfigurations](../../../../reference/application-discovery/latest/apireference/api-listconfigurations.md).

- [AssociateConfigurationItemsToApplication](../../../../reference/application-discovery/latest/apireference/api-associateconfigurationitemstoapplication.md)

- [CreateTags](../../../../reference/application-discovery/latest/apireference/api-createtags.md)

- [DeleteApplications](../../../../reference/application-discovery/latest/apireference/api-deleteapplications.md)

- [DeleteTags](../../../../reference/application-discovery/latest/apireference/api-deletetags.md)

- [DescribeBatchDeleteConfigurationTask](../../../../reference/application-discovery/latest/apireference/api-describebatchdeleteconfigurationtask.md)

- [DescribeImportTasks](../../../../reference/application-discovery/latest/apireference/api-describeimporttasks.md)

- [DisassociateConfigurationItemsFromApplication](../../../../reference/application-discovery/latest/apireference/api-disassociateconfigurationitemsfromapplication.md)

- [UpdateApplication](../../../../reference/application-discovery/latest/apireference/api-updateapplication.md)

Suggestions for managing eventual consistency:

- When you invoke the read operations [StartExportTask](../../../../reference/application-discovery/latest/apireference/api-startexporttask.md), [DescribeConfigurations](../../../../reference/application-discovery/latest/apireference/api-describeconfigurations.md), or [ListConfigurations](../../../../reference/application-discovery/latest/apireference/api-listconfigurations.md) (or their corresponding AWS CLI commands), use an
exponential backoff algorithm to allow enough time for any previous update operation to
propagate through the system. To do this, run the read operation repeatedly, starting
with a two-second wait time, and increasing gradually up to five minutes of wait
time.

- Add wait time between subsequent operations, even if an update operation returns a
200 - OK response. Apply an exponential backoff algorithm starting with a couple of
seconds of wait time, and increase gradually up to about five minutes of wait
time.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Grouping servers

AWS PrivateLink

All content copied from https://docs.aws.amazon.com/.
