# GetSpotPlacementScores

Calculates the Spot placement score for a Region or Availability Zone based on the
specified target capacity and compute requirements.

You can specify your compute requirements either by using
`InstanceRequirementsWithMetadata` and letting Amazon EC2 choose the optimal
instance types to fulfill your Spot request, or you can specify the instance types by using
`InstanceTypes`.

For more information, see [Spot placement score](../../../../services/ec2/latest/userguide/spot-placement-score.md) in
the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceRequirementsWithMetadata**

The attributes for the instance types. When you specify instance attributes, Amazon EC2 will
identify instance types with those attributes.

If you specify `InstanceRequirementsWithMetadata`, you can't specify
`InstanceTypes`.

Type: [InstanceRequirementsWithMetadataRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceRequirementsWithMetadataRequest.html) object

Required: No

**InstanceType.N**

The instance types. We recommend that you specify at least three instance types. If you
specify one or two instance types, or specify variations of a single instance type (for
example, an `m3.xlarge` with and without instance storage), the returned
placement score will always be low.

If you specify `InstanceTypes`, you can't specify
`InstanceRequirementsWithMetadata`.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 1000 items.

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 10. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**RegionName.N**

The Regions used to narrow down the list of Regions to be scored. Enter the Region code,
for example, `us-east-1`.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

**SingleAvailabilityZone**

Specify `true` so that the response returns a list of scored Availability Zones.
Otherwise, the response returns a list of scored Regions.

A list of scored Availability Zones is useful if you want to launch all of your Spot
capacity into a single Availability Zone.

Type: Boolean

Required: No

**TargetCapacity**

The target capacity.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 2000000000.

Required: Yes

**TargetCapacityUnitType**

The unit for the target capacity.

Type: String

Valid Values: `vcpu | memory-mib | units`

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**spotPlacementScoreSet**

The Spot placement score for the top 10 Regions or Availability Zones, scored on a scale
from 1 to 10. Each score  reflects how likely it is that each Region or Availability Zone
will succeed at fulfilling the specified target capacity  _at the time of the Spot_
_placement score request_. A score of `10` means that your Spot
capacity request is highly likely to succeed in that Region or Availability Zone.

If you request a Spot placement score for Regions, a high score assumes that your fleet
request will be configured to use all Availability Zones and the
`capacity-optimized` allocation strategy. If you request a Spot placement
score for Availability Zones, a high score assumes that your fleet request will be
configured to use a single Availability Zone and the `capacity-optimized`
allocation strategy.

Different  Regions or Availability Zones might return the same score.

###### Note

The Spot placement score serves as a recommendation only. No score guarantees that your
Spot request will be fully or partially fulfilled.

Type: Array of [SpotPlacementScore](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotPlacementScore.html) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetSpotPlacementScores)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetSpotPlacementScores)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetSnapshotBlockPublicAccessState

GetSubnetCidrReservations
