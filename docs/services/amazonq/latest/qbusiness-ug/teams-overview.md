# Microsoft Teams connector overview

The following table shows the Amazon Q Business
Microsoft Teams connector features and capabilities.

CategoryFeatureLatest ConnectorLegacy Connector**Security****Authentication type**OAuth 2.0 with Client Credentials FlowOAuth 2.0 with Client Credentials Flow**Authentication credentials**

- Microsoft Teams Client ID

- Microsoft Teams Client secret

- Microsoft Teams Client ID

- Microsoft Teams Client secret

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/teams-user-management.html).Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/teams-user-management.html).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**YesYes**[VPC](connector-concepts.md#connector-vpc)**NoYes**Customer Managed Key (CMK) support**NoYes**Crawl features****Custom metadata**NoYes**Entities**Chat messages, Channel postsYes. The following entities are supported:

- Chat messages

- Chat attachments

- Channel posts

- Channel file attachments

- Wiki

- Meeting chats

- Meeting details

- Meeting notes

- Meeting files

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**NoYes. Supports both default and custom field mappings. For more
information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/teams-field-mappings.html).**Filters**Only Date rangeYes. The following filters are supported:

- Include/exclude using user email

- Include/exclude using team name

- Include/exclude using channel name

- Include/exclude using file name

- Include/exclude using file type

- Chat message

- Chat attachment

- Channel post

- Channel attachment

- Channel wiki

- Calendar meeting

- Meeting chat

- Meeting file

- Meeting note

**[Sync mode](connector-concepts.md#connector-sync-mode)**Full sync onlySupports full and incremental sync

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
