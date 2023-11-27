package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"io"
	"reflect"
	"strings"
	"time"
)

var EsClient *elasticsearch.Client

func CreateIndex(indexName string) error {
	// 执行构建操作,十秒之后会调用执行函数
	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	select {
	case <-withTimeout.Done():
		// 如果超时，取消操作并执行超时后的操作
		return fmt.Errorf("create index timeout")
	default:
		// 执行正常操作
		withContext := EsClient.Indices.Create.WithContext(withTimeout)
		response, err := EsClient.Indices.Create(
			indexName,
			withContext,
		)
		if err != nil {
			return fmt.Errorf("create index error: %s", err)
		}
		if response.StatusCode != 200 {
			return fmt.Errorf("create index error: %s", response.Status())
		}
		return err
	}
}

func getIdByReflect(doc interface{}) (string, error) {
	// 反射获取id
	idFiled := reflect.ValueOf(doc).FieldByName("Id")
	if !idFiled.IsValid() {
		return "", fmt.Errorf("invalid id field")
	}
	kind := idFiled.Kind()
	if kind != reflect.String {
		return "", fmt.Errorf("invalid id is not string type")
	}
	return idFiled.String(), nil
}

func InsertDocument(indexName string, doc interface{}) error {
	// 将结构体转换为 JSON 的字节切片
	marshal, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	id, err := getIdByReflect(doc)
	if err != nil {
		return fmt.Errorf("invalid id field")
	}
	response, err := EsClient.Index(
		indexName,
		// 输入流
		bytes.NewReader(marshal),
		// 上下文
		EsClient.Index.WithContext(context.Background()),
		// 自定义Id
		EsClient.Index.WithDocumentID(id),
	)
	if err != nil {
		return fmt.Errorf("insert document error: %s", err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("insert document error: %s", response.Status())
	}
	return nil
}
func SearchDocuments(indexName string, query string) ([]interface{}, error) {
	//query是查询的条件"{\n  \"query\": {\n    \"match_all\": {}\n  }\n}\n"
	res, err := EsClient.Search(
		// 上下文
		EsClient.Search.WithContext(context.Background()),
		// 索引名字
		EsClient.Search.WithIndex(indexName),
		// 额外的查询条件
		EsClient.Search.WithBody(bytes.NewReader([]byte(query))),
	)
	if err != nil {
		return nil, fmt.Errorf("search document error: %s", err)
	}
	// 关闭输入流
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.IsError() {
		return nil, fmt.Errorf("search document error: %s", res.Status())
	}
	// 结果集
	var result map[string]interface{}
	// 封装结果集
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error in decoding result: %s", err)
	}
	// 结果集中获取命中,强转成数组
	r := result["hits"].(map[string]interface{})["hits"].([]interface{})

	var data []interface{}
	// 只留下_source里面的数据
	for j := 0; j < len(r); j++ {
		i := r[j].(map[string]interface{})["_source"]
		data = append(data, i)
	}
	return data, nil
}

func UpdateDocument(indexName string, docID string, doc interface{}) error {
	// 编码文档数据为 JSON 格式
	// 这个是部分更新
	marshal, err := json.Marshal(map[string]interface{}{"doc": doc})
	if err != nil {
		return fmt.Errorf("marshal document error: %s", err)
	}
	// 执行更新操作，不再指定文档类型
	update, err := EsClient.Update(
		indexName,
		docID,
		bytes.NewReader(marshal),
		EsClient.Update.WithContext(context.Background()))
	if err != nil {
		return fmt.Errorf("update document error: %s", err)
	}
	if update.StatusCode != 200 {
		return fmt.Errorf("update document error: %s", update.Status())
	}
	return nil
}

func DeleteDocument(indexName string, docID string) error {
	response, err := EsClient.Delete(
		indexName,
		docID,
		EsClient.Delete.WithContext(context.Background()))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("delete document error: %s", response.Status())
	}
	return nil
}
func BulkAddDocument(indexName string, data []interface{}) error {
	datas, err := convertToBulkData(indexName, data)
	if err != nil {
		return fmt.Errorf("convert to bulk data error: %s", err)
	}
	res, err := EsClient.Bulk(
		strings.NewReader(datas),
		EsClient.Bulk.WithContext(context.Background()))
	if err != nil {
		return fmt.Errorf("bulk add document error: %s", err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("bulk add document error: %s", res.Status())
	}
	return nil
}

func convertToBulkData(indexName string, data []interface{}) (string, error) {
	var buf bytes.Buffer
	for _, item := range data {
		id, err := getIdByReflect(item)
		if err != nil {
			return "", err
		}
		meta := map[string]interface{}{
			"index": map[string]string{
				"_index": indexName,
				"_id":    id,
			},
		}

		metaLine, err := json.Marshal(meta)
		if err != nil {
			return "", fmt.Errorf("marshal meta error: %s", err)
		}

		dataLine, err := json.Marshal(item)
		if err != nil {
			return "", fmt.Errorf("marshal meta error: %s", err)
		}
		// Append to buffer with newline
		buf.WriteString(string(metaLine) + "\n")
		buf.WriteString(string(dataLine) + "\n")
	}

	return buf.String(), nil
}
