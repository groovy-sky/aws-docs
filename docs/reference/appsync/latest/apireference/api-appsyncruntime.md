# AppSyncRuntime

Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note
that if a runtime is specified, code must also be specified.

## Contents

**name**

The `name` of the runtime to use. Currently, the only allowed value is
`APPSYNC_JS`.

Type: String

Valid Values: `APPSYNC_JS`

Required: Yes

**runtimeVersion**

The `version` of the runtime to use. Currently, the only allowed version is
`1.0.0`.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/appsyncruntime.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/appsyncruntime.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/appsyncruntime.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKey

AuthMode

All content copied from https://docs.aws.amazon.com/.
