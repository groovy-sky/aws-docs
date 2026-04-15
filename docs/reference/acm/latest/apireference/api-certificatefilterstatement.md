---
title: "CertificateFilterStatement"
---

# CertificateFilterStatement

A filter statement used to search for certificates. Can contain AND, OR, NOT logical operators or a single filter.

## Contents

###### Note

In the following list, the required parameters are described first.

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**And**

A list of filter statements that must all be true.

Type: Array of [CertificateFilterStatement](api-certificatefilterstatement.md) objects

Array Members: Minimum number of 1 item. Maximum number of 15 items.

Required: No

**Filter**

A single certificate filter.

Type: [CertificateFilter](api-certificatefilter.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**Not**

A filter statement that must not be true.

Type: [CertificateFilterStatement](api-certificatefilterstatement.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**Or**

A list of filter statements where at least one must be true.

Type: Array of [CertificateFilterStatement](api-certificatefilterstatement.md) objects

Array Members: Minimum number of 1 item. Maximum number of 15 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/certificatefilterstatement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/certificatefilterstatement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/certificatefilterstatement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateFilter

CertificateMetadata

All content copied from https://docs.aws.amazon.com/.
