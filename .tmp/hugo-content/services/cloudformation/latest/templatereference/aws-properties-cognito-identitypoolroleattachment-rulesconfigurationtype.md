This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPoolRoleAttachment RulesConfigurationType

`RulesConfigurationType` is a subproperty of the [RoleMapping](../userguide/aws-properties-cognito-identitypoolroleattachment-rolemapping.md) property that defines the rules to be used for mapping users to
roles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ MappingRule, ... ]
}

```

### YAML

```yaml

  Rules:
    - MappingRule

```

## Properties

`Rules`

The rules. You can specify up to 25 rules per identity provider.

_Required_: Yes

_Type_: Array of [MappingRule](aws-properties-cognito-identitypoolroleattachment-mappingrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RoleMapping

AWS::Cognito::LogDeliveryConfiguration

All content copied from https://docs.aws.amazon.com/.
