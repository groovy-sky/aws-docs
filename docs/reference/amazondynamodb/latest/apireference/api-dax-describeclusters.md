# DescribeClusters

Returns information about all provisioned DAX clusters if no cluster identifier is
specified, or about a specific DAX cluster if a cluster identifier is
supplied.

If the cluster is in the CREATING state, only cluster level information will be
displayed until all of the nodes are successfully provisioned.

If the cluster is in the DELETING state, only cluster level information will be
displayed.

If nodes are currently being added to the DAX cluster, node endpoint information
and creation time for the additional nodes will not be displayed until they are
completely provisioned. When the DAX cluster state is
_available_, the cluster is ready for use.

If nodes are currently being removed from the DAX cluster, no
endpoint information for the removed nodes is displayed.

## Request Syntax

```nohighlight

{
   "ClusterNames": [ "string" ],
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ClusterNames](#API_dax_DescribeClusters_RequestSyntax)**

The names of the DAX clusters being described.

Type: Array of strings

Required: No

**[MaxResults](#API_dax_DescribeClusters_RequestSyntax)**

The maximum number of results to include in the response. If more results exist
than the specified `MaxResults` value, a token is included in the response so
that the remaining results can be retrieved.

The value for `MaxResults` must be between 20 and 100.

Type: Integer

Required: No

**[NextToken](#API_dax_DescribeClusters_RequestSyntax)**

An optional token returned from a prior request. Use this token for pagination of
results from this action. If this parameter is specified, the response includes only
results beyond the token, up to the value specified by
`MaxResults`.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "Clusters": [
      {
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
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Clusters](#API_dax_DescribeClusters_ResponseSyntax)**

The descriptions of your DAX clusters, in response to a
_DescribeClusters_ request.

Type: Array of [Cluster](api-dax-cluster.md) objects

**[NextToken](#API_dax_DescribeClusters_ResponseSyntax)**

Provides an identifier to allow retrieval of paginated results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ClusterNotFoundFault**

The requested cluster ID does not refer to an existing DAX
cluster.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/describeclusters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/describeclusters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/describeclusters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/describeclusters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/describeclusters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/describeclusters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/describeclusters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/describeclusters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/describeclusters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/describeclusters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteSubnetGroup

DescribeDefaultParameters

All content copied from https://docs.aws.amazon.com/.
