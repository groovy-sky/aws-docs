---
title: "DistinguishedName"
---

# DistinguishedName

Contains X.500 distinguished name information.

## Contents

###### Note

In the following list, the required parameters are described first.

**CommonName**

The common name (CN) attribute.

Type: String

Required: No

**Country**

The country (C) attribute.

Type: String

Required: No

**CustomAttributes**

A list of custom attributes in the distinguished name. Each custom attribute contains an object identifier (OID) and its corresponding value.

Type: Array of [CustomAttribute](api-customattribute.md) objects

Required: No

**DistinguishedNameQualifier**

The distinguished name qualifier attribute.

Type: String

Required: No

**DomainComponents**

The domain component attributes.

Type: Array of strings

Required: No

**GenerationQualifier**

The generation qualifier attribute.

Type: String

Required: No

**GivenName**

The given name attribute.

Type: String

Required: No

**Initials**

The initials attribute.

Type: String

Required: No

**Locality**

The locality (L) attribute.

Type: String

Required: No

**Organization**

The organization (O) attribute.

Type: String

Required: No

**OrganizationalUnit**

The organizational unit (OU) attribute.

Type: String

Required: No

**Pseudonym**

The pseudonym attribute.

Type: String

Required: No

**SerialNumber**

The serial number attribute.

Type: String

Required: No

**State**

The state or province (ST) attribute.

Type: String

Required: No

**Surname**

The surname attribute.

Type: String

Required: No

**Title**

The title attribute.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/DistinguishedName)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/DistinguishedName)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/DistinguishedName)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomAttribute

DnsNameFilter

All content copied from https://docs.aws.amazon.com/.
