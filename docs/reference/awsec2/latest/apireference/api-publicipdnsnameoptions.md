# PublicIpDnsNameOptions

Public hostname type options. For more information, see [EC2 instance hostnames, DNS names, and domains](../../../../services/ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon EC2 User Guide_.

## Contents

**dnsHostnameType**

The public hostname type. For more information, see [EC2 instance hostnames, DNS names, and domains](../../../../services/ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**publicDualStackDnsName**

A dual-stack public hostname for a network interface. Requests from within the VPC resolve to both the private IPv4 address and the IPv6 Global Unicast Address of the network interface. Requests from the internet resolve to both the public IPv4 and the IPv6 GUA address of the network interface.

Type: String

Required: No

**publicIpv4DnsName**

An IPv4-enabled public hostname for a network interface. Requests from within the VPC resolve to the private primary IPv4 address of the network interface. Requests from the internet resolve to the public IPv4 address of the network interface.

Type: String

Required: No

**publicIpv6DnsName**

An IPv6-enabled public hostname for a network interface. Requests from within the VPC or from the internet resolve to the IPv6 GUA of the network interface.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/publicipdnsnameoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/publicipdnsnameoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/publicipdnsnameoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PtrUpdateStatus

PublicIpv4Pool
