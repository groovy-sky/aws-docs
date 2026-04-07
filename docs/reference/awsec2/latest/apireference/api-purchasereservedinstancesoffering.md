# PurchaseReservedInstancesOffering

Purchases a Reserved Instance for use with your account. With Reserved Instances, you pay
a lower hourly rate compared to On-Demand instance pricing.

Use [DescribeReservedInstancesOfferings](api-describereservedinstancesofferings.md) to get a list of Reserved
Instance offerings that match your specifications. After you've purchased a Reserved Instance,
you can check for your new Reserved Instance with [DescribeReservedInstances](api-describereservedinstances.md).

To queue a purchase for a future date and time, specify a purchase time. If you do not
specify a purchase time, the default is the current time.

For more information, see [Reserved\
Instances](../../../../services/ec2/latest/userguide/concepts-on-demand-reserved-instances.md) and [Sell in the Reserved Instance\
Marketplace](../../../../services/ec2/latest/userguide/ri-market-general.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making
the request, and provides an error response. If you have the required permissions, the error
response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceCount**

The number of Reserved Instances to purchase.

Type: Integer

Required: Yes

**LimitPrice**

Specified for Reserved Instance Marketplace offerings to limit the total order and ensure
that the Reserved Instances are not purchased at unexpected prices.

Type: [ReservedInstanceLimitPrice](api-reservedinstancelimitprice.md) object

Required: No

**PurchaseTime**

The time at which to purchase the Reserved Instance, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

Required: No

**ReservedInstancesOfferingId**

The ID of the Reserved Instance offering to purchase.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**reservedInstancesId**

The IDs of the purchased Reserved Instances. If your purchase crosses into a discounted
pricing tier, the final Reserved Instances IDs might change. For more information, see [Crossing pricing tiers](../../../../services/ec2/latest/userguide/concepts-reserved-instances-application.md#crossing-pricing-tiers) in the _Amazon EC2 User Guide_.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example uses a limit price to limit the total purchase order of Standard Reserved
Instances from the Reserved Instance Marketplace.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=PurchaseReservedInstancesOffering
&ReservedInstancesOfferingId=4b2293b4-5813-4cc8-9ce3-1957fEXAMPLE
&LimitPrice.Amount=200
&InstanceCount=2
&AUTHPARAMS
```

#### Sample Response

```

<PurchaseReservedInstancesOfferingResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <reservedInstancesId>e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE</reservedInstancesId>
</PurchaseReservedInstancesOfferingResponse>
```

### Example 2

This example illustrates a purchase of a Reserved Instances offering.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=PurchaseReservedInstancesOffering
&ReservedInstancesOfferingId=4b2293b4-5813-4cc8-9ce3-1957fEXAMPLE
&InstanceCount=2
&AUTHPARAMS
```

#### Sample Response

```

<PurchaseReservedInstancesOfferingResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <reservedInstancesId>e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE</reservedInstancesId>
</PurchaseReservedInstancesOfferingResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/PurchaseReservedInstancesOffering)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/PurchaseReservedInstancesOffering)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/purchasereservedinstancesoffering.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/PurchaseReservedInstancesOffering)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/purchasereservedinstancesoffering.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PurchaseHostReservation

PurchaseScheduledInstances
