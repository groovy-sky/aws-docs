# Dropbox connector overview

The following table gives an overview of the Amazon Q Business
Dropbox connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**OAuth 2.0 short‑lived access token and refresh token (offline access)**Authentication credentials**OAuth 2.0 short‑lived access token and refresh token (offline access)

- App key

- App secret

- Access token

- Refresh token (recommended)

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](dropbox-user-management.md).
**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Files

- Dropbox Paper

- Dropbox Paper Templates

- Shortcuts

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports default and custom field mappings. For more
information, see [Field mappings](dropbox-field-mappings.md).**Filters**Yes. The following filters are supported:

- Include/ exclude **Files** **Dropbox Paper**, **Dropbox Paper**
**templates**, and
**Shortcuts**.

- Include/exclude content by file name, file type, and file
path.

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
