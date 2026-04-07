# AllowedPrincipal

Describes a principal.

## Contents

**principal**

The Amazon Resource Name (ARN) of the principal.

Type: String

Required: No

**principalType**

The type of principal.

Type: String

Valid Values: `All | Service | OrganizationUnit | Account | User | Role`

Required: No

**serviceId**

The ID of the service.

Type: String

Required: No

**servicePermissionId**

The ID of the service permission.

Type: String

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AllowedPrincipal)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AllowedPrincipal)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AllowedPrincipal)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AddressTransfer

AlternatePathHint
