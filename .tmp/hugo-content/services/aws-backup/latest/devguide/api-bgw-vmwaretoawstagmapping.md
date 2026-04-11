# VmwareToAwsTagMapping

This displays the mapping of VMware tags to the
corresponding AWS tags.

## Contents

**AwsTagKey**

The key part of the AWS tag's key-value pair.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: Yes

**AwsTagValue**

The value part of the AWS tag's key-value pair.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `[^\x00]*`

Required: Yes

**VmwareCategory**

The is the category of VMware.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 80.

Required: Yes

**VmwareTagName**

This is the user-defined name of a VMware tag.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 80.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/vmwaretoawstagmapping.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/vmwaretoawstagmapping.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/vmwaretoawstagmapping.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VmwareTag

AWS Backup search

All content copied from https://docs.aws.amazon.com/.
