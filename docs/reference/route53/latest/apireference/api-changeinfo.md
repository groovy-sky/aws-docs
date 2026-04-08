# ChangeInfo

A complex type that describes change information about changes made to your hosted
zone.

## Contents

**Id**

This element contains an ID that you use when performing a [GetChange](api-getchange.md) action to get
detailed information about the change.

Type: String

Length Constraints: Maximum length of 32.

Required: Yes

**Status**

The current state of the request. `PENDING` indicates that this request has
not yet been applied to all Amazon Route 53 DNS servers.

Type: String

Valid Values: `PENDING | INSYNC`

Required: Yes

**SubmittedAt**

The date and time that the change request was submitted in [ISO 8601 format](https://en.wikipedia.org/wiki/ISO_8601) and Coordinated
Universal Time (UTC). For example, the value `2017-03-27T17:48:16.751Z`
represents March 27, 2017 at 17:48:16.751 UTC.

Type: Timestamp

Required: Yes

**Comment**

A comment you can provide.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/changeinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/changeinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/changeinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ChangeBatch

CidrBlockSummary
