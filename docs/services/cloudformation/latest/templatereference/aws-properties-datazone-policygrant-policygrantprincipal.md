This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant PolicyGrantPrincipal

The policy grant principal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainUnit" : DomainUnitPolicyGrantPrincipal,
  "Group" : GroupPolicyGrantPrincipal,
  "Project" : ProjectPolicyGrantPrincipal,
  "User" : UserPolicyGrantPrincipal
}

```

### YAML

```yaml

  DomainUnit:
    DomainUnitPolicyGrantPrincipal
  Group:
    GroupPolicyGrantPrincipal
  Project:
    ProjectPolicyGrantPrincipal
  User:
    UserPolicyGrantPrincipal

```

## Properties

`DomainUnit`

The domain unit of the policy grant principal.

_Required_: No

_Type_: [DomainUnitPolicyGrantPrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-domainunitpolicygrantprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Group`

The group of the policy grant principal.

_Required_: No

_Type_: [GroupPolicyGrantPrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-grouppolicygrantprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Project`

The project of the policy grant principal.

_Required_: No

_Type_: [ProjectPolicyGrantPrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-projectpolicygrantprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`User`

The user of the policy grant principal.

_Required_: No

_Type_: [UserPolicyGrantPrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-policygrant-userpolicygrantprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PolicyGrantDetail

ProjectGrantFilter
