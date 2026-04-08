# ModifyDBClusterEndpoint

Modifies the properties of an endpoint in an Amazon Aurora DB cluster.

###### Note

This operation only applies to Aurora DB clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterEndpointIdentifier**

The identifier of the endpoint to modify. This parameter is stored as a lowercase string.

Type: String

Required: Yes

**EndpointType**

The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.

Type: String

Required: No

**ExcludedMembers.member.N**

List of DB instance identifiers that aren't part of the custom endpoint group.
All other eligible instances are reachable through the custom endpoint.
Only relevant if the list of static members is empty.

Type: Array of strings

Required: No

**StaticMembers.member.N**

List of DB instance identifiers that are part of the custom endpoint group.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**CustomEndpointType**

The type associated with a custom endpoint. One of: `READER`,
`WRITER`, `ANY`.

Type: String

**DBClusterEndpointArn**

The Amazon Resource Name (ARN) for the endpoint.

Type: String

**DBClusterEndpointIdentifier**

The identifier associated with the endpoint. This parameter is stored as a lowercase string.

Type: String

**DBClusterEndpointResourceIdentifier**

A unique system-generated identifier for an endpoint. It remains the same for the whole life of the endpoint.

Type: String

**DBClusterIdentifier**

The DB cluster identifier of the DB cluster associated with the endpoint. This parameter is
stored as a lowercase string.

Type: String

**Endpoint**

The DNS address of the endpoint.

Type: String

**EndpointType**

The type of the endpoint. One of: `READER`, `WRITER`, `CUSTOM`.

Type: String

**ExcludedMembers.member.N**

List of DB instance identifiers that aren't part of the custom endpoint group.
All other eligible instances are reachable through the custom endpoint.
Only relevant if the list of static members is empty.

Type: Array of strings

**StaticMembers.member.N**

List of DB instance identifiers that are part of the custom endpoint group.

Type: Array of strings

**Status**

The current status of the endpoint. One of: `creating`, `available`, `deleting`, `inactive`, `modifying`. The `inactive` state applies to an endpoint that can't be used for a certain kind of cluster,
such as a `writer` endpoint for a read-only secondary cluster in a global database.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterEndpointNotFoundFault**

The specified custom endpoint doesn't exist.

HTTP Status Code: 400

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBClusterEndpointStateFault**

The requested operation can't be performed on the endpoint while the endpoint is in this state.

HTTP Status Code: 400

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbclusterendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbclusterendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBCluster

ModifyDBClusterParameterGroup

All content copied from https://docs.aws.amazon.com/.
