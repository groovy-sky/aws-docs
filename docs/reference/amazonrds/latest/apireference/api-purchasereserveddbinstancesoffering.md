# PurchaseReservedDBInstancesOffering

Purchases a reserved DB instance offering.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ReservedDBInstancesOfferingId**

The ID of the Reserved DB instance offering to purchase.

Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706

Type: String

Required: Yes

**DBInstanceCount**

The number of instances to reserve.

Default: `1`

Type: Integer

Required: No

**ReservedDBInstanceId**

Customer-specified identifier to track this reservation.

Example: myreservationID

Type: String

Required: No

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**ReservedDBInstance**

This data type is used as a response element in the
`DescribeReservedDBInstances` and
`PurchaseReservedDBInstancesOffering` actions.

Type: [ReservedDBInstance](api-reserveddbinstance.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ReservedDBInstanceAlreadyExists**

User already has a reservation with the given identifier.

HTTP Status Code: 404

**ReservedDBInstanceQuotaExceeded**

Request would exceed the user's DB Instance quota.

HTTP Status Code: 400

**ReservedDBInstancesOfferingNotFound**

Specified offering does not exist.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of PurchaseReservedDBInstancesOffering.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=PurchaseReservedDBInstancesOffering
   &ReservedDBInstanceId=myreservationID
   &ReservedDBInstancesOfferingId=438012d3-4052-4cc7-b2e3-8d3372e0e706
   &DBInstanceCount=10
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140415/us-east-1/rds/aws4_request
   &X-Amz-Date=20140415T232655Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=c2ac761e8c8f54a8c0727f5a87ad0a766fbb0024510b9aa34ea6d1f7df52fb11

```

#### Sample Response

```

<PurchaseReservedDBInstancesOfferingResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <PurchaseReservedDBInstancesOfferingResult>
    <ReservedDBInstance>
      <OfferingType>Partial Upfront</OfferingType>
      <CurrencyCode>USD</CurrencyCode>
      <RecurringCharges/>
      <ProductDescription>mysql</ProductDescription>
      <ReservedDBInstancesOfferingId>438012d3-4052-4cc7-b2e3-8d3372e0e706</ReservedDBInstancesOfferingId>
      <MultiAZ>true</MultiAZ>
      <State>payment-pending</State>
      <ReservedDBInstanceId>myreservationID</ReservedDBInstanceId>
      <DBInstanceCount>10</DBInstanceCount>
      <StartTime>2014-05-18T23:24:56.577Z</StartTime>
      <Duration>31536000</Duration>
      <FixedPrice>123.0</FixedPrice>
      <UsagePrice>0.123</UsagePrice>
      <DBInstanceClass>db.m1.small</DBInstanceClass>
    </ReservedDBInstance>
  </PurchaseReservedDBInstancesOfferingResult>
  <ResponseMetadata>
    <RequestId>7f099901-29cf-11e1-bd06-6fe008f046c3</RequestId>
  </ResponseMetadata>
</PurchaseReservedDBInstancesOfferingResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/purchasereserveddbinstancesoffering.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromoteReadReplicaDBCluster

RebootDBCluster

All content copied from https://docs.aws.amazon.com/.
