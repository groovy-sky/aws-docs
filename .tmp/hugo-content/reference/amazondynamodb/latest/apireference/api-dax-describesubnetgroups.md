# DescribeSubnetGroups

Returns a list of subnet group descriptions. If a subnet group name is specified,
the list will contain only the description of that group.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "SubnetGroupNames": [ "string" ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[MaxResults](#API_dax_DescribeSubnetGroups_RequestSyntax)**

The maximum number of results to include in the response. If more results exist
than the specified `MaxResults` value, a token is included in the response so
that the remaining results can be retrieved.

The value for `MaxResults` must be between 20 and 100.

Type: Integer

Required: No

**[NextToken](#API_dax_DescribeSubnetGroups_RequestSyntax)**

An optional token returned from a prior request. Use this token for pagination of
results from this action. If this parameter is specified, the response includes only
results beyond the token, up to the value specified by
`MaxResults`.

Type: String

Required: No

**[SubnetGroupNames](#API_dax_DescribeSubnetGroups_RequestSyntax)**

The name of the subnet group.

Type: Array of strings

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "SubnetGroups": [
      {
         "Description": "string",
         "SubnetGroupName": "string",
         "Subnets": [
            {
               "SubnetAvailabilityZone": "string",
               "SubnetIdentifier": "string",
               "SupportedNetworkTypes": [ "string" ]
            }
         ],
         "SupportedNetworkTypes": [ "string" ],
         "VpcId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_dax_DescribeSubnetGroups_ResponseSyntax)**

Provides an identifier to allow retrieval of paginated results.

Type: String

**[SubnetGroups](#API_dax_DescribeSubnetGroups_ResponseSyntax)**

An array of subnet groups. Each element in the array represents a single subnet
group.

Type: Array of [SubnetGroup](api-dax-subnetgroup.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**SubnetGroupNotFoundFault**

The requested subnet group name does not refer to an existing subnet
group.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/describesubnetgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/describesubnetgroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeParameters

IncreaseReplicationFactor

All content copied from https://docs.aws.amazon.com/.
