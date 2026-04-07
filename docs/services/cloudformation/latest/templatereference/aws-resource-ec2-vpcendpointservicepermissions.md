This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCEndpointServicePermissions

Grant or revoke permissions for service consumers (users, IAM roles, and AWS accounts) to connect to a VPC endpoint service.

If you grant permissions to all principals, the service is public. Any users who know
the name of a public service can send a request to attach an endpoint. If the service does
not require manual approval, attachments are automatically approved.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCEndpointServicePermissions",
  "Properties" : {
      "AllowedPrincipals" : [ String, ... ],
      "ServiceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCEndpointServicePermissions
Properties:
  AllowedPrincipals:
    - String
  ServiceId: String

```

## Properties

`AllowedPrincipals`

The Amazon Resource Names (ARN) of one or more principals (for example, users, IAM roles, and
AWS accounts). Permissions are granted to the principals in this list.
To grant permissions to all principals, specify an asterisk (\*). Permissions are revoked
for principals not in this list. If the list is empty, then all permissions are revoked.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceId`

The ID of the service.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPC endpoint service permissions.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::VPCGatewayAttachment
