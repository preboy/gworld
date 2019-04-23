package utils

import (
	"core/log"
	"gopkg.in/mgo.v2/bson"
)

// ============================================================================

func CloneBsonObject(obj interface{}) interface{} {
	buf, err := bson.Marshal(obj)
	if err != nil {
		log.Error("marshal object data failed:", err)
		return nil
	}

	out := make(map[string]interface{})
	err = bson.Unmarshal(buf, &out)
	if err != nil {
		log.Error("unmarshal object data failed:", err)
		return nil
	}

	return out
}

func CloneBsonArray(arr interface{}) interface{} {
	obj := struct {
		A interface{}
	}{arr}

	buf, err := bson.Marshal(obj)
	if err != nil {
		log.Error("marshal object data failed:", err)
		return nil
	}

	out := make(map[string]interface{})
	err = bson.Unmarshal(buf, &out)
	if err != nil {
		log.Error("unmarshal object data failed:", err)
		return nil
	}

	return out["a"]
}
