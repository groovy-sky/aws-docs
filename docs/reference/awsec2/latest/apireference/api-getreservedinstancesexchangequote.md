# GetReservedInstancesExchangeQuote

Returns a quote and exchange information for exchanging one or more specified Convertible
Reserved Instances for a new Convertible Reserved Instance. If the exchange cannot be
performed, the reason is returned in the response. Use [AcceptReservedInstancesExchangeQuote](api-acceptreservedinstancesexchangequote.md) to perform the exchange.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making
the request, and provides an error response. If you have the required permissions, the error
response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**ReservedInstanceId.N**

The IDs of the Convertible Reserved Instances to exchange.

Type: Array of strings

Required: Yes

**TargetConfiguration.N**

The configuration of the target Convertible Reserved Instance to exchange for your current
Convertible Reserved Instances.

Type: Array of [TargetConfigurationRequest](api-targetconfigurationrequest.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**currencyCode**

The currency of the transaction.

Type: String

**isValidExchange**

If `true`, the exchange is valid. If `false`, the exchange cannot be
completed.

Type: Boolean

**outputReservedInstancesWillExpireAt**

The new end date of the reservation term.

Type: Timestamp

**paymentDue**

The total true upfront charge for the exchange.

Type: String

**requestId**

The ID of the request.

Type: String

**reservedInstanceValueRollup**

The cost associated with the Reserved Instance.

Type: [ReservationValue](api-reservationvalue.md) object

**reservedInstanceValueSet**

The configuration of your Convertible Reserved Instances.

Type: Array of [ReservedInstanceReservationValue](api-reservedinstancereservationvalue.md) objects

**targetConfigurationValueRollup**

The cost associated with the Reserved Instance.

Type: [ReservationValue](api-reservationvalue.md) object

**targetConfigurationValueSet**

The values of the target Convertible Reserved Instances.

Type: Array of [TargetReservationValue](api-targetreservationvalue.md) objects

**validationFailureReason**

Describes the reason why the exchange cannot be completed.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the output of requesting whether a potential exchange is
valid.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetReservedInstancesExchangeQuote
&ReservedInstanceId.1=649fd0c8-7768-46b8-8f84-a6400EXAMPLE
&TargetConfiguration.1.OfferingId=24167194-6541-4041-9e31-bc7c5984aa53
&AUTHPARAMS
```

#### Sample Response

```

<GetReservedInstancesExchangeQuoteResponse>
  <requestId>d072f652-cc57-458c-89e0-e6c02EXAMPLE</requestId>
  <outputReservedInstancesWillExpireAt>2019-05-17T12:32:53Z</outputReservedInstancesWillExpireAt>
  <reservedInstanceValueSet>
    <item>
        <reservedInstancesId>649fd0c8-7768-46b8-8f84-a6400EXAMPLE</reservedInstanceId>
        <reservationValue>
            <remainingTotalValue>98.048402</remainingTotalValue>
            <hourlyPrice>0.018000</hourlyPrice>
            <remainingUpfrontValue>631.0</remainingUpfrontValue>
        </reservationValue>
    </item>
  </reservedInstanceValueSet>
  <targetConfigurationValueSet>
  <isValidExchange>false</isValidExchange>
  <paymentDue>-448.416438</paymentDue>
  <targetConfigurationValueRollup>
          <remainingTotalValue>0</remainingTotalValue>
          <hourlyPrice>0</hourlyPrice>
          <remainingUpfrontValue>0</remainingUpfrontValue>
  <targetConfigurationValueRollup>
  <reservedInstanceValueRollup>
          <remainingTotalValue>873.504438</remainingTotalValue>
          <hourlyPrice>0.018000</hourlyPrice>
          <remainingUpfrontValue>448.416438</remainingUpfrontValue>
  </reservedInstanceValueRollup>
  <currencyCode>USD</currencyCode>
  <validationFailureReason>The target configuration value is less than the input</validationFailureReason>
</GetReservedInstancesExchangeQuoteResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetReservedInstancesExchangeQuote)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetReservedInstancesExchangeQuote)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getreservedinstancesexchangequote.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetReservedInstancesExchangeQuote)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getreservedinstancesexchangequote.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetPasswordData

GetRouteServerAssociations
