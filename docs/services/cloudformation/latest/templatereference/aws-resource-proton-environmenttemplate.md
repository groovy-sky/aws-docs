This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Proton::EnvironmentTemplate

Create an environment template for AWS Proton. For more information, see [Environment Templates](../../../proton/latest/userguide/ag-templates.md) in the _AWS Proton User Guide_.

You can create an environment template in one of the two following ways:

- Register and publish a _standard_ environment template that instructs AWS Proton to deploy and manage environment
infrastructure.

- Register and publish a _customer managed_ environment template that connects AWS Proton to your existing provisioned
infrastructure that you manage. AWS Proton _doesn't_ manage your existing provisioned infrastructure. To create an environment
template for customer provisioned and managed infrastructure, include the `provisioning` parameter and set the value to
`CUSTOMER_MANAGED`. For more information, see [Register\
and publish an environment template](../../../proton/latest/userguide/template-create.md) in the _AWS Proton User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Proton::EnvironmentTemplate",
  "Properties" : {
      "Description" : String,
      "DisplayName" : String,
      "EncryptionKey" : String,
      "Name" : String,
      "Provisioning" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Proton::EnvironmentTemplate
Properties:
  Description: String
  DisplayName: String
  EncryptionKey: String
  Name: String
  Provisioning: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the environment template.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the environment template as displayed in the developer interface.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKey`

The customer provided encryption key for the environment template.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov):[a-zA-Z0-9-]+:[a-zA-Z0-9-]*:\d{12}:([\w+=,.@-]+[/:])*[\w+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the environment template.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z]+[0-9A-Za-z_\-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Provisioning`

When included, indicates that the environment template is for customer provisioned and managed infrastructure.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional list of metadata items that you can associate with the AWS Proton environment template. A tag is a key-value pair.

For more information, see [AWS Proton resources and tagging](../../../proton/latest/userguide/resources.md) in the
_AWS Proton User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-proton-environmenttemplate-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the environment template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the ARN of the environment template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
