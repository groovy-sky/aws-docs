This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::M2::Application

Specifies a new application with given parameters. Requires an existing runtime
environment and application definition file.

For information about application definitions, see the [AWS Mainframe Modernization User Guide](../../../m2/latest/userguide/applications-m2-definition.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::M2::Application",
  "Properties" : {
      "Definition" : Definition,
      "Description" : String,
      "EngineType" : String,
      "KmsKeyId" : String,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::M2::Application
Properties:
  Definition:
    Definition
  Description: String
  EngineType: String
  KmsKeyId: String
  Name: String
  RoleArn: String
  Tags:
    Key: Value

```

## Properties

`Definition`

The application definition for a particular application. You can specify either inline
JSON or an Amazon S3 bucket location.

For information about application definitions, see the [AWS Mainframe Modernization User Guide](../../../m2/latest/userguide/applications-m2-definition.md).

_Required_: No

_Type_: [Definition](aws-properties-m2-application-definition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the application.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineType`

The type of the target platform for this application.

_Required_: Yes

_Type_: String

_Allowed values_: `microfocus | bluage`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The identifier of a customer managed key.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the role associated with the application.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov):[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:([a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]|):[0-9]{12}:[A-Za-z0-9/][A-Za-z0-9:_/+=,@.-]{0,1023}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:).+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application Amazon Resource Name (ARN), such as the
following:

`{ "Ref": “SampleApp” }`

Returns a value similar to the following:

`arn:aws:m2:us-west-2:1234567890:app/y3ca6bhaife2bcvxar3lpivfou`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationArn`

The Amazon Resource Name (ARN) of the application.

`ApplicationId`

The identifier of the application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Mainframe Modernization

Definition

All content copied from https://docs.aws.amazon.com/.
