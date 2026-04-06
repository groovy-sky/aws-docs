# Salesforce Online connector overview

The following table gives an overview of the Salesforce Online connector and its
supported features.

CategoryFeatureSupport**Security****Authentication type**OAuth 2.0 with Resource Owner Password Flow

Note that Require Proof Key for Code Exchange (PKCE) is not supported and must be
disabled.**Authentication credentials**

- Salesforce authentication URL

- Username Client secret

- Password username

- Security token

- Consumer key

- Consumer secret

**[Access Control List (ACL)](connector-concepts.md#connector-authorization) crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/salesforce-user-management). **[Identity crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Supported versions**

- API 30-56

- Lightning, Classic

- Sandbox

**Crawl features****Custom objects**Yes**Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Account

- Campaign

- Partner

- Pricebook

- Case

- Contact

- Contract

- Document

- Group

- Idea

- Lead

- Opportunity

- Product

- Profile

- Solution

- Task

- User

- Chatter

- Knowledge Articles

See [What is a document?](connector-doc-crawl.md) for more details on what each
connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more information, see
[Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/salesforce-field-mappings).**Filters**Yes. The following filters are supported:

- Attachment filter for supported entities

- Regex filters for entities

- Inclusion and exclusion filters on file type for Documents

- Inclusion and exclusion filters on File Name and File Type for
Attachments

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports incremental sync only if ACL is turned off, otherwise only full sync
will be used.**[File types](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/doc-types.html)**Supports all files supported by Amazon Q.**[Crawled as a document](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/doc-types.html#connector-doc-crawl)**

- Each account

- Each contact

- Each campaign

- Each contract

- Each case

- Each partner

- Each opportunity

- Each group

- Each lead

- Each user

- Each task

- Each idea

- Each profile

- Each solution

- Each chatter

- Each document

- Each custom entity

- Each knowledge article

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
