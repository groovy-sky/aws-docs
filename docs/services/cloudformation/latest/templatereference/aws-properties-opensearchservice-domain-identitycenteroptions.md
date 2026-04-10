This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain IdentityCenterOptions

Settings container for integrating IAM Identity Center with OpenSearch UI applications,
which enables enabling secure user authentication and access control across multiple data
sources. This setup supports single sign-on (SSO) through IAM Identity Center, allowing
centralized user management.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnabledAPIAccess" : Boolean,
  "IdentityCenterApplicationARN" : String,
  "IdentityCenterInstanceARN" : String,
  "IdentityStoreId" : String,
  "RolesKey" : String,
  "SubjectKey" : String
}

```

### YAML

```yaml

  EnabledAPIAccess: Boolean
  IdentityCenterApplicationARN: String
  IdentityCenterInstanceARN: String
  IdentityStoreId: String
  RolesKey: String
  SubjectKey: String

```

## Properties

`EnabledAPIAccess`

Indicates whether IAM Identity Center is enabled for the application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityCenterApplicationARN`

The ARN of the IAM Identity Center application that integrates with Amazon OpenSearch
Service.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\\-]*:[a-z]+:[a-z0-9\\-]*:[0-9]*:[a-z0-9\\-]+\/[a-z0-9\\-]+\/[a-z0-9\\-]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityCenterInstanceARN`

The Amazon Resource Name (ARN) of the IAM Identity Center instance.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\\-]*:[a-z]+:[a-z0-9\\-]*:[0-9]*:[a-z0-9\\-]+\/[a-z0-9\\-]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityStoreId`

The identifier of the IAM Identity Store.

_Required_: No

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$|^[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RolesKey`

Specifies the attribute that contains the backend role identifier (such as group name or
group ID) in IAM Identity Center.

_Required_: No

_Type_: String

_Allowed values_: `GroupName | GroupId`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectKey`

Specifies the attribute that contains the subject identifier (such as username, user ID, or
email) in IAM Identity Center.

_Required_: No

_Type_: String

_Allowed values_: `UserName | UserId | Email`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IAMFederationOptions

Idp

All content copied from https://docs.aws.amazon.com/.
