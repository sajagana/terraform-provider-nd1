package models

type SiteAnpEpgStaticLeaf struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewSchemaSiteAnpEpgStaticleaf(ops, path, paths string, port int) *SiteAnpEpgStaticLeaf {
	var externalepgMap map[string]interface{}
	externalepgMap = map[string]interface{}{
		"path":          paths,
		"portEncapVlan": port,
	}

	return &SiteAnpEpgStaticLeaf{
		Ops:   ops,
		Path:  path,
		Value: externalepgMap,
	}

}

func (externalepgAttributes *SiteAnpEpgStaticLeaf) ToMap() (map[string]interface{}, error) {
	externalepgAttributesMap := make(map[string]interface{})
	A(externalepgAttributesMap, "op", externalepgAttributes.Ops)
	A(externalepgAttributesMap, "path", externalepgAttributes.Path)
	if externalepgAttributes.Value != nil {
		A(externalepgAttributesMap, "value", externalepgAttributes.Value)
	}

	return externalepgAttributesMap, nil
}
