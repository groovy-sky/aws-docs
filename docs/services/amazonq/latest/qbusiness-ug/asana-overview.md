# Asana connector overview (Preview)

The following table gives an overview of the Amazon Q Business
Asana connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Service Account and PAT**Authentication credentials**

- Service Account Tokens

- Personal Access Token

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**No. For preview, this connector only scans public Asana projects. For
more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/zendesk-user-management.html). **[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**No**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**No**Entities**Yes. The following entities are supported:

- Project

- Tasks

- Comments

###### Note

Each instance of an entity is crawled as a single
document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. For more information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/zendesk-field-mappings.html).**Filters**Yes. The following filters are supported:

- Specific Projects

- Project inclusion regex pattern

- Project exclusion regex pattern

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[Crawled as a\**
**document](doc-types.md#connector-doc-crawl)**

- Project

- Task

- Comment

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
