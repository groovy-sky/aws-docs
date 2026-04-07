This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant

Adds a policy grant (an authorization policy) to a specified entity, including domain
units, environment blueprint configurations, or environment profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::PolicyGrant",
  "Properties" : {
      "Detail" : PolicyGrantDetail,
      "DomainIdentifier" : String,
      "EntityIdentifier" : String,
      "EntityType" : String,
      "PolicyType" : String,
      "Principal" : PolicyGrantPrincipal
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::PolicyGrant
Properties:
  Detail:
    PolicyGrantDetail
  DomainIdentifier: String
  EntityIdentifier: String
  EntityType: String
  PolicyType: String
  Principal:
    PolicyGrantPrincipal

```

## Properties

`Detail`

The details of the policy grant member.

_Required_: No

_Type_: [PolicyGrantDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-policygrantdetail.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainIdentifier`

The ID of the domain where you want to add a policy grant.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityIdentifier`

The ID of the entity (resource) to which you want to add a policy grant.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityType`

The type of entity (resource) to which the grant is added.

_Required_: Yes

_Type_: String

_Allowed values_: `DOMAIN_UNIT | ENVIRONMENT_BLUEPRINT_CONFIGURATION | ENVIRONMENT_PROFILE | ASSET_TYPE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyType`

The type of policy that you want to grant.

_Required_: Yes

_Type_: String

_Allowed values_: `CREATE_DOMAIN_UNIT | OVERRIDE_DOMAIN_UNIT_OWNERS | ADD_TO_PROJECT_MEMBER_POOL | OVERRIDE_PROJECT_OWNERS | CREATE_GLOSSARY | CREATE_FORM_TYPE | CREATE_ASSET_TYPE | CREATE_PROJECT | CREATE_ENVIRONMENT_PROFILE | DELEGATE_CREATE_ENVIRONMENT_PROFILE | CREATE_ENVIRONMENT | CREATE_ENVIRONMENT_FROM_BLUEPRINT | CREATE_PROJECT_FROM_PROJECT_PROFILE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

The principal of the policy grant member.

_Required_: No

_Type_: [PolicyGrantPrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-policygrantprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

Specifies the timestamp at which policy grant member was created.

`CreatedBy`

Specifies the user who created the policy grant member.

`GrantId`

The ID of the policy grant.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OwnerUserProperties

AddToProjectMemberPoolPolicyGrantDetail
