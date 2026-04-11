# Slack connector overview

The following table gives an overview of the Amazon Q Business
Slack connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Token based authentication**Authentication credentials**

- Slack workspace ID

- Either Slack Bot token or User token

User token lets you make API requests on behalf of the
user. Bot token lets you make API requests as a
Slack bot.

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](slack-user-management.md).
**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes **[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**No**Entities**Yes. The following entities are supported:

- Public channels

- Private channels

- Group messages

- Private messages

- Bot messages

- Archived messages

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports default and custom field mappings. For more
information, see [Field mappings](slack-field-mappings.md).**Filters**Yes. The following filters are supported:

- Crawl public channel

- Crawl private channel

- Crawl group messages

- Crawl private messages

- Crawl channel by type

- Crawl channel by name

- Including and excluding content by file type

- Including and excluding content based on file name

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.**[Crawled as a\**
**document](doc-types.md#connector-doc-crawl)**

- Each message

- Each message attachment

- Each channel post

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
