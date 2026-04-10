This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::CredentialLocker

Create a credential locker.

###### Note

This operation will not trigger the creation of all the manufacturing
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTManagedIntegrations::CredentialLocker",
  "Properties" : {
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::IoTManagedIntegrations::CredentialLocker
Properties:
  Name: String
  Tags:
    Key: Value

```

## Properties

`Name`

The name of the credential locker.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A set of key/value pairs that are used to manage the credential locker.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the credential locker.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the credential locker.

`CreatedAt`

The timestamp value of when the credential locker requset occurred.

`Id`

The identifier of the credential locker.

`Identifier`

The identifier of the credential locker.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Managed integrations for AWS IoT Device Management

AWS::IoTManagedIntegrations::ManagedThing

All content copied from https://docs.aws.amazon.com/.
