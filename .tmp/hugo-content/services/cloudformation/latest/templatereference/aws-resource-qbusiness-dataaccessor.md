This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor

Creates a new data accessor for an ISV to access data from a Amazon Q Business
application. The data accessor is an entity that represents the ISV's access to the
Amazon Q Business application's data. It includes the IAM role ARN for the ISV, a friendly
name, and a set of action configurations that define the specific actions the ISV is
allowed to perform and any associated data filters. When the data accessor is created,
an IAM Identity Center application is also created to manage the ISV's identity and
authentication for accessing the Amazon Q Business application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::DataAccessor",
  "Properties" : {
      "ActionConfigurations" : [ ActionConfiguration, ... ],
      "ApplicationId" : String,
      "AuthenticationDetail" : DataAccessorAuthenticationDetail,
      "DisplayName" : String,
      "Principal" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::DataAccessor
Properties:
  ActionConfigurations:
    - ActionConfiguration
  ApplicationId: String
  AuthenticationDetail:
    DataAccessorAuthenticationDetail
  DisplayName: String
  Principal: String
  Tags:
    - Tag

```

## Properties

`ActionConfigurations`

A list of action configurations specifying the allowed actions and any associated
filters.

_Required_: Yes

_Type_: Array of [ActionConfiguration](aws-properties-qbusiness-dataaccessor-actionconfiguration.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationId`

The unique identifier of the Amazon Q Business application.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthenticationDetail`

The authentication configuration details for the data accessor. This specifies how the
ISV authenticates when accessing data through this data accessor.

_Required_: No

_Type_: [DataAccessorAuthenticationDetail](aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationdetail.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The friendly name of the data accessor.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

The Amazon Resource Name (ARN) of the IAM role for the ISV associated with this data
accessor.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws:iam::[0-9]{12}:role/[a-zA-Z0-9_/+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to associate with the data accessor.

_Required_: No

_Type_: Array of [Tag](aws-properties-qbusiness-dataaccessor-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application and data accessor ID. For example:

`{"Ref": "ApplicationId|DataAccessorId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the data accessor was created.

`DataAccessorArn`

The Amazon Resource Name (ARN) of the data accessor.

`DataAccessorId`

The unique identifier of the data accessor.

`IdcApplicationArn`

The Amazon Resource Name (ARN) of the associated IAM Identity Center
application.

`UpdatedAt`

The timestamp when the data accessor was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ActionConfiguration

All content copied from https://docs.aws.amazon.com/.
