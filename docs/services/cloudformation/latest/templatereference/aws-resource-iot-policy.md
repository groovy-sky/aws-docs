This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Policy

Use the `AWS::IoT::Policy` resource to declare an AWS IoT
policy. For more information about working with AWS IoT policies, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Policy",
  "Properties" : {
      "PolicyDocument" : Json,
      "PolicyName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Policy
Properties:
  PolicyDocument: Json
  PolicyName: String
  Tags:
    - Tag

```

## Properties

`PolicyDocument`

The JSON document that describes the policy.

_Required_: Yes

_Type_: Json

_Minimum_: `1`

_Maximum_: `404600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The policy name.

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-policy-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the policy name. For example:

`{ "Ref": "MyPolicy" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the AWS IoT policy, such as
`arn:aws:iot:us-east-2:123456789012:policy/MyPolicy`.

`Id`

The name of this policy.

## Examples

The following example declares an AWS IoT policy. This example grants
permission to connect to AWS IoT with client ID client1.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyIoTPolicy": {
      "Type": "AWS::IoT::Policy",
      "Properties": {
        "PolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Action": [
                "iot:Connect"
              ],
              "Resource": [
                "arn:aws:iot:us-east-1:123456789012:client/client1"
              ]
            }
          ]
        },
        "PolicyName": "PolicyName"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyIoTPolicy:
    Type: AWS::IoT::Policy
    Properties:
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - iot:Connect
            Resource:
              - arn:aws:iot:us-east-1:123456789012:client/client1
      PolicyName: PolicyName

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateDeviceCertificateParams

Tag
