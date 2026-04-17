---
title: "CopyValueEntry"
---

# CopyValueEntry

This object defines one value to be copied with the [copyValue](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-copoyValue) processor.

## Contents

**source**

The key to copy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**target**

The key of the field to copy the value to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**overwriteIfExists**

Specifies whether to overwrite the value if the destination key already exists. If you
omit this, the default is `false`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/CopyValueEntry)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/CopyValueEntry)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/CopyValueEntry)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyValue

CSV

All content copied from https://docs.aws.amazon.com/.
