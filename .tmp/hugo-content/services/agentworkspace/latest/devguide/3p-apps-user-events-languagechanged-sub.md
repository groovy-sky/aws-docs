# Subscribe a callback function when an Amazon Connect Agent Workspace user changes languages

Subscribes a callback function to-be-invoked whenever a user LanguageChanged event
occurs in the Amazon Connect Agent Workspace.

**Signature**

```typescript

onLanguageChanged(handler: UserLanguageChangedHandler)

```

**Usage**

```typescript

const handler: UserLanguageChangedHandler = async (data: UserLanguageChanged) => {
    console.log("User LanguageChange occurred! " + data);
};

settingsClient.onLanguageChanged(handler);

// UserLanguageChanged Structure
{
  language: string;
  previous: {
    language: string;
  };
}

```

**Permissions required:**

```typescript

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getLanguage()

offLanguageChanged()

All content copied from https://docs.aws.amazon.com/.
