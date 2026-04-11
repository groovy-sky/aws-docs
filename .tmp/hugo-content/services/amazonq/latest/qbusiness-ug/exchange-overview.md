# Microsoft Exchange connector overview

The following table shows the Amazon Q Business Microsoft Exchange connector features and capabilities.

CategoryFeatureLatest ConnectorLegacy Connector**Security****Authentication type**OAuth 2.0 with Client Credentials FlowOAuth 2.0 with Client Credentials Flow**Authentication credentials**

- Microsoft Exchange Client ID

- Microsoft Exchange Client secret

- Microsoft Exchange Client ID

- Microsoft Exchange Client secret

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](exchange-connector.md#exchange-user-management).Yes. For more information, see [ACL crawling](exchange-connector.md#exchange-user-management).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**YesYes**[VPC](connector-concepts.md#connector-vpc)**NoYes**Customer Managed Key (CMK) support**NoYes**Crawl features****Custom metadata**NoNo**Entities**Mail only (automatic)Yes. The following entities are supported:

- Mail

- Calendar

- Attachment

- OneNotes

- Contacts

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes (Automatic)Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](exchange-field-mappings.md).**Filters**Date range onlyYes. The following filters are supported:

- Include/exclue Calendars

- Include/exclude OneNotes

- Include/exclude Contacts

- Include/exclude using file user email ID

- Include/exclude using date

- Include/exclude using email to, from, subjects,
domains

- Include/exclude by file name regex patterns

- Include/exclude by file type regex patterns

**[Sync mode](connector-concepts.md#connector-sync-mode)**Full sync onlySupports full and incremental sync**[File types](doc-types.md)**File typesNoSupports all files supported by Amazon Q.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
