# Cluster

Contains all of the attributes of a specific DAX cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**ActiveNodes**

The number of nodes in the cluster that are active (i.e., capable of serving
requests).

Type: Integer

Required: No

**ClusterArn**

The Amazon Resource Name (ARN) that uniquely identifies the cluster.

Type: String

Required: No

**ClusterDiscoveryEndpoint**

The endpoint for this DAX cluster, consisting of a DNS name, a port
number, and a URL. Applications should use the URL to configure the DAX
client to find their cluster.

Type: [Endpoint](api-dax-endpoint.md) object

Required: No

**ClusterEndpointEncryptionType**

The type of encryption supported by the cluster's endpoint. Values are:

- `NONE` for no encryption

`TLS` for Transport Layer Security

Type: String

Valid Values: `NONE | TLS`

Required: No

**ClusterName**

The name of the DAX cluster.

Type: String

Required: No

**Description**

The description of the cluster.

Type: String

Required: No

**IamRoleArn**

A valid Amazon Resource Name (ARN) that identifies an IAM role. At
runtime, DAX will assume this role and use the role's permissions to
access DynamoDB on your behalf.

Type: String

Required: No

**NetworkType**

The IP address type of the cluster. Values are:

- `ipv4` \- IPv4 addresses only

- `ipv6` \- IPv6 addresses only

- `dual_stack` \- Both IPv4 and IPv6 addresses

Type: String

Valid Values: `ipv4 | ipv6 | dual_stack`

Required: No

**NodeIdsToRemove**

A list of nodes to be removed from the cluster.

Type: Array of strings

Required: No

**Nodes**

A list of nodes that are currently in the cluster.

Type: Array of [Node](api-dax-node.md) objects

Required: No

**NodeType**

The node type for the nodes in the cluster. (All nodes in a DAX cluster are of the
same type.)

Type: String

Required: No

**NotificationConfiguration**

Describes a notification topic and its status. Notification topics are used for
publishing DAX events to subscribers using Amazon Simple Notification Service
(SNS).

Type: [NotificationConfiguration](api-dax-notificationconfiguration.md) object

Required: No

**ParameterGroup**

The parameter group being used by nodes in the cluster.

Type: [ParameterGroupStatus](api-dax-parametergroupstatus.md) object

Required: No

**PreferredMaintenanceWindow**

A range of time when maintenance of DAX cluster software will be performed. For
example: `sun:01:00-sun:09:00`. Cluster maintenance normally takes less than
30 minutes, and is performed automatically within the maintenance window.

Type: String

Required: No

**SecurityGroups**

A list of security groups, and the status of each, for the nodes in the
cluster.

Type: Array of [SecurityGroupMembership](api-dax-securitygroupmembership.md) objects

Required: No

**SSEDescription**

The description of the server-side encryption status on the specified DAX cluster.

Type: [SSEDescription](api-dax-ssedescription.md) object

Required: No

**Status**

The current status of the cluster.

Type: String

Required: No

**SubnetGroup**

The subnet group where the DAX cluster is running.

Type: String

Required: No

**TotalNodes**

The total number of nodes in the cluster.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/cluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/cluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/cluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDB Accelerator

Endpoint

All content copied from https://docs.aws.amazon.com/.
