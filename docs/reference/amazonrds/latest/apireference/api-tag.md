# Tag

Metadata assigned to an Amazon RDS resource consisting of a key-value pair.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**Key**

A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with `aws:` or `rds:`. The string can only contain only the set of Unicode letters, digits, white-space, '\_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-@\]\*)$").

Type: String

Required: No

**Value**

A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with `aws:` or `rds:`. The string can only contain only the set of Unicode letters, digits, white-space, '\_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-@\]\*)$").

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/Tag)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/Tag)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/Tag)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SwitchoverDetail

TagSpecification
