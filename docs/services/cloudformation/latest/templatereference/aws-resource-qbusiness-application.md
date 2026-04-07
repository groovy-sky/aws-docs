This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Application

Creates an Amazon Q Business application.

###### Note

There are new tiers for Amazon Q Business. Not all features in Amazon Q Business Pro are
also available in Amazon Q Business Lite. For information on what's included in
Amazon Q Business Lite and what's included in Amazon Q Business Pro, see [Amazon Q Business tiers](../../../amazonq/latest/qbusiness-ug/tiers.md#user-sub-tiers). You must use the Amazon Q Business console to assign
subscription tiers to users.

An Amazon Q Apps service linked role will be created if it's absent in the
AWS account when `QAppsConfiguration` is enabled in
the request. For more information, see [Using\
service-linked roles for Q Apps](../../../amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.md).

When you create an application, Amazon Q Business may securely transmit data for
processing from your selected AWS region, but within your geography.
For more information, see [Cross region\
inference in Amazon Q Business](../../../amazonq/latest/qbusiness-ug/cross-region-inference.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::Application",
  "Properties" : {
      "AttachmentsConfiguration" : AttachmentsConfiguration,
      "AutoSubscriptionConfiguration" : AutoSubscriptionConfiguration,
      "ClientIdsForOIDC" : [ String, ... ],
      "Description" : String,
      "DisplayName" : String,
      "EncryptionConfiguration" : EncryptionConfiguration,
      "IamIdentityProviderArn" : String,
      "IdentityCenterInstanceArn" : String,
      "IdentityType" : String,
      "PersonalizationConfiguration" : PersonalizationConfiguration,
      "QAppsConfiguration" : QAppsConfiguration,
      "QuickSightConfiguration" : QuickSightConfiguration,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::Application
Properties:
  AttachmentsConfiguration:
    AttachmentsConfiguration
  AutoSubscriptionConfiguration:
    AutoSubscriptionConfiguration
  ClientIdsForOIDC:
    - String
  Description: String
  DisplayName: String
  EncryptionConfiguration:
    EncryptionConfiguration
  IamIdentityProviderArn: String
  IdentityCenterInstanceArn: String
  IdentityType: String
  PersonalizationConfiguration:
    PersonalizationConfiguration
  QAppsConfiguration:
    QAppsConfiguration
  QuickSightConfiguration:
    QuickSightConfiguration
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`AttachmentsConfiguration`

Configuration information for the file upload during chat feature.

_Required_: No

_Type_: [AttachmentsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-attachmentsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoSubscriptionConfiguration`

Subscription configuration information for an Amazon Q Business application
using IAM identity federation for user management.

_Required_: No

_Type_: [AutoSubscriptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-autosubscriptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientIdsForOIDC`

The OIDC client ID for a Amazon Q Business application.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the Amazon Q Business application.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the Amazon Q Business application.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

Provides the identifier of the AWS KMS key used to encrypt data indexed by
Amazon Q Business. Amazon Q Business doesn't support asymmetric keys.

_Required_: No

_Type_: [EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-encryptionconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamIdentityProviderArn`

The Amazon Resource Name (ARN) of an identity provider being used by an Amazon Q
Business application.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:iam::\d{12}:(oidc-provider|saml-provider)/[a-zA-Z0-9_\.\/@\-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityCenterInstanceArn`

The Amazon Resource Name (ARN) of the IAM Identity Center instance you are either
creating for—or connecting to—your Amazon Q Business application.

_Required_: `Yes`

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityType`

The authentication type being used by a Amazon Q Business application.

_Required_: No

_Type_: String

_Allowed values_: `AWS_IAM_IDP_SAML | AWS_IAM_IDP_OIDC | AWS_IAM_IDC | AWS_QUICKSIGHT_IDP | ANONYMOUS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PersonalizationConfiguration`

Configuration information about chat response personalization. For more information,
see [Personalizing chat responses](../../../amazonq/latest/qbusiness-ug/personalizing-chat-responses.md).

_Required_: No

_Type_: [PersonalizationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-personalizationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QAppsConfiguration`

Configuration information about Amazon Q Apps.

_Required_: No

_Type_: [QAppsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-qappsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuickSightConfiguration`

The Amazon Quick configuration for an Amazon Q Business application that uses Quick as the identity provider.

_Required_: No

_Type_: [QuickSightConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-quicksightconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role with permissions to access your Amazon
CloudWatch logs and metrics. If this property is not specified, Amazon Q Business will create a [service linked role (SLR)](../../../amazonq/latest/qbusiness-ug/using-service-linked-roles.md#slr-permissions) and use it as the
application's role.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize your Amazon Q Business application.
You can also use tags to help control access to the application. Tag keys and values can
consist of Unicode letters, digits, white space, and any of the following symbols: \_ . :
/ = \+ \- @.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-application-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID. For example:

`{"Ref": "ApplicationId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationArn`

The Amazon Resource Name (ARN) of the Amazon Q Business application.

`ApplicationId`

The identifier for the Amazon Q Business application.

`CreatedAt`

The Unix timestamp when the Amazon Q Business application was created.

`IdentityCenterApplicationArn`

The Amazon Resource Name (ARN) of the AWSIAM Identity Center instance
attached to your Amazon Q Business application.

`Status`

The status of the Amazon Q Business application. The application is ready to use when the
status is `ACTIVE`.

`UpdatedAt`

The Unix timestamp when the Amazon Q Business application was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Q Business

AttachmentsConfiguration
