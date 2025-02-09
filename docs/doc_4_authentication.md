# Authenticate using the Atlas Go SDK

The `atlas-sdk-go` library uses Digest authentication. 
You can [create an API key](https://www.mongodb.com/docs/atlas/configure-api-access/#create-an-api-key-in-an-organization) through the Atlas UI or the Atlas CLI.

To learn more about API authentication, see  [Atlas Administration API Authentication](https://www.mongodb.com/docs/atlas/api/api-authentication).

### Use the Atlas Go SDK in Your Code

Construct a new Atlas SDK client, then use the services on the client to
access different parts of the Atlas Admin API. For example:

```go
import "go.mongodb.org/atlas-sdk/v20231115005/admin"

func example() {
	ctx := context.Background()

	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

	sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
	if err != nil {
		log.Fatalf("Error when instantiating new client: %v", err)
	}
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalf("Could not fetch projects: %v", err)
	}
	fmt.Printf("Response status: %v\n", response.Status)
}
```
