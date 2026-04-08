# DescribeInstanceTypeOfferings

Lists the instance types that are offered for the specified location. If no location is
specified, the default is to list the instance types that are offered in the current
Region.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. Filter names and values are case-sensitive.

- `instance-type` \- The instance type. For a list of possible values, see [Instance](api-instance.md).

- `location` \- The location. For a list of possible identifiers, see [Regions and Zones](../../../../services/ec2/latest/userguide/using-regions-availability-zones.md).

Type: Array of [Filter](api-filter.md) objects

Required: No

**LocationType**

The location type.

- `availability-zone` \- The Availability Zone. When you specify a location
filter, it must be an Availability Zone for the current Region.

- `availability-zone-id` \- The AZ ID. When you specify a location filter, it must
be an AZ ID for the current Region.

- `outpost` \- The Outpost ARN. When you specify a location filter, it must be an
Outpost ARN for the current Region.

- `region` \- The current Region. If you specify a location filter, it must match
the current Region.

Type: String

Valid Values: `region | availability-zone | availability-zone-id | outpost`

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**instanceTypeOfferingSet**

The instance types offered in the location.

Type: Array of [InstanceTypeOffering](api-instancetypeoffering.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeinstancetypeofferings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeinstancetypeofferings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeInstanceTopology

DescribeInstanceTypes
