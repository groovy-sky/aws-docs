# IpamScopeExternalAuthorityConfiguration

The configuration that links an Amazon VPC IPAM scope to an external authority system. It specifies the type of external system and the external resource identifier that identifies your account or instance in that system.

In IPAM, an external authority is a third-party IP address management system that provides CIDR blocks when you provision address space for top-level IPAM pools. This allows you to use your existing IP management system to control which address ranges are allocated to AWS while using Amazon VPC IPAM to manage subnets within those ranges.

## Contents

**externalResourceIdentifier**

The identifier for the external resource managing this scope. For Infoblox integrations, this is the Infoblox resource identifier in the format `<version>.identity.account.<entity_realm>.<entity_id>`.

Type: String

Required: No

**type**

The type of external authority managing this scope. Currently supports `Infoblox` for integration with Infoblox Universal DDI.

Type: String

Valid Values: `infoblox`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipamscopeexternalauthorityconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipamscopeexternalauthorityconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipamscopeexternalauthorityconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamScope

IpPermission
