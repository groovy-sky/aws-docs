---
title: "Set the language of a user in Connect Customer agent workspace"
---

# Set the language of a user in Connect Customer agent workspace

Sets the language preference for the user that's currently logged in to the
Connect Customer agent workspace. The promise resolves once the language change
has been persisted, with the resulting language echoed back in the result.

**Signature**

```typescript

setLanguage(language: Locale): Promise<SetLanguageResult>

```

**Usage**

```typescript

const result: SetLanguageResult = await settingsClient.setLanguage("en_US");

```

**Input**

**Parameter****Type****Description**language _Required_LocaleThe locale to set for the current user. One of
`en_US`, `de_DE`, `es_ES`,
`fr_FR`, `ja_JP`, `it_IT`,
`ko_KR`, `pt_BR`, `zh_CN`, or
`zh_TW`.

**Output - SetLanguageResult**

**Parameter****Type****Description**languageLocale (optional)The locale that was set, echoed from the request.

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getNetworkType()

setVisualMode()

All content copied from https://docs.aws.amazon.com/.
