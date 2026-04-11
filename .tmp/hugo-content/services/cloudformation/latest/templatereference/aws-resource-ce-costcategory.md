This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CE::CostCategory

The `AWS::CE::CostCategory` resource creates groupings of cost that you can use
across products in the AWS Billing and Cost Management console, such as Cost Explorer and AWS Budgets. For more information, see [Managing Your Costs with\
Cost Categories](../../../awsaccountbilling/latest/aboutv2/manage-cost-categories.md) in the _AWS Billing and Cost Management User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CE::CostCategory",
  "Properties" : {
      "DefaultValue" : String,
      "Name" : String,
      "Rules" : String,
      "RuleVersion" : String,
      "SplitChargeRules" : String,
      "Tags" : [ ResourceTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CE::CostCategory
Properties:
  DefaultValue: String
  Name: String
  Rules: String
  RuleVersion: String
  SplitChargeRules: String
  Tags:
    - ResourceTag

```

## Properties

`DefaultValue`

The default value for the cost category.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of the Cost Category.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

The array of CostCategoryRule in JSON array format.

###### Note

Rules are processed in order. If there are multiple rules that match the line item, then
the first rule to match is used to determine that Cost Category value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleVersion`

The rule schema version in this particular Cost Category.

_Required_: Yes

_Type_: String

_Allowed values_: `CostCategoryExpression.v1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SplitChargeRules`

The split charge rules that are used to allocate your charges between your cost category values.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-ce-costcategory-resourcetag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Arn of the Cost Category that is created by the
template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute
of this type. The following are the available attributes.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`Arn`

The unique identifier for your Cost Category.

`EffectiveStart`

The Cost Category's effective start date.

## Examples

### Cost Category Department with two rules

The following example creates a Cost Category "Department" with two rules.

#### JSON

```json

{
  "CostCategoryDepartment": {
    "Type": "AWS::CE::CostCategory",
    "Properties": {
      "Name": "Department",
      "RuleVersion": "CostCategoryExpression.v1",
      "Rules": "[
        {
          \"Value\":\"Engineering\",
          \"Rule\": {
            \"Dimensions\": {
            \"Key\": \"LINKED_ACCOUNT\",
            \"Values\": [ \"111111111111\" ]
            }
          }
        },
        {
          \"Value\": \"Marketing\",
          \"Rule\": {
            \"Dimensions\": {
            \"Key\": \"LINKED_ACCOUNT\",
            \"Values\": [ \"222222222222\" ]
            }
          }
        }
        ]"
      }
    }
  }

```

#### YAML

```yaml

  CostCategoryDepartment:
    Type: 'AWS::CE::CostCategory'
    Properties:
      Name: Department
      RuleVersion: CostCategoryExpression.v1
      Rules: '
        [
          {
            "Value":"Engineering",
            "Rule": {
              "Dimensions": {
                "Key": "LINKED_ACCOUNT",
                "Values": [ "111111111111" ]
          }
        }
      },
      {
        "Value":"Marketing",
        "Rule": {
          "Dimensions": {
          "Key": "LINKED_ACCOUNT",
          "Values": [ "222222222222" ]
          }
        }

    ]'

```

## See also

- [CostCategory](../../../../reference/aws-cost-management/latest/apireference/api-costcategory.md) in the _AWS Billing and Cost Management API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Subscriber

ResourceTag

All content copied from https://docs.aws.amazon.com/.
