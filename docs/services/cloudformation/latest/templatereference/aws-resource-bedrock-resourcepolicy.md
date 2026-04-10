This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::ResourcePolicy

Adds a resource policy for a Bedrock resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::ResourcePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::ResourcePolicy
Properties:
  PolicyDocument: Json
  ResourceArn: String

```

## Properties

`PolicyDocument`

Property description not available.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]+)*:bedrock:[a-z0-9-]+:[0-9]{12}:(guardrail|guardrail-profile)/[a-z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolSpecification

Next

All content copied from https://docs.aws.amazon.com/.
