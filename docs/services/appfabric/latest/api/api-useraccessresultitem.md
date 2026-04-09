# UserAccessResultItem

Contains information about a user's access to an application.

## Contents

**app**

The name of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**email**

The email address of the target user.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 320.

Pattern: ``[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*``

Required: No

**resultStatus**

The status of the user access result item.

The following states are possible:

- `IN_PROGRESS`: The user access task is in progress.

- `COMPLETED`: The user access task completed successfully.

- `FAILED`: The user access task failed.

- `EXPIRED`: The user access task expired.

Type: String

Valid Values: `IN_PROGRESS | COMPLETED | FAILED | EXPIRED`

Required: No

**taskError**

Contains information about an error returned from a user access task.

Type: [TaskError](api-taskerror.md) object

Required: No

**taskId**

The unique ID of the task.

Type: String

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**tenantDisplayName**

The display name of the tenant.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**tenantId**

The ID of the application tenant.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**userFirstName**

The first name of the user.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**userFullName**

The full name of the user.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**userId**

The unique ID of user.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**userLastName**

The last name of the user.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**userStatus**

The status of the user returned by the application.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/useraccessresultitem.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/useraccessresultitem.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/useraccessresultitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tenant

UserAccessTaskItem

All content copied from https://docs.aws.amazon.com/.
