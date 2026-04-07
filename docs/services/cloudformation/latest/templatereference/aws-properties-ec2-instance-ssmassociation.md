This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance SsmAssociation

Specifies the SSM document and parameter values in AWS Systems Manager to associate
with an instance.

`SsmAssociations` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociationParameters" : [ AssociationParameter, ... ],
  "DocumentName" : String
}

```

### YAML

```yaml

  AssociationParameters:
    - AssociationParameter
  DocumentName: String

```

## Properties

`AssociationParameters`

The input parameter values to use with the associated SSM document.

_Required_: No

_Type_: Array of [AssociationParameter](aws-properties-ec2-instance-associationparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentName`

The name of an SSM document to associate with the instance.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PrivateIpAddressSpecification

State
