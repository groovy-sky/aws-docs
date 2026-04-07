This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::Template

Specifies an email template. Email templates enable you to send personalized email to
one or more destinations in a single API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::Template",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "Template" : Template
    }
}

```

### YAML

```yaml

Type: AWS::SES::Template
Properties:
  Tags:
    - Tag
  Template:
    Template

```

## Properties

`Tags`

An array of objects that define the tags (keys and values) to associate with the email template.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-template-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Template`

The content of the email, composed of a subject line and either an HTML part or a
text-only part.

_Required_: No

_Type_: [Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-template-template.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Id`

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource ID. For example:

`{ "Ref": "MyTemplate" }`

For the Amazon SES template `MyTemplate`, `Ref` returns the
template ID.

## Examples

Specifies an email template, which is used when sending templated email
messages.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES Template Sample Template",
    "Parameters": {
        "TemplateName": {
            "Type": "String"
        },
        "SubjectPart": {
            "Type": "String"
        },
        "TextPart": {
            "Type": "String"
        },
        "HtmlPart": {
            "Type": "String"
        }
    },
    "Resources": {
        "Template": {
            "Type": "AWS::SES::Template",
            "Properties": {
                "Template": {
                    "TemplateName": {
                        "Ref": "TemplateName"
                    },
                    "SubjectPart": {
                        "Ref": "SubjectPart"
                    },
                    "TextPart": {
                        "Ref": "TextPart"
                    },
                    "HtmlPart": {
                        "Ref": "HtmlPart"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES Template Sample Template
Parameters:
  TemplateName:
    Type: String
  SubjectPart:
    Type: String
  TextPart:
    Type: String
  HtmlPart:
    Type: String
Resources:
  Template:
    Type: 'AWS::SES::Template'
    Properties:
      Template:
        TemplateName: !Ref TemplateName
        SubjectPart: !Ref SubjectPart
        TextPart: !Ref TextPart
        HtmlPart: !Ref HtmlPart
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SES::ReceiptRuleSet

Tag
