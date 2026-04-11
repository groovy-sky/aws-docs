This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Parameter

The `AWS::SSM::Parameter` resource creates an SSM parameter in AWS Systems Manager Parameter Store.

###### Note

To create an SSM parameter, you must have the AWS Identity and Access Management (IAM) permissions `ssm:PutParameter` and
`ssm:AddTagsToResource`. On stack creation, AWS CloudFormation
adds the following three tags to the parameter:
`aws:cloudformation:stack-name`,
`aws:cloudformation:logical-id`, and
`aws:cloudformation:stack-id`, in addition to any custom tags you
specify.

To add, update, or remove tags during stack update, you must have IAM permissions for both `ssm:AddTagsToResource` and
`ssm:RemoveTagsFromResource`. For more information, see [Managing access using policies](../../../systems-manager/latest/userguide/security-iam.md#security_iam_access-manage) in the _AWS Systems Manager User Guide_.

For information about valid values for parameters, see [About requirements and constraints for parameter names](../../../systems-manager/latest/userguide/sysman-paramstore-su-create.md#sysman-parameter-name-constraints) in the
_AWS Systems Manager User Guide_ and [PutParameter](../../../../reference/systems-manager/latest/apireference/api-putparameter.md) in the _AWS Systems Manager API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::Parameter",
  "Properties" : {
      "AllowedPattern" : String,
      "DataType" : String,
      "Description" : String,
      "Name" : String,
      "Policies" : String,
      "Tags" : {Key: Value, ...},
      "Tier" : String,
      "Type" : String,
      "Value" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::Parameter
Properties:
  AllowedPattern: String
  DataType: String
  Description: String
  Name: String
  Policies: String
  Tags:
    Key: Value
  Tier: String
  Type: String
  Value: String

```

## Properties

`AllowedPattern`

A regular expression used to validate the parameter value. For example, for
`String` types with values restricted to numbers, you can specify the
following: `AllowedPattern=^\d+$`

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The data type of the parameter, such as `text` or `aws:ec2:image`. The
default is `text`.

_Required_: No

_Type_: String

_Allowed values_: `text | aws:ec2:image`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Information about the parameter.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the parameter.

###### Note

The reported maximum length of 2048 characters for a parameter name includes 1037
characters that are reserved for internal use by Systems Manager. The maximum
length for a parameter name that you specify is 1011 characters.

This count of 1011 characters includes the characters in the ARN that precede the
name you specify. This ARN length will vary depending on your partition and Region.
For example, the following 45 characters count toward the 1011 character maximum for
a parameter created in the US East (Ohio) Region:
`arn:aws:ssm:us-east-2:111122223333:parameter/`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policies`

Information about the policies assigned to a parameter.

[Assigning parameter\
policies](../../../systems-manager/latest/userguide/parameter-store-policies.md) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you assign to a resource in the form of an arbitrary set of
tags (key-value pairs). Tags enable you to categorize a resource in different ways, such
as by purpose, owner, or environment. For example, you might want to tag a Systems Manager parameter to identify the type of resource to which it applies, the
environment, or the type of configuration data referenced by the parameter.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The parameter tier.

_Required_: No

_Type_: String

_Allowed values_: `Standard | Advanced | Intelligent-Tiering`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of parameter.

###### Note

Parameters of type `SecureString` are not supported by AWS CloudFormation.

_Required_: Yes

_Type_: String

_Allowed values_: `String | StringList`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The parameter value.

###### Note

If type is `StringList`, the system returns a comma-separated string with no
spaces between commas in the `Value` field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the SSM parameter. For example,
`ssm-myparameter-ABCNPH3XCAO6`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

###### Note

Due to eventual consistency of the underlying API, a `{ Fn::GetValue }` of an SSM
Parameter that was just created may fail. Either avoid using `{ Fn::GetAtt }` on an
SSM Parameter, or be aware that stack creation may sometimes fail and you will need
to retry.

`Type`

Returns the type of the parameter. Valid values are `String` or
`StringList`.

`Value`

Returns the value of the parameter.

## Examples

- [Create a String-type parameter](#aws-resource-ssm-parameter--examples--Create_a_String-type_parameter)

- [Create a StringList-type parameter](#aws-resource-ssm-parameter--examples--Create_a_StringList-type_parameter)

- [Create an advanced tier parameter and assign a policy](#aws-resource-ssm-parameter--examples--Create_an_advanced_tier_parameter_and_assign_a_policy)

### Create a String-type parameter

The following example creates a Systems Manager parameter named command
with a `String` type and adds the tag key-value pair
`"Environment":"Dev"`.

#### JSON

```json

{
    "Resources": {
        "BasicParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Name": "command",
                "Type": "String",
                "Value": "date",
                "Description": "SSM Parameter for running date command.",
                "AllowedPattern": "^[a-zA-Z]{1,10}$",
                "Tags": {
                    "Environment": "DEV"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  BasicParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Name: command
      Type: String
      Value: date
      Description: SSM Parameter for running date command.
      AllowedPattern: "^[a-zA-Z]{1,10}$"
      Tags:
        Environment: DEV
```

### Create a StringList-type parameter

The following example creates a Systems Manager parameter named commands
with a `StringList` type.

#### JSON

```json

{
    "Resources": {
        "BasicParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Name": "commands",
                "Type": "StringList",
                "Value": "date,ls",
                "Description": "SSM Parameter of type StringList.",
                "AllowedPattern": "^[a-zA-Z]{1,10}$"
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  BasicParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Name: commands
      Type: StringList
      Value: date,ls
      Description: SSM parameter of type StringList.
      AllowedPattern: "^[a-zA-Z]{1,10}$"
```

### Create an advanced tier parameter and assign a policy

The following example creates a Systems Manager advanced tier parameter
named 'command' with a `String` type and a parameter policy.

#### JSON

```json

{
    "Resources": {
        "BasicParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Name": "command",
                "Type": "String",
                "Value": "date",
                "Tier": "Advanced",
                "Policies": "[{\"Type\":\"Expiration\",\"Version\":\"1.0\",\"Attributes\":{\"Timestamp\":\"2020-05-13T00:00:00.000Z\"}},{\"Type\":\"ExpirationNotification\",\"Version\":\"1.0\",\"Attributes\":{\"Before\":\"5\",\"Unit\":\"Days\"}},{\"Type\":\"NoChangeNotification\",\"Version\":\"1.0\",\"Attributes\":{\"After\":\"60\",\"Unit\":\"Days\"}}]",
                "Description": "SSM Parameter for running date command.",
                "AllowedPattern": "^[a-zA-Z]{1,10}$",
                "Tags": {
                    "Environment": "DEV"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  BasicParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Name: command
      Type: String
      Value: date
      Tier: Advanced
      Policies: '[{"Type":"Expiration","Version":"1.0","Attributes":{"Timestamp":"2020-05-13T00:00:00.000Z"}},{"Type":"ExpirationNotification","Version":"1.0","Attributes":{"Before":"5","Unit":"Days"}},{"Type":"NoChangeNotification","Version":"1.0","Attributes":{"After":"60","Unit":"Days"}}]'
      Description: SSM parameter for running date command.
      AllowedPattern: "^[a-zA-Z]{1,10}$"
      Tags:
        Environment: DEV
```

## See also

- [AWS Systems Manager Parameter Store](../../../systems-manager/latest/userguide/systems-manager-parameter-store.md)

- [Managing parameters tiers](../../../systems-manager/latest/userguide/parameter-store-advanced-parameters.md)

- [Assigning parameter policies](../../../systems-manager/latest/userguide/parameter-store-policies.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TaskInvocationParameters

AWS::SSM::PatchBaseline

All content copied from https://docs.aws.amazon.com/.
