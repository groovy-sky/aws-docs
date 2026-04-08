# ReservedDBInstancesOffering

This data type is used as a response element in the `DescribeReservedDBInstancesOfferings` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**CurrencyCode**

The currency code for the reserved DB instance offering.

Type: String

Required: No

**DBInstanceClass**

The DB instance class for the reserved DB instance.

Type: String

Required: No

**Duration**

The duration of the offering in seconds.

Type: Integer

Required: No

**FixedPrice**

The fixed price charged for this offering.

Type: Double

Required: No

**MultiAZ**

Indicates whether the offering applies to Multi-AZ deployments.

Type: Boolean

Required: No

**OfferingType**

The offering type.

Type: String

Required: No

**ProductDescription**

The database engine used by the offering.

Type: String

Required: No

**RecurringCharges.RecurringCharge.N**

The recurring price charged to run this reserved DB instance.

Type: Array of [RecurringCharge](api-recurringcharge.md) objects

Required: No

**ReservedDBInstancesOfferingId**

The offering identifier.

Type: String

Required: No

**UsagePrice**

The hourly price charged for this offering.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/reserveddbinstancesoffering.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/reserveddbinstancesoffering.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/reserveddbinstancesoffering.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReservedDBInstance

ResourcePendingMaintenanceActions

All content copied from https://docs.aws.amazon.com/.
