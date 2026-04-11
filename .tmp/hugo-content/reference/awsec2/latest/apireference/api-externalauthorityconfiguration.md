# ExternalAuthorityConfiguration

The configuration that links an Amazon VPC IPAM scope to an external authority system. It specifies the type of external system and the external resource identifier that identifies your account or instance in that system.

For more information, see [Integrate VPC IPAM with Infoblox infrastructure](../../../../services/vpc/latest/ipam/integrate-infoblox-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Contents

**ExternalResourceIdentifier**

The identifier for the external resource managing this scope. For Infoblox integrations, this is the Infoblox resource identifier in the format `<version>.identity.account.<entity_realm>.<entity_id>`.

Type: String

Required: No

**Type**

The type of external authority.

Type: String

Valid Values: `infoblox`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/externalauthorityconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/externalauthorityconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/externalauthorityconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportToS3TaskSpecification

FailedCapacityReservationFleetCancellationResult
