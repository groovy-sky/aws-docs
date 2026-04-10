This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::PartnerApp

The `AWS::SageMaker::PartnerApp` resource creates an Amazon SageMaker Partner AI App. For more
information, see [Partner AI\
Apps](../../../sagemaker/latest/dg/partner-apps.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::PartnerApp",
  "Properties" : {
      "ApplicationConfig" : PartnerAppConfig,
      "AppVersion" : String,
      "AuthType" : String,
      "EnableAutoMinorVersionUpgrade" : Boolean,
      "EnableIamSessionBasedIdentity" : Boolean,
      "ExecutionRoleArn" : String,
      "KmsKeyId" : String,
      "MaintenanceConfig" : PartnerAppMaintenanceConfig,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Tier" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::PartnerApp
Properties:
  ApplicationConfig:
    PartnerAppConfig
  AppVersion: String
  AuthType: String
  EnableAutoMinorVersionUpgrade: Boolean
  EnableIamSessionBasedIdentity: Boolean
  ExecutionRoleArn: String
  KmsKeyId: String
  MaintenanceConfig:
    PartnerAppMaintenanceConfig
  Name: String
  Tags:
    - Tag
  Tier: String
  Type: String

```

## Properties

`ApplicationConfig`

Configuration settings for the Partner AI App.

_Required_: No

_Type_: [PartnerAppConfig](aws-properties-sagemaker-partnerapp-partnerappconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppVersion`

The version of the Partner AI App. This property allows you to specify a particular version of the application
to deploy.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

Defines the authentication type used for the Partner AI App.

_Required_: Yes

_Type_: String

_Allowed values_: `IAM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableAutoMinorVersionUpgrade`

When set to `TRUE`, the SageMaker Partner AI App is automatically upgraded to the latest minor version during the next scheduled maintenance window, if one is available. Default is `FALSE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableIamSessionBasedIdentity`

Enables IAM Session based Identity for PartnerApp.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role of the user.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The AWS KMS customer managed key used to encrypt the data associated with the PartnerApp.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceConfig`

A collection of settings that specify the maintenance schedule for the PartnerApp.

_Required_: No

_Type_: [PartnerAppMaintenanceConfig](aws-properties-sagemaker-partnerapp-partnerappmaintenanceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the Partner AI App. This name must be unique within your account and region.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to apply to the PartnerApp.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-partnerapp-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

Specifies the tier or level of the Partner AI App. The tier size impacts the speed and capabilities of the
application. For more information, see [Set up Partner AI Apps](../../../sagemaker/latest/dg/partner-app-onboard.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of Partner AI App being created.

_Required_: Yes

_Type_: String

_Allowed values_: `lakera-guard | comet | deepchecks-llm-evaluation | fiddler`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the Partner AI App.

`BaseUrl`

The AppServerUrl based on app and account-info.

`CurrentVersionEolDate`

The end-of-life date for the current version of the Partner AI App.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotebookInstanceLifecycleHook

PartnerAppConfig

All content copied from https://docs.aws.amazon.com/.
