This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Ruleset

Specifies a new ruleset that can be used in a profile job to validate the data quality
of a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Ruleset",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Rules" : [ Rule, ... ],
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Ruleset
Properties:
  Description: String
  Name: String
  Rules:
    - Rule
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`Description`

The description of the ruleset.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the ruleset.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

Contains metadata about the ruleset.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-databrew-ruleset-rule.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see
[Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-databrew-ruleset-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is associated
with.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, Ref returns the resource name. For example, `{ "Ref": "myRuleset"
                }`.

For an AWS Glue DataBrew ruleset named `myRuleset`, `Ref`
returns the name of the ruleset.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ColumnSelector

All content copied from https://docs.aws.amazon.com/.
