# CancelReservedInstancesListing

Cancels the specified Reserved Instance listing in the Reserved Instance
Marketplace.

For more information, see [Sell in the Reserved Instance\
Marketplace](../../../../services/ec2/latest/userguide/ri-market-general.md) in the _Amazon EC2 User Guide_.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ReservedInstancesListingId**

The ID of the Reserved Instance listing.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**reservedInstancesListingsSet**

The Reserved Instance listing.

Type: Array of [ReservedInstancesListing](api-reservedinstanceslisting.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request cancels a Reserved Instance listing in the Reserved Instance
Marketplace. The response shows that the status is cancelled.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CancelReservedInstancesListing
&ReservedInstancesListingId=3ebe97b5-f273-43b6-a204-7a18cEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<CancelReservedInstancesListingResponse>
    <requestId>bec2cf62-98ef-434a-8a15-886fcexample</requestId>
    <reservedInstancesListingsSet>
        <item>
            <reservedInstancesListingId>3ebe97b5-f273-43b6-a204-7a18cEXAMPLE</reservedInstancesListingId>
            <reservedInstancesId>e5a2ff3b-7d14-494f-90af-0b5d0EXAMPLE</reservedInstancesId>
            <createDate>2012-07-12T16:55:28.000Z</createDate>
            <updateDate>2012-07-12T16:55:28.000Z</updateDate>
            <status>cancelled</status>
            <statusMessage>CANCELLED</statusMessage>
            <instanceCounts>
                <item>
                    <state>Available</state>
                    <instanceCount>0</instanceCount>
                </item>
                <item>
                    <state>Sold</state>
                    <instanceCount>0</instanceCount>
                </item>
                <item>
                    <state>Cancelled</state>
                    <instanceCount>1</instanceCount>
                </item>
                <item>
                    <state>Pending</state>
                    <instanceCount>0</instanceCount>
                </item>
            </instanceCounts>
            <priceSchedules>
                <item>
                    <term>5</term>
                    <price>166.64</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>4</term>
                    <price>133.32</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>3</term>
                    <price>99.99</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>2</term>
                    <price>66.66</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
                <item>
                    <term>1</term>
                    <price>33.33</price>
                    <currencyCode>USD</currencyCode>
                    <active>false</active>
                </item>
            </priceSchedules>
            <tagSet/>
            <clientToken>XqJIt1342112125076</clientToken>
        </item>
    </reservedInstancesListingsSet>
</CancelReservedInstancesListingResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/cancelreservedinstanceslisting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelreservedinstanceslisting.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelImportTask

CancelSpotFleetRequests
