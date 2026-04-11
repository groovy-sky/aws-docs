# SharePoint (Online) connector overview

The following table gives an overview of the Amazon Q Business
SharePoint (Online) connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**

In order from most to least secure authentication type:

- Microsoft Entra ID (formerly Azure AD) App-Only (OAuth 2.0 Certificate) (This is the
preferred method of connection)

- SharePoint App-Only with Client Credentials Flow

- OAuth 2.0 with Resource Owner Password Flow

- Basic authentication

**Authentication credentials**

**Azure App-Only (OAuth 2.0**
**Certificate)**

- Tenant ID

- Certificate path

- Client ID

- Private key

**SharePoint App-Only (OAuth 2.0 with**
**Client Credentials Flow)**

- Tenant ID

- Microsoft Entra ID (formerly Azure AD) Client ID

- Microsoft Entra ID (formerly Azure AD) Client secret

- SharePoint App-Only Client ID

- SharePoint App-Only Client secret

**OAuth 2.0 with Resource Owner Password**
**Flow**

- SharePoint Tenant ID

- SharePoint admin username

- SharePoint admin password

- Client ID

- Client secret

**Basic**

- SharePoint (Online) admin username

- SharePoint (Online) admin password

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](sharepoint-cloud-user-management.md).
**Integration with Identity Provider**
**(IdP)**Yes. Microsoft Entra ID (formerly Azure AD).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom metadata**Yes. Supports custom metadata for **File** entity
only.**Entities**Yes. The following entities are supported:

- Files

- Attachments

- Link

- Pages

- Events

- Comments

- List Data

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](sharepoint-cloud-field-mappings.md).**Filters**Yes. The following filters are supported:

- Include/exclude by **Links**

- Include/exclude by **Pages**

- Include/exclude by **Events**

- Include/exclude by file name

- Include/exclude by file path

- Include/exclude by file type

- Include/exclude by **OneNote Section**
name

- Include/exclude by **OneNote Page**
name

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.**[Crawled as a\**
**document](doc-types.md#connector-doc-crawl)**

- Each event

- Each page

- Each file

- Each link

- Each file attachment

- Each comment

- Each OneNote

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites

All content copied from https://docs.aws.amazon.com/.
