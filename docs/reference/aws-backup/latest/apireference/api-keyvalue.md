---
title: "KeyValue"
---

# KeyValue

Pair of two related strings. Allowed characters
are letters, white space, and numbers that can be
represented in UTF-8 and the following characters:
` + - = . _ : /`

## Contents

**Key**

The tag key (String). The key can't start with
`aws:`.

Length Constraints: Minimum length of 1. Maximum
length of 128.

Pattern: `^(?![aA]{1}[wW]{1}[sS]{1}:)([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Type: String

Required: Yes

**Value**

The value of the key.

Length Constraints: Maximum length of 256.

Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/backup-2018-11-15/KeyValue)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/backup-2018-11-15/KeyValue)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/backup-2018-11-15/KeyValue)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IndexedRecoveryPoint

LatestMpaApprovalTeamUpdate

All content copied from https://docs.aws.amazon.com/.
