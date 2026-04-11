This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowEntitlement

The `AWS::MediaConnect::FlowEntitlement` resource defines the permission that an AWS account grants to another AWS account to allow access
to the content in a specific AWS Elemental MediaConnect flow. The
content originator grants an entitlement to a specific AWS account (the
subscriber). When an entitlement is granted, the subscriber can create a flow using the
originator's flow as the source. Each flow can have up to 50 entitlements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::FlowEntitlement",
  "Properties" : {
      "DataTransferSubscriberFeePercent" : Integer,
      "Description" : String,
      "Encryption" : Encryption,
      "EntitlementStatus" : String,
      "FlowArn" : String,
      "Name" : String,
      "Subscribers" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::FlowEntitlement
Properties:
  DataTransferSubscriberFeePercent: Integer
  Description: String
  Encryption:
    Encryption
  EntitlementStatus: String
  FlowArn: String
  Name: String
  Subscribers:
    - String
  Tags:
    - Tag

```

## Properties

`DataTransferSubscriberFeePercent`

The percentage of the entitlement data transfer fee that you want the subscriber
to be responsible for.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the entitlement. This description appears only on the
MediaConnect console and is not visible outside of the current AWS account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

Encryption information.

_Required_: No

_Type_: [Encryption](aws-properties-mediaconnect-flowentitlement-encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntitlementStatus`

An indication of whether the new entitlement should be enabled or disabled as soon
as it is created. If you don’t specify the entitlementStatus field in your request,
MediaConnect sets it to ENABLED.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the entitlement. This value must be unique within the current
flow.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subscribers`

The AWS account IDs that you want to share your content with. The receiving
accounts (subscribers) will be allowed to create their own flows using your content
as the source.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-flowentitlement-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the entitlement ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-west-2:111122223333:entitlement:1-11aa22bb11aa22bb-3333cccc4444:AnyCompany_Entitlement"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EntitlementArn`

The entitlement ARN.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcInterfaceAttachment

Encryption

All content copied from https://docs.aws.amazon.com/.
