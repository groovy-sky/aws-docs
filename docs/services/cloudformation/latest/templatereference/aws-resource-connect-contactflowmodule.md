This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlowModule

Specifies a flow module for an Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ContactFlowModule",
  "Properties" : {
      "Content" : String,
      "Description" : String,
      "ExternalInvocationConfiguration" : ExternalInvocationConfiguration,
      "InstanceArn" : String,
      "Name" : String,
      "Settings" : String,
      "State" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ContactFlowModule
Properties:
  Content: String
  Description: String
  ExternalInvocationConfiguration:
    ExternalInvocationConfiguration
  InstanceArn: String
  Name: String
  Settings: String
  State: String
  Tags:
    - Tag

```

## Properties

`Content`

The content of the flow module.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the flow module.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalInvocationConfiguration`

The external invocation configuration for the flow module

_Required_: No

_Type_: [ExternalInvocationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-contactflowmodule-externalinvocationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow module.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

The configuration settings for the flow module.

_Required_: No

_Type_: String

_Maximum_: `256000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the flow module.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-contactflowmodule-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the flow module. For example:

`{ "Ref": "myFlowModuleArn" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ContactFlowModuleArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the flow module.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`Status`

Property description not available.

## Examples

### Specify a flow module resource

The following example specifies a flow module resource for an Amazon Connect instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a flow module for an Amazon Connect instance
Resources:
  cf11:
    Type: 'AWS::Connect::ContactFlowModule'
    Properties:
      Name: ExampleFlowModule
      Description: flow module created using cfn
      InstanceArn: arn:aws:connect:region-name:aws-account-id:instance/instance-arn
      Content: ExampleFlowModule content(JSON) using Amazon Connect Flow Language.
      Tags:
        - Key: testkey
          Value: testValue
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ExternalInvocationConfiguration
