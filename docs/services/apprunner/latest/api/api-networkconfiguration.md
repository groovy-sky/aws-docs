# NetworkConfiguration

Describes configuration settings related to network traffic of an AWS App Runner service. Consists of embedded objects for each configurable network
feature.

## Contents

**EgressConfiguration**

Network configuration settings for outbound message traffic.

Type: [EgressConfiguration](api-egressconfiguration.md) object

Required: No

**IngressConfiguration**

Network configuration settings for inbound message traffic.

Type: [IngressConfiguration](api-ingressconfiguration.md) object

Required: No

**IpAddressType**

App Runner provides you with the option to choose between _IPv4_ and _dual stack_ (IPv4 and IPv6).
This is an optional parameter. If you do not specify an `IpAddressType`, it defaults to select IPv4.

Type: String

Valid Values: `IPV4 | DUAL_STACK`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/networkconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/networkconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/networkconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListVpcIngressConnectionsFilter

ObservabilityConfiguration

All content copied from https://docs.aws.amazon.com/.
