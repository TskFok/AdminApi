package controller

import (
	"encoding/json"
	"github.com/TskFok/AdminApi/global"
	"github.com/TskFok/AdminApi/model"
	"github.com/TskFok/AdminApi/utils/cache"
	"github.com/TskFok/AdminApi/utils/curl"
	"github.com/TskFok/AdminApi/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type resVicuna struct {
	Object string `json:"object,omitempty"`
	Data   []struct {
		Object    string    `json:"object,omitempty"`
		Index     int       `json:"index,omitempty"`
		Embedding []float64 `json:"embedding,omitempty"`
	} `json:"data,omitempty"`
	Model string `json:"model,omitempty"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens,omitempty"`
		TotalTokens  int `json:"total_tokens,omitempty"`
	} `json:"usage"`
	Corpus string
}

func CorpusVicunaList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")

	pg, err := strconv.Atoi(page)

	if err != nil {
		logger.Error(err.Error())
	}

	sz, err := strconv.Atoi(size)

	if err != nil {
		logger.Error(err.Error())
	}

	corpus := &model.CorpusVicuna{}

	ctx.JSON(http.StatusOK, corpus.List(pg, sz))
}

func AddCorpusVicuna(ctx *gin.Context) {
	corpus := ctx.DefaultPostForm("corpus", "")

	if corpus == "" {
		ctx.JSON(http.StatusBadRequest, "请输入语料")
		return
	}

	//使用语料库
	body := make(map[string]interface{})
	body["model"] = "vicuna-13b"
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body["input"] = corpus

	requestion := &res{}
	httpStatus := curl.Post(global.VicunaUrl+"/embeddings", body, header, requestion)

	if httpStatus != http.StatusOK {
		logger.Error("查询失败")
		ctx.JSON(http.StatusBadRequest, "查询失败")
		return
	}

	b, err := json.Marshal(requestion.Data[0].Embedding)

	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	corpusModel := &model.CorpusVicuna{}
	corpusModel.Data = string(b)
	corpusModel.Corpus = corpus
	id := corpusModel.Add(corpusModel)

	i := strconv.Itoa(int(id))

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})

	cacheInfo := make(map[string]interface{})
	cacheInfo["corpus"] = corpus
	cacheInfo["data"] = requestion.Data[0].Embedding

	bd, err := json.Marshal(cacheInfo)
	if err != nil {
		logger.Error(err.Error())
	}

	cache.Set("embeding_new:"+i, string(bd), -1)

}

func UpdateCorpusVicuna(ctx *gin.Context) {
	id := ctx.DefaultPostForm("id", "0")

	if id == "0" {
		ctx.JSON(http.StatusBadRequest, "无效的id")
		return
	}

	i, _ := strconv.Atoi(id)

	corpus := ctx.DefaultPostForm("corpus", "")

	if corpus == "" {
		ctx.JSON(http.StatusBadRequest, "无效的语料")
	}

	//使用语料库
	body := make(map[string]interface{})
	body["model"] = "vicuna-13b"
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body["input"] = corpus

	requestion := &res{}
	httpStatus := curl.Post(global.VicunaUrl+"/embeddings", body, header, requestion)

	if httpStatus != http.StatusOK {
		logger.Error("查询失败")
		ctx.JSON(http.StatusBadRequest, "查询失败")
		return
	}

	b, err := json.Marshal(requestion.Data[0].Embedding)

	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	corpusModel := &model.CorpusVicuna{}
	corpusModel.Id = uint32(i)

	update := make(map[string]interface{})
	update["corpus"] = corpus
	update["data"] = string(b)

	res := corpusModel.Update(corpusModel, update)

	if !res {
		ctx.JSON(http.StatusBadRequest, "修改失败")
		return
	}

	ctx.JSON(http.StatusOK, "修改成功")

	cacheInfo := make(map[string]interface{})
	cacheInfo["corpus"] = corpus
	cacheInfo["data"] = requestion.Data[0].Embedding

	bd, err := json.Marshal(cacheInfo)
	if err != nil {
		logger.Error(err.Error())
	}

	cache.Set("embeding_new:"+id, string(bd), -1)
}

func DelCorpusVicuna(ctx *gin.Context) {
	id := ctx.Param("id")
	logger.Error(id)

	if id == "" {
		ctx.JSON(http.StatusBadRequest, "记录不存在")
		return
	}

	i, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, "id获取失败")

		return
	}

	corpusModel := &model.CorpusVicuna{}
	corpusModel.Id = uint32(i)

	res := corpusModel.Delete()

	if !res {
		ctx.JSON(http.StatusBadRequest, "删除失败")
		return
	}

	ctx.JSON(http.StatusOK, "删除成功")
	cache.Del("embeding_new:" + id)
}
