# UserAccessTaskItem

Contains information about a user access task.

## Contents

**app**

The name of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**tenantId**

The ID of the application tenant.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**error**

Error from the task, if any.

Type: [TaskError](api-taskerror.md) object

Required: No

**taskId**

The unique ID of the task.

Type: String

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/useraccesstaskitem.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/useraccesstaskitem.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/useraccesstaskitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserAccessResultItem

ValidationExceptionField

All content copied from https://docs.aws.amazon.com/.
