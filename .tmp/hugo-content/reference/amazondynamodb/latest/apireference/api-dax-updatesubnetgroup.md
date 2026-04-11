# UpdateSubnetGroup

Modifies an existing subnet group.

## Request Syntax

```nohighlight

{
   "Description": "string",
   "SubnetGroupName": "string",
   "SubnetIds": [ "string" ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[SubnetGroupName](#API_dax_UpdateSubnetGroup_RequestSyntax)**

The name of the subnet group.

Type: String

Required: Yes

**[Description](#API_dax_UpdateSubnetGroup_RequestSyntax)**

A description of the subnet group.

Type: String

Required: No

**[SubnetIds](#API_dax_UpdateSubnetGroup_RequestSyntax)**

A list of subnet IDs in the subnet group.

Type: Array of strings

Required: No

## Response Syntax

```nohighlight

{
   "SubnetGroup": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[SubnetGroup](#API_dax_UpdateSubnetGroup_ResponseSyntax)**

The subnet group that has been modified.

Type: [SubnetGroup](api-dax-subnetgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidSubnet**

An invalid subnet identifier was specified.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**SubnetGroupNotFoundFault**

The requested subnet group name does not refer to an existing subnet
group.

HTTP Status Code: 400

**SubnetInUse**

The requested subnet is being used by another subnet group.

HTTP Status Code: 400

**SubnetNotAllowedFault**

The specified subnet can't be used for the requested network type. This error
occurs when either there aren't enough subnets of the required network type to create
the cluster, or when you try to use a subnet that doesn't support the requested network
type (for example, trying to create a dual-stack cluster with a subnet that doesn't have
IPv6 CIDR).

HTTP Status Code: 400

**SubnetQuotaExceededFault**

The request cannot be processed because it would exceed the allowed number of
subnets in a subnet group.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/updatesubnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/updatesubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateParameterGroup

Amazon DynamoDB Streams

All content copied from https://docs.aws.amazon.com/.
