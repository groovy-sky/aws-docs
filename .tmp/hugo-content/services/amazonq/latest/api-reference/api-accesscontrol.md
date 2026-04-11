# AccessControl

A list of principals. Each principal can be either a `USER` or a
`GROUP` and can be designated document access permissions of either
`ALLOW` or `DENY`.

## Contents

**principals**

Contains a list of principals, where a principal can be either a `USER` or
a `GROUP`. Each principal can be have the following type of document access:
`ALLOW` or `DENY`.

Type: Array of [Principal](api-principal.md) objects

Required: Yes

**memberRelation**

Describes the member relation within a principal list.

Type: String

Valid Values: `AND | OR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/accesscontrol.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/accesscontrol.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/accesscontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessConfiguration

ActionConfiguration

All content copied from https://docs.aws.amazon.com/.
