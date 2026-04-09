# CreateCluster

Creates a DAX cluster. All nodes in the cluster run the same DAX caching software.

## Request Syntax

```nohighlight

{
   "AvailabilityZones": [ "string" ],
   "ClusterEndpointEncryptionType": "string",
   "ClusterName": "string",
   "Description": "string",
   "IamRoleArn": "string",
   "NetworkType": "string",
   "NodeType": "string",
   "NotificationTopicArn": "string",
   "ParameterGroupName": "string",
   "PreferredMaintenanceWindow": "string",
   "ReplicationFactor": number,
   "SecurityGroupIds": [ "string" ],
   "SSESpecification": {
      "Enabled": boolean
   },
   "SubnetGroupName": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ClusterName](#API_dax_CreateCluster_RequestSyntax)**

The cluster identifier. This parameter is stored as a lowercase string.

**Constraints:**

- A name must contain from 1 to 20 alphanumeric characters or
hyphens.

- The first character must be a letter.

- A name cannot end with a hyphen or contain two consecutive
hyphens.

Type: String

Required: Yes

**[IamRoleArn](#API_dax_CreateCluster_RequestSyntax)**

A valid Amazon Resource Name (ARN) that identifies an IAM role. At
runtime, DAX will assume this role and use the role's permissions to
access DynamoDB on your behalf.

Type: String

Required: Yes

**[NodeType](#API_dax_CreateCluster_RequestSyntax)**

The compute and memory capacity of the nodes in the cluster.

Type: String

Required: Yes

**[ReplicationFactor](#API_dax_CreateCluster_RequestSyntax)**

The number of nodes in the DAX cluster. A replication factor of 1
will create a single-node cluster, without any read replicas. For additional fault
tolerance, you can create a multiple node cluster with one or more read replicas. To do
this, set `ReplicationFactor` to a number between 3 (one primary and two read
replicas) and 10 (one primary and nine read replicas). `If the
                AvailabilityZones` parameter is provided, its length must equal the
`ReplicationFactor`.

###### Note

AWS recommends that you have at least two read replicas per
cluster.

Type: Integer

Required: Yes

**[AvailabilityZones](#API_dax_CreateCluster_RequestSyntax)**

The Availability Zones (AZs) in which the cluster nodes will reside after the
cluster has been created or updated. If provided, the length of this list must equal the
`ReplicationFactor` parameter. If you omit this parameter, DAX will spread the nodes across Availability Zones for the highest
availability.

Type: Array of strings

Required: No

**[ClusterEndpointEncryptionType](#API_dax_CreateCluster_RequestSyntax)**

The type of encryption the cluster's endpoint should support. Values are:

- `NONE` for no encryption

- `TLS` for Transport Layer Security

Type: String

Valid Values: `NONE | TLS`

Required: No

**[Description](#API_dax_CreateCluster_RequestSyntax)**

A description of the cluster.

Type: String

Required: No

**[NetworkType](#API_dax_CreateCluster_RequestSyntax)**

Specifies the IP protocol(s) the cluster uses for network communications. Values
are:

- `ipv4` \- The cluster is accessible only through IPv4
addresses

- `ipv6` \- The cluster is accessible only through IPv6
addresses

- `dual_stack` \- The cluster is accessible through both IPv4 and
IPv6 addresses.

###### Note

If no explicit `NetworkType` is provided, the network type is
derived based on the subnet group's configuration.

Type: String

Valid Values: `ipv4 | ipv6 | dual_stack`

Required: No

**[NotificationTopicArn](#API_dax_CreateCluster_RequestSyntax)**

The Amazon Resource Name (ARN) of the Amazon SNS topic to which
notifications will be sent.

###### Note

The Amazon SNS topic owner must be same as the DAX
cluster owner.

Type: String

Required: No

**[ParameterGroupName](#API_dax_CreateCluster_RequestSyntax)**

The parameter group to be associated with the DAX cluster.

Type: String

Required: No

**[PreferredMaintenanceWindow](#API_dax_CreateCluster_RequestSyntax)**

Specifies the weekly time range during which maintenance on the DAX cluster is
performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock
UTC). The minimum maintenance window is a 60 minute period. Valid values for
`ddd` are:

- `sun`

- `mon`

- `tue`

- `wed`

- `thu`

- `fri`

- `sat`

Example: `sun:05:00-sun:09:00`

###### Note

If you don't specify a preferred maintenance window when you create or modify a
cache cluster, DAX assigns a 60-minute maintenance window on a
randomly selected day of the week.

Type: String

Required: No

**[SecurityGroupIds](#API_dax_CreateCluster_RequestSyntax)**

A list of security group IDs to be assigned to each node in the DAX
cluster. (Each of the security group ID is system-generated.)

If this parameter is not specified, DAX assigns the default VPC
security group to each node.

Type: Array of strings

Required: No

**[SSESpecification](#API_dax_CreateCluster_RequestSyntax)**

Represents the settings used to enable server-side encryption on the
cluster.

Type: [SSESpecification](api-dax-ssespecification.md) object

Required: No

**[SubnetGroupName](#API_dax_CreateCluster_RequestSyntax)**

The name of the subnet group to be used for the replication group.

###### Important

DAX clusters can only run in an Amazon VPC environment.
All of the subnets that you specify in a subnet group must exist in the same
VPC.

Type: String

Required: No

**[Tags](#API_dax_CreateCluster_RequestSyntax)**

A set of tags to associate with the DAX cluster.

Type: Array of [Tag](api-dax-tag.md) objects

Required: No

## Response Syntax

```nohighlight

