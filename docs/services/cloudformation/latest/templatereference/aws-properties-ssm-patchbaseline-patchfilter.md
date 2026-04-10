This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline PatchFilter

The `PatchFilter` property type defines a patch filter for an AWS Systems Manager patch baseline.

The `PatchFilters` property of the [PatchFilterGroup](../userguide/aws-properties-ssm-patchbaseline-patchfiltergroup.md) property type contains a list of `PatchFilter`
property types.

You can view lists of valid values for the patch properties by running the
`DescribePatchProperties` command. For more information, see [DescribePatchProperties](../../../../reference/systems-manager/latest/apireference/api-describepatchproperties.md) in the _AWS Systems Manager API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

The key for the filter.

For information about valid keys, see [PatchFilter](../../../../reference/systems-manager/latest/apireference/api-patchfilter.md)
in the _AWS Systems Manager API Reference_.

_Required_: No

_Type_: String

_Allowed values_: `ADVISORY_ID | ARCH | BUGZILLA_ID | CLASSIFICATION | CVE_ID | EPOCH | MSRC_SEVERITY | NAME | PATCH_ID | PATCH_SET | PRIORITY | PRODUCT | PRODUCT_FAMILY | RELEASE | REPOSITORY | SECTION | SECURITY | SEVERITY | VERSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value for the filter key.

For information about valid values for each key based on operating system type, see
[PatchFilter](../../../../reference/systems-manager/latest/apireference/api-patchfilter.md)
in the _AWS Systems Manager API Reference_.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `64 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSM::PatchBaseline

PatchFilterGroup

All content copied from https://docs.aws.amazon.com/.
