This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::Repository LifecyclePolicy

The `LifecyclePolicy` property type specifies a lifecycle policy. For
information about lifecycle policy syntax, see [Lifecycle policy\
template](../../../amazonecr/latest/userguide/lifecyclepolicies.md) in the _Amazon ECR User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LifecyclePolicyText" : String,
  "RegistryId" : String
}

```

### YAML

```yaml

  LifecyclePolicyText: String
  RegistryId: String

```

## Properties

`LifecyclePolicyText`

The JSON repository policy text to apply to the repository.

_Required_: No

_Type_: String

_Minimum_: `100`

_Maximum_: `30720`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryId`

The AWS account ID associated with the registry that contains the repository. If you
do  not specify a registry, the default registry is assumed.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Creating a lifecycle\
policy](../../../amazonecr/latest/userguide/lp-creation.md) in the _Amazon ECR User Guide_

- [PutLifecyclePolicy](../../../../reference/amazonecr/latest/apireference/api-putlifecyclepolicy.md) in the _Amazon ECR API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageTagMutabilityExclusionFilter

Tag

All content copied from https://docs.aws.amazon.com/.
