---
title: "MoveKeyEntry"
---

# MoveKeyEntry

This object defines one key that will be moved with the [moveKey](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-moveKey) processor.

## Contents

**source**

The key to move.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**target**

The key to move to.

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/MoveKeyEntry)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/MoveKeyEntry)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/MoveKeyEntry)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricTransformation

MoveKeys

All content copied from https://docs.aws.amazon.com/.
