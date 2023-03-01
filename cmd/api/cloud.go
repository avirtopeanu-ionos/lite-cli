package api

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"litectl/internal"
	"os"
)

var cloud = &cobra.Command{
	Use:   "cloud",
	Short: "Use the Compute API (Ionos Cloud V6)",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := os.ReadFile("api/cloudapi.json")
		if err != nil {
			return fmt.Errorf("failed reading file: %w", err)
		}

		var doc openapi2.T
		if err = json.Unmarshal(input, &doc); err != nil {
			return fmt.Errorf("failed unmarshalling json: %w", err)
		}
		fmt.Printf("valid spec\n")

		fmt.Printf("Sorted: %+v\n", paths.SortByPartCount(maps.Keys(doc.Paths)))

		//for p, pObj := range doc.Paths {
		//	fmt.Printf("%s\n", p)
		//	for op, opObj := range pObj.Operations() {
		//		fmt.Printf("\t->%s: %s\n", op, opObj.Summary)
		//		for _, param := range opObj.Parameters {
		//			if param.In == "path" {
		//				fmt.Printf("\t\t(param-path) %s\n", param.Name)
		//			} else {
		//				fmt.Printf("\t\t(param) %s\n", param.Ref)
		//			}
		//		}
		//		for resp, respObj := range opObj.Responses {
		//			fmt.Printf("\t\t(resp) %s: %s\n", resp, respObj.Ref)
		//		}
		//	}
		//	fmt.Println()
		//}

		return nil
	},
}

// TODO: Path splitter which returns path vars separate of path elems

// TODO: Sort the paths, in this order:
// 1. /datacenters/ <-- Score: 1
// 2. /datacenters/{path} <-- Score: 1.01
// 3. /datacenters/servers <-- Score: 2
// ...
// n. /datacenters/{datacenterId}/servers/{serverId}/upgrade <-- Score: 3.02

// TODO: Add commands for each operation, in the sorted order (to avoid defining e.g. /datacenters/servers before /datacenters which would lead to redefinition of command)

// TODO:
