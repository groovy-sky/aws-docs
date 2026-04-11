# Box connector overview

The following table gives an overview of the Amazon Q Business
Box connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Token with JWT Auth by Box**Authentication credentials**

- Client ID

- Client secret

- Public Key ID

- Private Key

- Pass Phrase

###### Important

Admin privileges required.

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](box-user-management.md).
**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Files

- Comments

- Tasks

- Web links

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](box-field-mappings.md).**Filters**Yes. The following filters are supported:

- Include web links

- Include comments

- Include tasks

- Include/exclude by file name

- Include/exclude by file type

- Include/exclude by file path

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
