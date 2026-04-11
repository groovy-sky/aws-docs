This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyTemplate

Creates a policy template. A template can use placeholders for the principal and
resource. A template must be instantiated into a policy by associating it with specific
principals and resources to use for the placeholders. That instantiated policy can then
be considered in authorization decisions. The instantiated policy works identically to
any other policy, except that it is dynamically linked to the template. If the template
changes, then any policies that are linked to that template are immediately updated as
well.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VerifiedPermissions::PolicyTemplate",
  "Properties" : {
      "Description" : String,
      "PolicyStoreId" : String,
      "Statement" : String
    }
}

```

### YAML

```yaml

Type: AWS::VerifiedPermissions::PolicyTemplate
Properties:
  Description: String
  PolicyStoreId: String
  Statement: String

```

## Properties

`Description`

The description to attach to the new or updated policy template.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyStoreId`

The unique identifier of the policy store that contains the template.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Statement`

Specifies the content that you want to use for the new policy template, written in the Cedar
policy language.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique id of the policy store followed by '\|' and the
unique id of the new or updated policy template. For example:

`{ "Ref": "POLICYSTOREabcde111111|POLICYTEMPLATEab111111" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PolicyTemplateId`

The unique identifier of the new or modified policy template.

## Examples

### Creating a policy template

The following example creates a policy template with the specified
statement.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation sample template for creating a policy template for Verified Permissions.",
    "Parameters": {
        "PolicyStoreId": {
            "Type": "String"
        },
        "Description": {
            "Type": "String"
        },
        "Statement": {
            "Type": "String"
        }
    },
    "Resources": {
        "PolicyTemplate": {
            "Type": "AWS::VerifiedPermissions::PolicyTemplate",
            "Properties": {
                "PolicyStoreId": {
                    "Ref": "PolicyStoreId"
                },
                "Description": {
                    "Ref": "Description"
                },
                "Statement": {
                    "Ref": "Statement"
                }
            }
        }
    },
    "Outputs": {
        "PolicyTemplateId": {
            "Value": {
                "Fn::GetAtt": [
                    "PolicyTemplate",
                    "PolicyTemplateId"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation sample template for creating a policy template for Verified Permissions."
Parameters:
  PolicyStoreId:
    Type: String
  Description:
    Type: String
  Statement:
    Type: String
Resources:
  PolicyTemplate:
    Type: AWS::VerifiedPermissions::PolicyTemplate
    Properties:
      PolicyStoreId: !Ref PolicyStoreId
      Description: !Ref Description
      Statement: !Ref Statement
Outputs:
  PolicyTemplateId:
    Value: !GetAtt PolicyTemplate.PolicyTemplateId
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValidationSettings

Next

All content copied from https://docs.aws.amazon.com/.
