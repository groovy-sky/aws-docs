This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::AccessKey

Creates a new AWS secret access key and corresponding AWS
access key ID for the specified user. The default status for new keys is
`Active`.

For information about quotas on the number of keys you can create, see [IAM and\
AWS STS quotas](../../../iam/latest/userguide/reference-iam-quotas.md) in the _IAM User Guide_.

###### Important

To ensure the security of your AWS account, the secret access key is
accessible only during key and user creation. You must save the key (for example, in a
text file) if you want to be able to access it again. If a secret key is lost, you can
rotate access keys by increasing the value of the `serial` property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::AccessKey",
  "Properties" : {
      "Serial" : Integer,
      "Status" : String,
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::AccessKey
Properties:
  Serial: Integer
  Status: String
  UserName: String

```

## Properties

`Serial`

This value is specific to CloudFormation and can only be
_incremented_. Incrementing this value notifies CloudFormation that you want to rotate your access key. When you update your stack,
CloudFormation will replace the existing access key with a new key.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The status of the access key. `Active` means that the key is valid for API
calls, while `Inactive` means it is not.

_Required_: No

_Type_: String

_Allowed values_: `Active | Inactive | Expired`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The name of the IAM user that the new key will belong to.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AccessKeyId`. For example:
`AWS_ACCESS_KEY_ID_REDACTED`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for this access key.

`SecretAccessKey`

Returns the secret access key for the specified AWS::IAM::AccessKey resource. For
example: wJalrXUtnFEMI/K7MDENG/bPxRfiCYzEXAMPLEKEY.

## See also

- To view `AWS::IAM::AccessKey` template example snippets, see [Declaring an IAM Access Key Resource](../userguide/quickref-iam.md#scenario-iam-accesskey).

- [CreateAccessKey](../../../../reference/iam/latest/apireference/api-createaccesskey.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Identity and Access Management

AWS::IAM::Group

All content copied from https://docs.aws.amazon.com/.
