# UpdateCluster

Modifies the settings for a DAX cluster. You can use this action to
change one or more cluster configuration parameters by specifying the parameters and the
new values.

## Request Syntax

```nohighlight

{
   "ClusterName": "string",
   "Description": "string",
   "NotificationTopicArn": "string",
   "NotificationTopicStatus": "string",
   "ParameterGroupName": "string",
   "PreferredMaintenanceWindow": "string",
   "SecurityGroupIds": [ "string" ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ClusterName](#API_dax_UpdateCluster_RequestSyntax)**

The name of the DAX cluster to be modified.

Type: String

Required: Yes

**[Description](#API_dax_UpdateCluster_RequestSyntax)**

A description of the changes being made to the cluster.

Type: String

Required: No

**[NotificationTopicArn](#API_dax_UpdateCluster_RequestSyntax)**

The Amazon Resource Name (ARN) that identifies the topic.

Type: String

Required: No

**[NotificationTopicStatus](#API_dax_UpdateCluster_RequestSyntax)**

The current state of the topic. A value of “active” means that notifications will
be sent to the topic. A value of “inactive” means that notifications will not be sent to
the topic.

Type: String

Required: No

**[ParameterGroupName](#API_dax_UpdateCluster_RequestSyntax)**

The name of a parameter group for this cluster.

Type: String

Required: No

**[PreferredMaintenanceWindow](#API_dax_UpdateCluster_RequestSyntax)**

A range of time when maintenance of DAX cluster software will be performed. For
example: `sun:01:00-sun:09:00`. Cluster maintenance normally takes less than
30 minutes, and is performed automatically within the maintenance window.

Type: String

Required: No

**[SecurityGroupIds](#API_dax_UpdateCluster_RequestSyntax)**

A list of user-specified security group IDs to be assigned to each node in the DAX
cluster. If this parameter is not specified, DAX assigns the default VPC security group
to each node.

Type: Array of strings

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

**[Cluster](#API_dax_UpdateCluster_ResponseSyntax)**

A description of the DAX cluster, after it has been modified.

Type: [Cluster](api-dax-cluster.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ClusterNotFoundFault**

The requested cluster ID does not refer to an existing DAX
cluster.

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

**ParameterGroupNotFoundFault**

The specified parameter group does not exist.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/updatecluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/updatecluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/updatecluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/updatecluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/updatecluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/updatecluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/updatecluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/updatecluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/updatecluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/updatecluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

UpdateParameterGroup

All content copied from https://docs.aws.amazon.com/.
