This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SystemsManagerSAP::Application ComponentInfo

This is information about the component of
your SAP application, such as Web Dispatcher.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentType" : String,
  "Ec2InstanceId" : String,
  "Sid" : String
}

```

### YAML

```yaml

  ComponentType: String
  Ec2InstanceId: String
  Sid: String

```

## Properties

`ComponentType`

This string is the type of the component.

Accepted value is `WD`.

_Required_: No

_Type_: String

_Allowed values_: `HANA | HANA_NODE | ABAP | ASCS | DIALOG | WEBDISP | WD | ERS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ec2InstanceId`

This is the Amazon EC2 instance on which your SAP component is running.

Accepted values are alphanumeric.

_Required_: No

_Type_: String

_Pattern_: `^i-[\w\d]{8}$|^i-[\w\d]{17}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Sid`

This string is the SAP System ID of the component.

Accepted values are alphanumeric.

_Required_: No

_Type_: String

_Pattern_: `[A-Z][A-Z0-9]{2}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SystemsManagerSAP::Application

Credential
