package models

type User struct {
}

func (this *User) View(params map[string]interface {}) (map[string]interface {}, error) {
	var data map[string]interface {}
	data = make(map[string]interface{}, 1)
	data["hello"] = 1

	return data, nil
}