{
   "Cluster": {
      "ActiveNodes": number,
      "ClusterArn": "string",
      "ClusterDiscoveryEndpoint": {
         "Address": "string",
         "Port": number,
         "URL": "string"
      },
      "ClusterEndpointEncryptionType": "string",
      "ClusterName": "string",
      "Description": "string",
      "IamRoleArn": "string",
      "NetworkType": "string",
      "NodeIdsToRemove": [ "string" ],
      "Nodes": [
         {
            "AvailabilityZone": "string",
            "Endpoint": {
               "Address": "string",
               "Port": number,
               "URL": "string"
            },
            "NodeCreateTime": number,
            "NodeId": "string",
            "NodeStatus": "string",
            "ParameterGroupStatus": "string"
         }
      ],
      "NodeType": "string",
      "NotificationConfiguration": {
         "TopicArn": "string",
         "TopicStatus": "string"
      },
      "ParameterGroup": {
         "NodeIdsToReboot": [ "string" ],
         "ParameterApplyStatus": "string",
         "ParameterGroupName": "string"
      },
      "PreferredMaintenanceWindow": "string",
      "SecurityGroups": [
         {
            "SecurityGroupIdentifier": "string",
            "Status": "string"
         }
      ],
      "SSEDescription": {
         "Status": "string"
      },
      "Status": "string",
      "SubnetGroup": "string",
      "TotalNodes": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Cluster](#API_dax_CreateCluster_ResponseSyntax)**

A description of the DAX cluster that you have created.

Type: [Cluster](api-dax-cluster.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ClusterAlreadyExistsFault**

You already have a DAX cluster with the given identifier.

HTTP Status Code: 400

**ClusterQuotaForCustomerExceededFault**

You have attempted to exceed the maximum number of DAX clusters for
your AWS account.

HTTP Status Code: 400

**InsufficientClusterCapacityFault**

There are not enough system resources to create the cluster you requested (or to
resize an already-existing cluster).

HTTP Status Code: 400

**InvalidClusterStateFault**

The requested DAX cluster is not in the
_available_ state.

HTTP Status Code: 400

**InvalidParameterCombinationException**

Two or more incompatible parameters were specified.

HTTP Status Code: 400

**InvalidParameterGroupStateFault**

One or more parameters in a parameter group are in an invalid state.

HTTP Status Code: 400

**InvalidParameterValueException**

The value for a parameter is invalid.

HTTP Status Code: 400

**InvalidVPCNetworkStateFault**

The VPC network is in an invalid state.

HTTP Status Code: 400

**NodeQuotaForClusterExceededFault**

You have attempted to exceed the maximum number of nodes for a DAX
cluster.

HTTP Status Code: 400

**NodeQuotaForCustomerExceededFault**

You have attempted to exceed the maximum number of nodes for your AWS account.

HTTP Status Code: 400

**ParameterGroupNotFoundFault**

The specified parameter group does not exist.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**ServiceQuotaExceededException**

You have reached the maximum number of x509 certificates that can be created for
encrypted clusters in a 30 day period. Contact AWS customer support to
discuss options for continuing to create encrypted clusters.

HTTP Status Code: 400

**SubnetGroupNotFoundFault**

The requested subnet group name does not refer to an existing subnet
group.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

You have exceeded the maximum number of tags for this DAX cluster.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/createcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/createcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/createcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/createcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/createcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/createcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/createcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/createcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/createcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/createcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDB Accelerator

CreateParameterGroup

All content copied from https://docs.aws.amazon.com/.
