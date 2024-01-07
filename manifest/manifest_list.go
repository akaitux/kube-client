package manifest

import "github.com/akaitux/kube-client/manifest/releaseutil"

func ListFromYamlDocs(rawManifests string) ([]Manifest, error) {
	var manifests []Manifest

	for _, doc := range releaseutil.SplitManifests(rawManifests) {
		m, err := NewFromYAML(doc)
		if err != nil {
			return nil, err
		}

		if m.HasBasicFields() {
			manifests = append(manifests, m)
		}
	}

	return manifests, nil
}
