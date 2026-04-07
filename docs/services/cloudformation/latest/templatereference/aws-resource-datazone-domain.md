This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Domain

The `AWS::DataZone::Domain` resource specifies an Amazon DataZone domain. You can use domains to
organize your assets, users, and their projects.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::Domain",
  "Properties" : {
      "Description" : String,
      "DomainExecutionRole" : String,
      "DomainVersion" : String,
      "KmsKeyIdentifier" : String,
      "Name" : String,
      "ServiceRole" : String,
      "SingleSignOn" : SingleSignOn,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::Domain
Properties:
  Description: String
  DomainExecutionRole: String
  DomainVersion: String
  KmsKeyIdentifier: String
  Name: String
  ServiceRole: String
  SingleSignOn:
    SingleSignOn
  Tags:
    - Tag

```

## Properties

`Description`

The description of the Amazon DataZone domain.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainExecutionRole`

The domain execution role that is created when an Amazon DataZone domain is created. The
domain execution role is created in the AWS account that houses the
Amazon DataZone domain.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainVersion`

The domain version.

_Required_: No

_Type_: String

_Allowed values_: `V1 | V2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyIdentifier`

The identifier of the AWS Key Management Service (KMS) key that is used
to encrypt the Amazon DataZone domain, metadata, and reporting data.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the Amazon DataZone domain.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRole`

The service role of the domain.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleSignOn`

The single sign-on details in Amazon DataZone.

_Required_: No

_Type_: [SingleSignOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-domain-singlesignon.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags specified for the Amazon DataZone domain.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-domain-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Amazon DataZone domain.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the Amazon DataZone domain.

`CreatedAt`

A timestamp of when a Amazon DataZone domain was created.

`Id`

The ID of the Amazon DataZone domain.

`LastUpdatedAt`

A timestamp of when a Amazon DataZone domain was last updated.

`ManagedAccountId`

The identifier of the AWS account that manages the domain.

`PortalUrl`

The data portal URL for the Amazon DataZone domain.

`RootDomainUnitId`

The ID of the root domain unit.

`Status`

The status of the Amazon DataZone domain.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScheduleConfiguration

SingleSignOn
