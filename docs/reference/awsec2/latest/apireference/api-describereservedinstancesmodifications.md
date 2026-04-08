# DescribeReservedInstancesModifications

Describes the modifications made to your Reserved Instances. If no parameter is specified,
information about all your Reserved Instances modification requests is returned. If a
modification ID is specified, only information about the specific modification is
returned.

For more information, see [Modify Reserved Instances](../../../../services/ec2/latest/userguide/ri-modifying.md) in the
_Amazon EC2 User Guide_.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Filter.N**

One or more filters.

- `client-token` \- The idempotency token for the modification request.

- `create-date` \- The time when the modification request was created.

- `effective-date` \- The time when the modification becomes effective.

- `modification-result.reserved-instances-id` \- The ID for the Reserved Instances
created as part of the modification request. This ID is only available when the status of
the modification is `fulfilled`.

- `modification-result.target-configuration.availability-zone` \- The Availability
Zone for the new Reserved Instances.

- `modification-result.target-configuration.availability-zone-id` \- The ID of the
Availability Zone for the new Reserved Instances.

- `modification-result.target-configuration.instance-count ` \- The number of new
Reserved Instances.

- `modification-result.target-configuration.instance-type` \- The instance type of
the new Reserved Instances.

- `reserved-instances-id` \- The ID of the Reserved Instances modified.

- `reserved-instances-modification-id` \- The ID of the modification
request.

- `status` \- The status of the Reserved Instances modification request
( `processing` \| `fulfilled` \| `failed`).

- `status-message` \- The reason for the status.

- `update-date` \- The time when the modification request was last updated.

Type: Array of [Filter](api-filter.md) objects

Required: No

**NextToken**

The token to retrieve the next page of results.

Type: String

Required: No

**ReservedInstancesModificationId.N**

IDs for the submitted modification request.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null`
when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**reservedInstancesModificationsSet**

The Reserved Instance modification information.

Type: Array of [ReservedInstancesModification](api-reservedinstancesmodification.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example illustrates one usage of DescribeReservedInstancesModifications.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeReservedInstancesModifications
&AUTHPARAMS
```

### Example 2

This example filters the response to include only Reserved Instances modification
requests with status `processing`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeReservedInstancesModifications
&Filter.1.Name=status
&Filter.1.Value.1=processing
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describereservedinstancesmodifications.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describereservedinstancesmodifications.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeReservedInstancesListings

DescribeReservedInstancesOfferings
