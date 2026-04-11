# Managing users and groups for an Amazon Q Business application using APIs

Amazon Q Business provides APIs to manage users and groups in your Amazon Q Business. You can't configure user management using the
console—Amazon Q Business automatically invokes these API operations
for you when you configure your data source connector connection. You can use these APIs
to implement your own user and group management solution if you create a Amazon Q Business application environment programmatically.

###### Note

As of Dec 17, 2024, Amazon Q Business will recognize all email addresses as
case-insensitive and recognize subaddresses as equivalent to the original email
address. For example, JohnDoe@example.com, johndoe@example.com, and
johndoe+work@example.com will be considered the same email address. For assistance
with applications or to report a concern, contact Support, sign into the
[AWS Support Center](https://console.aws.amazon.com/support/home) .

API actionAPI descriptionRelevant User Guide topic[CreateUser](../api-reference/api-createuser.md)Creates a universally unique identifier (UUID) mapped to a list of
local user ids within an application[User mapping](connector-principal-store.md#user-mapping)[GetUser](../api-reference/api-getuser.md)Describes the universally unique identifier (UUID) associated with a
local user in a data source[User mapping](connector-principal-store.md#user-mapping)[UpdateUser](../api-reference/api-updateuser.md)Updates information associated with a user id[User mapping](connector-principal-store.md#user-mapping)[PutGroup](../api-reference/api-putgroup.md)Creates, or updates, a mapping of users to groups[Group mapping](connector-principal-store.md#group-mapping)[DeleteGroup](../api-reference/api-deletegrooup.md)Deletes a group so that all users and sub groups that belong to the
group can no longer access documents only available to that
group[Group mapping](connector-principal-store.md#group-mapping)[GetGroup](../api-reference/api-getgroup.md)Describes a group by group name[Group mapping](connector-principal-store.md#group-mapping)[ListGroups](../api-reference/api-listgroups.md)Provides a list of groups that are mapped to users[Group mapping](connector-principal-store.md#group-mapping)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Chat and conversation
management

Subscription management

All content copied from https://docs.aws.amazon.com/.
