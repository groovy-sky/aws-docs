This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe

Specifies a new AWS Glue DataBrew transformation recipe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Recipe",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Steps" : [ RecipeStep, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Recipe
Properties:
  Description: String
  Name: String
  Steps:
    - RecipeStep
  Tags:
    - Tag

```

## Properties

`Description`

The description of the recipe.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name for the recipe.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Steps`

A list of steps that are defined by the recipe.

_Required_: Yes

_Type_: Array of [RecipeStep](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-recipe-recipestep.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata tags that have been applied to the recipe.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-recipe-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource name. For example:

`{ "Ref": "myRecipe" }`

For an AWS Glue DataBrew recipe named `myRecipe`, `Ref`
returns the name of the recipe.

## Examples

### Creating recipes

The following examples create new DataBrew recipes.

#### YAML

```yaml

Resources:
  TestDataBrewRecipe:
    Type: AWS::DataBrew::Recipe
    Properties:
      Name: recipe-name
      Description: This is the recipe description.
      Steps:
      - Action:
          Operation: EXTRACT_PATTERN
          Parameters:
            SourceColumn: Consulate
            Pattern: A
            TargetColumn: extract_pattern
        ConditionExpressions:
        - Condition : LESS_THAN_EQUAL
          Value: 5
          TargetColumn: Target
      Tags: [{Key: key00AtCreate, Value: value001AtCreate}]

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "This CloudFormation template specifies a DataBrew Recipe",
    "Resources": {
        "MyDataBrewRecipe": {
            "Type": "AWS::DataBrew::Recipe",
            "Properties": {
              "Name": "na-recipe-cf-test",
              "Description": "This is the recipe description.",
              "Steps":[
                {
                  "Action":{
                    "Operation":"EXTRACT_PATTERN",
                    "Parameters":{
                      "SourceColumn": "Consulate",
                      "Pattern": "A",
                      "TargetColumn": "extract_pattern"
                  }
                },
                  "ConditionExpressions":[
                    {
                      "Condition": "LESS_THAN",
                      "ConditionValue": "2",
                      "TargetColumn": "target"
                    },
                    {
                      "Condition": "GREATER_THAN",
                      "Value": "0",
                      "TargetColumn": "target"
                    }
                  ]
                }
              ],
              "Tags": [
                    {
                        "Key": "key00AtCreate",
                        "Value": "value001AtCreate"
                    }
                ]
            }
        }
    }
}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Action
