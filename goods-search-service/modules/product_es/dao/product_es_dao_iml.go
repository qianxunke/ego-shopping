package dao

import (
	"context"
	"ego-goods-search-service/common"
	"encoding/json"
	"github.com/go-log/log"
	"github.com/olivere/elastic/v7"
	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"github.com/qianxunke/ego-shopping/ego-plugins/elasticsearch"
)

func (dao *productDaoIml) FindById(id string) (product *productProto.Product, err error) {
	product = &productProto.Product{}
	client, err := elasticsearch.MasterEngine()
	if err != nil {
		log.Logf("[FindById] err: ", err.Error())
		return
	}
	esResponse, err := client.Index().Index("ego_shopping").Type("goods").Id(id).Do(context.TODO())
	if err != nil {
		log.Logf("[FindById] err: ", err.Error())
		return
	}
	err = json.Unmarshal([]byte(esResponse.Result), &product)
	return
}

func (dao *productDaoIml) Insert(product *productProto.Product) (err error) {
	client, err := elasticsearch.MasterEngine()
	if err != nil {
		return
	}
	_, err = client.Index().Index(common.ES_INDEX_SHOP).Type(common.ES_TYPE_PRODUCT).Id(product.Id).BodyJson(product).Do(context.Background())
	return
}

func (dao *productDaoIml) SimpleQuery(limit int, pages int, key string, startTime string, endTime string, order string) (rsp *productProto.Out_GetProducts, err error) {
	rsp = &productProto.Out_GetProducts{}
	client, err := elasticsearch.MasterEngine()
	if err != nil {
		return
	}
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewMatchQuery("publish_status", 1))
	//判断key空不空
	if len(key) > 0 {
		boolQuery2 := elastic.NewBoolQuery()
		boolQuery2.Should(elastic.NewMatchQuery("name", key)).
			Should(elastic.NewMatchQuery("description", key)).
			Should(elastic.NewMatchQuery("detail_desc", key)).
			Should(elastic.NewMatchQuery("detail_html", key)).
			Should(elastic.NewMatchQuery("detail_mobile_html", key)).
			Should(elastic.NewMatchQuery("brand_name", key)).
			Should(elastic.NewMatchQuery("product_category_name", key))
		boolQuery.Must(boolQuery2)
	}
	if len(startTime) > 0 && len(endTime) > 0 {
		boolQuery.Filter(elastic.NewRangeQuery("created_time").From(startTime).To(endTime))
	} else if len(startTime) > 0 {
		boolQuery.Filter(elastic.NewRangeQuery("created_time").Gt(startTime))
	} else if len(endTime) > 0 {
		boolQuery.Filter(elastic.NewRangeQuery("created_time").Lt(endTime))
	}
	if len(order) < 0 {
		order = "created_time"
	}
	//先统计
	rsp.Total, err = client.Count().Index(common.ES_INDEX_SHOP).Query(boolQuery).Do(context.TODO())
	if err != nil || rsp.Total <= 0 {
		return rsp, err
	}
	esRsp, err := client.Search(common.ES_INDEX_SHOP).Query(boolQuery).Size(limit).From((pages - 1) * limit).Do(context.TODO())
	if err != nil {
		return rsp, err
	}
	if len(esRsp.Hits.Hits) <= 0 {
		return
	}
	rsp.ProductList = make([]*productProto.Product, len(esRsp.Hits.Hits))
	for i := 0; i < len(esRsp.Hits.Hits); i++ {
		tempjson := esRsp.Hits.Hits[i].Source
		temp := &productProto.Product{}
		if err = json.Unmarshal([]byte(tempjson), &temp); err != nil {
			return
		} else {
			rsp.ProductList[i] = temp
		}
	}
	return
}

func (dao *productDaoIml) Delete(ids []string) (err error) {
	client, err := elasticsearch.MasterEngine()
	if err != nil {
		return
	}
	for i := 0; i < len(ids); i++ {
		_, err = client.Delete().Index(common.ES_INDEX_SHOP).Type(common.ES_TYPE_PRODUCT).Id(ids[i]).Do(context.TODO())
		if err != nil {
			return
		}
	}
	return
}

func (dao *productDaoIml) Update(id string, reqMap map[string]interface{}) (err error) {
	client, err := elasticsearch.MasterEngine()
	if err != nil {
		return
	}
	_, err = client.Update().Index(common.ES_INDEX_SHOP).Type(common.ES_TYPE_PRODUCT).Id(id).Doc(reqMap).Do(context.TODO())
	return
}
