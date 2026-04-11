This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::PredefinedAttribute

Textual or numeric value that describes an attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::PredefinedAttribute",
  "Properties" : {
      "AttributeConfiguration" : AttributeConfiguration,
      "InstanceArn" : String,
      "Name" : String,
      "Purposes" : [ String, ... ],
      "Values" : Values
    }
}

```

### YAML

```yaml

Type: AWS::Connect::PredefinedAttribute
Properties:
  AttributeConfiguration:
    AttributeConfiguration
  InstanceArn: String
  Name: String
  Purposes:
    - String
  Values:
    Values

```

## Properties

`AttributeConfiguration`

Custom metadata that is associated to predefined attributes to control behavior
in upstream services, such as controlling
how a predefined attribute should be displayed in the Amazon Connect admin website.

_Required_: No

_Type_: [AttributeConfiguration](aws-properties-connect-predefinedattribute-attributeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the predefined attribute.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Purposes`

Values that enable you to categorize your predefined attributes. You can use them in custom UI elements across the Amazon Connect admin website.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values of a predefined attribute.

_Required_: No

_Type_: [Values](aws-properties-connect-predefinedattribute-values.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the predefined attribute. For example:

`{ "Ref": "myPredefinedAttribute" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LastModifiedRegion`

Last modified region.

`LastModifiedTime`

Last modified time.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AttributeConfiguration

All content copied from https://docs.aws.amazon.com/.
