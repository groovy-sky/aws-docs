This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter PackageFilter

Contains information on the details of a package filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Architecture" : StringFilter,
  "Epoch" : NumberFilter,
  "FilePath" : StringFilter,
  "Name" : StringFilter,
  "Release" : StringFilter,
  "SourceLambdaLayerArn" : StringFilter,
  "SourceLayerHash" : StringFilter,
  "Version" : StringFilter
}

```

### YAML

```yaml

  Architecture:
    StringFilter
  Epoch:
    NumberFilter
  FilePath:
    StringFilter
  Name:
    StringFilter
  Release:
    StringFilter
  SourceLambdaLayerArn:
    StringFilter
  SourceLayerHash:
    StringFilter
  Version:
    StringFilter

```

## Properties

`Architecture`

An object that contains details on the package architecture type to filter on.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Epoch`

An object that contains details on the package epoch to filter on.

_Required_: No

_Type_: [NumberFilter](aws-properties-inspectorv2-filter-numberfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilePath`

Property description not available.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

An object that contains details on the name of the package to filter on.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Release`

An object that contains details on the package release to filter on.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLambdaLayerArn`

Property description not available.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLayerHash`

An object that contains details on the source layer hash to filter on.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The package version to filter on.

_Required_: No

_Type_: [StringFilter](aws-properties-inspectorv2-filter-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberFilter

PortRangeFilter

All content copied from https://docs.aws.amazon.com/.
