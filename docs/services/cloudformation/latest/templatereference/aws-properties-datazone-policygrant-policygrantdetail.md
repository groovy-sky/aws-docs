This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant PolicyGrantDetail

The details of the policy grant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddToProjectMemberPool" : AddToProjectMemberPoolPolicyGrantDetail,
  "CreateAssetType" : CreateAssetTypePolicyGrantDetail,
  "CreateDomainUnit" : CreateDomainUnitPolicyGrantDetail,
  "CreateEnvironment" : Json,
  "CreateEnvironmentFromBlueprint" : Json,
  "CreateEnvironmentProfile" : CreateEnvironmentProfilePolicyGrantDetail,
  "CreateFormType" : CreateFormTypePolicyGrantDetail,
  "CreateGlossary" : CreateGlossaryPolicyGrantDetail,
  "CreateProject" : CreateProjectPolicyGrantDetail,
  "CreateProjectFromProjectProfile" : CreateProjectFromProjectProfilePolicyGrantDetail,
  "DelegateCreateEnvironmentProfile" : Json,
  "OverrideDomainUnitOwners" : OverrideDomainUnitOwnersPolicyGrantDetail,
  "OverrideProjectOwners" : OverrideProjectOwnersPolicyGrantDetail
}

```

### YAML

```yaml

  AddToProjectMemberPool:
    AddToProjectMemberPoolPolicyGrantDetail
  CreateAssetType:
    CreateAssetTypePolicyGrantDetail
  CreateDomainUnit:
    CreateDomainUnitPolicyGrantDetail
  CreateEnvironment: Json
  CreateEnvironmentFromBlueprint: Json
  CreateEnvironmentProfile:
    CreateEnvironmentProfilePolicyGrantDetail
  CreateFormType:
    CreateFormTypePolicyGrantDetail
  CreateGlossary:
    CreateGlossaryPolicyGrantDetail
  CreateProject:
    CreateProjectPolicyGrantDetail
  CreateProjectFromProjectProfile:
    CreateProjectFromProjectProfilePolicyGrantDetail
  DelegateCreateEnvironmentProfile: Json
  OverrideDomainUnitOwners:
    OverrideDomainUnitOwnersPolicyGrantDetail
  OverrideProjectOwners:
    OverrideProjectOwnersPolicyGrantDetail

```

## Properties

`AddToProjectMemberPool`

Specifies that the policy grant is to be added to the members of the project.

_Required_: No

_Type_: [AddToProjectMemberPoolPolicyGrantDetail](aws-properties-datazone-policygrant-addtoprojectmemberpoolpolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateAssetType`

Specifies that this is a create asset type policy.

_Required_: No

_Type_: [CreateAssetTypePolicyGrantDetail](aws-properties-datazone-policygrant-createassettypepolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateDomainUnit`

Specifies that this is a create domain unit policy.

_Required_: No

_Type_: [CreateDomainUnitPolicyGrantDetail](aws-properties-datazone-policygrant-createdomainunitpolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateEnvironment`

Specifies that this is a create environment policy.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateEnvironmentFromBlueprint`

Specifies that this is a create environment from blueprint policy.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateEnvironmentProfile`

Specifies that this is a create environment profile policy.

_Required_: No

_Type_: [CreateEnvironmentProfilePolicyGrantDetail](aws-properties-datazone-policygrant-createenvironmentprofilepolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateFormType`

Specifies that this is a create form type policy.

_Required_: No

_Type_: [CreateFormTypePolicyGrantDetail](aws-properties-datazone-policygrant-createformtypepolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateGlossary`

Specifies that this is a create glossary policy.

_Required_: No

_Type_: [CreateGlossaryPolicyGrantDetail](aws-properties-datazone-policygrant-createglossarypolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateProject`

Specifies that this is a create project policy.

_Required_: No

_Type_: [CreateProjectPolicyGrantDetail](aws-properties-datazone-policygrant-createprojectpolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreateProjectFromProjectProfile`

Specifies whether to create a project from project profile.

_Required_: No

_Type_: [CreateProjectFromProjectProfilePolicyGrantDetail](aws-properties-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DelegateCreateEnvironmentProfile`

Specifies that this is the delegation of the create environment profile policy.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OverrideDomainUnitOwners`

Specifies whether to override domain unit owners.

_Required_: No

_Type_: [OverrideDomainUnitOwnersPolicyGrantDetail](aws-properties-datazone-policygrant-overridedomainunitownerspolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OverrideProjectOwners`

Specifies whether to override project owners.

_Required_: No

_Type_: [OverrideProjectOwnersPolicyGrantDetail](aws-properties-datazone-policygrant-overrideprojectownerspolicygrantdetail.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverrideProjectOwnersPolicyGrantDetail

PolicyGrantPrincipal

All content copied from https://docs.aws.amazon.com/.
