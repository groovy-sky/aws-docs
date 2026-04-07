This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::PrivacyBudgetTemplate

An object that defines the privacy budget template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::PrivacyBudgetTemplate",
  "Properties" : {
      "AutoRefresh" : String,
      "MembershipIdentifier" : String,
      "Parameters" : Parameters,
      "PrivacyBudgetType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::PrivacyBudgetTemplate
Properties:
  AutoRefresh: String
  MembershipIdentifier: String
  Parameters:
    Parameters
  PrivacyBudgetType: String
  Tags:
    - Tag

```

## Properties

`AutoRefresh`

How often the privacy budget refreshes.

###### Important

If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.

_Required_: Yes

_Type_: String

_Allowed values_: `CALENDAR_MONTH | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MembershipIdentifier`

The identifier for a membership resource.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Specifies the epsilon and noise parameters for the privacy budget template.

_Required_: Yes

_Type_: [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-privacybudgettemplate-parameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivacyBudgetType`

Specifies the type of the privacy budget template.

_Required_: Yes

_Type_: String

_Allowed values_: `DIFFERENTIAL_PRIVACY | ACCESS_BUDGET`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag consists
of a key and an optional value, both of which you define. When you use tagging, you can also
use tag-based access control in IAM policies to control access to this
resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-privacybudgettemplate-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myPrivacyBudgetTemplate" }`

For the Clean Rooms privacy budget template, `Ref` returns the ARN of the privacy budget template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the privacy budget template.

`CollaborationArn`

The ARN of the collaboration that contains this privacy budget template.

`CollaborationIdentifier`

The unique ID of the collaboration that contains this privacy budget template.

`MembershipArn`

The Amazon Resource Name (ARN) of the member who created the privacy budget template.

`PrivacyBudgetTemplateIdentifier`

A unique identifier for one of your memberships for a collaboration. The privacy budget template is created in the collaboration that this membership belongs to. Accepts a membership ID.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

BudgetParameter
