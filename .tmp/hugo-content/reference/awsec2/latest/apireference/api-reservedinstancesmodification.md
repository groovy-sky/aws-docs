# ReservedInstancesModification

Describes a Reserved Instance modification.

## Contents

**clientToken**

A unique, case-sensitive key supplied by the client to ensure that the request is
idempotent. For more information, see [Ensuring\
Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**createDate**

The time when the modification request was created.

Type: Timestamp

Required: No

**effectiveDate**

The time for the modification to become effective.

Type: Timestamp

Required: No

**ModificationResultSet.N**

Contains target configurations along with their corresponding new Reserved Instance
IDs.

Type: Array of [ReservedInstancesModificationResult](api-reservedinstancesmodificationresult.md) objects

Required: No

**reservedInstancesModificationId**

A unique ID for the Reserved Instance modification.

Type: String

Required: No

**ReservedInstancesSet.N**

The IDs of one or more Reserved Instances.

Type: Array of [ReservedInstancesId](api-reservedinstancesid.md) objects

Required: No

**status**

The status of the Reserved Instances modification request.

Type: String

Required: No

**statusMessage**

The reason for the status.

Type: String

Required: No

**updateDate**

The time when the modification request was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/reservedinstancesmodification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/reservedinstancesmodification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/reservedinstancesmodification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReservedInstancesListing

ReservedInstancesModificationResult
