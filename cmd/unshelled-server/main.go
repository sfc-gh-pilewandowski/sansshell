package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/snowflakedb/unshelled/auth/mtls"
	mtlsFlags "github.com/snowflakedb/unshelled/auth/mtls/flags"
	"github.com/snowflakedb/unshelled/server"


	// Import the server modules you want to expose, they automatically register
	_ "github.com/snowflakedb/unshelled/services/exec"
	_ "github.com/snowflakedb/unshelled/services/healthcheck"
	_ "github.com/snowflakedb/unshelled/services/localfile"
)

var (
	//go:embed default-policy.rego
	defaultPolicy string
	hostport      = flag.String("hostport", "localhost:50042", "Where to listen for connections.")
	policyFlag    = flag.String("policy", defaultPolicy, "Local OPA policy governing access.  If empty, use builtin policy.")
	policyFile    = flag.String("policy-file", "", "Path to a file with an OPA policy.  If empty, uses --policy.")
	credSource    = flag.String("credential-source", mtlsFlags.Name(), fmt.Sprintf("Method used to obtain mTLS credentials (one of [%s])", strings.Join(mtls.Loaders(), ",")))
)

// choosePolicy selects an OPA policy based on the flags, or calls log.Fatal if
// an invalid combination is provided.
func choosePolicy() string {
	if *policyFlag != defaultPolicy && *policyFile != "" {
		log.Fatal("Do not set both --policy and --policy-file.")
	}

	var policy string
	if *policyFile != "" {
		pff, err := ioutil.ReadFile(*policyFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Using policy from --policy-file")
		policy = string(pff)
	} else {
		if *policyFlag != defaultPolicy {
			log.Println("Using policy from --policy")
			policy = *policyFlag
		} else {
			log.Println("Using built-in policy")
			policy = defaultPolicy
		}
	}
	return policy
}

func main() {
	flag.Parse()

	policy := choosePolicy()
	ctx := context.Background()

	creds, err := mtls.LoadServerCredentials(ctx, *credSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(*hostport, creds, policy); err != nil {
		log.Fatal(err)
	}
}
