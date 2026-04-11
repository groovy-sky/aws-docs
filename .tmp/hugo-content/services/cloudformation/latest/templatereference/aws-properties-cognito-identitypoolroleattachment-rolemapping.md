This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPoolRoleAttachment RoleMapping

One of a set of `RoleMappings`, a property of the [AWS::Cognito::IdentityPoolRoleAttachment](../userguide/aws-resource-cognito-identitypoolroleattachment.md) resource that defines the
role-mapping attributes of an Amazon Cognito identity pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmbiguousRoleResolution" : String,
  "IdentityProvider" : String,
  "RulesConfiguration" : RulesConfigurationType,
  "Type" : String
}

```

### YAML

```yaml

  AmbiguousRoleResolution: String
  IdentityProvider: String
  RulesConfiguration:
    RulesConfigurationType
  Type: String

```

## Properties

`AmbiguousRoleResolution`

If you specify Token or Rules as the `Type`,
`AmbiguousRoleResolution` is required.

Specifies the action to be taken if either no rules match the claim value for the
`Rules` type, or there is no `cognito:preferred_role` claim
and there are multiple `cognito:roles` matches for the `Token`
type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProvider`

Identifier for the identity provider for which the role is mapped. For example:
`graph.facebook.com` or
`cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id
                (http://cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id)`.
This is the identity provider that is used by the user for authentication.

If the identity provider property isn't provided, the key of the entry in the
`RoleMappings` map is used as the identity provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RulesConfiguration`

The rules to be used for mapping users to roles. If you specify "Rules" as the
role-mapping type, RulesConfiguration is required.

_Required_: No

_Type_: [RulesConfigurationType](aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The role mapping type. Token will use `cognito:roles` and
`cognito:preferred_role` claims from the Cognito identity provider token
to map groups to roles. Rules will attempt to match claims from the token to map to a
role.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MappingRule

RulesConfigurationType

All content copied from https://docs.aws.amazon.com/.
