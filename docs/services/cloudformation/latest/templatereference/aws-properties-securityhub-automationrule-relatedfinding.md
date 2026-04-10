This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule RelatedFinding

Provides details about a list of findings that the current finding relates to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "ProductArn" : String
}

```

### YAML

```yaml

  Id: String
  ProductArn: String

```

## Properties

`Id`

The product-generated identifier for a related finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductArn`

The Amazon Resource Name (ARN) for the product that generated a related finding.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-iso-?[a-z]{0,2}):[A-Za-z0-9]{1,63}:[a-z]+-([a-z]{1,10}-)?[a-z]+-[0-9]+:([0-9]{12})?:.+$`

_Minimum_: `12`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberFilter

SeverityUpdate

All content copied from https://docs.aws.amazon.com/.
