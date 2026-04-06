# SessionSummary

Contains summary information about a session.

## Contents

**Description**

The session description.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**EngineVersion**

The engine version used by the session (for example, `PySpark engine version
                3`).

Type: [EngineVersion](https://docs.aws.amazon.com/athena/latest/APIReference/API_EngineVersion.html) object

Required: No

**NotebookVersion**

The notebook version.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**SessionId**

The session ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**Status**

Contains information about the session status.

Type: [SessionStatus](api-sessionstatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/SessionSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/SessionSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/SessionSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SessionStatus

TableMetadata
