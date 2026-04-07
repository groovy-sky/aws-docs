# GetInstanceTypesFromInstanceRequirements

Returns a list of instance types with the specified instance attributes. You can
use the response to preview the instance types without launching instances. Note
that the response does not consider capacity.

When you specify multiple parameters, you get instance types that satisfy all of the
specified parameters. If you specify multiple values for a parameter, you get instance
types that satisfy any of the specified values.

For more information, see [Preview instance types with specified attributes](../../../../services/ec2/latest/userguide/ec2-fleet-attribute-based-instance-type-selection.md#ec2fleet-get-instance-types-from-instance-requirements), [Specify attributes for instance type selection for EC2 Fleet or Spot Fleet](../../../../services/ec2/latest/userguide/ec2-fleet-attribute-based-instance-type-selection.md), and [Spot\
placement score](../../../../services/ec2/latest/userguide/spot-placement-score.md) in the _Amazon EC2 User Guide_, and [Creating\
mixed instance groups using attribute-based instance type selection](../../../../services/autoscaling/ec2/userguide/create-asg-instance-type-requirements.md) in the
_Amazon EC2 Auto Scaling User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ArchitectureType.N**

The processor architecture type.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 3 items.

Valid Values: `i386 | x86_64 | arm64 | x86_64_mac | arm64_mac`

Required: Yes

**Context**

Reserved.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceRequirements**

The attributes required for the instance types.

Type: [InstanceRequirementsRequest](api-instancerequirementsrequest.md) object

Required: Yes

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**VirtualizationType.N**

The virtualization type.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 2 items.

Valid Values: `hvm | paravirtual`

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceTypeSet**

The instance types with the specified instance attributes.

Type: Array of [InstanceTypeInfoFromInstanceRequirements](api-instancetypeinfofrominstancerequirements.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetInstanceTypesFromInstanceRequirements)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetInstanceTypesFromInstanceRequirements)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetInstanceTypesFromInstanceRequirements)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getinstancetypesfrominstancerequirements.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetInstanceTpmEkPub

GetInstanceUefiData
