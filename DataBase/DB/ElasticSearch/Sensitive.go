package ElasticSearch

type SKey struct {
	Keys []string `json:"keys"`
}

type KeyDocument struct {
	Key string `json:"key"`
}

var Keys string

func init() {
	Keys = `{
    "mappings": {
        "properties": {
            "key": {
                "type": "text",
                "analyzer": "ik_max_words"
            }
        }
    }
}`

}
