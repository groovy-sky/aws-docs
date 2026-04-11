This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::PolicyPrincipalAttachment

Use the `AWS::IoT::PolicyPrincipalAttachment` resource to attach an AWS IoT policy to a principal (an X.509 certificate or other credential).

For information about working with AWS IoT policies and principals, see
[Authorization](../../../iot/latest/developerguide/authorization.md) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::PolicyPrincipalAttachment",
  "Properties" : {
      "PolicyName" : String,
      "Principal" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::PolicyPrincipalAttachment
Properties:
  PolicyName: String
  Principal: String

```

## Properties

`PolicyName`

The name of the AWS IoT policy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

The principal, which can be a certificate ARN (as returned from the
`CreateCertificate` operation) or an Amazon Cognito ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

The following example attaches a policy to a principal.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyPolicyPrincipalAttachment": {
      "Type": "AWS::IoT::PolicyPrincipalAttachment",
      "Properties": {
        "PolicyName": {
          "Ref": "NameParameter"
        },
        "Principal": "arn:aws:iot:ap-southeast-2:123456789012:cert/a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2"
      }
    }
  },
  "Parameters": {
    "NameParameter": {
      "Type": "String"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyPolicyPrincipalAttachment:
    Type: AWS::IoT::PolicyPrincipalAttachment
    Properties:
      PolicyName:
        Ref: NameParameter
      Principal: arn:aws:iot:ap-southeast-2:123456789012:cert/a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2

Parameters:
  NameParameter:
    Type: String
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoT::ProvisioningTemplate

All content copied from https://docs.aws.amazon.com/.
