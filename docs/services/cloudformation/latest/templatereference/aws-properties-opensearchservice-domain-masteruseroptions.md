This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain MasterUserOptions

Specifies information about the master user.

Required if `InternalUserDatabaseEnabled` is true in [AdvancedSecurityOptions](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MasterUserARN" : String,
  "MasterUserName" : String,
  "MasterUserPassword" : String
}

```

### YAML

```yaml

  MasterUserARN: String
  MasterUserName: String
  MasterUserPassword: String

```

## Properties

`MasterUserARN`

Amazon Resource Name (ARN) for the master user. The ARN can point to an IAM user or role. This
property is required for Amazon Cognito to work, and it must match the role configured for
Cognito. Only specify if `InternalUserDatabaseEnabled` is false in [AdvancedSecurityOptionsInput](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserName`

Username for the master user. Only specify if `InternalUserDatabaseEnabled` is true
in [AdvancedSecurityOptionsInput](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

If you don't want to specify this value directly within the template, you can use a [dynamic reference](../userguide/dynamic-references.md) instead.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserPassword`

Password for the master user. Only specify if `InternalUserDatabaseEnabled` is true
in [AdvancedSecurityOptionsInput](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

If you don't want to specify this value directly within the template, you can use a [dynamic reference](../userguide/dynamic-references.md) instead.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `8`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogPublishingOption

NodeConfig

All content copied from https://docs.aws.amazon.com/.
