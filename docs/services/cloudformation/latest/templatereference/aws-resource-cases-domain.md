This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Domain

Creates a domain, which is a container for all case data, such as cases, fields, templates
and layouts. Each Amazon Connect instance can be associated with only one Cases
domain.

###### Important

This will not associate your connect instance to Cases domain. Instead, use the
Amazon Connect
[CreateIntegrationAssociation](../../../../reference/connect/latest/apireference/api-createintegrationassociation.md) API. You need specific IAM
permissions to successfully associate the Cases domain. For more information, see
[Onboard to Cases](../../../connect/latest/adminguide/required-permissions-iam-cases.md#onboard-cases-iam).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cases::Domain",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cases::Domain
Properties:
  Name: String
  Tags:
    - Tag

```

## Properties

`Name`

The name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[\S]$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: Updates are not supported.

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cases-domain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the domain. For example:

`arn:aws:cases:us-west-2:123456789012:domain/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

The timestamp when the Cases domain was created.

`DomainArn`

The Amazon Resource Name (ARN) for the Cases domain.

`DomainId`

The unique identifier of the Cases domain.

`DomainStatus`

The status of the Cases domain.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
