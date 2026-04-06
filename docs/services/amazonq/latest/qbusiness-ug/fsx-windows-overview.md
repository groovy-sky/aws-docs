# Amazon FSx (Windows) connector overview

The following table gives an overview of the Amazon FSx (Windows) connector and its
supported features.

CategoryFeatureSupport**Security****Authentication type**Ad Server authentication**Authentication credentials**

- Amazon FSx (Windows) username

- Amazon FSx (Windows) password

**[Access Control List (ACL)](connector-concepts.md#connector-authorization) crawling**Yes. For more information, see [ACL\
crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/fsx-user-management.html). **[Identity crawling](connector-concepts.md#connector-identity-crawler)**No**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more information, see
[Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/fsx-field-mappings.html).**Filters**Yes. The following filters are supported:

- Include/exclude by file name

- Include/exclude by file type

- Include/exclude by file path

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon FSx (Windows)

Prerequisites
