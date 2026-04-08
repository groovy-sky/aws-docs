# ResourceRecord

Contains a DNS record value that you can use to validate ownership or control of a
domain. This is used by the [DescribeCertificate](api-describecertificate.md) action.

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the DNS record to create in your domain. This is supplied by ACM.

Type: String

Required: Yes

**Type**

The type of DNS record. Currently this can be `CNAME`.

Type: String

Valid Values: `CNAME`

Required: Yes

**Value**

The value of the CNAME record to add to your DNS database. This is supplied by
ACM.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/resourcerecord.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/resourcerecord.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/resourcerecord.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RenewalSummary

Tag
