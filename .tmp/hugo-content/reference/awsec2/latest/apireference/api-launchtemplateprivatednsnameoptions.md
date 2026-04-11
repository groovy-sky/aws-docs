# LaunchTemplatePrivateDnsNameOptions

Describes the options for instance hostnames.

## Contents

**enableResourceNameDnsAAAARecord**

Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA
records.

Type: Boolean

Required: No

**enableResourceNameDnsARecord**

Indicates whether to respond to DNS queries for instance hostnames with DNS A
records.

Type: Boolean

Required: No

**hostnameType**

The type of hostname to assign to an instance.

Type: String

Valid Values: `ip-name | resource-name`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateprivatednsnameoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateprivatednsnameoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateprivatednsnameoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplatePlacementRequest

LaunchTemplatePrivateDnsNameOptionsRequest
