# Action

An action defines the tasks that the extension performs during the AWS AppConfig
workflow. Each action includes an action point, as shown in the following list:

- `PRE_CREATE_HOSTED_CONFIGURATION_VERSION`

- `PRE_START_DEPLOYMENT`

- `AT_DEPLOYMENT_TICK`

- `ON_DEPLOYMENT_START`

- `ON_DEPLOYMENT_STEP`

- `ON_DEPLOYMENT_BAKING`

- `ON_DEPLOYMENT_COMPLETE`

- `ON_DEPLOYMENT_ROLLED_BACK`

Each action also includes a name, a URI to an AWS Lambda function, and an
Amazon Resource Name (ARN) for an AWS Identity and Access Management assume role. You specify the name,
URI, and ARN for each _action point_ defined in the extension.

## Contents

**Description**

Information about the action.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**Name**

The action name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**RoleArn**

An Amazon Resource Name (ARN) for an AWS Identity and Access Management assume role.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

Required: No

**Uri**

The extension URI associated to the action point in the extension definition. The URI
can be an Amazon Resource Name (ARN) for one of the following: an AWS Lambda
function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/action.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/action.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/action.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AWS AppConfig

ActionInvocation
