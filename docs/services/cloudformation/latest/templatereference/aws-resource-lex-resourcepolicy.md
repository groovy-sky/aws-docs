This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::ResourcePolicy

###### Note

Amazon Lex V2 is the only supported version in CloudFormation.

Specifies a new resource policy with the specified policy
statements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lex::ResourcePolicy",
  "Properties" : {
      "Policy" : Json,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lex::ResourcePolicy
Properties:
  Policy: Json
  ResourceArn: String

```

## Properties

`Policy`

A resource policy to add to the resource. The policy is a JSON
structure that contains one or more statements that define the policy.
The policy must follow IAM syntax. If the policy isn't
valid, Amazon Lex returns a validation exception.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the bot or bot alias that the
resource policy is attached to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The identifier of the resource policy.

`RevisionId`

Specifies the current revision of a resource policy.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BotVersionLocaleSpecification

Next
