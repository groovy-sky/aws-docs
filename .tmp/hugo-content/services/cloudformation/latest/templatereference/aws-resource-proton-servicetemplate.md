This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Proton::ServiceTemplate

Create a service template. The administrator creates a service template to define
standardized infrastructure and an optional CI/CD service pipeline. Developers, in turn,
select the service template from AWS Proton. If the selected service template includes a
service pipeline definition, they provide a link to their source code repository. AWS Proton
then deploys and manages the infrastructure defined by the selected service template. For more
information, see [AWS Proton templates](../../../proton/latest/userguide/ag-templates.md) in the _AWS Proton User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Proton::ServiceTemplate",
  "Properties" : {
      "Description" : String,
      "DisplayName" : String,
      "EncryptionKey" : String,
      "Name" : String,
      "PipelineProvisioning" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Proton::ServiceTemplate
Properties:
  Description: String
  DisplayName: String
  EncryptionKey: String
  Name: String
  PipelineProvisioning: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the service template.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The service template name as displayed in the developer interface.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKey`

The customer provided service template encryption key that's used to encrypt data.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov):[a-zA-Z0-9-]+:[a-zA-Z0-9-]*:\d{12}:([\w+=,.@-]+[/:])*[\w+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the service template.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z]+[0-9A-Za-z_\-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PipelineProvisioning`

If `pipelineProvisioning` is `true`, a service pipeline is included
in the service template. Otherwise, a service pipeline _isn't_ included in
the service template.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An object that includes the template bundle S3 bucket path and name for the new version of
a service template.

_Required_: No

_Type_: Array of [Tag](aws-properties-proton-servicetemplate-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the service template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the service template ARN.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
