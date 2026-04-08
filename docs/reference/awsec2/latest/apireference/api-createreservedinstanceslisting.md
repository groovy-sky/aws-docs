# CreateReservedInstancesListing

Creates a listing for Amazon EC2 Standard Reserved Instances to be sold in the Reserved
Instance Marketplace. You can submit one Standard Reserved Instance listing at a time. To get
a list of your Standard Reserved Instances, you can use the [DescribeReservedInstances](api-describereservedinstances.md) operation.

###### Note

Only Standard Reserved Instances can be sold in the Reserved Instance Marketplace.
Convertible Reserved Instances cannot be sold.

The Reserved Instance Marketplace matches sellers who want to resell Standard Reserved
Instance capacity that they no longer need with buyers who want to purchase additional
capacity. Reserved Instances bought and sold through the Reserved Instance Marketplace work
like any other Reserved Instances.

To sell your Standard Reserved Instances, you must first register as a seller in the
Reserved Instance Marketplace. After completing the registration process, you can create a
Reserved Instance Marketplace listing of some or all of your Standard Reserved Instances, and
specify the upfront price to receive for them. Your Standard Reserved Instance listings then
become available for purchase. To view the details of your Standard Reserved Instance listing,
you can use the [DescribeReservedInstancesListings](api-describereservedinstanceslistings.md) operation.

For more information, see [Sell in the Reserved Instance\
Marketplace](../../../../services/ec2/latest/userguide/ri-market-general.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier you provide to ensure idempotency of your listings. This
helps avoid duplicate listings. For more information, see [Ensuring\
Idempotency](run-instance-idempotency.md).

Type: String

Required: Yes

**InstanceCount**

The number of instances that are a part of a Reserved Instance account to be listed in the
Reserved Instance Marketplace. This number should be less than or equal to the instance count
associated with the Reserved Instance ID specified in this call.

Type: Integer

Required: Yes

**PriceSchedules.N**

A list specifying the price of the Standard Reserved Instance for each month remaining in
the Reserved Instance term.

Type: Array of [PriceScheduleSpecification](api-priceschedulespecification.md) objects

Required: Yes

**ReservedInstancesId**

The ID of the active Standard Reserved Instance.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**reservedInstancesListingsSet**

Information about the Standard Reserved Instance listing.

Type: Array of [ReservedInstancesListing](api-reservedinstanceslisting.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a Reserved Instance Marketplace listing from the specified
Standard Reserved Instance, which has 11 months remaining in its term. In this example, we
set the upfront price at $2.50, and the price drops over the course of the 11-month term
if the instance is still not sold.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateReservedInstancesListing
&ClientToken=myIdempToken1
&InstanceCount=1
&PriceSchedules.1.Price=2.5
&PriceSchedules.1.Term=11
&PriceSchedules.2.Price=2.0
&PriceSchedules.2.Term=8
&PriceSchedules.3.Price=1.5
&PriceSchedules.3.Term=5
&PriceSchedules.4.Price=0.7
&PriceSchedules.4.Term=3
&PriceSchedules.5.Price=0.1
&PriceSchedules.5.Term=1
&ReservedInstancesId=e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<CreateReservedInstancesListingResponse>
    <requestId>a42481af-335a-4e9e-b291-bd18dexample</requestId>
    <reservedInstancesListingsSet>
        <item>
            <reservedInstancesListingId>5ec28771-05ff-4b9b-aa31-9e57dEXAMPLE</reservedInstancesListingId>
            <reservedInstancesId>e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE</reservedInstancesId>
            <createDate>2012-07-17T17:11:09.449Z</createDate>
            <updateDate>2012-07-17T17:11:09.468Z</updateDate>
            <status>active</status>
            <statusMessage>ACTIVE</statusMessage>
            <instanceCounts>
                <item>
                    <state>Available</state>
                    <instanceCount>1</instanceCount>
                </item>
                <item>
                    <state>Sold</state>
                    <instanceCount>0</instanceCount>
                </item>
                <item>
                    <state>Cancelled</state>
                    <instanceCount>0</instanceCount>
                </item>
                <item>
                    <state>Pending</state>
                    <instanceCount>0</instanceCount>
                </item>
            </instanceCounts>
            <priceSchedules>
                <item>
                    <term>11</term>
                    <price>2.5</price>
                    <currencyCode>USD</currencyCode>
                    <active>true</active>
                </item>
                <item>
                    <term>10</term>
                    <price>2.5</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>9</term>
                    <price>2.5</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>8</term>
                    <price>2.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>7</term>
                    <price>2.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>6</term>
                    <price>2.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>5</term>
                    <price>1.5</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>4</term>
                    <price>1.5</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>3</term>
                    <price>0.7</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>2</term>
                    <price>0.7</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>1</term>
                    <price>0.1</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
            </priceSchedules>
            <tagSet/>
            <clientToken>myIdempToken1</clientToken>
        </item>
    </reservedInstancesListingsSet>
</CreateReservedInstancesListingResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createreservedinstanceslisting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createreservedinstanceslisting.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateReplaceRootVolumeTask

CreateRestoreImageTask
