# DeleteCluster

Deletes a previously provisioned DAX cluster.
_DeleteCluster_ deletes all associated nodes, node endpoints and
the DAX cluster itself. When you receive a successful response from this
action, DAX immediately begins deleting the cluster; you cannot cancel or
revert this action.

## Request Syntax

```nohighlight

{
   "ClusterName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ClusterName](#API_dax_DeleteCluster_RequestSyntax)**

The name of the cluster to be deleted.

Type: String

Required: Yes

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

**[Cluster](#API_dax_DeleteCluster_ResponseSyntax)**

A description of the DAX cluster that is being deleted.

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

**InvalidParameterValueException**

The value for a parameter is invalid.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/deletecluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/deletecluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/deletecluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/deletecluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/deletecluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/deletecluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/deletecluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/deletecluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/deletecluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/deletecluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecreaseReplicationFactor

DeleteParameterGroup

All content copied from https://docs.aws.amazon.com/.
