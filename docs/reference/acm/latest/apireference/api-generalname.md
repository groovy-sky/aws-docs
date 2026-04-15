---
title: "GeneralName"
---

# GeneralName

Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280). Only one of the following naming options should be provided.

## Contents

###### Note

In the following list, the required parameters are described first.

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**DirectoryName**

Contains information about the certificate subject. The `Subject` field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The `Subject` must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.

Type: [DistinguishedName](api-distinguishedname.md) object

Required: No

**DnsName**

Represents `GeneralName` as a DNS name.

Type: String

Required: No

**IpAddress**

Represents `GeneralName` as an IPv4 or IPv6 address.

Type: String

Required: No

**OtherName**

Represents `GeneralName` using an `OtherName` object.

Type: [OtherName](api-othername.md) object

Required: No

**RegisteredId**

Represents `GeneralName` as an object identifier (OID).

Type: String

Required: No

**Rfc822Name**

Represents `GeneralName` as an [RFC 822](https://datatracker.ietf.org/doc/html/rfc822) email address.

Type: String

Required: No

**UniformResourceIdentifier**

Represents `GeneralName` as a URI.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/generalname.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/generalname.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/generalname.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filters

HttpRedirect

All content copied from https://docs.aws.amazon.com/.
