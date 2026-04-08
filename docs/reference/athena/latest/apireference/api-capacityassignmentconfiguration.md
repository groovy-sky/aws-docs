# CapacityAssignmentConfiguration

Assigns Athena workgroups (and hence their queries) to capacity
reservations. A capacity reservation can have only one capacity assignment
configuration, but the capacity assignment configuration can be made up of multiple
individual assignments. Each assignment specifies how Athena queries can
consume capacity from the capacity reservation that their workgroup is mapped to.

## Contents

**CapacityAssignments**

The list of assignments that make up the capacity assignment configuration.

Type: Array of [CapacityAssignment](api-capacityassignment.md) objects

Required: No

**CapacityReservationName**

The name of the reservation that the capacity assignment configuration is for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9._-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/capacityassignmentconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/capacityassignmentconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/capacityassignmentconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityAssignment

CapacityReservation
