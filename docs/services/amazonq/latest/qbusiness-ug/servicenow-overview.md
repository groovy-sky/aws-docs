# ServiceNow Online connector overview

The following table gives an overview of the Amazon Q Business
ServiceNow Online connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Basic, OAuth 2.0 with Resource Owner Password Flow**Authentication credentials**

**Basic**

- ServiceNow Online host URL

- User name

- Password

- ServiceNow Online version

**OAuth 2.0 with Resource Owner Password**
**Flow**

- ServiceNow Online host URL

- User name

- Password

- Client ID

- Client secret

- ServiceNow Online version

###### Important

ServiceNow Online admin privileges required.

**Supported versions**All ServiceNow versions**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](servicenow-user-management.md).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes **[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes. Supports custom fields for all entities.**Entities**Yes. The following entities are supported:

- Knowledge articles

- Knowledge article attachments

- Service catalogs

- Active service catalog items

- Inactive service catalog items

- Service catalog attachments

- Incidents

- Active incidents

- Inactive incidents

- All active incident types

- Resolved incidents

- Open incidents

- Open – Unassigned incidents

- Incident attachments

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](servicenow-field-mappings.md).**Filters**Yes. The following filters are supported:

- Crawl public knowledge articles

- Crawl knowledge articles with filter query

- Crawl knowledge article attachments

- Use regex filters for knowledge articles

- Crawl service catalog items

- Crawl service catalog item attachments

- Use regex filters for service catalog items

- Crawl incidents

- Crawl incident attachments

- Crawl incidents with filter query

- Use regex filters for active and inactive incidents

- Including and excluding content by file type

- Including and excluding content based on file name

- Crawl ACL for knowledge article, service catalogs, and
incidents

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
