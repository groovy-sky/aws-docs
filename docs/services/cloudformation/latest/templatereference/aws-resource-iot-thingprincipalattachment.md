This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingPrincipalAttachment

Use the `AWS::IoT::ThingPrincipalAttachment` resource to attach a principal
(an X.509 certificate or another credential) to a thing.

For more information about working with AWS IoT things and principals, see
[Authorization](../../../iot/latest/developerguide/authorization.md) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ThingPrincipalAttachment",
  "Properties" : {
      "Principal" : String,
      "ThingName" : String,
      "ThingPrincipalType" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ThingPrincipalAttachment
Properties:
  Principal: String
  ThingName: String
  ThingPrincipalType: String

```

## Properties

`Principal`

The principal, which can be a certificate ARN (as returned from the
`CreateCertificate` operation) or an Amazon Cognito ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingName`

The name of the AWS IoT thing.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingPrincipalType`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example attaches a principal to a thing.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "NameParameter": {
      "Type": "String",
      "Description": "Name of the IoT Thing to attach the principal to"
    }
  },
  "Resources": {
    "MyThingPrincipalAttachment": {
      "Type": "AWS::IoT::ThingPrincipalAttachment",
      "Properties": {
        "ThingName": {
          "Ref": "NameParameter"
        },
        "Principal": "arn:aws:iot:ap-southeast-2:123456789012:cert/a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyThingPrincipalAttachment:
    Type: AWS::IoT::ThingPrincipalAttachment
    Properties:
      ThingName:
        Ref: NameParameter
      Principal: arn:aws:iot:ap-southeast-2:123456789012:cert/a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2

Parameters:
  NameParameter:
    Type: String
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThingGroupProperties

AWS::IoT::ThingType

All content copied from https://docs.aws.amazon.com/.
