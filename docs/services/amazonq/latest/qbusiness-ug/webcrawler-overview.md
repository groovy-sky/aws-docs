# Web Crawler connector overview

The following table gives an overview of the Amazon Q Business
Web Crawler connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**

- Basic

- NTLM/Kerberos

- Form

- SAML

###### Note

You don't need authentication to crawl public websites you
have permission to crawl.

**Authentication credentials**

**Basic authentication**

- Website username

- Website password

**NTLM/Kerberos**
**authentication**

- NTLM/Kerberos username

- NTLM/Kerberos password

**Form authentication**

- Login page URL

- Website username

- Website password

- Username field Xpath

- Password field Xpath

- Password button Xpath

- (Optional) Username button Xpath

**SAML authentication**

- Login page URL

- Website username

- Website password

- Username field Xpath

- Password field Xpath

- Password button Xpath

- (Optional) Username button Xpath

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**No**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**No**Crawl features****Custom metadata**Yes**Visual content processing**Yes. Amazon Q Business can extract and index content from images
embedded in webpages and the following supported document types: PDF,
PowerPoint, Microsoft Word (DOCX), Google Slides, Google Docs**Entities**Yes. The following entities are supported:

- Web page

- Attachment

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. For more information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/web-crawler-field-mappings).**Filters**Yes. The following filters are supported:

- Sync specific domains and subdomains

- Include files linked on web pages

- Regex patterns to crawl and index specific URLs

- Regex patterns to crawl and index specific files

- Include web pages by crawl depth

- Specify maximum file size and links per page for Amazon Q to crawl

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and new, modified, or deleted content sync**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Q Web Crawler

Prerequisites
