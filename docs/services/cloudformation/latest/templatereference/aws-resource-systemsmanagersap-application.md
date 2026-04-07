This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SystemsManagerSAP::Application

An SAP application registered with AWS Systems Manager for
SAP.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SystemsManagerSAP::Application",
  "Properties" : {
      "ApplicationId" : String,
      "ApplicationType" : String,
      "ComponentsInfo" : [ ComponentInfo, ... ],
      "Credentials" : [ Credential, ... ],
      "DatabaseArn" : String,
      "Instances" : [ String, ... ],
      "SapInstanceNumber" : String,
      "Sid" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SystemsManagerSAP::Application
Properties:
  ApplicationId: String
  ApplicationType: String
  ComponentsInfo:
    - ComponentInfo
  Credentials:
    - Credential
  DatabaseArn: String
  Instances:
    - String
  SapInstanceNumber: String
  Sid: String
  Tags:
    - Tag

```

## Properties

`ApplicationId`

The ID of the application.

_Required_: Yes

_Type_: String

_Pattern_: `[\w\d\.-]{1,60}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationType`

The type of the application.

_Required_: Yes

_Type_: String

_Allowed values_: `HANA | SAP_ABAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentsInfo`

Property description not available.

_Required_: No

_Type_: Array of [ComponentInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-systemsmanagersap-application-componentinfo.html)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Credentials`

The credentials of the SAP application.

_Required_: No

_Type_: Array of [Credential](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-systemsmanagersap-application-credential.html)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseArn`

The Amazon Resource Name (ARN) of the database.

_Required_: No

_Type_: String

_Pattern_: `^arn:(.+:){2,4}.+$|^arn:(.+:){1,3}.+\/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Instances`

The Amazon EC2 instances on which your SAP application is running.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SapInstanceNumber`

The SAP instance number of the application.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{2}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Sid`

The System ID of the application.

_Required_: No

_Type_: String

_Pattern_: `[A-Z][A-Z0-9]{2}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags on the application.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-systemsmanagersap-application-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name of the SAP application.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Systems Manager for SAP

ComponentInfo
