This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::PermissionSet

Specifies a permission set within a specified IAM Identity Center instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::PermissionSet",
  "Properties" : {
      "CustomerManagedPolicyReferences" : [ CustomerManagedPolicyReference, ... ],
      "Description" : String,
      "InlinePolicy" : Json,
      "InstanceArn" : String,
      "ManagedPolicies" : [ String, ... ],
      "Name" : String,
      "PermissionsBoundary" : PermissionsBoundary,
      "RelayStateType" : String,
      "SessionDuration" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSO::PermissionSet
Properties:
  CustomerManagedPolicyReferences:
    - CustomerManagedPolicyReference
  Description: String
  InlinePolicy: Json
  InstanceArn: String
  ManagedPolicies:
    - String
  Name: String
  PermissionsBoundary:
    PermissionsBoundary
  RelayStateType: String
  SessionDuration: String
  Tags:
    - Tag

```

## Properties

`CustomerManagedPolicyReferences`

Specifies the names and paths of the customer managed policies that you have attached
to your permission set.

_Required_: No

_Type_: Array of [CustomerManagedPolicyReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sso-permissionset-customermanagedpolicyreference.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-sso-permissionset.html).

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007E\u00A1-\u00FF]*`

_Minimum_: `1`

_Maximum_: `700`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InlinePolicy`

The inline policy that is attached to the permission set.

###### Note

For `Length Constraints`, if a valid ARN is provided for a permission
set, it is possible for an empty inline policy to be returned.

_Required_: No

_Type_: Json

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `32768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The ARN of the IAM Identity Center instance under which the operation will be executed.
For more information about ARNs, see [Amazon Resource Names (ARNs) and\
AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedPolicies`

A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.

_Required_: No

_Type_: Array of String

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the permission set.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionsBoundary`

Specifies the configuration of the AWS managed or customer managed policy that you
want to set as a permissions boundary. Specify either
`CustomerManagedPolicyReference` to use the name and path of a customer
managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed
policy. A permissions boundary represents the maximum permissions that any policy can
grant your role. For more information, see [Permissions boundaries\
for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the _IAM User Guide_.

###### Important

Policies used as permissions boundaries don't provide permissions. You must also
attach an IAM policy to the role. To learn how the effective permissions for a
role are evaluated, see [IAM JSON\
policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the _IAM User_
_Guide_.

_Required_: No

_Type_: [PermissionsBoundary](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sso-permissionset-permissionsboundary.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelayStateType`

Used to redirect users within the application during the federation authentication
process.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9&amp;$@#\/%?=~\-_'&quot;|!:,.;*+\[\]\ \(\)\{\}]+`

_Minimum_: `1`

_Maximum_: `240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionDuration`

The length of time that the application user sessions are valid for in the ISO-8601
standard.

_Required_: No

_Type_: String

_Pattern_: `^(-?)P(?=\d|T\d)(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)([DW]))?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+(?:\.\d+)?)S)?)?$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to attach to the new [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-sso-permissionset.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sso-permissionset-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, such as
`permission-arn|sso-instance-arn`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PermissionSetArn`

The permission set ARN of the permission set, such as
`arn:aws:sso:::permissionSet/ins-instanceid/ps-permissionsetid`.

## Examples

- [Creating a new custom permission set for IAM Identity Center](#aws-resource-sso-permissionset--examples--Creating_a_new_custom_permission_set_for)

- [Creating a new custom permission set for IAM Identity Center with a customer managed policy as a permissions boundary](#aws-resource-sso-permissionset--examples--Creating_a_new_custom_permission_set_for_with_a_customer_managed_policy_as_a_permissions_boundary)

- [Creating a new custom permission set for IAM Identity Center with an AWS managed policy as a permissions boundary](#aws-resource-sso-permissionset--examples--Creating_a_new_custom_permission_set_for_with_an_managed_policy_as_a_permissions_boundary)

### Creating a new custom permission set for IAM Identity Center

The following example creates a custom permission set, `PermissionSet`,
with a managed policies attachment and inline policy.

#### JSON

```json

{
  "PermissionSet": {
    "Type": "AWS::SSO::PermissionSet",
    "Properties": {
      "InstanceArn": "arn:aws:sso:::instance/ssoins-instanceId",
      "Name": "PermissionSet",
      "Description": "This is a sample permission set.",
      "SessionDuration": "PT8H",
      "ManagedPolicies": [
         "arn:aws:iam::aws:policy/AdministratorAccess"
      ],
      "InlinePolicy": "Inline policy json string",
      "Tags": [
        {
          "Key": "tagKey",
          "Value": "tagValue"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

PermissionSet:
  Type: AWS::SSO::PermissionSet
    Properties:
    InstanceArn: 'arn:aws:sso:::instance/ssoins-instanceId'
    Name: 'PermissionSet'
    Description: 'This is a sample permission set.'
    SessionDuration: 'PT8H'
    ManagedPolicies:
      - 'arn:aws:iam::aws:policy/AdministratorAccess'
    InlinePolicy: 'Inline policy json string'
    Tags:
      - Key: tagKey
        Value: tagValue

```

### Creating a new custom permission set for IAM Identity Center with a customer managed policy as a permissions boundary

The following example creates a custom permission set,
`PermissionSetWithCmpPb`, with policies attached and a customer managed
policy as a permissions boundary.

#### JSON

```json

{
  "PermissionSetWithCustomerManagedPolicyReferenceForPermissionsBoundary": {
    "Type": "AWS::SSO::PermissionSet",
    "Properties": {
      "InstanceArn": "arn:aws:sso:::instance/ssoins-instanceId",
      "Name": "PermissionSetWithCmpPb",
      "Description": "This is a sample permission set.",
      "SessionDuration": "PT8H",
      "ManagedPolicies": [
        "arn:aws:iam::aws:policy/AdministratorAccess"
      ],
      "CustomerManagedPolicyReferences": [{
          "Name": "MyCustomPolicyName",
          "Path": "/myCustomPath/"
        },
        {
          "Name": "AnotherCustomPolicyName",
        },
        {
          "Name": "YetAnotherCustomPolicyName",
          "Path": "/"
        }
      ],
      "PermissionsBoundary": {
        "CustomerManagedPolicyReference": {
          "Name": "PolicyName",
          "Path": "/myPolicyPath/"
        }
      }
    }
  }
}

```

#### YAML

```yaml

PermissionSetWithCustomerManagedPolicyReferenceForPermissionsBoundary:
  Type: AWS::SSO::PermissionSet
   Properties:
     InstanceArn: 'arn:aws:sso:::instance/ssoins-instanceId'
     Name: 'PermissionSetWithCmpPb'
     Description: 'This is a sample permission set.'
     SessionDuration: 'PT8H'
     ManagedPolicies:
     - 'arn:aws:iam::aws:policy/AdministratorAccess'
     CustomerManagedPolicyReferences:
     - Name: 'MyCustomPolicyName'
       Path: '/myCustomPath/'
     - Name: 'AnotherCustomPolicyName'
     - Name: 'YetAnotherCustomPolicyName'
       Path: '/'
     PermissionsBoundary:
       CustomerManagedPolicyReference:
         Name: PolicyName
         Path: /myPolicyPath/

```

### Creating a new custom permission set for IAM Identity Center with an AWS managed policy as a permissions boundary

The following example creates a custom permission set,
`PermissionSetWithAmpPb`, with policies attached and an AWS managed policy as a permissions boundary.

#### JSON

```json

{
  "PermissionSetWithAWSManagedPolicyForPermissionsBoundary": {
    "Type": "AWS::SSO::PermissionSet",
    "Properties": {
      "InstanceArn": "arn:aws:sso:::instance/ssoins-instanceId",
      "Name": "PermissionSetWithAmpPb",
      "Description": "This is a sample permission set.",
      "SessionDuration": "PT8H",
      "ManagedPolicies": [
        "arn:aws:iam::aws:policy/AdministratorAccess"
      ],
      "CustomerManagedPolicyReferences": [{
          "Name": "MyCustomPolicyName",
          "Path": "/myCustomPath/"
        },
        {
          "Name": "AnotherCustomPolicyName",
        },
        {
          "Name": "YetAnotherCustomPolicyName",
          "Path": "/"
        }
      ],
      "PermissionsBoundary": {
        "ManagedPolicyArn": {
          "Fn::Sub": "arn:aws:iam::aws:policy/ReadOnlyAccess"
        }
      }
    }
  }
}

```

#### YAML

```yaml

PermissionSetWithAwsManagedPolicyForPermissionsBoundary:
  Type: AWS::SSO::PermissionSet
    Properties:
      InstanceArn: 'arn:aws:sso:::instance/ssoins-instanceId'
      Name: 'PermissionSetWithAmpPb'
      Description: 'This is a sample permission set.'
      SessionDuration: 'PT8H'
      ManagedPolicies:
         - 'arn:aws:iam::aws:policy/AdministratorAccess'
      CustomerManagedPolicyReferences:
        - Name: 'MyCustomPolicy'
          Path: '/myCustomPath/'
        - Name: 'AnotherCustomPolicy'
        - Name: YetAnotherCustomPolicyName
          Path: /
      PermissionsBoundary:
        ManagedPolicyArn: arn:aws:iam::aws:policy/ReadOnlyAccess'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessControlAttributeValue

CustomerManagedPolicyReference
