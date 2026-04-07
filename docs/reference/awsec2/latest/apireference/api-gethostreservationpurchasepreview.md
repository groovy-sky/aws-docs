# GetHostReservationPurchasePreview

Preview a reservation purchase with configurations that match those of your Dedicated
Host. You must have active Dedicated Hosts in your account before you purchase a
reservation.

This is a preview of the [PurchaseHostReservation](api-purchasehostreservation.md) action and does not
result in the offering being purchased.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**HostIdSet.N**

The IDs of the Dedicated Hosts with which the reservation is associated.

Type: Array of strings

Required: Yes

**OfferingId**

The offering ID of the reservation.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**currencyCode**

The currency in which the `totalUpfrontPrice` and
`totalHourlyPrice` amounts are specified. At this time, the only
supported currency is `USD`.

Type: String

Valid Values: `USD`

**purchase**

The purchase information of the Dedicated Host reservation and the Dedicated Hosts
associated with it.

Type: Array of [Purchase](api-purchase.md) objects

**requestId**

The ID of the request.

Type: String

**totalHourlyPrice**

The potential total hourly price of the reservation per hour.

Type: String

**totalUpfrontPrice**

The potential total upfront price. This is billed immediately.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example is a preview of the reservation to be purchased.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetHostReservationPurchasePreview
&OfferingId=hro-0eb3541dght849c2d
&HostIdSet=h-0fgr9ddb0ecd0a1cd
&AUTHPARAMS
```

#### Sample Response

```

<GetHostReservationPurchasePreviewResult xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-84c3-4967-gtyk-a9983EXAMPLE</requestId>
        <purchase>
            <item>
                <duration>31536000</duration>
                <upfrontPrice>7453.000</upfrontPrice>
                <paymentOption>PartialUpfront</paymentOption>
                <instanceFamily>m4</instanceFamily>
                <hourlyPrice>0.850</hourlyPrice>
                <hostIdSet>
                    <item>h-0fgr9ddb0ecd0a1cd</item>
                </hostIdSet>
            </item>
       </purchase>
       <totalHourlyPrice>0.850<totalHourlyPrice>
       <totalUpfrontPrice>7453.000<totalUpfrontPrice>
</GetHostReservationPurchasePreviewResult>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetHostReservationPurchasePreview)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetHostReservationPurchasePreview)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/gethostreservationpurchasepreview.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetHostReservationPurchasePreview)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/gethostreservationpurchasepreview.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetGroupsForCapacityReservation

GetImageAncestry
