# DescribeImageReferences

Describes your AWS resources that are referencing the specified images.

For more information, see [Identify your resources referencing\
specified AMIs](../../../../services/ec2/latest/userguide/ec2-ami-references.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageId.N**

The IDs of the images to check for resource references.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: Yes

**IncludeAllResourceTypes**

Specifies whether to check all supported AWS resource types for image references. When
specified, default values are applied for `ResourceTypeOptions`. For the default
values, see [How AMI\
reference checks work](../../../../services/ec2/latest/userguide/ec2-ami-references.md#how-ami-references-works) in the _Amazon EC2 User Guide_. If you also
specify `ResourceTypes` with `ResourceTypeOptions`, your specified
values override the default values.

Supported resource types: `ec2:Instance` \| `ec2:LaunchTemplate` \|
`ssm:Parameter` \| `imagebuilder:ImageRecipe` \|
`imagebuilder:ContainerRecipe`

Either `IncludeAllResourceTypes` or `ResourceTypes` must be
specified.

Type: Boolean

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**ResourceType.N**

The AWS resource types to check for image references.

Either `IncludeAllResourceTypes` or `ResourceTypes` must be
specified.

Type: Array of [ResourceTypeRequest](api-resourcetyperequest.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**imageReferenceSet**

The resources that are referencing the specified images.

Type: Array of [ImageReference](api-imagereference.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeimagereferences.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeimagereferences.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeImageAttribute

DescribeImages
