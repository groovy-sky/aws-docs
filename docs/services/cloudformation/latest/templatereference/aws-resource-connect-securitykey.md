This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityKey

The security key for the instance.

###### Note

Only two security keys are allowed per Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::SecurityKey",
  "Properties" : {
      "InstanceId" : String,
      "Key" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::SecurityKey
Properties:
  InstanceId: String
  Key: String

```

## Properties

`InstanceId`

The Amazon Resource Name (ARN) of the instance.

_Minimum_: `1`

_Maximum_: `100`

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Key`

A valid security key in PEM format. For example:

`"-----BEGIN PUBLIC KEY-----\ [a lot of characters] ----END PUBLIC KEY-----"`

_Minimum_: `1`

_Maximum_: `1024`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

An `AssociationId` is automatically generated when a storage config is
associated with an instance.

## Examples

### Specify a security key for an Amazon Connect instance

The following example specifies a security key for an Amazon Connect
instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a security key for an Amazon Connect instance
Resources:
  SecurityKey:
    Type: AWS::Connect::SecurityKey
    Properties:
      InstanceId: arn:aws:connect:region-name:aws-account-id:instance/instance-arn
      Key: "valid-pem-key"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateCaseAction

AWS::Connect::SecurityProfile
