package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go"
	"github.com/jmespath/go-jmespath"
	"golang.org/x/example/stringutil"
  "sigs.k8s.io/kustomize/kyaml"
  "sigs.k8s.io/kustomize/api"
)

func main() {
	fmt.Println(stringutil.Reverse("My new project!"))
}
