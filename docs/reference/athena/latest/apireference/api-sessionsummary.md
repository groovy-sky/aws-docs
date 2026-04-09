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

Type: [EngineVersion](api-engineversion.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/sessionsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/sessionsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/sessionsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionStatus

TableMetadata

All content copied from https://docs.aws.amazon.com/.
