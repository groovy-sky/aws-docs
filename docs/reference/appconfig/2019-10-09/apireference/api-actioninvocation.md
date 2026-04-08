# ActionInvocation

An extension that was invoked as part of a deployment event.

## Contents

**ActionName**

The name of the action.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ErrorCode**

The error code when an extension invocation fails.

Type: String

Required: No

**ErrorMessage**

The error message when an extension invocation fails.

Type: String

Required: No

**ExtensionIdentifier**

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**InvocationId**

A system-generated ID for this invocation.

Type: String

Pattern: `[a-z0-9]{4,7}`

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

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/actioninvocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/actioninvocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/actioninvocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Action

Application
