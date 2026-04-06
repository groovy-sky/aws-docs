# GitHub (Server) connector overview

The following table gives an overview of the Amazon Q Business
GitHub (Server) connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**OAuth token, Personal token**Authentication credentials**

- GitHub token

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](github-cloud-user-management.md).
**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Entities**Yes. The following entities are supported:

- Repository

- Repository Commit

- Issue Document

- Issue Comment

- Issue Attachment

- Pull Request Comment

- Pull request Document

- Pull Request Attachment

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports default and custom field mappings. For more
information, see [Field mappings](github-cloud-field-mappings.md).**Filters**Yes. The following filters are supported:

- Include select repositories

- Include content by specific entities.

- Include specific branched by name

- Include/exclude content by file name, file type, and file
path

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GitHub (Server)

Prerequisites
