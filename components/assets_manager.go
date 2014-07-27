package components

type AssetManagerType struct {
	assets []Asseter
}

var assetManagerInstance *AssetManagerType = &AssetManagerType{}
func AssetManager() *AssetManagerType {
	return assetManagerInstance
}

func (this *AssetManagerType) Add(assets ...Asseter) *AssetManagerType {
	this.assets = append(this.assets, assets...)
	return this
}

func (this *AssetManagerType) Js() (string, error) {
	result := ""
	for _, asset := range this.assets {
		content, err := asset.Js()
		if err != nil {
			return "", err
		}
		result += content
	}

	return result, nil
}

func (this *AssetManagerType) Css() (string, error) {
	result := ""

	for _, asset := range this.assets {
		content, err := asset.Css()
		if err != nil {
			return "", err
		}
		result += content
	}

	return result, nil
}
