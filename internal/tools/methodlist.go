package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/martinohmann/collections-go/internal/validation"
)

func main() {
	methodMap := validation.MethodMap

	signatures := make([]string, 0, len(methodMap))

	var recv validation.Parameter

	for name, sig := range methodMap {
		recv, sig.In = sig.In[0], sig.In[1:]
		s := fmt.Sprintf("func (%s) %s%s", recv, name, sig)
		signatures = append(signatures, s)
	}

	sort.Strings(signatures)

	fmt.Println(strings.Join(signatures, "\n"))
}
