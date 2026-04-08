# DescribeReservedDBInstances

Returns information about reserved DB instances for this account, or about a specified reserved DB instance.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceClass**

The DB instance class filter value. Specify this parameter to show only those reservations matching the specified DB instances class.

Type: String

Required: No

**Duration**

The duration filter value, specified in years or seconds. Specify this parameter to show only reservations for this duration.

Valid Values: `1 | 3 | 31536000 | 94608000`

Type: String

Required: No

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**LeaseId**

The lease identifier filter value. Specify this parameter to show only the reservation that matches the specified lease ID.

###### Note

AWS Support might request the lease ID for an issue related to a reserved DB instance.

Type: String

Required: No

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more than the `MaxRecords` value is available, a pagination token called a marker is
included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

**MultiAZ**

Specifies whether to show only those reservations that support Multi-AZ.

Type: Boolean

Required: No

**OfferingType**

The offering type filter value. Specify this parameter to show only the available offerings matching the specified offering type.

Valid Values: `"Partial Upfront" | "All Upfront" | "No Upfront" `

Type: String

Required: No

**ProductDescription**

The product description filter value. Specify this parameter to show only those reservations matching the specified product description.

Type: String

Required: No

**ReservedDBInstanceId**

The reserved DB instance identifier filter value. Specify this parameter to show only the reservation that matches the specified reservation ID.

Type: String

Required: No

**ReservedDBInstancesOfferingId**

The offering identifier filter value. Specify this parameter to show only purchased reservations matching the specified offering identifier.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

**ReservedDBInstances.ReservedDBInstance.N**

A list of reserved DB instances.

Type: Array of [ReservedDBInstance](api-reserveddbinstance.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ReservedDBInstanceNotFound**

The specified reserved DB Instance not found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeReservedDBInstances.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeReservedDBInstances
   &ReservedDBInstanceId=customerSpecifiedID
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140420/us-west-2/rds/aws4_request
   &X-Amz-Date=20140420T162211Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=3312d17a4c43bcd209bc22a0778dd23e73f8434254abbd7ac53b89ade3dae88e

```

#### Sample Response

```

<DescribeReservedDBInstancesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeReservedDBInstancesResult>
    <ReservedDBInstances>
      <ReservedDBInstance>
        <OfferingType>Partial Upfront</OfferingType>
        <CurrencyCode>USD</CurrencyCode>
        <RecurringCharges/>
        <ProductDescription>mysql</ProductDescription>
        <ReservedDBInstancesOfferingId>649fd0c8-cf6d-47a0-bfa6-060f8e75e95f</ReservedDBInstancesOfferingId>
        <MultiAZ>false</MultiAZ>
        <State>active</State>
        <ReservedDBInstanceId>myreservationid</ReservedDBInstanceId>
        <DBInstanceCount>1</DBInstanceCount>
        <StartTime>2014-05-15T00:25:14.131Z</StartTime>
        <Duration>31536000</Duration>
        <FixedPrice>227.5</FixedPrice>
        <UsagePrice>0.046</UsagePrice>
        <DBInstanceClass>db.m1.small</DBInstanceClass>
      </ReservedDBInstance>
  </DescribeReservedDBInstancesResult>
  <ResponseMetadata>
    <RequestId>c695119b-2961-11e1-bd06-6fe008f046c3</RequestId>
  </ResponseMetadata>
</DescribeReservedDBInstancesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describereserveddbinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describereserveddbinstances.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribePendingMaintenanceActions

DescribeReservedDBInstancesOfferings

All content copied from https://docs.aws.amazon.com/.
