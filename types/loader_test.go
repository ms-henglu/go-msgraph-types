package types

import (
	"log"
	"testing"

	"github.com/ms-henglu/go-msgraph-types/embed"
)

func availableAPIVersions() []string {
	return []string{"v1.0", "beta"}
}

func Test_DefaultMSGraphSchemaLoader(t *testing.T) {
	msgraphTypes := DefaultMSGraphSchemaLoader()
	for _, version := range availableAPIVersions() {
		if msgraphTypes.GetSchema(version) == nil {
			t.Errorf("failed to load azure schema version %s", version)
		}
	}
}

func Test_NewMSGraphSchemaLoader(t *testing.T) {
	msgraphTypes := NewMSGraphSchemaLoader(embed.StaticFiles)
	for _, version := range availableAPIVersions() {
		if msgraphTypes.GetSchema(version) == nil {
			t.Errorf("failed to load azure schema version %s", version)
		}
	}
}

func Test_ListResources(t *testing.T) {
	msgraphTypes := DefaultMSGraphSchemaLoader()
	for _, version := range availableAPIVersions() {
		if len(msgraphTypes.ListResources(version)) == 0 {
			t.Errorf("expect multiple resources but got 0 for version %s", version)
		}
	}
}

func Test_ListAPIVersions(t *testing.T) {
	msgraphTypes := DefaultMSGraphSchemaLoader()
	actual := msgraphTypes.ListAPIVersions()
	expected := availableAPIVersions()
	if len(actual) != len(expected) {
		t.Errorf("expect %d api versions but got %d", len(expected), len(actual))
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expect %s but got %s", expected[i], v)
		}
	}
}

func Test_GetResourceDefinition(t *testing.T) {
	msgraphTypes := DefaultMSGraphSchemaLoader()

	cases := []struct {
		url string
	}{
		{"applications"},
	}

	for _, c := range cases {
		def := msgraphTypes.GetResourceDefinition("v1.0", c.url)
		if def == nil {
			t.Errorf("failed to load resource definition for %s api-version %s", c.url, "v1.0")
		}
	}
}

func Test_AllMSGraphTypes(t *testing.T) {
	msgraphTypes := DefaultMSGraphSchemaLoader()
	for _, apiVersion := range msgraphTypes.ListAPIVersions() {
		for _, url := range msgraphTypes.ListResources(apiVersion) {
			log.Printf("loading resource definition for %s api-version %s", url, apiVersion)
			def := msgraphTypes.GetResourceDefinition(apiVersion, url)
			if def == nil {
				t.Errorf("failed to load resource definition for %s api-version %s", url, apiVersion)
			}
		}
	}
}
