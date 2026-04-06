# AccessControl

A list of principals. Each principal can be either a `USER` or a
`GROUP` and can be designated document access permissions of either
`ALLOW` or `DENY`.

## Contents

**principals**

Contains a list of principals, where a principal can be either a `USER` or
a `GROUP`. Each principal can be have the following type of document access:
`ALLOW` or `DENY`.

Type: Array of [Principal](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_Principal.html) objects

Required: Yes

**memberRelation**

Describes the member relation within a principal list.

Type: String

Valid Values: `AND | OR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/AccessControl)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/AccessControl)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/AccessControl)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessConfiguration

ActionConfiguration
