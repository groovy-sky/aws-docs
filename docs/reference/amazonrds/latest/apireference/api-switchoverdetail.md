# SwitchoverDetail

Contains the details about a blue/green deployment.

For more information, see [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User_
_Guide_ and [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**SourceMember**

The Amazon Resource Name (ARN) of a resource in the blue environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[A-Za-z][0-9A-Za-z-:._]*`

Required: No

**Status**

The switchover status of a resource in a blue/green deployment.

Values:

- `PROVISIONING` \- The resource is being prepared to switch
over.

- `AVAILABLE` \- The resource is ready to switch over.

- `SWITCHOVER_IN_PROGRESS` \- The resource is being switched
over.

- `SWITCHOVER_COMPLETED` \- The resource has been switched
over.

- `SWITCHOVER_FAILED` \- The resource attempted to switch over but
failed.

- `MISSING_SOURCE` \- The source resource has been deleted.

- `MISSING_TARGET` \- The target resource has been deleted.

Type: String

Required: No

**TargetMember**

The Amazon Resource Name (ARN) of a resource in the green environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[A-Za-z][0-9A-Za-z-:._]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/switchoverdetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/switchoverdetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/switchoverdetail.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SupportedEngineLifecycle

Tag

All content copied from https://docs.aws.amazon.com/.
