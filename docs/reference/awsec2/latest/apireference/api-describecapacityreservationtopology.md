# DescribeCapacityReservationTopology

Describes a tree-based hierarchy that represents the physical host placement of your
pending or active Capacity Reservations within an Availability Zone or Local Zone. You
can use this information to determine the relative proximity of your capacity within the
AWS network before it is launched and use this information to allocate capacity
together to support your tightly coupled workloads.

Capacity Reservation topology is supported for specific instance types only. For more
information, see [Prerequisites\
for Amazon EC2 instance topology](../../../../services/ec2/latest/userguide/ec2-instance-topology-prerequisites.md) in the
_Amazon EC2 User Guide_.

###### Note

The Amazon EC2 API follows an eventual consistency model due to the
distributed nature of the system supporting it. As a result, when you call the
DescribeCapacityReservationTopology API command immediately after launching
instances, the response might return a `null` value for
`capacityBlockId` because the data might not have fully propagated
across all subsystems. For more information, see [Eventual consistency in the\
Amazon EC2 API](../../../../services/ec2/latest/devguide/eventual-consistency.md) in the _Amazon EC2 Developer_
_Guide_.

For more information, see [Amazon EC2 topology](../../../../services/ec2/latest/userguide/ec2-instance-topology.md) in
the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId.N**

The Capacity Reservation IDs.

Default: Describes all your Capacity Reservations.

Constraints: Maximum 100 explicitly specified Capacity Reservation IDs.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `availability-zone` \- The name of the Availability Zone (for
example, `us-west-2a`) or Local Zone (for example,
`us-west-2-lax-1b`) that the Capacity Reservation is in.

- `instance-type` \- The instance type (for example,
`p4d.24xlarge`) or instance family (for example,
`p4d*`). You can use the `*` wildcard to match zero or
more characters, or the `?` wildcard to match zero or one
character.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

You can't specify this parameter and the Capacity Reservation IDs parameter in the
same request.

Default: `10`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 10.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityReservationSet**

Information about the topology of each Capacity Reservation.

Type: Array of [CapacityReservationTopology](api-capacityreservationtopology.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describecapacityreservationtopology.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacityreservationtopology.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCapacityReservations

DescribeCarrierGateways
