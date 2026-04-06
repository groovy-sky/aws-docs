# Zendesk connector overview

The following table gives an overview of the Zendesk connector and its
supported features.

CategoryFeatureSupport**Security****Authentication type**OAuth 2.0 with implicit grant authentication, OAuth 2.0 authentication**Authentication credentials**

- OAuth 2.0 with Resource Owner Password Flow - Implicit grant token

- OAuth 2.0 authentication - Zendesk Client ID

- OAuth 2.0 authentication - Zendesk Client secret

- OAuth 2.0 authentication - Zendesk Username

- OAuth 2.0 authentication - Zendesk Password

**[Access Control List (ACL)](connector-concepts.md#connector-authorization) crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/zendesk-user-management.html). **[Identity crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Ticket

- Ticket comment

- Ticket comment attachment

- Community topic

- Community post

- Community post comment

- Article

- Article attachment

- Article comment

###### Note

Each instance of an entity is crawled as a single document.

See [What is a document?](connector-doc-crawl.md) for more details on what each
connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more information, see
[Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/zendesk-field-mappings.html).**Filters**Yes. The following filters are supported:

- Organization name filter

- Crawl tickets

- Crawl ticket comments

- Crawl ticket comment attachments

- Crawl articles

- Crawl article attachments

- Crawl article comments

- Crawl community topics

- Crawl community posts

- Crawl community post comments

- Including and excluding content by file type

- Including content based on a specific date

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.**[Crawled as a document](doc-types.md#connector-doc-crawl)**

- Each ticket

- Each ticket comment

- Each ticket comment attachment

- Each article

- Each article attachment

- Each article comment

- Each community topic

- Each community post

- Each community post comment

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
