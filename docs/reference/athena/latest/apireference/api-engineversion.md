# EngineVersion

The Athena engine version for running queries, or the PySpark engine
version for running sessions.

## Contents

**EffectiveEngineVersion**

Read only. The engine version on which the query runs. If the user requests a valid
engine version other than Auto, the effective engine version is the same as the engine
version that the user requested. If the user requests Auto, the effective engine version
is chosen by Athena. When a request to update the engine version is made by
a `CreateWorkGroup` or `UpdateWorkGroup` operation, the
`EffectiveEngineVersion` field is ignored.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**SelectedEngineVersion**

The engine version requested by the user. Possible values are determined by the output
of `ListEngineVersions`, including AUTO. The default is AUTO.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/engineversion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/engineversion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/engineversion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EngineConfiguration

ExecutorsSummary

All content copied from https://docs.aws.amazon.com/.
