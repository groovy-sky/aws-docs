This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::SecurityConfig IamIdentityCenterConfigOptions

Describes IAM Identity Center options for an OpenSearch Serverless security
configuration in the form of a key-value map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationArn" : String,
  "ApplicationDescription" : String,
  "ApplicationName" : String,
  "GroupAttribute" : String,
  "InstanceArn" : String,
  "UserAttribute" : String
}

```

### YAML

```yaml

  ApplicationArn: String
  ApplicationDescription: String
  ApplicationName: String
  GroupAttribute: String
  InstanceArn: String
  UserAttribute: String

```

## Properties

`ApplicationArn`

The ARN of the IAM Identity Center application used to integrate with OpenSearch
Serverless.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationDescription`

The description of the IAM Identity Center application used to integrate with
OpenSearch Serverless.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationName`

The name of the IAM Identity Center application used to integrate with OpenSearch
Serverless.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupAttribute`

The group attribute for this IAM Identity Center integration. Defaults to
`GroupId`.

_Required_: No

_Type_: String

_Allowed values_: `GroupId | GroupName`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The ARN of the IAM Identity Center instance used to integrate with OpenSearch
Serverless.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserAttribute`

The user attribute for this IAM Identity Center integration. Defaults to
`UserId`

_Required_: No

_Type_: String

_Allowed values_: `UserId | UserName | Email`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IamFederationConfigOptions

SamlConfigOptions

All content copied from https://docs.aws.amazon.com/.
