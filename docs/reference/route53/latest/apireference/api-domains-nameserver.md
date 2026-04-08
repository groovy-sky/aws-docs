# Nameserver

Name server includes the following elements.

## Contents

**Name**

The fully qualified host name of the name server.

Constraint: Maximum 255 characters

Type: String

Length Constraints: Maximum length of 255.

Pattern: `[a-zA-Z0-9_\-.]*`

Required: Yes

**GlueIps**

Glue IP address of a name server entry. Glue IP addresses are required only when the
name of the name server is a subdomain of the domain. For example, if your domain is
example.com and the name server for the domain is ns.example.com, you need to specify
the IP address for ns.example.com.

Constraints: The list can contain only one IPv4 and one IPv6 address.

Type: Array of strings

Length Constraints: Maximum length of 45.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/nameserver.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/nameserver.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/nameserver.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FilterCondition

OperationSummary
