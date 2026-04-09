# Tag

A key-value pair you can use to manage, filter, and search for your resources. Allowed
characters include UTF-8 letters, numbers, and the following characters: + - = . \_ :
/. Spaces are not allowed in tag values.

## Contents

**Key**

The key part of a tag's key-value pair. The key can't start with `aws:`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: Yes

**Value**

The value part of a tag's key-value pair.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `[^\x00]*`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/tag.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/tag.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/tag.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceStartTime

VirtualMachine

All content copied from https://docs.aws.amazon.com/.
