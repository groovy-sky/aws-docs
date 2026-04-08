# DescribeReservedInstancesListings

Describes your account's Reserved Instance listings in the Reserved Instance
Marketplace.

The Reserved Instance Marketplace matches sellers who want to resell Reserved Instance
capacity that they no longer need with buyers who want to purchase additional capacity.
Reserved Instances bought and sold through the Reserved Instance Marketplace work like any
other Reserved Instances.

As a seller, you choose to list some or all of your Reserved Instances, and you specify
the upfront price to receive for them. Your Reserved Instances are then listed in the Reserved
Instance Marketplace and are available for purchase.

As a buyer, you specify the configuration of the Reserved Instance to purchase, and the
Marketplace matches what you're searching for with what's available. The Marketplace first
sells the lowest priced Reserved Instances to you, and continues to sell available Reserved
Instance listings to you until your demand is met. You are charged based on the total price of
all of the listings that you purchase.

For more information, see [Sell in the Reserved Instance\
Marketplace](../../../../services/ec2/latest/userguide/ri-market-general.md) in the _Amazon EC2 User Guide_.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Filter.N**

One or more filters.

- `reserved-instances-id` \- The ID of the Reserved Instances.

- `reserved-instances-listing-id` \- The ID of the Reserved Instances
listing.

- `status` \- The status of the Reserved Instance listing ( `pending` \|
`active` \| `cancelled` \| `closed`).

- `status-message` \- The reason for the status.

Type: Array of [Filter](api-filter.md) objects

Required: No

**ReservedInstancesId**

One or more Reserved Instance IDs.

Type: String

Required: No

**ReservedInstancesListingId**

One or more Reserved Instance listing IDs.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**reservedInstancesListingsSet**

Information about the Reserved Instance listing.

Type: Array of [ReservedInstancesListing](api-reservedinstanceslisting.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example shows all the listings associated with your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeReservedInstancesListings
&AUTHPARAMS
```

#### Sample Response

```

<DescribeReservedInstancesListingsResponse>
    <requestId>cec5c904-8f3a-4de5-8f5a-ff7f9EXAMPLE</requestId>
    <reservedInstancesListingsSet>
        <item>
            <reservedInstancesListingId>253dfbf9-c335-4808-b956-d942cEXAMPLE</reservedInstancesListingId>
            <reservedInstancesId>e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE</reservedInstancesId>
            <createDate>2012-07-06T19:35:29.000Z</createDate>
            <updateDate>2012-07-06T19:35:30.000Z</updateDate>
            <status>active</status>
            <statusMessage>ACTIVE</statusMessage>
            <instanceCounts>
                <item>
                    <state>Available</state>
                    <instanceCount>20</instanceCount>
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
                    <term>8</term>
                    <price>480.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>7</term>
                    <price>420.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>6</term>
                    <price>360.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>active</active>
                </item>
                <item>
                    <term>5</term>
                    <price>300.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>4</term>
                    <price>240.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>3</term>
                    <price>180.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>2</term>
                    <price>120.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>1</term>
                    <price>60.0</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
            </priceSchedules>
            <tagSet/>
            <clientToken>myclienttoken1</clientToken>
        </item>
    </reservedInstancesListingsSet>
</DescribeReservedInstancesListingsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describereservedinstanceslistings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describereservedinstanceslistings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeReservedInstances

DescribeReservedInstancesModifications
