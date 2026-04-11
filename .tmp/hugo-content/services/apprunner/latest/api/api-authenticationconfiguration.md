# AuthenticationConfiguration

Describes resources needed to authenticate access to some source repositories. The specific resource depends on the repository provider.

## Contents

**AccessRoleArn**

The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository. It's required for ECR image repositories
(but not for ECR Public repositories).

Type: String

Length Constraints: Minimum length of 29. Maximum length of 1024.

Pattern: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:(role|role\/service-role)\/[\w+=,.@\-/]{1,1000}`

Required: No

**ConnectionArn**

The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository. It's required for GitHub code
repositories.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/authenticationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/authenticationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/authenticationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

AutoScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
