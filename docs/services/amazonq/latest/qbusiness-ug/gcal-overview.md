# Google Calendar connector overview (Preview)

The following table gives an overview of the Google Calendar connector and its
supported features.

CategoryFeatureSupport**Security****Authentication type**Google Service Account and OAuth 2.0**Authentication credentials**

**Google service account**

- Google service account

- Admin account email

- Client email

- Private key

OAuth 2.0

- Client ID

- Client secret

- Refresh token

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**No**Entities**Yes. The following entities are supported:

- Calendar

- Events

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports default field mappings. For more information, see
[Field mappings](gmail-field-mappings.md).**Filters**Yes. The following filters are supported:

- Include/Exclude calendars by email address

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
