This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ValidationStrategy

The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to `LENIENT`, validation is skipped for specific errors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

The mode of validation for the asset to be created or updated. When you set this value to `STRICT`, strict validation for every error is enforced. When you set this value to `LENIENT`, validation is skipped for specific UI errors.

_Required_: Yes

_Type_: String

_Allowed values_: `STRICT | LENIENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UniqueValuesComputation

VisibleRangeOptions
