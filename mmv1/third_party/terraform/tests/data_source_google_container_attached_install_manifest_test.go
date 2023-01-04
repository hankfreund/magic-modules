package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceGoogleContainerAttachedInstallManifest(t *testing.T) {
	t.Parallel()

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleContainerAttachedInstallManifestConfig(randString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceGoogleContainerAttachedInstallManifestCheck("data.google_container_attached_install_manifest.manifest"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleContainerAttachedInstallManifestCheck(data_source_name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[data_source_name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", data_source_name)
		}

		manifest, ok := ds.Primary.Attributes["manifest"]
		if !ok {
			return fmt.Errorf("cannot find 'manifest' attribute")
		}
		if manifest == "" {
			return fmt.Errorf("install manifest data is empty")
		}
		return nil
	}
}

func testAccDataSourceGoogleContainerAttachedInstallManifestConfig(suffix string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

data "google_container_attached_versions" "versions" {
	location       = "us-west1"
	project        = data.google_project.project.project_id
}

data "google_container_attached_install_manifest" "manifest" {
	location = "us-west1"
	parent = "projects/${data.google_project.project.project_id}/locations/us-west1"
	cluster_id = "test-cluster-%s"
	platform_version = data.google_container_attached_versions.versions.valid_versions[0]
}
`, suffix)
}
