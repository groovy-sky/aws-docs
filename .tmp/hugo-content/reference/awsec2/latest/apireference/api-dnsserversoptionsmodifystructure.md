# DnsServersOptionsModifyStructure

Information about the DNS server to be used.

## Contents

**CustomDnsServers.N**

The IPv4 address range, in CIDR notation, of the DNS servers to be used. You can specify up to
two DNS servers. Ensure that the DNS servers can be reached by the clients. The specified values
overwrite the existing values.

Type: Array of strings

Required: No

**Enabled**

Indicates whether DNS servers should be used. Specify `False` to delete the existing DNS
servers.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/dnsserversoptionsmodifystructure.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/dnsserversoptionsmodifystructure.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/dnsserversoptionsmodifystructure.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DnsOptionsSpecification

EbsBlockDevice
