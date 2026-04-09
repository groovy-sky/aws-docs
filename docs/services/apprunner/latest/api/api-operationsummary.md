# OperationSummary

Provides summary information for an operation that occurred on an AWS App Runner service.

## Contents

**EndedAt**

The time when the operation ended. It's in the Unix time stamp format.

Type: Timestamp

Required: No

**Id**

A unique ID of this operation. It's unique in the scope of the App Runner service.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`

Required: No

**StartedAt**

The time when the operation started. It's in the Unix time stamp format.

Type: Timestamp

Required: No

**Status**

The current state of the operation.

Type: String

Valid Values: `PENDING | IN_PROGRESS | FAILED | SUCCEEDED | ROLLBACK_IN_PROGRESS | ROLLBACK_FAILED | ROLLBACK_SUCCEEDED`

Required: No

**TargetArn**

The Amazon Resource Name (ARN) of the resource that the operation acted on (for example, an App Runner service).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**Type**

The type of operation. It indicates a specific action that occured.

Type: String

Valid Values: `START_DEPLOYMENT | CREATE_SERVICE | PAUSE_SERVICE | RESUME_SERVICE | DELETE_SERVICE | UPDATE_SERVICE`

Required: No

**UpdatedAt**

The time when the operation was last updated. It's in the Unix time stamp format.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/operationsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/operationsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/operationsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObservabilityConfigurationSummary

Service

All content copied from https://docs.aws.amazon.com/.
