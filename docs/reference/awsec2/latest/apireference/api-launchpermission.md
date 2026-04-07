# LaunchPermission

Describes a launch permission.

## Contents

**Group** (request), **group** (response)

The name of the group.

Type: String

Valid Values: `all`

Required: No

**OrganizationalUnitArn** (request), **organizationalUnitArn** (response)

The Amazon Resource Name (ARN) of an organizational unit (OU).

Type: String

Required: No

**OrganizationArn** (request), **organizationArn** (response)

The Amazon Resource Name (ARN) of an organization.

Type: String

Required: No

**UserId** (request), **userId** (response)

The AWS account ID.

Constraints: Up to 10 000 account IDs can be specified in a single request.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchpermission.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchpermission.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchpermission.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LastError

LaunchPermissionModifications
